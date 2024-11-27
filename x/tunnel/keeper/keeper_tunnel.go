package keeper

import (
	"fmt"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	feedstypes "github.com/bandprotocol/chain/v3/x/feeds/types"
	"github.com/bandprotocol/chain/v3/x/tunnel/types"
)

// AddTunnel adds a new tunnel
func (k Keeper) AddTunnel(
	ctx sdk.Context,
	route types.RouteI,
	encoder feedstypes.Encoder,
	signalDeviations []types.SignalDeviation,
	interval uint64,
	creator sdk.AccAddress,
) (*types.Tunnel, error) {
	id := k.GetTunnelCount(ctx)
	newID := id + 1

	// generate a new tunnel account as a fee payer
	feePayer, err := k.GenerateTunnelAccount(ctx, fmt.Sprintf("%d", newID))
	if err != nil {
		return nil, err
	}

	// set the prices info
	k.SetLatestPrices(ctx, types.NewLatestPrices(newID, []feedstypes.Price{}, 0))

	// create a new tunnel
	tunnel, err := types.NewTunnel(
		newID,
		0,
		route,
		encoder,
		feePayer.String(),
		signalDeviations,
		interval,
		sdk.NewCoins(),
		false,
		ctx.BlockTime().Unix(),
		creator.String(),
	)
	if err != nil {
		return nil, err
	}
	k.SetTunnel(ctx, tunnel)

	// increment the tunnel count
	k.SetTunnelCount(ctx, newID)

	// Emit an event
	event := sdk.NewEvent(
		types.EventTypeCreateTunnel,
		sdk.NewAttribute(types.AttributeKeyTunnelID, fmt.Sprintf("%d", tunnel.ID)),
		sdk.NewAttribute(types.AttributeKeyInterval, fmt.Sprintf("%d", tunnel.Interval)),
		sdk.NewAttribute(types.AttributeKeyRoute, tunnel.Route.String()),
		sdk.NewAttribute(types.AttributeKeyEncoder, tunnel.Encoder.String()),
		sdk.NewAttribute(types.AttributeKeyFeePayer, tunnel.FeePayer),
		sdk.NewAttribute(types.AttributeKeyIsActive, fmt.Sprintf("%t", tunnel.IsActive)),
		sdk.NewAttribute(types.AttributeKeyCreatedAt, fmt.Sprintf("%d", tunnel.CreatedAt)),
		sdk.NewAttribute(types.AttributeKeyCreator, tunnel.Creator),
	)
	for _, sd := range tunnel.SignalDeviations {
		event = event.AppendAttributes(
			sdk.NewAttribute(types.AttributeKeySignalID, sd.SignalID),
			sdk.NewAttribute(types.AttributeKeySoftDeviationBPS, fmt.Sprintf("%d", sd.SoftDeviationBPS)),
			sdk.NewAttribute(types.AttributeKeyHardDeviationBPS, fmt.Sprintf("%d", sd.HardDeviationBPS)),
		)
	}
	ctx.EventManager().EmitEvent(event)

	return &tunnel, nil
}

// UpdateAndResetTunnel edits a tunnel and reset latest signal price interval.
func (k Keeper) UpdateAndResetTunnel(
	ctx sdk.Context,
	tunnelID uint64,
	signalDeviations []types.SignalDeviation,
	interval uint64,
) error {
	tunnel, err := k.GetTunnel(ctx, tunnelID)
	if err != nil {
		return err
	}

	// edit the tunnel
	tunnel.SignalDeviations = signalDeviations
	tunnel.Interval = interval
	k.SetTunnel(ctx, tunnel)

	// edit the prices info
	k.SetLatestPrices(ctx, types.NewLatestPrices(tunnelID, []feedstypes.Price{}, 0))

	event := sdk.NewEvent(
		types.EventTypeUpdateAndResetTunnel,
		sdk.NewAttribute(types.AttributeKeyTunnelID, fmt.Sprintf("%d", tunnel.ID)),
		sdk.NewAttribute(types.AttributeKeyInterval, fmt.Sprintf("%d", tunnel.Interval)),
	)
	for _, sd := range signalDeviations {
		event = event.AppendAttributes(
			sdk.NewAttribute(types.AttributeKeySignalID, sd.SignalID),
			sdk.NewAttribute(types.AttributeKeySoftDeviationBPS, fmt.Sprintf("%d", sd.SoftDeviationBPS)),
			sdk.NewAttribute(types.AttributeKeyHardDeviationBPS, fmt.Sprintf("%d", sd.HardDeviationBPS)),
		)
	}
	ctx.EventManager().EmitEvent(event)

	return nil
}

// SetTunnelCount sets the tunnel count in the store
func (k Keeper) SetTunnelCount(ctx sdk.Context, count uint64) {
	ctx.KVStore(k.storeKey).Set(types.TunnelCountStoreKey, sdk.Uint64ToBigEndian(count))
}

// GetTunnelCount returns the current number of all tunnels ever existed
func (k Keeper) GetTunnelCount(ctx sdk.Context) uint64 {
	return sdk.BigEndianToUint64(ctx.KVStore(k.storeKey).Get(types.TunnelCountStoreKey))
}

// SetTunnel sets a tunnel in the store
func (k Keeper) SetTunnel(ctx sdk.Context, tunnel types.Tunnel) {
	ctx.KVStore(k.storeKey).Set(types.TunnelStoreKey(tunnel.ID), k.cdc.MustMarshal(&tunnel))
}

// GetTunnel retrieves a tunnel by its ID
func (k Keeper) GetTunnel(ctx sdk.Context, tunnelID uint64) (types.Tunnel, error) {
	bz := ctx.KVStore(k.storeKey).Get(types.TunnelStoreKey(tunnelID))
	if bz == nil {
		return types.Tunnel{}, types.ErrTunnelNotFound.Wrapf("tunnelID: %d", tunnelID)
	}

	var tunnel types.Tunnel
	k.cdc.MustUnmarshal(bz, &tunnel)
	return tunnel, nil
}

// MustGetTunnel retrieves a tunnel by its ID. Panics if the tunnel does not exist.
func (k Keeper) MustGetTunnel(ctx sdk.Context, tunnelID uint64) types.Tunnel {
	tunnel, err := k.GetTunnel(ctx, tunnelID)
	if err != nil {
		panic(err)
	}
	return tunnel
}

// GetTunnels returns all tunnels
func (k Keeper) GetTunnels(ctx sdk.Context) []types.Tunnel {
	var tunnels []types.Tunnel
	iterator := storetypes.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.TunnelStoreKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var tunnel types.Tunnel
		k.cdc.MustUnmarshal(iterator.Value(), &tunnel)
		tunnels = append(tunnels, tunnel)
	}
	return tunnels
}

// SetActiveTunnelID sets the active tunnel ID in the store
func (k Keeper) SetActiveTunnelID(ctx sdk.Context, tunnelID uint64) {
	ctx.KVStore(k.storeKey).Set(types.ActiveTunnelIDStoreKey(tunnelID), []byte{0x01})
}

// DeleteActiveTunnelID deletes the active tunnel ID from the store
func (k Keeper) DeleteActiveTunnelID(ctx sdk.Context, tunnelID uint64) {
	ctx.KVStore(k.storeKey).Delete(types.ActiveTunnelIDStoreKey(tunnelID))
}

// GetActiveTunnelIDs retrieves the active tunnel IDs from the store
func (k Keeper) GetActiveTunnelIDs(ctx sdk.Context) []uint64 {
	var ids []uint64
	iterator := storetypes.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.ActiveTunnelIDStoreKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		id := sdk.BigEndianToUint64(iterator.Key()[1:])
		ids = append(ids, id)
	}
	return ids
}

// ActivateTunnel activates a tunnel
func (k Keeper) ActivateTunnel(ctx sdk.Context, tunnelID uint64) error {
	tunnel, err := k.GetTunnel(ctx, tunnelID)
	if err != nil {
		return err
	}

	// verify if the total deposit meets or exceeds the minimum required deposit
	minDeposit := k.GetParams(ctx).MinDeposit
	if !tunnel.TotalDeposit.IsAllGTE(minDeposit) {
		return types.ErrInsufficientDeposit
	}

	// add the tunnel ID to the active tunnel IDs
	k.SetActiveTunnelID(ctx, tunnelID)

	// set the last interval timestamp to the current block time
	tunnel.IsActive = true
	k.SetTunnel(ctx, tunnel)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeActivateTunnel,
		sdk.NewAttribute(types.AttributeKeyTunnelID, fmt.Sprintf("%d", tunnelID)),
		sdk.NewAttribute(types.AttributeKeyIsActive, fmt.Sprintf("%t", true)),
	))

	return nil
}

// DeactivateTunnel deactivates a tunnel
func (k Keeper) DeactivateTunnel(ctx sdk.Context, tunnelID uint64) error {
	tunnel, err := k.GetTunnel(ctx, tunnelID)
	if err != nil {
		return err
	}

	// remove the tunnel ID from the active tunnel IDs
	k.DeleteActiveTunnelID(ctx, tunnelID)

	// set the last interval timestamp to the current block time
	tunnel.IsActive = false
	k.SetTunnel(ctx, tunnel)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDeactivateTunnel,
		sdk.NewAttribute(types.AttributeKeyTunnelID, fmt.Sprintf("%d", tunnelID)),
		sdk.NewAttribute(types.AttributeKeyIsActive, fmt.Sprintf("%t", tunnel.IsActive)),
	))

	return nil
}

// SetTotalFees sets the total fees in the store
func (k Keeper) SetTotalFees(ctx sdk.Context, totalFee types.TotalFees) {
	ctx.KVStore(k.storeKey).Set(types.TotalFeeStoreKey, k.cdc.MustMarshal(&totalFee))
}

// GetTotalFees retrieves the total fees from the store
func (k Keeper) GetTotalFees(ctx sdk.Context) types.TotalFees {
	bz := ctx.KVStore(k.storeKey).Get(types.TotalFeeStoreKey)
	if bz == nil {
		return types.TotalFees{}
	}

	var totalFee types.TotalFees
	k.cdc.MustUnmarshal(bz, &totalFee)
	return totalFee
}

// HasEnoughFundToCreatePacket checks if the fee payer has enough balance to create a packet
func (k Keeper) HasEnoughFundToCreatePacket(ctx sdk.Context, tunnelID uint64) (bool, error) {
	tunnel, err := k.GetTunnel(ctx, tunnelID)
	if err != nil {
		return false, err
	}

	// get the route fee from the tunnel
	route, err := tunnel.GetRouteValue()
	if err != nil {
		return false, err
	}
	routeFee, err := k.GetRouteFee(ctx, route)
	if err != nil {
		return false, err
	}

	// get the base packet fee and calculate total fee
	basePacketFee := k.GetParams(ctx).BasePacketFee
	totalFee := basePacketFee.Add(routeFee...)

	// compare the fee payer's balance with the total fee
	feePayer, err := sdk.AccAddressFromBech32(tunnel.FeePayer)
	if err != nil {
		return false, err
	}
	balances := k.bankKeeper.SpendableCoins(ctx, feePayer)
	return balances.IsAllGTE(totalFee), nil
}

func (k Keeper) GenerateTunnelAccount(ctx sdk.Context, key string) (sdk.AccAddress, error) {
	header := ctx.BlockHeader()

	buf := []byte(key)
	buf = append(buf, header.AppHash...)
	buf = append(buf, header.DataHash...)

	moduleCred, err := authtypes.NewModuleCredential(types.ModuleName, []byte(types.TunnelAccountsKey), buf)
	if err != nil {
		return nil, err
	}

	tunnelAccAddr := sdk.AccAddress(moduleCred.Address())

	if acc := k.authKeeper.GetAccount(ctx, tunnelAccAddr); acc != nil {
		// this should not happen
		return nil, types.ErrAccountAlreadyExist.Wrapf(
			"existing account for newly generated key account address %s",
			tunnelAccAddr.String(),
		)
	}

	tunnelAcc, err := authtypes.NewBaseAccountWithPubKey(moduleCred)
	if err != nil {
		return nil, err
	}

	k.authKeeper.NewAccount(ctx, tunnelAcc)
	k.authKeeper.SetAccount(ctx, tunnelAcc)

	return tunnelAccAddr, nil
}

// GetRouteFee returns the fee of the given route
func (k Keeper) GetRouteFee(ctx sdk.Context, route types.RouteI) (sdk.Coins, error) {
	switch route.(type) {
	case *types.TSSRoute:
		return sdk.Coins{}, nil
	case *types.IBCRoute:
		return sdk.Coins{}, nil
	default:
		return sdk.Coins{}, types.ErrInvalidRoute
	}
}

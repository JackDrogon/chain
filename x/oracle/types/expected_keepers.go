package types

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"

	"github.com/bandprotocol/chain/v2/pkg/tss"
	tsstypes "github.com/bandprotocol/chain/v2/x/tss/types"
)

// AccountKeeper defines the expected account keeper.
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
}

// BankKeeper defines the expected bank keeper.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
}

// StakingKeeper defines the expected staking keeper.
type StakingKeeper interface {
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI
	IterateBondedValidatorsByPower(
		ctx sdk.Context,
		fn func(index int64, validator stakingtypes.ValidatorI) (stop bool),
	)
	Validator(ctx sdk.Context, address sdk.ValAddress) stakingtypes.ValidatorI
}

// DistrKeeper defines the expected distribution keeper.
type DistrKeeper interface {
	GetCommunityTax(ctx sdk.Context) (percent sdk.Dec)
	GetFeePool(ctx sdk.Context) (feePool distrtypes.FeePool)
	SetFeePool(ctx sdk.Context, feePool distrtypes.FeePool)
	AllocateTokensToValidator(ctx sdk.Context, val stakingtypes.ValidatorI, tokens sdk.DecCoins)
}

// ChannelKeeper defines the expected IBC channel keeper
type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel ibcchanneltypes.Channel, found bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacket(
		ctx sdk.Context,
		chanCap *capabilitytypes.Capability,
		sourcePort string,
		sourceChannel string,
		timeoutHeight ibcclienttypes.Height,
		timeoutTimestamp uint64,
		data []byte,
	) (sequence uint64, err error)
}

// PortKeeper defines the expected IBC port keeper
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
}

// AuthzKeeper defines the expected authz keeper. for query and testing only don't use to create/remove grant on deliver tx
type AuthzKeeper interface {
	DispatchActions(ctx sdk.Context, grantee sdk.AccAddress, msgs []sdk.Msg) ([][]byte, error)
	GetAuthorization(
		ctx sdk.Context,
		grantee sdk.AccAddress,
		granter sdk.AccAddress,
		msgType string,
	) (authz.Authorization, *time.Time)
	GetAuthorizations(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress) ([]authz.Authorization, error)
	SaveGrant(
		ctx sdk.Context,
		grantee, granter sdk.AccAddress,
		authorization authz.Authorization,
		expiration *time.Time,
	) error
	DeleteGrant(ctx sdk.Context, grantee sdk.AccAddress, granter sdk.AccAddress, msgType string) error
	GranterGrants(c context.Context, req *authz.QueryGranterGrantsRequest) (*authz.QueryGranterGrantsResponse, error)
}

// RollingseedKeeper defines the expected rollingseed keeper
type RollingseedKeeper interface {
	GetRollingSeed(ctx sdk.Context) []byte
}

// TSSKeeper defines the expected tss keeper.
type TSSKeeper interface {
	HandleRequestSign(
		ctx sdk.Context,
		groupID tss.GroupID,
		content tsstypes.Content,
		feePayer sdk.AccAddress,
		feeLimit sdk.Coins,
	) (tss.SigningID, error)
	GetSigning(ctx sdk.Context, signingID tss.SigningID) (tsstypes.Signing, error)
}

package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bandprotocol/chain/v2/x/feeds/types"
)

func (k Keeper) ValidateSubmitPricesRequest(ctx sdk.Context, blockTime int64, req *types.MsgSubmitPrices) error {
	vals := k.stakingKeeper.GetBondedValidatorsByPower(ctx)
	isInTop := false
	for _, val := range vals {
		if req.Validator == val.GetOperator().String() {
			isInTop = true
			break
		}
	}
	if !isInTop {
		return types.ErrNotTopValidator
	}

	val, err := sdk.ValAddressFromBech32(req.Validator)
	if err != nil {
		return err
	}

	status := k.oracleKeeper.GetValidatorStatus(ctx, val)
	if !status.IsActive {
		return types.ErrOracleStatusNotActive.Wrapf("val: %s", val.String())
	}

	if types.AbsInt64(req.Timestamp-blockTime) > k.GetParams(ctx).AllowDiffTime {
		return types.ErrInvalidTimestamp.Wrapf(
			"block_time: %d, timestamp: %d",
			blockTime,
			req.Timestamp,
		)
	}
	return nil
}

func (k Keeper) NewPriceValidator(
	ctx sdk.Context,
	blockTime int64,
	price types.SubmitPrice,
	val sdk.ValAddress,
	transitionTime int64,
) (types.PriceValidator, error) {
	s, err := k.GetSymbol(ctx, price.Symbol)
	if err != nil {
		return types.PriceValidator{}, err
	}

	priceVal, err := k.GetPriceValidator(ctx, price.Symbol, val)
	if err == nil && blockTime < priceVal.Timestamp+s.Interval-transitionTime {
		return types.PriceValidator{}, types.ErrPriceTooFast.Wrapf(
			"symbol: %s, old: %d, new: %d, interval: %d",
			price.Symbol,
			priceVal.Timestamp,
			blockTime,
			s.Interval,
		)
	}

	return types.PriceValidator{
		Validator: val.String(),
		Symbol:    price.Symbol,
		Price:     price.Price,
		Timestamp: blockTime,
	}, nil
}

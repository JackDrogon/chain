package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	feedstypes "github.com/bandprotocol/chain/v3/x/feeds/types"
)

// IBCPacket defines the packet sent over the IBC channel
func NewIBCPacketResult(
	tunnelID uint64,
	sequence uint64,
	prices []feedstypes.Price,
	created_at int64,
) IBCPacketResult {
	return IBCPacketResult{
		TunnelID:  tunnelID,
		Sequence:  sequence,
		Prices:    prices,
		CreatedAt: created_at,
	}
}

// GetBytes returns the IBCPacketResult bytes
func (p IBCPacketResult) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}

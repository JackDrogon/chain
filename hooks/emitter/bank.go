package emitter

import (
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/bandprotocol/chain/v3/hooks/common"
)

// handleMsgSend implements emitter handler for MsgSend.
func (h *Hook) handleMsgSend(msg *types.MsgSend) {
	h.AddAccountsInTx(msg.ToAddress)
}

// handleMsgMultiSend implements emitter handler for MsgMultiSend.
func (h *Hook) handleMsgMultiSend(msg *types.MsgMultiSend) {
	for _, output := range msg.Outputs {
		h.AddAccountsInTx(output.Address)
	}
}

// handleMsgEventTypeTransfer implements emitter handler for Msg that has EventTypeTransfer.
func (h *Hook) handleMsgEventTypeTransfer(evMap common.EvMap) {
	h.AddAccountsInTx(
		evMap[types.EventTypeTransfer+"."+types.AttributeKeyRecipient][0],
		evMap[types.EventTypeTransfer+"."+types.AttributeKeySender][0],
	)
}

// handleEventTypeTransfer implements emitter handler for Event that has EventTypeTransfer.
func (h *Hook) handleEventTypeTransfer(evMap common.EvMap) {
	h.AddAccountsInBlock(
		evMap[types.EventTypeTransfer+"."+types.AttributeKeyRecipient][0],
		evMap[types.EventTypeTransfer+"."+types.AttributeKeySender][0],
	)
}

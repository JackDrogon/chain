package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	tsstypes "github.com/bandprotocol/chain/v2/x/tss/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/oracle interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgRequestData{}, "oracle/Request")
	legacy.RegisterAminoMsg(cdc, &MsgReportData{}, "oracle/Report")
	legacy.RegisterAminoMsg(cdc, &MsgCreateDataSource{}, "oracle/CreateDataSource")
	legacy.RegisterAminoMsg(cdc, &MsgEditDataSource{}, "oracle/EditDataSource")
	legacy.RegisterAminoMsg(cdc, &MsgCreateOracleScript{}, "oracle/CreateOracleScript")
	legacy.RegisterAminoMsg(cdc, &MsgEditOracleScript{}, "oracle/EditOracleScript")
	legacy.RegisterAminoMsg(cdc, &MsgActivate{}, "oracle/Activate")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "oracle/UpdateParams")

	cdc.RegisterConcrete(&OracleResultSignatureOrder{}, "oracle/OracleResultSignatureOrder", nil)

}

// RegisterInterfaces register the oracle module interfaces to protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestData{},
		&MsgReportData{},
		&MsgCreateDataSource{},
		&MsgEditDataSource{},
		&MsgCreateOracleScript{},
		&MsgEditOracleScript{},
		&MsgActivate{},
		&MsgUpdateParams{},
	)

	registry.RegisterImplementations(
		(*tsstypes.Content)(nil),
		&OracleResultSignatureOrder{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/oracle module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding.
	//
	// The actual codec used for serialization should be provided to x/oracle and
	// defined at the application level.
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

	// AminoCdc is a amino codec created to support amino json compatible msgs.
	AminoCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

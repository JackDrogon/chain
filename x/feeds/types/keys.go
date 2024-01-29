package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "feeds"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the feeds module
	QuerierRoute = ModuleName
)

var (
	GlobalStoreKeyPrefix = []byte{0x00}

	PriceServiceStoreKey = append(GlobalStoreKeyPrefix, []byte("PriceService")...)

	SymbolStoreKeyPrefix          = []byte{0x01}
	PriceValidatorStoreKeyPrefix  = []byte{0x02}
	PriceStoreKeyPrefix           = []byte{0x03}
	DelegatorSignalStoreKeyPrefix = []byte{0x10}
	SymbolPowerStoreKeyPrefix     = []byte{0x11}

	ParamsKey = []byte{0x05}
)

func DelegatorSignalStoreKey(delegator sdk.AccAddress) []byte {
	return append(DelegatorSignalStoreKeyPrefix, delegator...)
}

func SymbolPowerStoreKey(symbol string) []byte {
	return append(SymbolPowerStoreKeyPrefix, []byte(symbol)...)
}

func SymbolStoreKey(symbol string) []byte {
	return append(SymbolStoreKeyPrefix, []byte(symbol)...)
}

func PriceValidatorsStoreKey(symbol string) []byte {
	return append(PriceValidatorStoreKeyPrefix, []byte(symbol)...)
}

func PriceValidatorStoreKey(symbol string, validator sdk.ValAddress) []byte {
	return append(PriceValidatorsStoreKey(symbol), validator...)
}

func PriceStoreKey(symbol string) []byte {
	return append(PriceStoreKeyPrefix, []byte(symbol)...)
}

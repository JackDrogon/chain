package upgrades

import (
	"github.com/bandprotocol/chain/v2/app/keepers"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// Upgrade defines a struct containing necessary fields that a SoftwareUpgradeProposal
// must have written, in order for the state migration to go smoothly.
// An upgrade must implement this struct, and then set it in the app.go.
// The app.go will then define the handler.
type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`.
	UpgradeName string

	// CreateUpgradeHandler defines the function that creates an upgrade handler.
	CreateUpgradeHandler func(*module.Manager, module.Configurator, AppManager, *keepers.AppKeepers) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades store.StoreUpgrades
}

// BaseAppParamManager defines an interrace that BaseApp is expected to fullfil
// that allows upgrade handlers to modify BaseApp parameters.
type AppManager interface {
	GetConsensusParams(ctx sdk.Context) *abci.ConsensusParams
	StoreConsensusParams(ctx sdk.Context, cp *abci.ConsensusParams)
	GetSubspace(moduleName string) paramstypes.Subspace
}

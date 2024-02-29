package grogucontext

import (
	"sync"
	"time"

	rpcclient "github.com/cometbft/cometbft/rpc/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"

	band "github.com/bandprotocol/chain/v2/app"
	"github.com/bandprotocol/chain/v2/grogu/priceservice"
	"github.com/bandprotocol/chain/v2/x/feeds/types"
)

// Config data structure for grogu daemon.
type Config struct {
	ChainID          string `mapstructure:"chain-id"`          // ChainID of the target chain
	NodeURI          string `mapstructure:"node"`              // Remote RPC URI of BandChain node to connect to
	Validator        string `mapstructure:"validator"`         // The validator address that I'm responsible for
	GasPrices        string `mapstructure:"gas-prices"`        // Gas prices of the transaction
	LogLevel         string `mapstructure:"log-level"`         // Log level of the logger
	PriceService     string `mapstructure:"price-service"`     // PriceService name and URL (example: "PriceService name:URL")
	BroadcastTimeout string `mapstructure:"broadcast-timeout"` // The time that Grogu will wait for tx commit
	RPCPollInterval  string `mapstructure:"rpc-poll-interval"` // The duration of rpc poll interval
	MaxTry           uint64 `mapstructure:"max-try"`           // The maximum number of tries to submit a report transaction
}

// Global instances.
var (
	Cfg              Config
	Kb               keyring.Keyring
	DefaultGroguHome string
	// Proto codec for encoding/decoding proto message
	Cdc = band.MakeEncodingConfig().Marshaler
)

type Context struct {
	Client           rpcclient.Client
	Validator        sdk.ValAddress
	GasPrices        string
	Keys             []*keyring.Record
	PriceService     priceservice.PriceService
	BroadcastTimeout time.Duration
	MaxTry           uint64
	RPCPollInterval  time.Duration

	PendingSymbols    chan []string
	PendingPrices     chan []types.SubmitPrice
	InProgressSymbols *sync.Map
	FreeKeys          chan int64

	Home string
}

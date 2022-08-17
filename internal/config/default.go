package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
	"time"
)

// Default values of configuration options
const (
	// this defines default application name
	defApplicationName = "Motif GraphQL API Server (custom)"

	// defSelfAddress is a default address used as a placeholder
	// for actual API server identification.
	// Please make sure to configure your real key for your API server on the wild.
	defSelfAddress    = "0x2f5538324b488C815db40D02028f01F3D67E726B"
	defSelfPrivateKey = ""

	// EmptyAddress defines an empty address
	EmptyAddress = "0x0000000000000000000000000000000000000000"

	// defServerBind holds default API server binding address
	defServerBind = "localhost:16761"

	// default set of timeouts for the server
	defReadTimeout     = 2
	defWriteTimeout    = 15
	defIdleTimeout     = 1
	defHeaderTimeout   = 1
	defResolverTimeout = 30

	// defServerDomain holds default API server domain address
	defServerDomain = "localhost:16761"

	// defLoggingLevel holds default Logging level
	// See `godoc.org/github.com/op/go-logging` for the full format specification
	// See `golang.org/pkg/time/` for time format specification
	defLoggingLevel = "INFO"

	// defLoggingFormat holds default format of the Logger output
	defLoggingFormat = "%{color}%{level:-8s} %{shortpkg}/%{shortfunc}%{color:reset}: %{message}"

	// defLachesisUrl holds default Lachesis connection string
	defLachesisUrl = "~/.lachesis/data/lachesis.ipc"

	// defMongoUrl holds default MongoDB connection string
	defMongoUrl = "mongodb://localhost:27017"

	// defMongoDatabase holds the default name of the API persistent database
	defMongoDatabase = "motif"

	// defCacheEvictionTime holds default time for in-memory eviction periods
	defCacheEvictionTime = 15 * time.Minute

	// defCacheMax size represents the default max size of the cache in MB
	defCacheMaxSize = 4096

	// defSolCompilerPath represents the default SOL compiler path
	defSolCompilerPath = "/usr/bin/solc"

	// defApiStateOrigin represents the default origin used for API state syncing
	defApiStateOrigin = "https://localhost"

	// defSfcContract is the default address of the SFC contract
	defSfcContract = "0xFC00FACE00000000000000000000000000000000"

	// defStiContract holds deployment address of the Staker Info smart contract.
	defStiContract = "0xbedcb01a0192e90cdae79227ccd5c8195b17d683"

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiFMintAddressProvider = "0x08aa8b88721853c62cc8192da8f24aeb94aaeb66"

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiUniswapCore = EmptyAddress

	// defDefiFMintAddressProvider represents the address of the fMintAddressProvider
	defDefiUniswapRouter = EmptyAddress

	// defTokenLogoFilePath represents the default path to the tokens map file
	defTokenLogoFilePath = "tokens.json"

	// defBlockScanRescanDepth represents the amount of blocks re-scanned on server start
	defBlockScanRescanDepth = 200
)

// default list of API peers
var defApiPeers = []string{"https://localhost:16761/api"}

// defCorsAllowOrigins holds CORS default allowed origins.
var defCorsAllowOrigins = []string{"*"}

// default list of API peers
var defVotingSources = make([]string, 0)

// defERC20Logo defines default no-URL value for ERC20 logo list
var defERC20Logo = map[common.Address]string{
	common.HexToAddress(EmptyAddress): "https://i.ibb.co/RNLvGqm/symbol.png",
}

// applyDefaults sets default values for configuration options.
func applyDefaults(cfg *viper.Viper) {
	// set simple details
	cfg.SetDefault(keyAppName, defApplicationName)
	cfg.SetDefault(keyBindAddress, defServerBind)
	cfg.SetDefault(keyDomainAddress, defServerDomain)
	cfg.SetDefault(keySignatureAddress, defSelfAddress)
	cfg.SetDefault(keySignaturePrivateKey, defSelfPrivateKey)
	cfg.SetDefault(keyLoggingLevel, defLoggingLevel)
	cfg.SetDefault(keyLoggingFormat, defLoggingFormat)
	cfg.SetDefault(keyLachesisUrl, defLachesisUrl)
	cfg.SetDefault(keyMongoUrl, defMongoUrl)
	cfg.SetDefault(keyMongoDatabase, defMongoDatabase)
	cfg.SetDefault(keySolCompilerPath, defSolCompilerPath)
	cfg.SetDefault(keyApiPeers, defApiPeers)
	cfg.SetDefault(keyApiStateOrigin, defApiStateOrigin)
	cfg.SetDefault(keyErc20TokenMapFilePath, defTokenLogoFilePath)
	cfg.SetDefault(keyErc20Logos, defERC20Logo)

	// in-memory cache
	cfg.SetDefault(keyCacheEvictionTime, defCacheEvictionTime)
	cfg.SetDefault(keyCacheMaxSize, defCacheMaxSize)

	// server timeouts
	cfg.SetDefault(keyTimeoutRead, defReadTimeout)
	cfg.SetDefault(keyTimeoutWrite, defWriteTimeout)
	cfg.SetDefault(keyTimeoutHeader, defHeaderTimeout)
	cfg.SetDefault(keyTimeoutIdle, defIdleTimeout)
	cfg.SetDefault(keyTimeoutResolver, defResolverTimeout)

	// no voting sources by default
	cfg.SetDefault(keyVotingSources, defVotingSources)

	// cors
	cfg.SetDefault(keyCorsAllowOrigins, defCorsAllowOrigins)

	// staking configuration defaults
	cfg.SetDefault(keyStakingSfcContract, defSfcContract)
	cfg.SetDefault(keyStakingStiContract, defStiContract)
	cfg.SetDefault(keyStakingTokenizerContract, EmptyAddress)
	cfg.SetDefault(keyStakingERC20Token, EmptyAddress)

	// DeFi configuration
	cfg.SetDefault(keyDefiFMintAddressProvider, defDefiFMintAddressProvider)
	cfg.SetDefault(keyDefiUniswapCore, defDefiUniswapCore)
	cfg.SetDefault(keyDefiUniswapRouter, defDefiUniswapRouter)
}

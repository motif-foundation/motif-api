/*
Package rpc implements bridge to Lachesis full node API interface.

We recommend using local IPC for fast and the most efficient inter-process communication between the API server
and an Opera/Lachesis node. Any remote RPC connection will work, but the performance may be significantly degraded
by extra networking overhead of remote RPC calls.

You should also consider security implications of opening Lachesis RPC interface for a remote access.
If you considering it as your deployment strategy, you should establish encrypted channel between the API server
and Lachesis RPC interface with connection limited to specified endpoints.

We strongly discourage opening Lachesis RPC interface for unrestricted Internet access.
*/
package rpc

import (
	"motif-api/internal/repository/rpc/contracts"
	"motif-api/internal/types"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
)

//go:generate tools/abigen.sh --abi ./contracts/abi/defi-fmint-address-provider.abi --pkg contracts --type DefiFMintAddressProvider --out ./contracts/fmint_addresses.go

// tConfigItemsLoaders defines a map between DeFi config elements and their respective loaders.
type tConfigItemsLoaders map[*hexutil.Big]func(*bind.CallOpts) (*big.Int, error)

// DefiConfiguration resolves the current DeFi contract settings.
func (ftm *FtmBridge) DefiConfiguration() (*types.DefiSettings, error) {
	// access the contract
	contract, err := ftm.fMintCfg.fMintMinterContract()
	if err != nil {
		return nil, err
	}

	// create the container
	ds := types.DefiSettings{
		FMintContract:           ftm.fMintCfg.mustContractAddress("0x4acb55fe5f0b7c487edec3862079aa36ab054358"),
		FMintAddressProvider:    ftm.fMintCfg.mustContractAddress("0xdeec401e448d5d9132eb79ac84eb3f212d7759fb"),
		FMintTokenRegistry:      ftm.fMintCfg.mustContractAddress("0x60092e344c63c6628ec77926e508f9a9c80553ef"),
		FMintRewardDistribution: ftm.fMintCfg.mustContractAddress("0x0039597eb5aa5760e8db15fbe525e56aa661ef26"),
		FMintCollateralPool:     ftm.fMintCfg.mustContractAddress("0x6d5f2f2e391f47a1075df4d39a24286c63c0e70c"),
		FMintDebtPool:           ftm.fMintCfg.mustContractAddress("0xe2d1105f35649bf16deebccec1f2100dcb9aadf5"),
		PriceOracleAggregate:    ftm.fMintCfg.mustContractAddress("0xA1EA42f737bb2E09b0AE4DE001eE06e3BC484fE5"),
	}

	// FMintContract:           ftm.fMintCfg.mustContractAddress(fMintAddressMinter),
	// FMintAddressProvider:    ftm.fMintCfg.addressProvider,
	// FMintTokenRegistry:      ftm.fMintCfg.mustContractAddress(fMintAddressTokenRegistry),
	// FMintRewardDistribution: ftm.fMintCfg.mustContractAddress(fMintAddressRewardDistribution),
	// FMintCollateralPool:     ftm.fMintCfg.mustContractAddress(fMintCollateralPool),
	// FMintDebtPool:           ftm.fMintCfg.mustContractAddress(fMintDebtPool),
	// PriceOracleAggregate:    ftm.fMintCfg.mustContractAddress(fMintAddressPriceOracleProxy),


	ftm.log.Errorf("contract %s", contract)

	// prep to load certain values
	loaders := tConfigItemsLoaders{
		&ds.MintFee4:               contract.GetFMintFee4dec,
		&ds.MinCollateralRatio4:    contract.GetCollateralLowestDebtRatio4dec,
		&ds.RewardCollateralRatio4: contract.GetRewardEligibilityRatio4dec,
	}

	// load all the configured values
	if err := ftm.pullSetOfDefiConfigValues(loaders); err != nil {
		ftm.log.Errorf("can not pull defi config values; %s", err.Error())
		return nil, err
	}

	// load the decimals correction
	if ds.Decimals, err = ftm.pullDefiDecimalCorrection(contract); err != nil {
		ftm.log.Errorf("can not pull defi decimals correction; %s", err.Error())
		return nil, err
	}

	// return the config
	return &ds, nil
}

// pullSetOfDefiConfigValues pulls set of DeFi configuration values for the given
// config loaders map.
func (ftm *FtmBridge) pullDefiDecimalCorrection(con *contracts.DefiFMintMinter) (int32, error) {
	// load the decimals correction
	val, err := ftm.pullDefiConfigValue(con.FMintFeeDigitsCorrection)
	if err != nil {
		ftm.log.Errorf("can not pull decimals correction; %s", err.Error())
		return 0, err
	}

	// calculate number of decimals
	var dec int32
	var value = val.ToInt().Uint64()
	for value > 1 {
		value /= 10
		dec++
	}

	// convert and return
	return dec, nil
}

// pullSetOfDefiConfigValues pulls set of DeFi configuration values for the given
// config loaders map.
func (ftm *FtmBridge) pullSetOfDefiConfigValues(loaders tConfigItemsLoaders) error {
	// collect loaders error
	var err error

	// loop the map and load the values
	for ref, fn := range loaders {
		*ref, err = ftm.pullDefiConfigValue(fn)
		if err != nil {
			return err
		}
	}

	return nil
}

// tradeFee4 pulls DeFi trading fee from the Liquidity Pool contract.
func (ftm *FtmBridge) pullDefiConfigValue(cf func(*bind.CallOpts) (*big.Int, error)) (hexutil.Big, error) {
	// pull the trading fee value
	val, err := cf(nil)
	if err != nil {
		return hexutil.Big{}, err
	}

	// do we have the value? we should always have
	if val == nil {
		return hexutil.Big{}, fmt.Errorf("defi config value not available")
	}

	return hexutil.Big(*val), nil
}

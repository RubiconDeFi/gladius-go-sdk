// Package binds provides wrapper functions for creating instances
// of 'gladius-protocol' contracts and other helper contracts.
package binds

import (
	"github.com/RubiconDeFi/gladius-go-sdk/constants"
	"github.com/RubiconDeFi/gladius-go-sdk/contracts"
	"github.com/ethereum/go-ethereum/common"
)

// Returns an instance of 'RubiconFeeController' contract.
// Can use either 'Address' or 'ChainID' in 'bopt'
func (bopt *BindOpts) FeeController() (ctrl *contracts.FeeController, err error) {
	if bopt.Address == "" {
		bopt.Address = constants.FEE_CTRL_MAPPING[bopt.ChainID]
	}

	ctrl, err = contracts.NewFeeController(
		common.HexToAddress(bopt.Address), bopt.EthClient)

	return
}

// Returns an instance of 'GladiusReactor' contract.
// Can use either 'Address' or 'ChainID' in 'bopt'
func (bopt *BindOpts) GladiusReactor() (rctr *contracts.GladiusReactor, err error) {
	if bopt.Address == "" {
		bopt.Address = constants.
			REACTOR_ADDRESS_MAPPING[bopt.ChainID][constants.Dutch]
	}

	rctr, err = contracts.NewGladiusReactor(
		common.HexToAddress(bopt.Address), bopt.EthClient)

	return
}

// Returns an instance of 'Erc20' contract.
// Can use only 'ChainID' in 'bopt'.
func (bopt *BindOpts) Erc20() (erc20 *contracts.Erc20, err error) {
	if bopt.Address == "" {
		err = EmptyAddress
		return
	}

	erc20, err = contracts.NewErc20(
		common.HexToAddress(bopt.Address), bopt.EthClient)

	return
}

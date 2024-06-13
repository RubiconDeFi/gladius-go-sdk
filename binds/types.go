package binds

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// Arguments for binding functions, required to
// create an instance of the target contract.
type BindOpts struct {
	// Initialized ethclient, that will be bound
	// to the contract, to execute RPC calls.
	EthClient *ethclient.Client

	// Note: provide either 'Address' or 'ChainID,
	// if one is provided the other won't be used.
	// But, only 'gladius-protocol' contracts are
	// able to get addresses from 'ChainID'
	// other ones require 'Address'.

	// Address to which contract's interface will be bound.
	Address string
	// ID of a chain that can be used to get an address of
	// a 'gladius-protocol' contract on that chain.
	ChainID int
}

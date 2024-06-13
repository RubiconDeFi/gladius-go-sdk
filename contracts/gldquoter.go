// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// InputToken is an auto generated low-level Go binding around an user-defined struct.
type InputToken struct {
	Token     common.Address
	Amount    *big.Int
	MaxAmount *big.Int
}

// OrderInfo is an auto generated low-level Go binding around an user-defined struct.
type OrderInfo struct {
	Reactor                      common.Address
	Swapper                      common.Address
	Nonce                        *big.Int
	Deadline                     *big.Int
	AdditionalValidationContract common.Address
	AdditionalValidationData     []byte
}

// OutputToken is an auto generated low-level Go binding around an user-defined struct.
type OutputToken struct {
	Token     common.Address
	Amount    *big.Int
	Recipient common.Address
}

// ResolvedOrder is an auto generated low-level Go binding around an user-defined struct.
type ResolvedOrder struct {
	Info    OrderInfo
	Input   InputToken
	Outputs []OutputToken
	Sig     []byte
	Hash    [32]byte
}

// GladiusOrderQuoterMetaData contains all meta data concerning the GladiusOrderQuoter contract.
var GladiusOrderQuoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OrdersLengthIncorrect\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"}],\"name\":\"getReactor\",\"outputs\":[{\"internalType\":\"contractIGladiusReactor\",\"name\":\"reactor\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIReactor\",\"name\":\"reactor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"contractIValidationCallback\",\"name\":\"additionalValidationContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalValidationData\",\"type\":\"bytes\"}],\"internalType\":\"structOrderInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"internalType\":\"structInputToken\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structOutputToken[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structResolvedOrder\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIReactor\",\"name\":\"reactor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"contractIValidationCallback\",\"name\":\"additionalValidationContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalValidationData\",\"type\":\"bytes\"}],\"internalType\":\"structOrderInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"internalType\":\"structInputToken\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structOutputToken[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structResolvedOrder[]\",\"name\":\"resolvedOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"reactorCallback\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// GladiusOrderQuoterABI is the input ABI used to generate the binding from.
// Deprecated: Use GladiusOrderQuoterMetaData.ABI instead.
var GladiusOrderQuoterABI = GladiusOrderQuoterMetaData.ABI

// GladiusOrderQuoter is an auto generated Go binding around an Ethereum contract.
type GladiusOrderQuoter struct {
	GladiusOrderQuoterCaller     // Read-only binding to the contract
	GladiusOrderQuoterTransactor // Write-only binding to the contract
	GladiusOrderQuoterFilterer   // Log filterer for contract events
}

// GladiusOrderQuoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type GladiusOrderQuoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GladiusOrderQuoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GladiusOrderQuoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GladiusOrderQuoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GladiusOrderQuoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GladiusOrderQuoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GladiusOrderQuoterSession struct {
	Contract     *GladiusOrderQuoter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GladiusOrderQuoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GladiusOrderQuoterCallerSession struct {
	Contract *GladiusOrderQuoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// GladiusOrderQuoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GladiusOrderQuoterTransactorSession struct {
	Contract     *GladiusOrderQuoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// GladiusOrderQuoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type GladiusOrderQuoterRaw struct {
	Contract *GladiusOrderQuoter // Generic contract binding to access the raw methods on
}

// GladiusOrderQuoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GladiusOrderQuoterCallerRaw struct {
	Contract *GladiusOrderQuoterCaller // Generic read-only contract binding to access the raw methods on
}

// GladiusOrderQuoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GladiusOrderQuoterTransactorRaw struct {
	Contract *GladiusOrderQuoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGladiusOrderQuoter creates a new instance of GladiusOrderQuoter, bound to a specific deployed contract.
func NewGladiusOrderQuoter(address common.Address, backend bind.ContractBackend) (*GladiusOrderQuoter, error) {
	contract, err := bindGladiusOrderQuoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GladiusOrderQuoter{GladiusOrderQuoterCaller: GladiusOrderQuoterCaller{contract: contract}, GladiusOrderQuoterTransactor: GladiusOrderQuoterTransactor{contract: contract}, GladiusOrderQuoterFilterer: GladiusOrderQuoterFilterer{contract: contract}}, nil
}

// NewGladiusOrderQuoterCaller creates a new read-only instance of GladiusOrderQuoter, bound to a specific deployed contract.
func NewGladiusOrderQuoterCaller(address common.Address, caller bind.ContractCaller) (*GladiusOrderQuoterCaller, error) {
	contract, err := bindGladiusOrderQuoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GladiusOrderQuoterCaller{contract: contract}, nil
}

// NewGladiusOrderQuoterTransactor creates a new write-only instance of GladiusOrderQuoter, bound to a specific deployed contract.
func NewGladiusOrderQuoterTransactor(address common.Address, transactor bind.ContractTransactor) (*GladiusOrderQuoterTransactor, error) {
	contract, err := bindGladiusOrderQuoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GladiusOrderQuoterTransactor{contract: contract}, nil
}

// NewGladiusOrderQuoterFilterer creates a new log filterer instance of GladiusOrderQuoter, bound to a specific deployed contract.
func NewGladiusOrderQuoterFilterer(address common.Address, filterer bind.ContractFilterer) (*GladiusOrderQuoterFilterer, error) {
	contract, err := bindGladiusOrderQuoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GladiusOrderQuoterFilterer{contract: contract}, nil
}

// bindGladiusOrderQuoter binds a generic wrapper to an already deployed contract.
func bindGladiusOrderQuoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GladiusOrderQuoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GladiusOrderQuoter *GladiusOrderQuoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GladiusOrderQuoter.Contract.GladiusOrderQuoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GladiusOrderQuoter *GladiusOrderQuoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GladiusOrderQuoter.Contract.GladiusOrderQuoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GladiusOrderQuoter *GladiusOrderQuoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GladiusOrderQuoter.Contract.GladiusOrderQuoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GladiusOrderQuoter *GladiusOrderQuoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GladiusOrderQuoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GladiusOrderQuoter *GladiusOrderQuoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GladiusOrderQuoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GladiusOrderQuoter *GladiusOrderQuoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GladiusOrderQuoter.Contract.contract.Transact(opts, method, params...)
}

// GetReactor is a free data retrieval call binding the contract method 0x7671d07b.
//
// Solidity: function getReactor(bytes order) pure returns(address reactor)
func (_GladiusOrderQuoter *GladiusOrderQuoterCaller) GetReactor(opts *bind.CallOpts, order []byte) (common.Address, error) {
	var out []interface{}
	err := _GladiusOrderQuoter.contract.Call(opts, &out, "getReactor", order)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReactor is a free data retrieval call binding the contract method 0x7671d07b.
//
// Solidity: function getReactor(bytes order) pure returns(address reactor)
func (_GladiusOrderQuoter *GladiusOrderQuoterSession) GetReactor(order []byte) (common.Address, error) {
	return _GladiusOrderQuoter.Contract.GetReactor(&_GladiusOrderQuoter.CallOpts, order)
}

// GetReactor is a free data retrieval call binding the contract method 0x7671d07b.
//
// Solidity: function getReactor(bytes order) pure returns(address reactor)
func (_GladiusOrderQuoter *GladiusOrderQuoterCallerSession) GetReactor(order []byte) (common.Address, error) {
	return _GladiusOrderQuoter.Contract.GetReactor(&_GladiusOrderQuoter.CallOpts, order)
}

// ReactorCallback is a free data retrieval call binding the contract method 0x585da628.
//
// Solidity: function reactorCallback(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32)[] resolvedOrders, bytes ) pure returns()
func (_GladiusOrderQuoter *GladiusOrderQuoterCaller) ReactorCallback(opts *bind.CallOpts, resolvedOrders []ResolvedOrder, arg1 []byte) error {
	var out []interface{}
	err := _GladiusOrderQuoter.contract.Call(opts, &out, "reactorCallback", resolvedOrders, arg1)

	if err != nil {
		return err
	}

	return err

}

// ReactorCallback is a free data retrieval call binding the contract method 0x585da628.
//
// Solidity: function reactorCallback(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32)[] resolvedOrders, bytes ) pure returns()
func (_GladiusOrderQuoter *GladiusOrderQuoterSession) ReactorCallback(resolvedOrders []ResolvedOrder, arg1 []byte) error {
	return _GladiusOrderQuoter.Contract.ReactorCallback(&_GladiusOrderQuoter.CallOpts, resolvedOrders, arg1)
}

// ReactorCallback is a free data retrieval call binding the contract method 0x585da628.
//
// Solidity: function reactorCallback(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32)[] resolvedOrders, bytes ) pure returns()
func (_GladiusOrderQuoter *GladiusOrderQuoterCallerSession) ReactorCallback(resolvedOrders []ResolvedOrder, arg1 []byte) error {
	return _GladiusOrderQuoter.Contract.ReactorCallback(&_GladiusOrderQuoter.CallOpts, resolvedOrders, arg1)
}

// Quote is a paid mutator transaction binding the contract method 0x630dca1b.
//
// Solidity: function quote(bytes order, bytes sig, uint256 quantity) returns(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32) result)
func (_GladiusOrderQuoter *GladiusOrderQuoterTransactor) Quote(opts *bind.TransactOpts, order []byte, sig []byte, quantity *big.Int) (*types.Transaction, error) {
	return _GladiusOrderQuoter.contract.Transact(opts, "quote", order, sig, quantity)
}

// Quote is a paid mutator transaction binding the contract method 0x630dca1b.
//
// Solidity: function quote(bytes order, bytes sig, uint256 quantity) returns(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32) result)
func (_GladiusOrderQuoter *GladiusOrderQuoterSession) Quote(order []byte, sig []byte, quantity *big.Int) (*types.Transaction, error) {
	return _GladiusOrderQuoter.Contract.Quote(&_GladiusOrderQuoter.TransactOpts, order, sig, quantity)
}

// Quote is a paid mutator transaction binding the contract method 0x630dca1b.
//
// Solidity: function quote(bytes order, bytes sig, uint256 quantity) returns(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32) result)
func (_GladiusOrderQuoter *GladiusOrderQuoterTransactorSession) Quote(order []byte, sig []byte, quantity *big.Int) (*types.Transaction, error) {
	return _GladiusOrderQuoter.Contract.Quote(&_GladiusOrderQuoter.TransactOpts, order, sig, quantity)
}

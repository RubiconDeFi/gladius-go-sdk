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

// SignedOrder is an auto generated low-level Go binding around an user-defined struct.
type SignedOrder struct {
	Order []byte
	Sig   []byte
}

// GladiusReactorMetaData contains all meta data concerning the GladiusReactor contract.
var GladiusReactorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineBeforeEndTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlinePassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"duplicateToken\",\"type\":\"address\"}],\"name\":\"DuplicateFeeOutput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EndTimeBeforeStartTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectAmounts\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputAndOutputDecay\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEth\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"name\":\"InvalidFeeToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOutLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReactor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NativeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoExclusiveOverride\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderEndTimeBeforeStartTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillUnderflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuantityLtThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RelativeErrTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"filler\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Fill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFeeController\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeController\",\"type\":\"address\"}],\"name\":\"ProtocolFeeControllerSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"quantities\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"executeBatchWithCallback\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"order\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"}],\"internalType\":\"structSignedOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"executeWithCallback\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeController\",\"outputs\":[{\"internalType\":\"contractIProtocolFeeController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPermit2\",\"name\":\"_permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2\",\"outputs\":[{\"internalType\":\"contractIPermit2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFeeController\",\"type\":\"address\"}],\"name\":\"setProtocolFeeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// GladiusReactorABI is the input ABI used to generate the binding from.
// Deprecated: Use GladiusReactorMetaData.ABI instead.
var GladiusReactorABI = GladiusReactorMetaData.ABI

// GladiusReactor is an auto generated Go binding around an Ethereum contract.
type GladiusReactor struct {
	GladiusReactorCaller     // Read-only binding to the contract
	GladiusReactorTransactor // Write-only binding to the contract
	GladiusReactorFilterer   // Log filterer for contract events
}

// GladiusReactorCaller is an auto generated read-only Go binding around an Ethereum contract.
type GladiusReactorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GladiusReactorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GladiusReactorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GladiusReactorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GladiusReactorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GladiusReactorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GladiusReactorSession struct {
	Contract     *GladiusReactor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GladiusReactorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GladiusReactorCallerSession struct {
	Contract *GladiusReactorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GladiusReactorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GladiusReactorTransactorSession struct {
	Contract     *GladiusReactorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GladiusReactorRaw is an auto generated low-level Go binding around an Ethereum contract.
type GladiusReactorRaw struct {
	Contract *GladiusReactor // Generic contract binding to access the raw methods on
}

// GladiusReactorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GladiusReactorCallerRaw struct {
	Contract *GladiusReactorCaller // Generic read-only contract binding to access the raw methods on
}

// GladiusReactorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GladiusReactorTransactorRaw struct {
	Contract *GladiusReactorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGladiusReactor creates a new instance of GladiusReactor, bound to a specific deployed contract.
func NewGladiusReactor(address common.Address, backend bind.ContractBackend) (*GladiusReactor, error) {
	contract, err := bindGladiusReactor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GladiusReactor{GladiusReactorCaller: GladiusReactorCaller{contract: contract}, GladiusReactorTransactor: GladiusReactorTransactor{contract: contract}, GladiusReactorFilterer: GladiusReactorFilterer{contract: contract}}, nil
}

// NewGladiusReactorCaller creates a new read-only instance of GladiusReactor, bound to a specific deployed contract.
func NewGladiusReactorCaller(address common.Address, caller bind.ContractCaller) (*GladiusReactorCaller, error) {
	contract, err := bindGladiusReactor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GladiusReactorCaller{contract: contract}, nil
}

// NewGladiusReactorTransactor creates a new write-only instance of GladiusReactor, bound to a specific deployed contract.
func NewGladiusReactorTransactor(address common.Address, transactor bind.ContractTransactor) (*GladiusReactorTransactor, error) {
	contract, err := bindGladiusReactor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GladiusReactorTransactor{contract: contract}, nil
}

// NewGladiusReactorFilterer creates a new log filterer instance of GladiusReactor, bound to a specific deployed contract.
func NewGladiusReactorFilterer(address common.Address, filterer bind.ContractFilterer) (*GladiusReactorFilterer, error) {
	contract, err := bindGladiusReactor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GladiusReactorFilterer{contract: contract}, nil
}

// bindGladiusReactor binds a generic wrapper to an already deployed contract.
func bindGladiusReactor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GladiusReactorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GladiusReactor *GladiusReactorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GladiusReactor.Contract.GladiusReactorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GladiusReactor *GladiusReactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GladiusReactor.Contract.GladiusReactorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GladiusReactor *GladiusReactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GladiusReactor.Contract.GladiusReactorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GladiusReactor *GladiusReactorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GladiusReactor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GladiusReactor *GladiusReactorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GladiusReactor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GladiusReactor *GladiusReactorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GladiusReactor.Contract.contract.Transact(opts, method, params...)
}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_GladiusReactor *GladiusReactorCaller) FeeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GladiusReactor.contract.Call(opts, &out, "feeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_GladiusReactor *GladiusReactorSession) FeeController() (common.Address, error) {
	return _GladiusReactor.Contract.FeeController(&_GladiusReactor.CallOpts)
}

// FeeController is a free data retrieval call binding the contract method 0x6999b377.
//
// Solidity: function feeController() view returns(address)
func (_GladiusReactor *GladiusReactorCallerSession) FeeController() (common.Address, error) {
	return _GladiusReactor.Contract.FeeController(&_GladiusReactor.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_GladiusReactor *GladiusReactorCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GladiusReactor.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_GladiusReactor *GladiusReactorSession) Initialized() (bool, error) {
	return _GladiusReactor.Contract.Initialized(&_GladiusReactor.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_GladiusReactor *GladiusReactorCallerSession) Initialized() (bool, error) {
	return _GladiusReactor.Contract.Initialized(&_GladiusReactor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GladiusReactor *GladiusReactorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GladiusReactor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GladiusReactor *GladiusReactorSession) Owner() (common.Address, error) {
	return _GladiusReactor.Contract.Owner(&_GladiusReactor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GladiusReactor *GladiusReactorCallerSession) Owner() (common.Address, error) {
	return _GladiusReactor.Contract.Owner(&_GladiusReactor.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_GladiusReactor *GladiusReactorCaller) Permit2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GladiusReactor.contract.Call(opts, &out, "permit2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_GladiusReactor *GladiusReactorSession) Permit2() (common.Address, error) {
	return _GladiusReactor.Contract.Permit2(&_GladiusReactor.CallOpts)
}

// Permit2 is a free data retrieval call binding the contract method 0x12261ee7.
//
// Solidity: function permit2() view returns(address)
func (_GladiusReactor *GladiusReactorCallerSession) Permit2() (common.Address, error) {
	return _GladiusReactor.Contract.Permit2(&_GladiusReactor.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xdc081cc9.
//
// Solidity: function execute((bytes,bytes) order, uint256 quantity) payable returns()
func (_GladiusReactor *GladiusReactorTransactor) Execute(opts *bind.TransactOpts, order SignedOrder, quantity *big.Int) (*types.Transaction, error) {
	return _GladiusReactor.contract.Transact(opts, "execute", order, quantity)
}

// Execute is a paid mutator transaction binding the contract method 0xdc081cc9.
//
// Solidity: function execute((bytes,bytes) order, uint256 quantity) payable returns()
func (_GladiusReactor *GladiusReactorSession) Execute(order SignedOrder, quantity *big.Int) (*types.Transaction, error) {
	return _GladiusReactor.Contract.Execute(&_GladiusReactor.TransactOpts, order, quantity)
}

// Execute is a paid mutator transaction binding the contract method 0xdc081cc9.
//
// Solidity: function execute((bytes,bytes) order, uint256 quantity) payable returns()
func (_GladiusReactor *GladiusReactorTransactorSession) Execute(order SignedOrder, quantity *big.Int) (*types.Transaction, error) {
	return _GladiusReactor.Contract.Execute(&_GladiusReactor.TransactOpts, order, quantity)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x6411e816.
//
// Solidity: function executeBatch((bytes,bytes)[] orders, uint256[] quantities) payable returns()
func (_GladiusReactor *GladiusReactorTransactor) ExecuteBatch(opts *bind.TransactOpts, orders []SignedOrder, quantities []*big.Int) (*types.Transaction, error) {
	return _GladiusReactor.contract.Transact(opts, "executeBatch", orders, quantities)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x6411e816.
//
// Solidity: function executeBatch((bytes,bytes)[] orders, uint256[] quantities) payable returns()
func (_GladiusReactor *GladiusReactorSession) ExecuteBatch(orders []SignedOrder, quantities []*big.Int) (*types.Transaction, error) {
	return _GladiusReactor.Contract.ExecuteBatch(&_GladiusReactor.TransactOpts, orders, quantities)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x6411e816.
//
// Solidity: function executeBatch((bytes,bytes)[] orders, uint256[] quantities) payable returns()
func (_GladiusReactor *GladiusReactorTransactorSession) ExecuteBatch(orders []SignedOrder, quantities []*big.Int) (*types.Transaction, error) {
	return _GladiusReactor.Contract.ExecuteBatch(&_GladiusReactor.TransactOpts, orders, quantities)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0xbccd8d09.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, uint256[] quantities, bytes callbackData) payable returns()
func (_GladiusReactor *GladiusReactorTransactor) ExecuteBatchWithCallback(opts *bind.TransactOpts, orders []SignedOrder, quantities []*big.Int, callbackData []byte) (*types.Transaction, error) {
	return _GladiusReactor.contract.Transact(opts, "executeBatchWithCallback", orders, quantities, callbackData)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0xbccd8d09.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, uint256[] quantities, bytes callbackData) payable returns()
func (_GladiusReactor *GladiusReactorSession) ExecuteBatchWithCallback(orders []SignedOrder, quantities []*big.Int, callbackData []byte) (*types.Transaction, error) {
	return _GladiusReactor.Contract.ExecuteBatchWithCallback(&_GladiusReactor.TransactOpts, orders, quantities, callbackData)
}

// ExecuteBatchWithCallback is a paid mutator transaction binding the contract method 0xbccd8d09.
//
// Solidity: function executeBatchWithCallback((bytes,bytes)[] orders, uint256[] quantities, bytes callbackData) payable returns()
func (_GladiusReactor *GladiusReactorTransactorSession) ExecuteBatchWithCallback(orders []SignedOrder, quantities []*big.Int, callbackData []byte) (*types.Transaction, error) {
	return _GladiusReactor.Contract.ExecuteBatchWithCallback(&_GladiusReactor.TransactOpts, orders, quantities, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0xe767b794.
//
// Solidity: function executeWithCallback((bytes,bytes) order, uint256 quantity, bytes callbackData) payable returns()
func (_GladiusReactor *GladiusReactorTransactor) ExecuteWithCallback(opts *bind.TransactOpts, order SignedOrder, quantity *big.Int, callbackData []byte) (*types.Transaction, error) {
	return _GladiusReactor.contract.Transact(opts, "executeWithCallback", order, quantity, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0xe767b794.
//
// Solidity: function executeWithCallback((bytes,bytes) order, uint256 quantity, bytes callbackData) payable returns()
func (_GladiusReactor *GladiusReactorSession) ExecuteWithCallback(order SignedOrder, quantity *big.Int, callbackData []byte) (*types.Transaction, error) {
	return _GladiusReactor.Contract.ExecuteWithCallback(&_GladiusReactor.TransactOpts, order, quantity, callbackData)
}

// ExecuteWithCallback is a paid mutator transaction binding the contract method 0xe767b794.
//
// Solidity: function executeWithCallback((bytes,bytes) order, uint256 quantity, bytes callbackData) payable returns()
func (_GladiusReactor *GladiusReactorTransactorSession) ExecuteWithCallback(order SignedOrder, quantity *big.Int, callbackData []byte) (*types.Transaction, error) {
	return _GladiusReactor.Contract.ExecuteWithCallback(&_GladiusReactor.TransactOpts, order, quantity, callbackData)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _permit2, address _owner) returns()
func (_GladiusReactor *GladiusReactorTransactor) Initialize(opts *bind.TransactOpts, _permit2 common.Address, _owner common.Address) (*types.Transaction, error) {
	return _GladiusReactor.contract.Transact(opts, "initialize", _permit2, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _permit2, address _owner) returns()
func (_GladiusReactor *GladiusReactorSession) Initialize(_permit2 common.Address, _owner common.Address) (*types.Transaction, error) {
	return _GladiusReactor.Contract.Initialize(&_GladiusReactor.TransactOpts, _permit2, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _permit2, address _owner) returns()
func (_GladiusReactor *GladiusReactorTransactorSession) Initialize(_permit2 common.Address, _owner common.Address) (*types.Transaction, error) {
	return _GladiusReactor.Contract.Initialize(&_GladiusReactor.TransactOpts, _permit2, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_GladiusReactor *GladiusReactorTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _GladiusReactor.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_GladiusReactor *GladiusReactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _GladiusReactor.Contract.SetOwner(&_GladiusReactor.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_GladiusReactor *GladiusReactorTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _GladiusReactor.Contract.SetOwner(&_GladiusReactor.TransactOpts, owner_)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_GladiusReactor *GladiusReactorTransactor) SetProtocolFeeController(opts *bind.TransactOpts, _newFeeController common.Address) (*types.Transaction, error) {
	return _GladiusReactor.contract.Transact(opts, "setProtocolFeeController", _newFeeController)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_GladiusReactor *GladiusReactorSession) SetProtocolFeeController(_newFeeController common.Address) (*types.Transaction, error) {
	return _GladiusReactor.Contract.SetProtocolFeeController(&_GladiusReactor.TransactOpts, _newFeeController)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address _newFeeController) returns()
func (_GladiusReactor *GladiusReactorTransactorSession) SetProtocolFeeController(_newFeeController common.Address) (*types.Transaction, error) {
	return _GladiusReactor.Contract.SetProtocolFeeController(&_GladiusReactor.TransactOpts, _newFeeController)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GladiusReactor *GladiusReactorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GladiusReactor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GladiusReactor *GladiusReactorSession) Receive() (*types.Transaction, error) {
	return _GladiusReactor.Contract.Receive(&_GladiusReactor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GladiusReactor *GladiusReactorTransactorSession) Receive() (*types.Transaction, error) {
	return _GladiusReactor.Contract.Receive(&_GladiusReactor.TransactOpts)
}

// GladiusReactorFillIterator is returned from FilterFill and is used to iterate over the raw logs and unpacked data for Fill events raised by the GladiusReactor contract.
type GladiusReactorFillIterator struct {
	Event *GladiusReactorFill // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GladiusReactorFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GladiusReactorFill)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GladiusReactorFill)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GladiusReactorFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GladiusReactorFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GladiusReactorFill represents a Fill event raised by the GladiusReactor contract.
type GladiusReactorFill struct {
	OrderHash [32]byte
	Filler    common.Address
	Swapper   common.Address
	Nonce     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFill is a free log retrieval operation binding the contract event 0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66.
//
// Solidity: event Fill(bytes32 indexed orderHash, address indexed filler, address indexed swapper, uint256 nonce)
func (_GladiusReactor *GladiusReactorFilterer) FilterFill(opts *bind.FilterOpts, orderHash [][32]byte, filler []common.Address, swapper []common.Address) (*GladiusReactorFillIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}
	var swapperRule []interface{}
	for _, swapperItem := range swapper {
		swapperRule = append(swapperRule, swapperItem)
	}

	logs, sub, err := _GladiusReactor.contract.FilterLogs(opts, "Fill", orderHashRule, fillerRule, swapperRule)
	if err != nil {
		return nil, err
	}
	return &GladiusReactorFillIterator{contract: _GladiusReactor.contract, event: "Fill", logs: logs, sub: sub}, nil
}

// WatchFill is a free log subscription operation binding the contract event 0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66.
//
// Solidity: event Fill(bytes32 indexed orderHash, address indexed filler, address indexed swapper, uint256 nonce)
func (_GladiusReactor *GladiusReactorFilterer) WatchFill(opts *bind.WatchOpts, sink chan<- *GladiusReactorFill, orderHash [][32]byte, filler []common.Address, swapper []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var fillerRule []interface{}
	for _, fillerItem := range filler {
		fillerRule = append(fillerRule, fillerItem)
	}
	var swapperRule []interface{}
	for _, swapperItem := range swapper {
		swapperRule = append(swapperRule, swapperItem)
	}

	logs, sub, err := _GladiusReactor.contract.WatchLogs(opts, "Fill", orderHashRule, fillerRule, swapperRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GladiusReactorFill)
				if err := _GladiusReactor.contract.UnpackLog(event, "Fill", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFill is a log parse operation binding the contract event 0x78ad7ec0e9f89e74012afa58738b6b661c024cb0fd185ee2f616c0a28924bd66.
//
// Solidity: event Fill(bytes32 indexed orderHash, address indexed filler, address indexed swapper, uint256 nonce)
func (_GladiusReactor *GladiusReactorFilterer) ParseFill(log types.Log) (*GladiusReactorFill, error) {
	event := new(GladiusReactorFill)
	if err := _GladiusReactor.contract.UnpackLog(event, "Fill", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GladiusReactorLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the GladiusReactor contract.
type GladiusReactorLogSetOwnerIterator struct {
	Event *GladiusReactorLogSetOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GladiusReactorLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GladiusReactorLogSetOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GladiusReactorLogSetOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GladiusReactorLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GladiusReactorLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GladiusReactorLogSetOwner represents a LogSetOwner event raised by the GladiusReactor contract.
type GladiusReactorLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_GladiusReactor *GladiusReactorFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*GladiusReactorLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GladiusReactor.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GladiusReactorLogSetOwnerIterator{contract: _GladiusReactor.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_GladiusReactor *GladiusReactorFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *GladiusReactorLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GladiusReactor.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GladiusReactorLogSetOwner)
				if err := _GladiusReactor.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogSetOwner is a log parse operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_GladiusReactor *GladiusReactorFilterer) ParseLogSetOwner(log types.Log) (*GladiusReactorLogSetOwner, error) {
	event := new(GladiusReactorLogSetOwner)
	if err := _GladiusReactor.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GladiusReactorProtocolFeeControllerSetIterator is returned from FilterProtocolFeeControllerSet and is used to iterate over the raw logs and unpacked data for ProtocolFeeControllerSet events raised by the GladiusReactor contract.
type GladiusReactorProtocolFeeControllerSetIterator struct {
	Event *GladiusReactorProtocolFeeControllerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GladiusReactorProtocolFeeControllerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GladiusReactorProtocolFeeControllerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GladiusReactorProtocolFeeControllerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GladiusReactorProtocolFeeControllerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GladiusReactorProtocolFeeControllerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GladiusReactorProtocolFeeControllerSet represents a ProtocolFeeControllerSet event raised by the GladiusReactor contract.
type GladiusReactorProtocolFeeControllerSet struct {
	OldFeeController common.Address
	NewFeeController common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeControllerSet is a free log retrieval operation binding the contract event 0xb904ae9529e373e48bc82df4326cceaf1b4c472babf37f5b7dec46fecc6b53e0.
//
// Solidity: event ProtocolFeeControllerSet(address oldFeeController, address newFeeController)
func (_GladiusReactor *GladiusReactorFilterer) FilterProtocolFeeControllerSet(opts *bind.FilterOpts) (*GladiusReactorProtocolFeeControllerSetIterator, error) {

	logs, sub, err := _GladiusReactor.contract.FilterLogs(opts, "ProtocolFeeControllerSet")
	if err != nil {
		return nil, err
	}
	return &GladiusReactorProtocolFeeControllerSetIterator{contract: _GladiusReactor.contract, event: "ProtocolFeeControllerSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeControllerSet is a free log subscription operation binding the contract event 0xb904ae9529e373e48bc82df4326cceaf1b4c472babf37f5b7dec46fecc6b53e0.
//
// Solidity: event ProtocolFeeControllerSet(address oldFeeController, address newFeeController)
func (_GladiusReactor *GladiusReactorFilterer) WatchProtocolFeeControllerSet(opts *bind.WatchOpts, sink chan<- *GladiusReactorProtocolFeeControllerSet) (event.Subscription, error) {

	logs, sub, err := _GladiusReactor.contract.WatchLogs(opts, "ProtocolFeeControllerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GladiusReactorProtocolFeeControllerSet)
				if err := _GladiusReactor.contract.UnpackLog(event, "ProtocolFeeControllerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProtocolFeeControllerSet is a log parse operation binding the contract event 0xb904ae9529e373e48bc82df4326cceaf1b4c472babf37f5b7dec46fecc6b53e0.
//
// Solidity: event ProtocolFeeControllerSet(address oldFeeController, address newFeeController)
func (_GladiusReactor *GladiusReactorFilterer) ParseProtocolFeeControllerSet(log types.Log) (*GladiusReactorProtocolFeeControllerSet, error) {
	event := new(GladiusReactorProtocolFeeControllerSet)
	if err := _GladiusReactor.contract.UnpackLog(event, "ProtocolFeeControllerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

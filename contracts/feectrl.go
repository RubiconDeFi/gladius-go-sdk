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

// NOTE: structs commented out, because they're in './gldquoter.go' already!
// InputToken is an auto generated low-level Go binding around an user-defined struct.
/*type InputToken struct {
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
}*/

// FeeControllerMetaData contains all meta data concerning the FeeController contract.
var FeeControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"applyFee\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"contractIReactor\",\"name\":\"reactor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapper\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"contractIValidationCallback\",\"name\":\"additionalValidationContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalValidationData\",\"type\":\"bytes\"}],\"internalType\":\"structOrderInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmount\",\"type\":\"uint256\"}],\"internalType\":\"structInputToken\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structOutputToken[]\",\"name\":\"outputs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"sig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structResolvedOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"getFeeOutputs\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structOutputToken[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getPairHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"applyFee\",\"type\":\"bool\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FeeControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeControllerMetaData.ABI instead.
var FeeControllerABI = FeeControllerMetaData.ABI

// FeeController is an auto generated Go binding around an Ethereum contract.
type FeeController struct {
	FeeControllerCaller     // Read-only binding to the contract
	FeeControllerTransactor // Write-only binding to the contract
	FeeControllerFilterer   // Log filterer for contract events
}

// FeeControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeControllerSession struct {
	Contract     *FeeController    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeControllerCallerSession struct {
	Contract *FeeControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FeeControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeControllerTransactorSession struct {
	Contract     *FeeControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FeeControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeControllerRaw struct {
	Contract *FeeController // Generic contract binding to access the raw methods on
}

// FeeControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeControllerCallerRaw struct {
	Contract *FeeControllerCaller // Generic read-only contract binding to access the raw methods on
}

// FeeControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeControllerTransactorRaw struct {
	Contract *FeeControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeController creates a new instance of FeeController, bound to a specific deployed contract.
func NewFeeController(address common.Address, backend bind.ContractBackend) (*FeeController, error) {
	contract, err := bindFeeController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeController{FeeControllerCaller: FeeControllerCaller{contract: contract}, FeeControllerTransactor: FeeControllerTransactor{contract: contract}, FeeControllerFilterer: FeeControllerFilterer{contract: contract}}, nil
}

// NewFeeControllerCaller creates a new read-only instance of FeeController, bound to a specific deployed contract.
func NewFeeControllerCaller(address common.Address, caller bind.ContractCaller) (*FeeControllerCaller, error) {
	contract, err := bindFeeController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeControllerCaller{contract: contract}, nil
}

// NewFeeControllerTransactor creates a new write-only instance of FeeController, bound to a specific deployed contract.
func NewFeeControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeControllerTransactor, error) {
	contract, err := bindFeeController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeControllerTransactor{contract: contract}, nil
}

// NewFeeControllerFilterer creates a new log filterer instance of FeeController, bound to a specific deployed contract.
func NewFeeControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeControllerFilterer, error) {
	contract, err := bindFeeController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeControllerFilterer{contract: contract}, nil
}

// bindFeeController binds a generic wrapper to an already deployed contract.
func bindFeeController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeController *FeeControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeController.Contract.FeeControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeController *FeeControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeController.Contract.FeeControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeController *FeeControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeController.Contract.FeeControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeController *FeeControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeController.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeController *FeeControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeController.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeController *FeeControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeController.Contract.contract.Transact(opts, method, params...)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_FeeController *FeeControllerCaller) BASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeController.contract.Call(opts, &out, "BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_FeeController *FeeControllerSession) BASEFEE() (*big.Int, error) {
	return _FeeController.Contract.BASEFEE(&_FeeController.CallOpts)
}

// BASEFEE is a free data retrieval call binding the contract method 0x3d18651e.
//
// Solidity: function BASE_FEE() view returns(uint256)
func (_FeeController *FeeControllerCallerSession) BASEFEE() (*big.Int, error) {
	return _FeeController.Contract.BASEFEE(&_FeeController.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FeeController *FeeControllerCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeController.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FeeController *FeeControllerSession) FeeRecipient() (common.Address, error) {
	return _FeeController.Contract.FeeRecipient(&_FeeController.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_FeeController *FeeControllerCallerSession) FeeRecipient() (common.Address, error) {
	return _FeeController.Contract.FeeRecipient(&_FeeController.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0xcdb5661f.
//
// Solidity: function fees(bytes32 ) view returns(bool applyFee, uint256 fee)
func (_FeeController *FeeControllerCaller) Fees(opts *bind.CallOpts, arg0 [32]byte) (struct {
	ApplyFee bool
	Fee      *big.Int
}, error) {
	var out []interface{}
	err := _FeeController.contract.Call(opts, &out, "fees", arg0)

	outstruct := new(struct {
		ApplyFee bool
		Fee      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ApplyFee = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Fees is a free data retrieval call binding the contract method 0xcdb5661f.
//
// Solidity: function fees(bytes32 ) view returns(bool applyFee, uint256 fee)
func (_FeeController *FeeControllerSession) Fees(arg0 [32]byte) (struct {
	ApplyFee bool
	Fee      *big.Int
}, error) {
	return _FeeController.Contract.Fees(&_FeeController.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0xcdb5661f.
//
// Solidity: function fees(bytes32 ) view returns(bool applyFee, uint256 fee)
func (_FeeController *FeeControllerCallerSession) Fees(arg0 [32]byte) (struct {
	ApplyFee bool
	Fee      *big.Int
}, error) {
	return _FeeController.Contract.Fees(&_FeeController.CallOpts, arg0)
}

// GetFeeOutputs is a free data retrieval call binding the contract method 0x8aa6cf03.
//
// Solidity: function getFeeOutputs(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32) order) view returns((address,uint256,address)[] result)
func (_FeeController *FeeControllerCaller) GetFeeOutputs(opts *bind.CallOpts, order ResolvedOrder) ([]OutputToken, error) {
	var out []interface{}
	err := _FeeController.contract.Call(opts, &out, "getFeeOutputs", order)

	if err != nil {
		return *new([]OutputToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]OutputToken)).(*[]OutputToken)

	return out0, err

}

// GetFeeOutputs is a free data retrieval call binding the contract method 0x8aa6cf03.
//
// Solidity: function getFeeOutputs(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32) order) view returns((address,uint256,address)[] result)
func (_FeeController *FeeControllerSession) GetFeeOutputs(order ResolvedOrder) ([]OutputToken, error) {
	return _FeeController.Contract.GetFeeOutputs(&_FeeController.CallOpts, order)
}

// GetFeeOutputs is a free data retrieval call binding the contract method 0x8aa6cf03.
//
// Solidity: function getFeeOutputs(((address,address,uint256,uint256,address,bytes),(address,uint256,uint256),(address,uint256,address)[],bytes,bytes32) order) view returns((address,uint256,address)[] result)
func (_FeeController *FeeControllerCallerSession) GetFeeOutputs(order ResolvedOrder) ([]OutputToken, error) {
	return _FeeController.Contract.GetFeeOutputs(&_FeeController.CallOpts, order)
}

// GetPairHash is a free data retrieval call binding the contract method 0xd7b7961c.
//
// Solidity: function getPairHash(address tokenIn, address tokenOut) pure returns(bytes32 hash)
func (_FeeController *FeeControllerCaller) GetPairHash(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) ([32]byte, error) {
	var out []interface{}
	err := _FeeController.contract.Call(opts, &out, "getPairHash", tokenIn, tokenOut)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPairHash is a free data retrieval call binding the contract method 0xd7b7961c.
//
// Solidity: function getPairHash(address tokenIn, address tokenOut) pure returns(bytes32 hash)
func (_FeeController *FeeControllerSession) GetPairHash(tokenIn common.Address, tokenOut common.Address) ([32]byte, error) {
	return _FeeController.Contract.GetPairHash(&_FeeController.CallOpts, tokenIn, tokenOut)
}

// GetPairHash is a free data retrieval call binding the contract method 0xd7b7961c.
//
// Solidity: function getPairHash(address tokenIn, address tokenOut) pure returns(bytes32 hash)
func (_FeeController *FeeControllerCallerSession) GetPairHash(tokenIn common.Address, tokenOut common.Address) ([32]byte, error) {
	return _FeeController.Contract.GetPairHash(&_FeeController.CallOpts, tokenIn, tokenOut)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeController *FeeControllerCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeeController.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeController *FeeControllerSession) Initialized() (bool, error) {
	return _FeeController.Contract.Initialized(&_FeeController.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeController *FeeControllerCallerSession) Initialized() (bool, error) {
	return _FeeController.Contract.Initialized(&_FeeController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeController *FeeControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeController.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeController *FeeControllerSession) Owner() (common.Address, error) {
	return _FeeController.Contract.Owner(&_FeeController.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeController *FeeControllerCallerSession) Owner() (common.Address, error) {
	return _FeeController.Contract.Owner(&_FeeController.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _feeRecipient) returns()
func (_FeeController *FeeControllerTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FeeController.contract.Transact(opts, "initialize", _owner, _feeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _feeRecipient) returns()
func (_FeeController *FeeControllerSession) Initialize(_owner common.Address, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FeeController.Contract.Initialize(&_FeeController.TransactOpts, _owner, _feeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _feeRecipient) returns()
func (_FeeController *FeeControllerTransactorSession) Initialize(_owner common.Address, _feeRecipient common.Address) (*types.Transaction, error) {
	return _FeeController.Contract.Initialize(&_FeeController.TransactOpts, _owner, _feeRecipient)
}

// SetFee is a paid mutator transaction binding the contract method 0x9815f084.
//
// Solidity: function setFee(address tokenIn, address tokenOut, uint256 fee, bool applyFee) returns()
func (_FeeController *FeeControllerTransactor) SetFee(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, applyFee bool) (*types.Transaction, error) {
	return _FeeController.contract.Transact(opts, "setFee", tokenIn, tokenOut, fee, applyFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x9815f084.
//
// Solidity: function setFee(address tokenIn, address tokenOut, uint256 fee, bool applyFee) returns()
func (_FeeController *FeeControllerSession) SetFee(tokenIn common.Address, tokenOut common.Address, fee *big.Int, applyFee bool) (*types.Transaction, error) {
	return _FeeController.Contract.SetFee(&_FeeController.TransactOpts, tokenIn, tokenOut, fee, applyFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x9815f084.
//
// Solidity: function setFee(address tokenIn, address tokenOut, uint256 fee, bool applyFee) returns()
func (_FeeController *FeeControllerTransactorSession) SetFee(tokenIn common.Address, tokenOut common.Address, fee *big.Int, applyFee bool) (*types.Transaction, error) {
	return _FeeController.Contract.SetFee(&_FeeController.TransactOpts, tokenIn, tokenOut, fee, applyFee)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_FeeController *FeeControllerTransactor) SetFeeRecipient(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _FeeController.contract.Transact(opts, "setFeeRecipient", recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_FeeController *FeeControllerSession) SetFeeRecipient(recipient common.Address) (*types.Transaction, error) {
	return _FeeController.Contract.SetFeeRecipient(&_FeeController.TransactOpts, recipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address recipient) returns()
func (_FeeController *FeeControllerTransactorSession) SetFeeRecipient(recipient common.Address) (*types.Transaction, error) {
	return _FeeController.Contract.SetFeeRecipient(&_FeeController.TransactOpts, recipient)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_FeeController *FeeControllerTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _FeeController.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_FeeController *FeeControllerSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _FeeController.Contract.SetOwner(&_FeeController.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address owner_) returns()
func (_FeeController *FeeControllerTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _FeeController.Contract.SetOwner(&_FeeController.TransactOpts, owner_)
}

// FeeControllerLogSetOwnerIterator is returned from FilterLogSetOwner and is used to iterate over the raw logs and unpacked data for LogSetOwner events raised by the FeeController contract.
type FeeControllerLogSetOwnerIterator struct {
	Event *FeeControllerLogSetOwner // Event containing the contract specifics and raw log

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
func (it *FeeControllerLogSetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeControllerLogSetOwner)
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
		it.Event = new(FeeControllerLogSetOwner)
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
func (it *FeeControllerLogSetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeControllerLogSetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeControllerLogSetOwner represents a LogSetOwner event raised by the FeeController contract.
type FeeControllerLogSetOwner struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogSetOwner is a free log retrieval operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_FeeController *FeeControllerFilterer) FilterLogSetOwner(opts *bind.FilterOpts, owner []common.Address) (*FeeControllerLogSetOwnerIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FeeController.contract.FilterLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return &FeeControllerLogSetOwnerIterator{contract: _FeeController.contract, event: "LogSetOwner", logs: logs, sub: sub}, nil
}

// WatchLogSetOwner is a free log subscription operation binding the contract event 0xce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed94.
//
// Solidity: event LogSetOwner(address indexed owner)
func (_FeeController *FeeControllerFilterer) WatchLogSetOwner(opts *bind.WatchOpts, sink chan<- *FeeControllerLogSetOwner, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FeeController.contract.WatchLogs(opts, "LogSetOwner", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeControllerLogSetOwner)
				if err := _FeeController.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
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
func (_FeeController *FeeControllerFilterer) ParseLogSetOwner(log types.Log) (*FeeControllerLogSetOwner, error) {
	event := new(FeeControllerLogSetOwner)
	if err := _FeeController.contract.UnpackLog(event, "LogSetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

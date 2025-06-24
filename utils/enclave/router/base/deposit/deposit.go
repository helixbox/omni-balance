// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package base_deposit

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

// BaseDepositMetaData contains all meta data concerning the BaseDeposit contract.
var BaseDepositMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"localToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20BridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ERC20WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHBridgeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHDepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"name\":\"ETHWithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MESSENGER\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OTHER_BRIDGE\",\"outputs\":[{\"internalType\":\"contractStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"bridgeETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositERC20To\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_minGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"depositETHTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_localToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_remoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeBridgeETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l2Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeERC20Withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"finalizeETHWithdrawal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"_messenger\",\"type\":\"address\"},{\"internalType\":\"contractISuperchainConfig\",\"name\":\"_superchainConfig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2TokenBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"contractICrossDomainMessenger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"otherBridge\",\"outputs\":[{\"internalType\":\"contractStandardBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superchainConfig\",\"outputs\":[{\"internalType\":\"contractISuperchainConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BaseDepositABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseDepositMetaData.ABI instead.
var BaseDepositABI = BaseDepositMetaData.ABI

// BaseDeposit is an auto generated Go binding around an Ethereum contract.
type BaseDeposit struct {
	BaseDepositCaller     // Read-only binding to the contract
	BaseDepositTransactor // Write-only binding to the contract
	BaseDepositFilterer   // Log filterer for contract events
}

// BaseDepositCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseDepositCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseDepositTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseDepositTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseDepositFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseDepositFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseDepositSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseDepositSession struct {
	Contract     *BaseDeposit      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseDepositCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseDepositCallerSession struct {
	Contract *BaseDepositCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BaseDepositTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseDepositTransactorSession struct {
	Contract     *BaseDepositTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BaseDepositRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseDepositRaw struct {
	Contract *BaseDeposit // Generic contract binding to access the raw methods on
}

// BaseDepositCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseDepositCallerRaw struct {
	Contract *BaseDepositCaller // Generic read-only contract binding to access the raw methods on
}

// BaseDepositTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseDepositTransactorRaw struct {
	Contract *BaseDepositTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseDeposit creates a new instance of BaseDeposit, bound to a specific deployed contract.
func NewBaseDeposit(address common.Address, backend bind.ContractBackend) (*BaseDeposit, error) {
	contract, err := bindBaseDeposit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseDeposit{BaseDepositCaller: BaseDepositCaller{contract: contract}, BaseDepositTransactor: BaseDepositTransactor{contract: contract}, BaseDepositFilterer: BaseDepositFilterer{contract: contract}}, nil
}

// NewBaseDepositCaller creates a new read-only instance of BaseDeposit, bound to a specific deployed contract.
func NewBaseDepositCaller(address common.Address, caller bind.ContractCaller) (*BaseDepositCaller, error) {
	contract, err := bindBaseDeposit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseDepositCaller{contract: contract}, nil
}

// NewBaseDepositTransactor creates a new write-only instance of BaseDeposit, bound to a specific deployed contract.
func NewBaseDepositTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseDepositTransactor, error) {
	contract, err := bindBaseDeposit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseDepositTransactor{contract: contract}, nil
}

// NewBaseDepositFilterer creates a new log filterer instance of BaseDeposit, bound to a specific deployed contract.
func NewBaseDepositFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseDepositFilterer, error) {
	contract, err := bindBaseDeposit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseDepositFilterer{contract: contract}, nil
}

// bindBaseDeposit binds a generic wrapper to an already deployed contract.
func bindBaseDeposit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BaseDepositMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseDeposit *BaseDepositRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseDeposit.Contract.BaseDepositCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseDeposit *BaseDepositRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BaseDepositTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseDeposit *BaseDepositRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BaseDepositTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseDeposit *BaseDepositCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseDeposit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseDeposit *BaseDepositTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseDeposit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseDeposit *BaseDepositTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseDeposit.Contract.contract.Transact(opts, method, params...)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_BaseDeposit *BaseDepositCaller) MESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_BaseDeposit *BaseDepositSession) MESSENGER() (common.Address, error) {
	return _BaseDeposit.Contract.MESSENGER(&_BaseDeposit.CallOpts)
}

// MESSENGER is a free data retrieval call binding the contract method 0x927ede2d.
//
// Solidity: function MESSENGER() view returns(address)
func (_BaseDeposit *BaseDepositCallerSession) MESSENGER() (common.Address, error) {
	return _BaseDeposit.Contract.MESSENGER(&_BaseDeposit.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_BaseDeposit *BaseDepositCaller) OTHERBRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "OTHER_BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_BaseDeposit *BaseDepositSession) OTHERBRIDGE() (common.Address, error) {
	return _BaseDeposit.Contract.OTHERBRIDGE(&_BaseDeposit.CallOpts)
}

// OTHERBRIDGE is a free data retrieval call binding the contract method 0x7f46ddb2.
//
// Solidity: function OTHER_BRIDGE() view returns(address)
func (_BaseDeposit *BaseDepositCallerSession) OTHERBRIDGE() (common.Address, error) {
	return _BaseDeposit.Contract.OTHERBRIDGE(&_BaseDeposit.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_BaseDeposit *BaseDepositCaller) Deposits(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "deposits", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_BaseDeposit *BaseDepositSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BaseDeposit.Contract.Deposits(&_BaseDeposit.CallOpts, arg0, arg1)
}

// Deposits is a free data retrieval call binding the contract method 0x8f601f66.
//
// Solidity: function deposits(address , address ) view returns(uint256)
func (_BaseDeposit *BaseDepositCallerSession) Deposits(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BaseDeposit.Contract.Deposits(&_BaseDeposit.CallOpts, arg0, arg1)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_BaseDeposit *BaseDepositCaller) L2TokenBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "l2TokenBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_BaseDeposit *BaseDepositSession) L2TokenBridge() (common.Address, error) {
	return _BaseDeposit.Contract.L2TokenBridge(&_BaseDeposit.CallOpts)
}

// L2TokenBridge is a free data retrieval call binding the contract method 0x91c49bf8.
//
// Solidity: function l2TokenBridge() view returns(address)
func (_BaseDeposit *BaseDepositCallerSession) L2TokenBridge() (common.Address, error) {
	return _BaseDeposit.Contract.L2TokenBridge(&_BaseDeposit.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_BaseDeposit *BaseDepositCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_BaseDeposit *BaseDepositSession) Messenger() (common.Address, error) {
	return _BaseDeposit.Contract.Messenger(&_BaseDeposit.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_BaseDeposit *BaseDepositCallerSession) Messenger() (common.Address, error) {
	return _BaseDeposit.Contract.Messenger(&_BaseDeposit.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_BaseDeposit *BaseDepositCaller) OtherBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "otherBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_BaseDeposit *BaseDepositSession) OtherBridge() (common.Address, error) {
	return _BaseDeposit.Contract.OtherBridge(&_BaseDeposit.CallOpts)
}

// OtherBridge is a free data retrieval call binding the contract method 0xc89701a2.
//
// Solidity: function otherBridge() view returns(address)
func (_BaseDeposit *BaseDepositCallerSession) OtherBridge() (common.Address, error) {
	return _BaseDeposit.Contract.OtherBridge(&_BaseDeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseDeposit *BaseDepositCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseDeposit *BaseDepositSession) Paused() (bool, error) {
	return _BaseDeposit.Contract.Paused(&_BaseDeposit.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BaseDeposit *BaseDepositCallerSession) Paused() (bool, error) {
	return _BaseDeposit.Contract.Paused(&_BaseDeposit.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_BaseDeposit *BaseDepositCaller) SuperchainConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "superchainConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_BaseDeposit *BaseDepositSession) SuperchainConfig() (common.Address, error) {
	return _BaseDeposit.Contract.SuperchainConfig(&_BaseDeposit.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_BaseDeposit *BaseDepositCallerSession) SuperchainConfig() (common.Address, error) {
	return _BaseDeposit.Contract.SuperchainConfig(&_BaseDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BaseDeposit *BaseDepositCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BaseDeposit.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BaseDeposit *BaseDepositSession) Version() (string, error) {
	return _BaseDeposit.Contract.Version(&_BaseDeposit.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_BaseDeposit *BaseDepositCallerSession) Version() (string, error) {
	return _BaseDeposit.Contract.Version(&_BaseDeposit.CallOpts)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactor) BridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "bridgeERC20", _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BridgeERC20(&_BaseDeposit.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0x87087623.
//
// Solidity: function bridgeERC20(address _localToken, address _remoteToken, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactorSession) BridgeERC20(_localToken common.Address, _remoteToken common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BridgeERC20(&_BaseDeposit.TransactOpts, _localToken, _remoteToken, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactor) BridgeERC20To(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "bridgeERC20To", _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BridgeERC20To(&_BaseDeposit.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeERC20To is a paid mutator transaction binding the contract method 0x540abf73.
//
// Solidity: function bridgeERC20To(address _localToken, address _remoteToken, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactorSession) BridgeERC20To(_localToken common.Address, _remoteToken common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BridgeERC20To(&_BaseDeposit.TransactOpts, _localToken, _remoteToken, _to, _amount, _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactor) BridgeETH(opts *bind.TransactOpts, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "bridgeETH", _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositSession) BridgeETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BridgeETH(&_BaseDeposit.TransactOpts, _minGasLimit, _extraData)
}

// BridgeETH is a paid mutator transaction binding the contract method 0x09fc8843.
//
// Solidity: function bridgeETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactorSession) BridgeETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BridgeETH(&_BaseDeposit.TransactOpts, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactor) BridgeETHTo(opts *bind.TransactOpts, _to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "bridgeETHTo", _to, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositSession) BridgeETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BridgeETHTo(&_BaseDeposit.TransactOpts, _to, _minGasLimit, _extraData)
}

// BridgeETHTo is a paid mutator transaction binding the contract method 0xe11013dd.
//
// Solidity: function bridgeETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactorSession) BridgeETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.BridgeETHTo(&_BaseDeposit.TransactOpts, _to, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactor) DepositERC20(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "depositERC20", _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.DepositERC20(&_BaseDeposit.TransactOpts, _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x58a997f6.
//
// Solidity: function depositERC20(address _l1Token, address _l2Token, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactorSession) DepositERC20(_l1Token common.Address, _l2Token common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.DepositERC20(&_BaseDeposit.TransactOpts, _l1Token, _l2Token, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactor) DepositERC20To(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "depositERC20To", _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.DepositERC20To(&_BaseDeposit.TransactOpts, _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// DepositERC20To is a paid mutator transaction binding the contract method 0x838b2520.
//
// Solidity: function depositERC20To(address _l1Token, address _l2Token, address _to, uint256 _amount, uint32 _minGasLimit, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactorSession) DepositERC20To(_l1Token common.Address, _l2Token common.Address, _to common.Address, _amount *big.Int, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.DepositERC20To(&_BaseDeposit.TransactOpts, _l1Token, _l2Token, _to, _amount, _minGasLimit, _extraData)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactor) DepositETH(opts *bind.TransactOpts, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "depositETH", _minGasLimit, _extraData)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositSession) DepositETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.DepositETH(&_BaseDeposit.TransactOpts, _minGasLimit, _extraData)
}

// DepositETH is a paid mutator transaction binding the contract method 0xb1a1a882.
//
// Solidity: function depositETH(uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactorSession) DepositETH(_minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.DepositETH(&_BaseDeposit.TransactOpts, _minGasLimit, _extraData)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactor) DepositETHTo(opts *bind.TransactOpts, _to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "depositETHTo", _to, _minGasLimit, _extraData)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositSession) DepositETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.DepositETHTo(&_BaseDeposit.TransactOpts, _to, _minGasLimit, _extraData)
}

// DepositETHTo is a paid mutator transaction binding the contract method 0x9a2ac6d5.
//
// Solidity: function depositETHTo(address _to, uint32 _minGasLimit, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactorSession) DepositETHTo(_to common.Address, _minGasLimit uint32, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.DepositETHTo(&_BaseDeposit.TransactOpts, _to, _minGasLimit, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactor) FinalizeBridgeERC20(opts *bind.TransactOpts, _localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "finalizeBridgeERC20", _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.FinalizeBridgeERC20(&_BaseDeposit.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeERC20 is a paid mutator transaction binding the contract method 0x0166a07a.
//
// Solidity: function finalizeBridgeERC20(address _localToken, address _remoteToken, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactorSession) FinalizeBridgeERC20(_localToken common.Address, _remoteToken common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.FinalizeBridgeERC20(&_BaseDeposit.TransactOpts, _localToken, _remoteToken, _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactor) FinalizeBridgeETH(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "finalizeBridgeETH", _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositSession) FinalizeBridgeETH(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.FinalizeBridgeETH(&_BaseDeposit.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeBridgeETH is a paid mutator transaction binding the contract method 0x1635f5fd.
//
// Solidity: function finalizeBridgeETH(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactorSession) FinalizeBridgeETH(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.FinalizeBridgeETH(&_BaseDeposit.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactor) FinalizeERC20Withdrawal(opts *bind.TransactOpts, _l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "finalizeERC20Withdrawal", _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.FinalizeERC20Withdrawal(&_BaseDeposit.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// FinalizeERC20Withdrawal is a paid mutator transaction binding the contract method 0xa9f9e675.
//
// Solidity: function finalizeERC20Withdrawal(address _l1Token, address _l2Token, address _from, address _to, uint256 _amount, bytes _extraData) returns()
func (_BaseDeposit *BaseDepositTransactorSession) FinalizeERC20Withdrawal(_l1Token common.Address, _l2Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.FinalizeERC20Withdrawal(&_BaseDeposit.TransactOpts, _l1Token, _l2Token, _from, _to, _amount, _extraData)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactor) FinalizeETHWithdrawal(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "finalizeETHWithdrawal", _from, _to, _amount, _extraData)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositSession) FinalizeETHWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.FinalizeETHWithdrawal(&_BaseDeposit.TransactOpts, _from, _to, _amount, _extraData)
}

// FinalizeETHWithdrawal is a paid mutator transaction binding the contract method 0x1532ec34.
//
// Solidity: function finalizeETHWithdrawal(address _from, address _to, uint256 _amount, bytes _extraData) payable returns()
func (_BaseDeposit *BaseDepositTransactorSession) FinalizeETHWithdrawal(_from common.Address, _to common.Address, _amount *big.Int, _extraData []byte) (*types.Transaction, error) {
	return _BaseDeposit.Contract.FinalizeETHWithdrawal(&_BaseDeposit.TransactOpts, _from, _to, _amount, _extraData)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _superchainConfig) returns()
func (_BaseDeposit *BaseDepositTransactor) Initialize(opts *bind.TransactOpts, _messenger common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _BaseDeposit.contract.Transact(opts, "initialize", _messenger, _superchainConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _superchainConfig) returns()
func (_BaseDeposit *BaseDepositSession) Initialize(_messenger common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _BaseDeposit.Contract.Initialize(&_BaseDeposit.TransactOpts, _messenger, _superchainConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messenger, address _superchainConfig) returns()
func (_BaseDeposit *BaseDepositTransactorSession) Initialize(_messenger common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _BaseDeposit.Contract.Initialize(&_BaseDeposit.TransactOpts, _messenger, _superchainConfig)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseDeposit *BaseDepositTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseDeposit.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseDeposit *BaseDepositSession) Receive() (*types.Transaction, error) {
	return _BaseDeposit.Contract.Receive(&_BaseDeposit.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BaseDeposit *BaseDepositTransactorSession) Receive() (*types.Transaction, error) {
	return _BaseDeposit.Contract.Receive(&_BaseDeposit.TransactOpts)
}

// BaseDepositERC20BridgeFinalizedIterator is returned from FilterERC20BridgeFinalized and is used to iterate over the raw logs and unpacked data for ERC20BridgeFinalized events raised by the BaseDeposit contract.
type BaseDepositERC20BridgeFinalizedIterator struct {
	Event *BaseDepositERC20BridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BaseDepositERC20BridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositERC20BridgeFinalized)
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
		it.Event = new(BaseDepositERC20BridgeFinalized)
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
func (it *BaseDepositERC20BridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositERC20BridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositERC20BridgeFinalized represents a ERC20BridgeFinalized event raised by the BaseDeposit contract.
type BaseDepositERC20BridgeFinalized struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeFinalized is a free log retrieval operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) FilterERC20BridgeFinalized(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*BaseDepositERC20BridgeFinalizedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BaseDepositERC20BridgeFinalizedIterator{contract: _BaseDeposit.contract, event: "ERC20BridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeFinalized is a free log subscription operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) WatchERC20BridgeFinalized(opts *bind.WatchOpts, sink chan<- *BaseDepositERC20BridgeFinalized, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "ERC20BridgeFinalized", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositERC20BridgeFinalized)
				if err := _BaseDeposit.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
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

// ParseERC20BridgeFinalized is a log parse operation binding the contract event 0xd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd.
//
// Solidity: event ERC20BridgeFinalized(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) ParseERC20BridgeFinalized(log types.Log) (*BaseDepositERC20BridgeFinalized, error) {
	event := new(BaseDepositERC20BridgeFinalized)
	if err := _BaseDeposit.contract.UnpackLog(event, "ERC20BridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseDepositERC20BridgeInitiatedIterator is returned from FilterERC20BridgeInitiated and is used to iterate over the raw logs and unpacked data for ERC20BridgeInitiated events raised by the BaseDeposit contract.
type BaseDepositERC20BridgeInitiatedIterator struct {
	Event *BaseDepositERC20BridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BaseDepositERC20BridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositERC20BridgeInitiated)
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
		it.Event = new(BaseDepositERC20BridgeInitiated)
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
func (it *BaseDepositERC20BridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositERC20BridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositERC20BridgeInitiated represents a ERC20BridgeInitiated event raised by the BaseDeposit contract.
type BaseDepositERC20BridgeInitiated struct {
	LocalToken  common.Address
	RemoteToken common.Address
	From        common.Address
	To          common.Address
	Amount      *big.Int
	ExtraData   []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterERC20BridgeInitiated is a free log retrieval operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) FilterERC20BridgeInitiated(opts *bind.FilterOpts, localToken []common.Address, remoteToken []common.Address, from []common.Address) (*BaseDepositERC20BridgeInitiatedIterator, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BaseDepositERC20BridgeInitiatedIterator{contract: _BaseDeposit.contract, event: "ERC20BridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20BridgeInitiated is a free log subscription operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) WatchERC20BridgeInitiated(opts *bind.WatchOpts, sink chan<- *BaseDepositERC20BridgeInitiated, localToken []common.Address, remoteToken []common.Address, from []common.Address) (event.Subscription, error) {

	var localTokenRule []interface{}
	for _, localTokenItem := range localToken {
		localTokenRule = append(localTokenRule, localTokenItem)
	}
	var remoteTokenRule []interface{}
	for _, remoteTokenItem := range remoteToken {
		remoteTokenRule = append(remoteTokenRule, remoteTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "ERC20BridgeInitiated", localTokenRule, remoteTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositERC20BridgeInitiated)
				if err := _BaseDeposit.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
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

// ParseERC20BridgeInitiated is a log parse operation binding the contract event 0x7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf.
//
// Solidity: event ERC20BridgeInitiated(address indexed localToken, address indexed remoteToken, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) ParseERC20BridgeInitiated(log types.Log) (*BaseDepositERC20BridgeInitiated, error) {
	event := new(BaseDepositERC20BridgeInitiated)
	if err := _BaseDeposit.contract.UnpackLog(event, "ERC20BridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseDepositERC20DepositInitiatedIterator is returned from FilterERC20DepositInitiated and is used to iterate over the raw logs and unpacked data for ERC20DepositInitiated events raised by the BaseDeposit contract.
type BaseDepositERC20DepositInitiatedIterator struct {
	Event *BaseDepositERC20DepositInitiated // Event containing the contract specifics and raw log

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
func (it *BaseDepositERC20DepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositERC20DepositInitiated)
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
		it.Event = new(BaseDepositERC20DepositInitiated)
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
func (it *BaseDepositERC20DepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositERC20DepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositERC20DepositInitiated represents a ERC20DepositInitiated event raised by the BaseDeposit contract.
type BaseDepositERC20DepositInitiated struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20DepositInitiated is a free log retrieval operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) FilterERC20DepositInitiated(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*BaseDepositERC20DepositInitiatedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "ERC20DepositInitiated", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BaseDepositERC20DepositInitiatedIterator{contract: _BaseDeposit.contract, event: "ERC20DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchERC20DepositInitiated is a free log subscription operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) WatchERC20DepositInitiated(opts *bind.WatchOpts, sink chan<- *BaseDepositERC20DepositInitiated, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "ERC20DepositInitiated", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositERC20DepositInitiated)
				if err := _BaseDeposit.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
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

// ParseERC20DepositInitiated is a log parse operation binding the contract event 0x718594027abd4eaed59f95162563e0cc6d0e8d5b86b1c7be8b1b0ac3343d0396.
//
// Solidity: event ERC20DepositInitiated(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) ParseERC20DepositInitiated(log types.Log) (*BaseDepositERC20DepositInitiated, error) {
	event := new(BaseDepositERC20DepositInitiated)
	if err := _BaseDeposit.contract.UnpackLog(event, "ERC20DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseDepositERC20WithdrawalFinalizedIterator is returned from FilterERC20WithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ERC20WithdrawalFinalized events raised by the BaseDeposit contract.
type BaseDepositERC20WithdrawalFinalizedIterator struct {
	Event *BaseDepositERC20WithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *BaseDepositERC20WithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositERC20WithdrawalFinalized)
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
		it.Event = new(BaseDepositERC20WithdrawalFinalized)
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
func (it *BaseDepositERC20WithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositERC20WithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositERC20WithdrawalFinalized represents a ERC20WithdrawalFinalized event raised by the BaseDeposit contract.
type BaseDepositERC20WithdrawalFinalized struct {
	L1Token   common.Address
	L2Token   common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20WithdrawalFinalized is a free log retrieval operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) FilterERC20WithdrawalFinalized(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*BaseDepositERC20WithdrawalFinalizedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "ERC20WithdrawalFinalized", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &BaseDepositERC20WithdrawalFinalizedIterator{contract: _BaseDeposit.contract, event: "ERC20WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchERC20WithdrawalFinalized is a free log subscription operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) WatchERC20WithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *BaseDepositERC20WithdrawalFinalized, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "ERC20WithdrawalFinalized", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositERC20WithdrawalFinalized)
				if err := _BaseDeposit.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
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

// ParseERC20WithdrawalFinalized is a log parse operation binding the contract event 0x3ceee06c1e37648fcbb6ed52e17b3e1f275a1f8c7b22a84b2b84732431e046b3.
//
// Solidity: event ERC20WithdrawalFinalized(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) ParseERC20WithdrawalFinalized(log types.Log) (*BaseDepositERC20WithdrawalFinalized, error) {
	event := new(BaseDepositERC20WithdrawalFinalized)
	if err := _BaseDeposit.contract.UnpackLog(event, "ERC20WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseDepositETHBridgeFinalizedIterator is returned from FilterETHBridgeFinalized and is used to iterate over the raw logs and unpacked data for ETHBridgeFinalized events raised by the BaseDeposit contract.
type BaseDepositETHBridgeFinalizedIterator struct {
	Event *BaseDepositETHBridgeFinalized // Event containing the contract specifics and raw log

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
func (it *BaseDepositETHBridgeFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositETHBridgeFinalized)
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
		it.Event = new(BaseDepositETHBridgeFinalized)
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
func (it *BaseDepositETHBridgeFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositETHBridgeFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositETHBridgeFinalized represents a ETHBridgeFinalized event raised by the BaseDeposit contract.
type BaseDepositETHBridgeFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHBridgeFinalized is a free log retrieval operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) FilterETHBridgeFinalized(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaseDepositETHBridgeFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "ETHBridgeFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseDepositETHBridgeFinalizedIterator{contract: _BaseDeposit.contract, event: "ETHBridgeFinalized", logs: logs, sub: sub}, nil
}

// WatchETHBridgeFinalized is a free log subscription operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) WatchETHBridgeFinalized(opts *bind.WatchOpts, sink chan<- *BaseDepositETHBridgeFinalized, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "ETHBridgeFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositETHBridgeFinalized)
				if err := _BaseDeposit.contract.UnpackLog(event, "ETHBridgeFinalized", log); err != nil {
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

// ParseETHBridgeFinalized is a log parse operation binding the contract event 0x31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d.
//
// Solidity: event ETHBridgeFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) ParseETHBridgeFinalized(log types.Log) (*BaseDepositETHBridgeFinalized, error) {
	event := new(BaseDepositETHBridgeFinalized)
	if err := _BaseDeposit.contract.UnpackLog(event, "ETHBridgeFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseDepositETHBridgeInitiatedIterator is returned from FilterETHBridgeInitiated and is used to iterate over the raw logs and unpacked data for ETHBridgeInitiated events raised by the BaseDeposit contract.
type BaseDepositETHBridgeInitiatedIterator struct {
	Event *BaseDepositETHBridgeInitiated // Event containing the contract specifics and raw log

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
func (it *BaseDepositETHBridgeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositETHBridgeInitiated)
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
		it.Event = new(BaseDepositETHBridgeInitiated)
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
func (it *BaseDepositETHBridgeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositETHBridgeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositETHBridgeInitiated represents a ETHBridgeInitiated event raised by the BaseDeposit contract.
type BaseDepositETHBridgeInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHBridgeInitiated is a free log retrieval operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) FilterETHBridgeInitiated(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaseDepositETHBridgeInitiatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "ETHBridgeInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseDepositETHBridgeInitiatedIterator{contract: _BaseDeposit.contract, event: "ETHBridgeInitiated", logs: logs, sub: sub}, nil
}

// WatchETHBridgeInitiated is a free log subscription operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) WatchETHBridgeInitiated(opts *bind.WatchOpts, sink chan<- *BaseDepositETHBridgeInitiated, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "ETHBridgeInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositETHBridgeInitiated)
				if err := _BaseDeposit.contract.UnpackLog(event, "ETHBridgeInitiated", log); err != nil {
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

// ParseETHBridgeInitiated is a log parse operation binding the contract event 0x2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5.
//
// Solidity: event ETHBridgeInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) ParseETHBridgeInitiated(log types.Log) (*BaseDepositETHBridgeInitiated, error) {
	event := new(BaseDepositETHBridgeInitiated)
	if err := _BaseDeposit.contract.UnpackLog(event, "ETHBridgeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseDepositETHDepositInitiatedIterator is returned from FilterETHDepositInitiated and is used to iterate over the raw logs and unpacked data for ETHDepositInitiated events raised by the BaseDeposit contract.
type BaseDepositETHDepositInitiatedIterator struct {
	Event *BaseDepositETHDepositInitiated // Event containing the contract specifics and raw log

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
func (it *BaseDepositETHDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositETHDepositInitiated)
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
		it.Event = new(BaseDepositETHDepositInitiated)
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
func (it *BaseDepositETHDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositETHDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositETHDepositInitiated represents a ETHDepositInitiated event raised by the BaseDeposit contract.
type BaseDepositETHDepositInitiated struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHDepositInitiated is a free log retrieval operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) FilterETHDepositInitiated(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaseDepositETHDepositInitiatedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "ETHDepositInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseDepositETHDepositInitiatedIterator{contract: _BaseDeposit.contract, event: "ETHDepositInitiated", logs: logs, sub: sub}, nil
}

// WatchETHDepositInitiated is a free log subscription operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) WatchETHDepositInitiated(opts *bind.WatchOpts, sink chan<- *BaseDepositETHDepositInitiated, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "ETHDepositInitiated", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositETHDepositInitiated)
				if err := _BaseDeposit.contract.UnpackLog(event, "ETHDepositInitiated", log); err != nil {
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

// ParseETHDepositInitiated is a log parse operation binding the contract event 0x35d79ab81f2b2017e19afb5c5571778877782d7a8786f5907f93b0f4702f4f23.
//
// Solidity: event ETHDepositInitiated(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) ParseETHDepositInitiated(log types.Log) (*BaseDepositETHDepositInitiated, error) {
	event := new(BaseDepositETHDepositInitiated)
	if err := _BaseDeposit.contract.UnpackLog(event, "ETHDepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseDepositETHWithdrawalFinalizedIterator is returned from FilterETHWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for ETHWithdrawalFinalized events raised by the BaseDeposit contract.
type BaseDepositETHWithdrawalFinalizedIterator struct {
	Event *BaseDepositETHWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *BaseDepositETHWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositETHWithdrawalFinalized)
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
		it.Event = new(BaseDepositETHWithdrawalFinalized)
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
func (it *BaseDepositETHWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositETHWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositETHWithdrawalFinalized represents a ETHWithdrawalFinalized event raised by the BaseDeposit contract.
type BaseDepositETHWithdrawalFinalized struct {
	From      common.Address
	To        common.Address
	Amount    *big.Int
	ExtraData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterETHWithdrawalFinalized is a free log retrieval operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) FilterETHWithdrawalFinalized(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BaseDepositETHWithdrawalFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "ETHWithdrawalFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BaseDepositETHWithdrawalFinalizedIterator{contract: _BaseDeposit.contract, event: "ETHWithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchETHWithdrawalFinalized is a free log subscription operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) WatchETHWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *BaseDepositETHWithdrawalFinalized, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "ETHWithdrawalFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositETHWithdrawalFinalized)
				if err := _BaseDeposit.contract.UnpackLog(event, "ETHWithdrawalFinalized", log); err != nil {
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

// ParseETHWithdrawalFinalized is a log parse operation binding the contract event 0x2ac69ee804d9a7a0984249f508dfab7cb2534b465b6ce1580f99a38ba9c5e631.
//
// Solidity: event ETHWithdrawalFinalized(address indexed from, address indexed to, uint256 amount, bytes extraData)
func (_BaseDeposit *BaseDepositFilterer) ParseETHWithdrawalFinalized(log types.Log) (*BaseDepositETHWithdrawalFinalized, error) {
	event := new(BaseDepositETHWithdrawalFinalized)
	if err := _BaseDeposit.contract.UnpackLog(event, "ETHWithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BaseDepositInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BaseDeposit contract.
type BaseDepositInitializedIterator struct {
	Event *BaseDepositInitialized // Event containing the contract specifics and raw log

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
func (it *BaseDepositInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BaseDepositInitialized)
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
		it.Event = new(BaseDepositInitialized)
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
func (it *BaseDepositInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BaseDepositInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BaseDepositInitialized represents a Initialized event raised by the BaseDeposit contract.
type BaseDepositInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BaseDeposit *BaseDepositFilterer) FilterInitialized(opts *bind.FilterOpts) (*BaseDepositInitializedIterator, error) {

	logs, sub, err := _BaseDeposit.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BaseDepositInitializedIterator{contract: _BaseDeposit.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BaseDeposit *BaseDepositFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BaseDepositInitialized) (event.Subscription, error) {

	logs, sub, err := _BaseDeposit.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BaseDepositInitialized)
				if err := _BaseDeposit.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BaseDeposit *BaseDepositFilterer) ParseInitialized(log types.Log) (*BaseDepositInitialized, error) {
	event := new(BaseDepositInitialized)
	if err := _BaseDeposit.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdraw

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

// WithdrawMetaData contains all meta data concerning the Withdraw contract.
var WithdrawMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDefaultGateway\",\"type\":\"address\"}],\"name\":\"DefaultGatewayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userFrom\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"TransferRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"TxToL1\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1ERC20\",\"type\":\"address\"}],\"name\":\"calculateL2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpartGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeInboundTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getOutboundCalldata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_counterpartGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultGateway\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"l1TokenToGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newL2DefaultGateway\",\"type\":\"address\"}],\"name\":\"setDefaultGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_l1Token\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_gateway\",\"type\":\"address[]\"}],\"name\":\"setGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WithdrawABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawMetaData.ABI instead.
var WithdrawABI = WithdrawMetaData.ABI

// Withdraw is an auto generated Go binding around an Ethereum contract.
type Withdraw struct {
	WithdrawCaller     // Read-only binding to the contract
	WithdrawTransactor // Write-only binding to the contract
	WithdrawFilterer   // Log filterer for contract events
}

// WithdrawCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawSession struct {
	Contract     *Withdraw         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawCallerSession struct {
	Contract *WithdrawCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WithdrawTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawTransactorSession struct {
	Contract     *WithdrawTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WithdrawRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawRaw struct {
	Contract *Withdraw // Generic contract binding to access the raw methods on
}

// WithdrawCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawCallerRaw struct {
	Contract *WithdrawCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawTransactorRaw struct {
	Contract *WithdrawTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdraw creates a new instance of Withdraw, bound to a specific deployed contract.
func NewWithdraw(address common.Address, backend bind.ContractBackend) (*Withdraw, error) {
	contract, err := bindWithdraw(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdraw{WithdrawCaller: WithdrawCaller{contract: contract}, WithdrawTransactor: WithdrawTransactor{contract: contract}, WithdrawFilterer: WithdrawFilterer{contract: contract}}, nil
}

// NewWithdrawCaller creates a new read-only instance of Withdraw, bound to a specific deployed contract.
func NewWithdrawCaller(address common.Address, caller bind.ContractCaller) (*WithdrawCaller, error) {
	contract, err := bindWithdraw(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawCaller{contract: contract}, nil
}

// NewWithdrawTransactor creates a new write-only instance of Withdraw, bound to a specific deployed contract.
func NewWithdrawTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawTransactor, error) {
	contract, err := bindWithdraw(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawTransactor{contract: contract}, nil
}

// NewWithdrawFilterer creates a new log filterer instance of Withdraw, bound to a specific deployed contract.
func NewWithdrawFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawFilterer, error) {
	contract, err := bindWithdraw(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawFilterer{contract: contract}, nil
}

// bindWithdraw binds a generic wrapper to an already deployed contract.
func bindWithdraw(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WithdrawMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdraw *WithdrawRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdraw.Contract.WithdrawCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdraw *WithdrawRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdraw.Contract.WithdrawTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdraw *WithdrawRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdraw.Contract.WithdrawTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdraw *WithdrawCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdraw.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdraw *WithdrawTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdraw.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdraw *WithdrawTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdraw.Contract.contract.Transact(opts, method, params...)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_Withdraw *WithdrawCaller) CalculateL2TokenAddress(opts *bind.CallOpts, l1ERC20 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Withdraw.contract.Call(opts, &out, "calculateL2TokenAddress", l1ERC20)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_Withdraw *WithdrawSession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _Withdraw.Contract.CalculateL2TokenAddress(&_Withdraw.CallOpts, l1ERC20)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_Withdraw *WithdrawCallerSession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _Withdraw.Contract.CalculateL2TokenAddress(&_Withdraw.CallOpts, l1ERC20)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_Withdraw *WithdrawCaller) CounterpartGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdraw.contract.Call(opts, &out, "counterpartGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_Withdraw *WithdrawSession) CounterpartGateway() (common.Address, error) {
	return _Withdraw.Contract.CounterpartGateway(&_Withdraw.CallOpts)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_Withdraw *WithdrawCallerSession) CounterpartGateway() (common.Address, error) {
	return _Withdraw.Contract.CounterpartGateway(&_Withdraw.CallOpts)
}

// DefaultGateway is a free data retrieval call binding the contract method 0x03295802.
//
// Solidity: function defaultGateway() view returns(address)
func (_Withdraw *WithdrawCaller) DefaultGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdraw.contract.Call(opts, &out, "defaultGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultGateway is a free data retrieval call binding the contract method 0x03295802.
//
// Solidity: function defaultGateway() view returns(address)
func (_Withdraw *WithdrawSession) DefaultGateway() (common.Address, error) {
	return _Withdraw.Contract.DefaultGateway(&_Withdraw.CallOpts)
}

// DefaultGateway is a free data retrieval call binding the contract method 0x03295802.
//
// Solidity: function defaultGateway() view returns(address)
func (_Withdraw *WithdrawCallerSession) DefaultGateway() (common.Address, error) {
	return _Withdraw.Contract.DefaultGateway(&_Withdraw.CallOpts)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _token) view returns(address gateway)
func (_Withdraw *WithdrawCaller) GetGateway(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Withdraw.contract.Call(opts, &out, "getGateway", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _token) view returns(address gateway)
func (_Withdraw *WithdrawSession) GetGateway(_token common.Address) (common.Address, error) {
	return _Withdraw.Contract.GetGateway(&_Withdraw.CallOpts, _token)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _token) view returns(address gateway)
func (_Withdraw *WithdrawCallerSession) GetGateway(_token common.Address) (common.Address, error) {
	return _Withdraw.Contract.GetGateway(&_Withdraw.CallOpts, _token)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes)
func (_Withdraw *WithdrawCaller) GetOutboundCalldata(opts *bind.CallOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	var out []interface{}
	err := _Withdraw.contract.Call(opts, &out, "getOutboundCalldata", _token, _from, _to, _amount, _data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes)
func (_Withdraw *WithdrawSession) GetOutboundCalldata(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _Withdraw.Contract.GetOutboundCalldata(&_Withdraw.CallOpts, _token, _from, _to, _amount, _data)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes)
func (_Withdraw *WithdrawCallerSession) GetOutboundCalldata(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _Withdraw.Contract.GetOutboundCalldata(&_Withdraw.CallOpts, _token, _from, _to, _amount, _data)
}

// L1TokenToGateway is a free data retrieval call binding the contract method 0xed08fdc6.
//
// Solidity: function l1TokenToGateway(address ) view returns(address)
func (_Withdraw *WithdrawCaller) L1TokenToGateway(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Withdraw.contract.Call(opts, &out, "l1TokenToGateway", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1TokenToGateway is a free data retrieval call binding the contract method 0xed08fdc6.
//
// Solidity: function l1TokenToGateway(address ) view returns(address)
func (_Withdraw *WithdrawSession) L1TokenToGateway(arg0 common.Address) (common.Address, error) {
	return _Withdraw.Contract.L1TokenToGateway(&_Withdraw.CallOpts, arg0)
}

// L1TokenToGateway is a free data retrieval call binding the contract method 0xed08fdc6.
//
// Solidity: function l1TokenToGateway(address ) view returns(address)
func (_Withdraw *WithdrawCallerSession) L1TokenToGateway(arg0 common.Address) (common.Address, error) {
	return _Withdraw.Contract.L1TokenToGateway(&_Withdraw.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Withdraw *WithdrawCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Withdraw.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Withdraw *WithdrawSession) Router() (common.Address, error) {
	return _Withdraw.Contract.Router(&_Withdraw.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Withdraw *WithdrawCallerSession) Router() (common.Address, error) {
	return _Withdraw.Contract.Router(&_Withdraw.CallOpts)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address , address , address , uint256 , bytes ) payable returns()
func (_Withdraw *WithdrawTransactor) FinalizeInboundTransfer(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Withdraw.contract.Transact(opts, "finalizeInboundTransfer", arg0, arg1, arg2, arg3, arg4)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address , address , address , uint256 , bytes ) payable returns()
func (_Withdraw *WithdrawSession) FinalizeInboundTransfer(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Withdraw.Contract.FinalizeInboundTransfer(&_Withdraw.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address , address , address , uint256 , bytes ) payable returns()
func (_Withdraw *WithdrawTransactorSession) FinalizeInboundTransfer(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Withdraw.Contract.FinalizeInboundTransfer(&_Withdraw.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpartGateway, address _defaultGateway) returns()
func (_Withdraw *WithdrawTransactor) Initialize(opts *bind.TransactOpts, _counterpartGateway common.Address, _defaultGateway common.Address) (*types.Transaction, error) {
	return _Withdraw.contract.Transact(opts, "initialize", _counterpartGateway, _defaultGateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpartGateway, address _defaultGateway) returns()
func (_Withdraw *WithdrawSession) Initialize(_counterpartGateway common.Address, _defaultGateway common.Address) (*types.Transaction, error) {
	return _Withdraw.Contract.Initialize(&_Withdraw.TransactOpts, _counterpartGateway, _defaultGateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _counterpartGateway, address _defaultGateway) returns()
func (_Withdraw *WithdrawTransactorSession) Initialize(_counterpartGateway common.Address, _defaultGateway common.Address) (*types.Transaction, error) {
	return _Withdraw.Contract.Initialize(&_Withdraw.TransactOpts, _counterpartGateway, _defaultGateway)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0x7b3a3c8b.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, bytes _data) payable returns(bytes)
func (_Withdraw *WithdrawTransactor) OutboundTransfer(opts *bind.TransactOpts, _l1Token common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Withdraw.contract.Transact(opts, "outboundTransfer", _l1Token, _to, _amount, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0x7b3a3c8b.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, bytes _data) payable returns(bytes)
func (_Withdraw *WithdrawSession) OutboundTransfer(_l1Token common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Withdraw.Contract.OutboundTransfer(&_Withdraw.TransactOpts, _l1Token, _to, _amount, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0x7b3a3c8b.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, bytes _data) payable returns(bytes)
func (_Withdraw *WithdrawTransactorSession) OutboundTransfer(_l1Token common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _Withdraw.Contract.OutboundTransfer(&_Withdraw.TransactOpts, _l1Token, _to, _amount, _data)
}

// OutboundTransfer0 is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_Withdraw *WithdrawTransactor) OutboundTransfer0(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _Withdraw.contract.Transact(opts, "outboundTransfer0", _token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransfer0 is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_Withdraw *WithdrawSession) OutboundTransfer0(_token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _Withdraw.Contract.OutboundTransfer0(&_Withdraw.TransactOpts, _token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransfer0 is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_Withdraw *WithdrawTransactorSession) OutboundTransfer0(_token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _Withdraw.Contract.OutboundTransfer0(&_Withdraw.TransactOpts, _token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_Withdraw *WithdrawTransactor) PostUpgradeInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdraw.contract.Transact(opts, "postUpgradeInit")
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_Withdraw *WithdrawSession) PostUpgradeInit() (*types.Transaction, error) {
	return _Withdraw.Contract.PostUpgradeInit(&_Withdraw.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_Withdraw *WithdrawTransactorSession) PostUpgradeInit() (*types.Transaction, error) {
	return _Withdraw.Contract.PostUpgradeInit(&_Withdraw.TransactOpts)
}

// SetDefaultGateway is a paid mutator transaction binding the contract method 0xf7c9362f.
//
// Solidity: function setDefaultGateway(address newL2DefaultGateway) returns()
func (_Withdraw *WithdrawTransactor) SetDefaultGateway(opts *bind.TransactOpts, newL2DefaultGateway common.Address) (*types.Transaction, error) {
	return _Withdraw.contract.Transact(opts, "setDefaultGateway", newL2DefaultGateway)
}

// SetDefaultGateway is a paid mutator transaction binding the contract method 0xf7c9362f.
//
// Solidity: function setDefaultGateway(address newL2DefaultGateway) returns()
func (_Withdraw *WithdrawSession) SetDefaultGateway(newL2DefaultGateway common.Address) (*types.Transaction, error) {
	return _Withdraw.Contract.SetDefaultGateway(&_Withdraw.TransactOpts, newL2DefaultGateway)
}

// SetDefaultGateway is a paid mutator transaction binding the contract method 0xf7c9362f.
//
// Solidity: function setDefaultGateway(address newL2DefaultGateway) returns()
func (_Withdraw *WithdrawTransactorSession) SetDefaultGateway(newL2DefaultGateway common.Address) (*types.Transaction, error) {
	return _Withdraw.Contract.SetDefaultGateway(&_Withdraw.TransactOpts, newL2DefaultGateway)
}

// SetGateway is a paid mutator transaction binding the contract method 0x4201f985.
//
// Solidity: function setGateway(address[] _l1Token, address[] _gateway) returns()
func (_Withdraw *WithdrawTransactor) SetGateway(opts *bind.TransactOpts, _l1Token []common.Address, _gateway []common.Address) (*types.Transaction, error) {
	return _Withdraw.contract.Transact(opts, "setGateway", _l1Token, _gateway)
}

// SetGateway is a paid mutator transaction binding the contract method 0x4201f985.
//
// Solidity: function setGateway(address[] _l1Token, address[] _gateway) returns()
func (_Withdraw *WithdrawSession) SetGateway(_l1Token []common.Address, _gateway []common.Address) (*types.Transaction, error) {
	return _Withdraw.Contract.SetGateway(&_Withdraw.TransactOpts, _l1Token, _gateway)
}

// SetGateway is a paid mutator transaction binding the contract method 0x4201f985.
//
// Solidity: function setGateway(address[] _l1Token, address[] _gateway) returns()
func (_Withdraw *WithdrawTransactorSession) SetGateway(_l1Token []common.Address, _gateway []common.Address) (*types.Transaction, error) {
	return _Withdraw.Contract.SetGateway(&_Withdraw.TransactOpts, _l1Token, _gateway)
}

// WithdrawDefaultGatewayUpdatedIterator is returned from FilterDefaultGatewayUpdated and is used to iterate over the raw logs and unpacked data for DefaultGatewayUpdated events raised by the Withdraw contract.
type WithdrawDefaultGatewayUpdatedIterator struct {
	Event *WithdrawDefaultGatewayUpdated // Event containing the contract specifics and raw log

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
func (it *WithdrawDefaultGatewayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawDefaultGatewayUpdated)
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
		it.Event = new(WithdrawDefaultGatewayUpdated)
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
func (it *WithdrawDefaultGatewayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawDefaultGatewayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawDefaultGatewayUpdated represents a DefaultGatewayUpdated event raised by the Withdraw contract.
type WithdrawDefaultGatewayUpdated struct {
	NewDefaultGateway common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDefaultGatewayUpdated is a free log retrieval operation binding the contract event 0x3a8f8eb961383a94d41d193e16a3af73eaddfd5764a4c640257323a1603ac331.
//
// Solidity: event DefaultGatewayUpdated(address newDefaultGateway)
func (_Withdraw *WithdrawFilterer) FilterDefaultGatewayUpdated(opts *bind.FilterOpts) (*WithdrawDefaultGatewayUpdatedIterator, error) {

	logs, sub, err := _Withdraw.contract.FilterLogs(opts, "DefaultGatewayUpdated")
	if err != nil {
		return nil, err
	}
	return &WithdrawDefaultGatewayUpdatedIterator{contract: _Withdraw.contract, event: "DefaultGatewayUpdated", logs: logs, sub: sub}, nil
}

// WatchDefaultGatewayUpdated is a free log subscription operation binding the contract event 0x3a8f8eb961383a94d41d193e16a3af73eaddfd5764a4c640257323a1603ac331.
//
// Solidity: event DefaultGatewayUpdated(address newDefaultGateway)
func (_Withdraw *WithdrawFilterer) WatchDefaultGatewayUpdated(opts *bind.WatchOpts, sink chan<- *WithdrawDefaultGatewayUpdated) (event.Subscription, error) {

	logs, sub, err := _Withdraw.contract.WatchLogs(opts, "DefaultGatewayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawDefaultGatewayUpdated)
				if err := _Withdraw.contract.UnpackLog(event, "DefaultGatewayUpdated", log); err != nil {
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

// ParseDefaultGatewayUpdated is a log parse operation binding the contract event 0x3a8f8eb961383a94d41d193e16a3af73eaddfd5764a4c640257323a1603ac331.
//
// Solidity: event DefaultGatewayUpdated(address newDefaultGateway)
func (_Withdraw *WithdrawFilterer) ParseDefaultGatewayUpdated(log types.Log) (*WithdrawDefaultGatewayUpdated, error) {
	event := new(WithdrawDefaultGatewayUpdated)
	if err := _Withdraw.contract.UnpackLog(event, "DefaultGatewayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawGatewaySetIterator is returned from FilterGatewaySet and is used to iterate over the raw logs and unpacked data for GatewaySet events raised by the Withdraw contract.
type WithdrawGatewaySetIterator struct {
	Event *WithdrawGatewaySet // Event containing the contract specifics and raw log

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
func (it *WithdrawGatewaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawGatewaySet)
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
		it.Event = new(WithdrawGatewaySet)
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
func (it *WithdrawGatewaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawGatewaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawGatewaySet represents a GatewaySet event raised by the Withdraw contract.
type WithdrawGatewaySet struct {
	L1Token common.Address
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGatewaySet is a free log retrieval operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address indexed l1Token, address indexed gateway)
func (_Withdraw *WithdrawFilterer) FilterGatewaySet(opts *bind.FilterOpts, l1Token []common.Address, gateway []common.Address) (*WithdrawGatewaySetIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _Withdraw.contract.FilterLogs(opts, "GatewaySet", l1TokenRule, gatewayRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawGatewaySetIterator{contract: _Withdraw.contract, event: "GatewaySet", logs: logs, sub: sub}, nil
}

// WatchGatewaySet is a free log subscription operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address indexed l1Token, address indexed gateway)
func (_Withdraw *WithdrawFilterer) WatchGatewaySet(opts *bind.WatchOpts, sink chan<- *WithdrawGatewaySet, l1Token []common.Address, gateway []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _Withdraw.contract.WatchLogs(opts, "GatewaySet", l1TokenRule, gatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawGatewaySet)
				if err := _Withdraw.contract.UnpackLog(event, "GatewaySet", log); err != nil {
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

// ParseGatewaySet is a log parse operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address indexed l1Token, address indexed gateway)
func (_Withdraw *WithdrawFilterer) ParseGatewaySet(log types.Log) (*WithdrawGatewaySet, error) {
	event := new(WithdrawGatewaySet)
	if err := _Withdraw.contract.UnpackLog(event, "GatewaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawTransferRoutedIterator is returned from FilterTransferRouted and is used to iterate over the raw logs and unpacked data for TransferRouted events raised by the Withdraw contract.
type WithdrawTransferRoutedIterator struct {
	Event *WithdrawTransferRouted // Event containing the contract specifics and raw log

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
func (it *WithdrawTransferRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawTransferRouted)
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
		it.Event = new(WithdrawTransferRouted)
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
func (it *WithdrawTransferRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawTransferRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawTransferRouted represents a TransferRouted event raised by the Withdraw contract.
type WithdrawTransferRouted struct {
	Token    common.Address
	UserFrom common.Address
	UserTo   common.Address
	Gateway  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferRouted is a free log retrieval operation binding the contract event 0x85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5.
//
// Solidity: event TransferRouted(address indexed token, address indexed _userFrom, address indexed _userTo, address gateway)
func (_Withdraw *WithdrawFilterer) FilterTransferRouted(opts *bind.FilterOpts, token []common.Address, _userFrom []common.Address, _userTo []common.Address) (*WithdrawTransferRoutedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var _userFromRule []interface{}
	for _, _userFromItem := range _userFrom {
		_userFromRule = append(_userFromRule, _userFromItem)
	}
	var _userToRule []interface{}
	for _, _userToItem := range _userTo {
		_userToRule = append(_userToRule, _userToItem)
	}

	logs, sub, err := _Withdraw.contract.FilterLogs(opts, "TransferRouted", tokenRule, _userFromRule, _userToRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawTransferRoutedIterator{contract: _Withdraw.contract, event: "TransferRouted", logs: logs, sub: sub}, nil
}

// WatchTransferRouted is a free log subscription operation binding the contract event 0x85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5.
//
// Solidity: event TransferRouted(address indexed token, address indexed _userFrom, address indexed _userTo, address gateway)
func (_Withdraw *WithdrawFilterer) WatchTransferRouted(opts *bind.WatchOpts, sink chan<- *WithdrawTransferRouted, token []common.Address, _userFrom []common.Address, _userTo []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var _userFromRule []interface{}
	for _, _userFromItem := range _userFrom {
		_userFromRule = append(_userFromRule, _userFromItem)
	}
	var _userToRule []interface{}
	for _, _userToItem := range _userTo {
		_userToRule = append(_userToRule, _userToItem)
	}

	logs, sub, err := _Withdraw.contract.WatchLogs(opts, "TransferRouted", tokenRule, _userFromRule, _userToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawTransferRouted)
				if err := _Withdraw.contract.UnpackLog(event, "TransferRouted", log); err != nil {
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

// ParseTransferRouted is a log parse operation binding the contract event 0x85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5.
//
// Solidity: event TransferRouted(address indexed token, address indexed _userFrom, address indexed _userTo, address gateway)
func (_Withdraw *WithdrawFilterer) ParseTransferRouted(log types.Log) (*WithdrawTransferRouted, error) {
	event := new(WithdrawTransferRouted)
	if err := _Withdraw.contract.UnpackLog(event, "TransferRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawTxToL1Iterator is returned from FilterTxToL1 and is used to iterate over the raw logs and unpacked data for TxToL1 events raised by the Withdraw contract.
type WithdrawTxToL1Iterator struct {
	Event *WithdrawTxToL1 // Event containing the contract specifics and raw log

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
func (it *WithdrawTxToL1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawTxToL1)
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
		it.Event = new(WithdrawTxToL1)
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
func (it *WithdrawTxToL1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawTxToL1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawTxToL1 represents a TxToL1 event raised by the Withdraw contract.
type WithdrawTxToL1 struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTxToL1 is a free log retrieval operation binding the contract event 0x2b986d32a0536b7e19baa48ab949fec7b903b7fad7730820b20632d100cc3a68.
//
// Solidity: event TxToL1(address indexed _from, address indexed _to, uint256 indexed _id, bytes _data)
func (_Withdraw *WithdrawFilterer) FilterTxToL1(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _id []*big.Int) (*WithdrawTxToL1Iterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Withdraw.contract.FilterLogs(opts, "TxToL1", _fromRule, _toRule, _idRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawTxToL1Iterator{contract: _Withdraw.contract, event: "TxToL1", logs: logs, sub: sub}, nil
}

// WatchTxToL1 is a free log subscription operation binding the contract event 0x2b986d32a0536b7e19baa48ab949fec7b903b7fad7730820b20632d100cc3a68.
//
// Solidity: event TxToL1(address indexed _from, address indexed _to, uint256 indexed _id, bytes _data)
func (_Withdraw *WithdrawFilterer) WatchTxToL1(opts *bind.WatchOpts, sink chan<- *WithdrawTxToL1, _from []common.Address, _to []common.Address, _id []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _Withdraw.contract.WatchLogs(opts, "TxToL1", _fromRule, _toRule, _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawTxToL1)
				if err := _Withdraw.contract.UnpackLog(event, "TxToL1", log); err != nil {
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

// ParseTxToL1 is a log parse operation binding the contract event 0x2b986d32a0536b7e19baa48ab949fec7b903b7fad7730820b20632d100cc3a68.
//
// Solidity: event TxToL1(address indexed _from, address indexed _to, uint256 indexed _id, bytes _data)
func (_Withdraw *WithdrawFilterer) ParseTxToL1(log types.Log) (*WithdrawTxToL1, error) {
	event := new(WithdrawTxToL1)
	if err := _Withdraw.contract.UnpackLog(event, "TxToL1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

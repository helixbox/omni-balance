// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package layerzeroMessager

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

// LayerzeroMessagerMetaData contains all meta data concerning the LayerzeroMessager contract.
var LayerzeroMessagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_endpoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"lzRemoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"successed\",\"type\":\"bool\"}],\"name\":\"CallResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"lzRemoteChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srcAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"remoteAppAddress\",\"type\":\"address\"}],\"name\":\"CallerUnMatched\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"appAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"authoriseAppCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"callerWhiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endpoint\",\"outputs\":[{\"internalType\":\"contractILayerZeroEndpoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nativeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zroFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"lzReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_remoteBridge\",\"type\":\"address\"}],\"name\":\"registerRemoteReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_remoteBridge\",\"type\":\"address\"}],\"name\":\"registerRemoteSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"remoteAppReceivers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"remoteAppSenders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remoteMessagers\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"lzRemoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"messager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_appRemoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_lzRemoteChainId\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_remoteMessager\",\"type\":\"address\"}],\"name\":\"setRemoteMessager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"trustedRemotes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LayerzeroMessagerABI is the input ABI used to generate the binding from.
// Deprecated: Use LayerzeroMessagerMetaData.ABI instead.
var LayerzeroMessagerABI = LayerzeroMessagerMetaData.ABI

// LayerzeroMessager is an auto generated Go binding around an Ethereum contract.
type LayerzeroMessager struct {
	LayerzeroMessagerCaller     // Read-only binding to the contract
	LayerzeroMessagerTransactor // Write-only binding to the contract
	LayerzeroMessagerFilterer   // Log filterer for contract events
}

// LayerzeroMessagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LayerzeroMessagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerzeroMessagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LayerzeroMessagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerzeroMessagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LayerzeroMessagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LayerzeroMessagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LayerzeroMessagerSession struct {
	Contract     *LayerzeroMessager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LayerzeroMessagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LayerzeroMessagerCallerSession struct {
	Contract *LayerzeroMessagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LayerzeroMessagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LayerzeroMessagerTransactorSession struct {
	Contract     *LayerzeroMessagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LayerzeroMessagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LayerzeroMessagerRaw struct {
	Contract *LayerzeroMessager // Generic contract binding to access the raw methods on
}

// LayerzeroMessagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LayerzeroMessagerCallerRaw struct {
	Contract *LayerzeroMessagerCaller // Generic read-only contract binding to access the raw methods on
}

// LayerzeroMessagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LayerzeroMessagerTransactorRaw struct {
	Contract *LayerzeroMessagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLayerzeroMessager creates a new instance of LayerzeroMessager, bound to a specific deployed contract.
func NewLayerzeroMessager(address common.Address, backend bind.ContractBackend) (*LayerzeroMessager, error) {
	contract, err := bindLayerzeroMessager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LayerzeroMessager{LayerzeroMessagerCaller: LayerzeroMessagerCaller{contract: contract}, LayerzeroMessagerTransactor: LayerzeroMessagerTransactor{contract: contract}, LayerzeroMessagerFilterer: LayerzeroMessagerFilterer{contract: contract}}, nil
}

// NewLayerzeroMessagerCaller creates a new read-only instance of LayerzeroMessager, bound to a specific deployed contract.
func NewLayerzeroMessagerCaller(address common.Address, caller bind.ContractCaller) (*LayerzeroMessagerCaller, error) {
	contract, err := bindLayerzeroMessager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LayerzeroMessagerCaller{contract: contract}, nil
}

// NewLayerzeroMessagerTransactor creates a new write-only instance of LayerzeroMessager, bound to a specific deployed contract.
func NewLayerzeroMessagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LayerzeroMessagerTransactor, error) {
	contract, err := bindLayerzeroMessager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LayerzeroMessagerTransactor{contract: contract}, nil
}

// NewLayerzeroMessagerFilterer creates a new log filterer instance of LayerzeroMessager, bound to a specific deployed contract.
func NewLayerzeroMessagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LayerzeroMessagerFilterer, error) {
	contract, err := bindLayerzeroMessager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LayerzeroMessagerFilterer{contract: contract}, nil
}

// bindLayerzeroMessager binds a generic wrapper to an already deployed contract.
func bindLayerzeroMessager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LayerzeroMessagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LayerzeroMessager *LayerzeroMessagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LayerzeroMessager.Contract.LayerzeroMessagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LayerzeroMessager *LayerzeroMessagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.LayerzeroMessagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LayerzeroMessager *LayerzeroMessagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.LayerzeroMessagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LayerzeroMessager *LayerzeroMessagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LayerzeroMessager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LayerzeroMessager *LayerzeroMessagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LayerzeroMessager *LayerzeroMessagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.contract.Transact(opts, method, params...)
}

// CallerWhiteList is a free data retrieval call binding the contract method 0x393e806f.
//
// Solidity: function callerWhiteList(address ) view returns(bool)
func (_LayerzeroMessager *LayerzeroMessagerCaller) CallerWhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "callerWhiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CallerWhiteList is a free data retrieval call binding the contract method 0x393e806f.
//
// Solidity: function callerWhiteList(address ) view returns(bool)
func (_LayerzeroMessager *LayerzeroMessagerSession) CallerWhiteList(arg0 common.Address) (bool, error) {
	return _LayerzeroMessager.Contract.CallerWhiteList(&_LayerzeroMessager.CallOpts, arg0)
}

// CallerWhiteList is a free data retrieval call binding the contract method 0x393e806f.
//
// Solidity: function callerWhiteList(address ) view returns(bool)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) CallerWhiteList(arg0 common.Address) (bool, error) {
	return _LayerzeroMessager.Contract.CallerWhiteList(&_LayerzeroMessager.CallOpts, arg0)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerSession) Dao() (common.Address, error) {
	return _LayerzeroMessager.Contract.Dao(&_LayerzeroMessager.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) Dao() (common.Address, error) {
	return _LayerzeroMessager.Contract.Dao(&_LayerzeroMessager.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCaller) Endpoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "endpoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerSession) Endpoint() (common.Address, error) {
	return _LayerzeroMessager.Contract.Endpoint(&_LayerzeroMessager.CallOpts)
}

// Endpoint is a free data retrieval call binding the contract method 0x5e280f11.
//
// Solidity: function endpoint() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) Endpoint() (common.Address, error) {
	return _LayerzeroMessager.Contract.Endpoint(&_LayerzeroMessager.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xb63b150b.
//
// Solidity: function fee(uint256 _remoteChainId, bytes _message) view returns(uint256 nativeFee, uint256 zroFee)
func (_LayerzeroMessager *LayerzeroMessagerCaller) Fee(opts *bind.CallOpts, _remoteChainId *big.Int, _message []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "fee", _remoteChainId, _message)

	outstruct := new(struct {
		NativeFee *big.Int
		ZroFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NativeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ZroFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Fee is a free data retrieval call binding the contract method 0xb63b150b.
//
// Solidity: function fee(uint256 _remoteChainId, bytes _message) view returns(uint256 nativeFee, uint256 zroFee)
func (_LayerzeroMessager *LayerzeroMessagerSession) Fee(_remoteChainId *big.Int, _message []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _LayerzeroMessager.Contract.Fee(&_LayerzeroMessager.CallOpts, _remoteChainId, _message)
}

// Fee is a free data retrieval call binding the contract method 0xb63b150b.
//
// Solidity: function fee(uint256 _remoteChainId, bytes _message) view returns(uint256 nativeFee, uint256 zroFee)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) Fee(_remoteChainId *big.Int, _message []byte) (struct {
	NativeFee *big.Int
	ZroFee    *big.Int
}, error) {
	return _LayerzeroMessager.Contract.Fee(&_LayerzeroMessager.CallOpts, _remoteChainId, _message)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerSession) Operator() (common.Address, error) {
	return _LayerzeroMessager.Contract.Operator(&_LayerzeroMessager.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) Operator() (common.Address, error) {
	return _LayerzeroMessager.Contract.Operator(&_LayerzeroMessager.CallOpts)
}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCaller) PendingDao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "pendingDao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerSession) PendingDao() (common.Address, error) {
	return _LayerzeroMessager.Contract.PendingDao(&_LayerzeroMessager.CallOpts)
}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) PendingDao() (common.Address, error) {
	return _LayerzeroMessager.Contract.PendingDao(&_LayerzeroMessager.CallOpts)
}

// RemoteAppReceivers is a free data retrieval call binding the contract method 0x5a4220ad.
//
// Solidity: function remoteAppReceivers(bytes32 ) view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCaller) RemoteAppReceivers(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "remoteAppReceivers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteAppReceivers is a free data retrieval call binding the contract method 0x5a4220ad.
//
// Solidity: function remoteAppReceivers(bytes32 ) view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerSession) RemoteAppReceivers(arg0 [32]byte) (common.Address, error) {
	return _LayerzeroMessager.Contract.RemoteAppReceivers(&_LayerzeroMessager.CallOpts, arg0)
}

// RemoteAppReceivers is a free data retrieval call binding the contract method 0x5a4220ad.
//
// Solidity: function remoteAppReceivers(bytes32 ) view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) RemoteAppReceivers(arg0 [32]byte) (common.Address, error) {
	return _LayerzeroMessager.Contract.RemoteAppReceivers(&_LayerzeroMessager.CallOpts, arg0)
}

// RemoteAppSenders is a free data retrieval call binding the contract method 0x31c3ac91.
//
// Solidity: function remoteAppSenders(bytes32 ) view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCaller) RemoteAppSenders(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "remoteAppSenders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteAppSenders is a free data retrieval call binding the contract method 0x31c3ac91.
//
// Solidity: function remoteAppSenders(bytes32 ) view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerSession) RemoteAppSenders(arg0 [32]byte) (common.Address, error) {
	return _LayerzeroMessager.Contract.RemoteAppSenders(&_LayerzeroMessager.CallOpts, arg0)
}

// RemoteAppSenders is a free data retrieval call binding the contract method 0x31c3ac91.
//
// Solidity: function remoteAppSenders(bytes32 ) view returns(address)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) RemoteAppSenders(arg0 [32]byte) (common.Address, error) {
	return _LayerzeroMessager.Contract.RemoteAppSenders(&_LayerzeroMessager.CallOpts, arg0)
}

// RemoteMessagers is a free data retrieval call binding the contract method 0xced4e6e0.
//
// Solidity: function remoteMessagers(uint256 ) view returns(uint16 lzRemoteChainId, address messager)
func (_LayerzeroMessager *LayerzeroMessagerCaller) RemoteMessagers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LzRemoteChainId uint16
	Messager        common.Address
}, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "remoteMessagers", arg0)

	outstruct := new(struct {
		LzRemoteChainId uint16
		Messager        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LzRemoteChainId = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Messager = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RemoteMessagers is a free data retrieval call binding the contract method 0xced4e6e0.
//
// Solidity: function remoteMessagers(uint256 ) view returns(uint16 lzRemoteChainId, address messager)
func (_LayerzeroMessager *LayerzeroMessagerSession) RemoteMessagers(arg0 *big.Int) (struct {
	LzRemoteChainId uint16
	Messager        common.Address
}, error) {
	return _LayerzeroMessager.Contract.RemoteMessagers(&_LayerzeroMessager.CallOpts, arg0)
}

// RemoteMessagers is a free data retrieval call binding the contract method 0xced4e6e0.
//
// Solidity: function remoteMessagers(uint256 ) view returns(uint16 lzRemoteChainId, address messager)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) RemoteMessagers(arg0 *big.Int) (struct {
	LzRemoteChainId uint16
	Messager        common.Address
}, error) {
	return _LayerzeroMessager.Contract.RemoteMessagers(&_LayerzeroMessager.CallOpts, arg0)
}

// TrustedRemotes is a free data retrieval call binding the contract method 0x2b5d23d2.
//
// Solidity: function trustedRemotes(uint16 ) view returns(bytes32)
func (_LayerzeroMessager *LayerzeroMessagerCaller) TrustedRemotes(opts *bind.CallOpts, arg0 uint16) ([32]byte, error) {
	var out []interface{}
	err := _LayerzeroMessager.contract.Call(opts, &out, "trustedRemotes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TrustedRemotes is a free data retrieval call binding the contract method 0x2b5d23d2.
//
// Solidity: function trustedRemotes(uint16 ) view returns(bytes32)
func (_LayerzeroMessager *LayerzeroMessagerSession) TrustedRemotes(arg0 uint16) ([32]byte, error) {
	return _LayerzeroMessager.Contract.TrustedRemotes(&_LayerzeroMessager.CallOpts, arg0)
}

// TrustedRemotes is a free data retrieval call binding the contract method 0x2b5d23d2.
//
// Solidity: function trustedRemotes(uint16 ) view returns(bytes32)
func (_LayerzeroMessager *LayerzeroMessagerCallerSession) TrustedRemotes(arg0 uint16) ([32]byte, error) {
	return _LayerzeroMessager.Contract.TrustedRemotes(&_LayerzeroMessager.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) AcceptOwnership() (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.AcceptOwnership(&_LayerzeroMessager.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.AcceptOwnership(&_LayerzeroMessager.TransactOpts)
}

// AuthoriseAppCaller is a paid mutator transaction binding the contract method 0x2fc8b880.
//
// Solidity: function authoriseAppCaller(address appAddress, bool enable) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) AuthoriseAppCaller(opts *bind.TransactOpts, appAddress common.Address, enable bool) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "authoriseAppCaller", appAddress, enable)
}

// AuthoriseAppCaller is a paid mutator transaction binding the contract method 0x2fc8b880.
//
// Solidity: function authoriseAppCaller(address appAddress, bool enable) returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) AuthoriseAppCaller(appAddress common.Address, enable bool) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.AuthoriseAppCaller(&_LayerzeroMessager.TransactOpts, appAddress, enable)
}

// AuthoriseAppCaller is a paid mutator transaction binding the contract method 0x2fc8b880.
//
// Solidity: function authoriseAppCaller(address appAddress, bool enable) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) AuthoriseAppCaller(appAddress common.Address, enable bool) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.AuthoriseAppCaller(&_LayerzeroMessager.TransactOpts, appAddress, enable)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 , bytes _payload) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) LzReceive(opts *bind.TransactOpts, _srcChainId uint16, _srcAddress []byte, arg2 uint64, _payload []byte) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "lzReceive", _srcChainId, _srcAddress, arg2, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 , bytes _payload) returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) LzReceive(_srcChainId uint16, _srcAddress []byte, arg2 uint64, _payload []byte) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.LzReceive(&_LayerzeroMessager.TransactOpts, _srcChainId, _srcAddress, arg2, _payload)
}

// LzReceive is a paid mutator transaction binding the contract method 0x001d3567.
//
// Solidity: function lzReceive(uint16 _srcChainId, bytes _srcAddress, uint64 , bytes _payload) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) LzReceive(_srcChainId uint16, _srcAddress []byte, arg2 uint64, _payload []byte) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.LzReceive(&_LayerzeroMessager.TransactOpts, _srcChainId, _srcAddress, arg2, _payload)
}

// RegisterRemoteReceiver is a paid mutator transaction binding the contract method 0xc2fdb804.
//
// Solidity: function registerRemoteReceiver(uint256 _remoteChainId, address _remoteBridge) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) RegisterRemoteReceiver(opts *bind.TransactOpts, _remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "registerRemoteReceiver", _remoteChainId, _remoteBridge)
}

// RegisterRemoteReceiver is a paid mutator transaction binding the contract method 0xc2fdb804.
//
// Solidity: function registerRemoteReceiver(uint256 _remoteChainId, address _remoteBridge) returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) RegisterRemoteReceiver(_remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.RegisterRemoteReceiver(&_LayerzeroMessager.TransactOpts, _remoteChainId, _remoteBridge)
}

// RegisterRemoteReceiver is a paid mutator transaction binding the contract method 0xc2fdb804.
//
// Solidity: function registerRemoteReceiver(uint256 _remoteChainId, address _remoteBridge) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) RegisterRemoteReceiver(_remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.RegisterRemoteReceiver(&_LayerzeroMessager.TransactOpts, _remoteChainId, _remoteBridge)
}

// RegisterRemoteSender is a paid mutator transaction binding the contract method 0xc07d7713.
//
// Solidity: function registerRemoteSender(uint256 _remoteChainId, address _remoteBridge) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) RegisterRemoteSender(opts *bind.TransactOpts, _remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "registerRemoteSender", _remoteChainId, _remoteBridge)
}

// RegisterRemoteSender is a paid mutator transaction binding the contract method 0xc07d7713.
//
// Solidity: function registerRemoteSender(uint256 _remoteChainId, address _remoteBridge) returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) RegisterRemoteSender(_remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.RegisterRemoteSender(&_LayerzeroMessager.TransactOpts, _remoteChainId, _remoteBridge)
}

// RegisterRemoteSender is a paid mutator transaction binding the contract method 0xc07d7713.
//
// Solidity: function registerRemoteSender(uint256 _remoteChainId, address _remoteBridge) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) RegisterRemoteSender(_remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.RegisterRemoteSender(&_LayerzeroMessager.TransactOpts, _remoteChainId, _remoteBridge)
}

// SendMessage is a paid mutator transaction binding the contract method 0x6ea9cec9.
//
// Solidity: function sendMessage(uint256 _remoteChainId, bytes _message, bytes _params) payable returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) SendMessage(opts *bind.TransactOpts, _remoteChainId *big.Int, _message []byte, _params []byte) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "sendMessage", _remoteChainId, _message, _params)
}

// SendMessage is a paid mutator transaction binding the contract method 0x6ea9cec9.
//
// Solidity: function sendMessage(uint256 _remoteChainId, bytes _message, bytes _params) payable returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) SendMessage(_remoteChainId *big.Int, _message []byte, _params []byte) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.SendMessage(&_LayerzeroMessager.TransactOpts, _remoteChainId, _message, _params)
}

// SendMessage is a paid mutator transaction binding the contract method 0x6ea9cec9.
//
// Solidity: function sendMessage(uint256 _remoteChainId, bytes _message, bytes _params) payable returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) SendMessage(_remoteChainId *big.Int, _message []byte, _params []byte) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.SendMessage(&_LayerzeroMessager.TransactOpts, _remoteChainId, _message, _params)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.SetOperator(&_LayerzeroMessager.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.SetOperator(&_LayerzeroMessager.TransactOpts, _operator)
}

// SetRemoteMessager is a paid mutator transaction binding the contract method 0x937217c0.
//
// Solidity: function setRemoteMessager(uint256 _appRemoteChainId, uint16 _lzRemoteChainId, address _remoteMessager) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) SetRemoteMessager(opts *bind.TransactOpts, _appRemoteChainId *big.Int, _lzRemoteChainId uint16, _remoteMessager common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "setRemoteMessager", _appRemoteChainId, _lzRemoteChainId, _remoteMessager)
}

// SetRemoteMessager is a paid mutator transaction binding the contract method 0x937217c0.
//
// Solidity: function setRemoteMessager(uint256 _appRemoteChainId, uint16 _lzRemoteChainId, address _remoteMessager) returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) SetRemoteMessager(_appRemoteChainId *big.Int, _lzRemoteChainId uint16, _remoteMessager common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.SetRemoteMessager(&_LayerzeroMessager.TransactOpts, _appRemoteChainId, _lzRemoteChainId, _remoteMessager)
}

// SetRemoteMessager is a paid mutator transaction binding the contract method 0x937217c0.
//
// Solidity: function setRemoteMessager(uint256 _appRemoteChainId, uint16 _lzRemoteChainId, address _remoteMessager) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) SetRemoteMessager(_appRemoteChainId *big.Int, _lzRemoteChainId uint16, _remoteMessager common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.SetRemoteMessager(&_LayerzeroMessager.TransactOpts, _appRemoteChainId, _lzRemoteChainId, _remoteMessager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactor) TransferOwnership(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.contract.Transact(opts, "transferOwnership", _dao)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_LayerzeroMessager *LayerzeroMessagerSession) TransferOwnership(_dao common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.TransferOwnership(&_LayerzeroMessager.TransactOpts, _dao)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_LayerzeroMessager *LayerzeroMessagerTransactorSession) TransferOwnership(_dao common.Address) (*types.Transaction, error) {
	return _LayerzeroMessager.Contract.TransferOwnership(&_LayerzeroMessager.TransactOpts, _dao)
}

// LayerzeroMessagerCallResultIterator is returned from FilterCallResult and is used to iterate over the raw logs and unpacked data for CallResult events raised by the LayerzeroMessager contract.
type LayerzeroMessagerCallResultIterator struct {
	Event *LayerzeroMessagerCallResult // Event containing the contract specifics and raw log

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
func (it *LayerzeroMessagerCallResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroMessagerCallResult)
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
		it.Event = new(LayerzeroMessagerCallResult)
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
func (it *LayerzeroMessagerCallResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroMessagerCallResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroMessagerCallResult represents a CallResult event raised by the LayerzeroMessager contract.
type LayerzeroMessagerCallResult struct {
	LzRemoteChainId uint16
	SrcAddress      []byte
	Successed       bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCallResult is a free log retrieval operation binding the contract event 0x79bc1c4c6abc86eaef83dd00bb54b7e16310d2a4efb0f8007a5cf453b05e5e4d.
//
// Solidity: event CallResult(uint16 lzRemoteChainId, bytes srcAddress, bool successed)
func (_LayerzeroMessager *LayerzeroMessagerFilterer) FilterCallResult(opts *bind.FilterOpts) (*LayerzeroMessagerCallResultIterator, error) {

	logs, sub, err := _LayerzeroMessager.contract.FilterLogs(opts, "CallResult")
	if err != nil {
		return nil, err
	}
	return &LayerzeroMessagerCallResultIterator{contract: _LayerzeroMessager.contract, event: "CallResult", logs: logs, sub: sub}, nil
}

// WatchCallResult is a free log subscription operation binding the contract event 0x79bc1c4c6abc86eaef83dd00bb54b7e16310d2a4efb0f8007a5cf453b05e5e4d.
//
// Solidity: event CallResult(uint16 lzRemoteChainId, bytes srcAddress, bool successed)
func (_LayerzeroMessager *LayerzeroMessagerFilterer) WatchCallResult(opts *bind.WatchOpts, sink chan<- *LayerzeroMessagerCallResult) (event.Subscription, error) {

	logs, sub, err := _LayerzeroMessager.contract.WatchLogs(opts, "CallResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroMessagerCallResult)
				if err := _LayerzeroMessager.contract.UnpackLog(event, "CallResult", log); err != nil {
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

// ParseCallResult is a log parse operation binding the contract event 0x79bc1c4c6abc86eaef83dd00bb54b7e16310d2a4efb0f8007a5cf453b05e5e4d.
//
// Solidity: event CallResult(uint16 lzRemoteChainId, bytes srcAddress, bool successed)
func (_LayerzeroMessager *LayerzeroMessagerFilterer) ParseCallResult(log types.Log) (*LayerzeroMessagerCallResult, error) {
	event := new(LayerzeroMessagerCallResult)
	if err := _LayerzeroMessager.contract.UnpackLog(event, "CallResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LayerzeroMessagerCallerUnMatchedIterator is returned from FilterCallerUnMatched and is used to iterate over the raw logs and unpacked data for CallerUnMatched events raised by the LayerzeroMessager contract.
type LayerzeroMessagerCallerUnMatchedIterator struct {
	Event *LayerzeroMessagerCallerUnMatched // Event containing the contract specifics and raw log

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
func (it *LayerzeroMessagerCallerUnMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LayerzeroMessagerCallerUnMatched)
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
		it.Event = new(LayerzeroMessagerCallerUnMatched)
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
func (it *LayerzeroMessagerCallerUnMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LayerzeroMessagerCallerUnMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LayerzeroMessagerCallerUnMatched represents a CallerUnMatched event raised by the LayerzeroMessager contract.
type LayerzeroMessagerCallerUnMatched struct {
	LzRemoteChainId  uint16
	SrcAddress       []byte
	RemoteAppAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCallerUnMatched is a free log retrieval operation binding the contract event 0xbbba031a77046c6975d96a4b1da9941e1b00cd6662158415c0abd01ea765664f.
//
// Solidity: event CallerUnMatched(uint16 lzRemoteChainId, bytes srcAddress, address remoteAppAddress)
func (_LayerzeroMessager *LayerzeroMessagerFilterer) FilterCallerUnMatched(opts *bind.FilterOpts) (*LayerzeroMessagerCallerUnMatchedIterator, error) {

	logs, sub, err := _LayerzeroMessager.contract.FilterLogs(opts, "CallerUnMatched")
	if err != nil {
		return nil, err
	}
	return &LayerzeroMessagerCallerUnMatchedIterator{contract: _LayerzeroMessager.contract, event: "CallerUnMatched", logs: logs, sub: sub}, nil
}

// WatchCallerUnMatched is a free log subscription operation binding the contract event 0xbbba031a77046c6975d96a4b1da9941e1b00cd6662158415c0abd01ea765664f.
//
// Solidity: event CallerUnMatched(uint16 lzRemoteChainId, bytes srcAddress, address remoteAppAddress)
func (_LayerzeroMessager *LayerzeroMessagerFilterer) WatchCallerUnMatched(opts *bind.WatchOpts, sink chan<- *LayerzeroMessagerCallerUnMatched) (event.Subscription, error) {

	logs, sub, err := _LayerzeroMessager.contract.WatchLogs(opts, "CallerUnMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LayerzeroMessagerCallerUnMatched)
				if err := _LayerzeroMessager.contract.UnpackLog(event, "CallerUnMatched", log); err != nil {
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

// ParseCallerUnMatched is a log parse operation binding the contract event 0xbbba031a77046c6975d96a4b1da9941e1b00cd6662158415c0abd01ea765664f.
//
// Solidity: event CallerUnMatched(uint16 lzRemoteChainId, bytes srcAddress, address remoteAppAddress)
func (_LayerzeroMessager *LayerzeroMessagerFilterer) ParseCallerUnMatched(log types.Log) (*LayerzeroMessagerCallerUnMatched, error) {
	event := new(LayerzeroMessagerCallerUnMatched)
	if err := _LayerzeroMessager.contract.UnpackLog(event, "CallerUnMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


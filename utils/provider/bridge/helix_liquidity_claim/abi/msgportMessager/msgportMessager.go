// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package msgportMessager

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

// MsgportMessagerMetaData contains all meta data concerning the MsgportMessager contract.
var MsgportMessagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_msgport\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAppChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"CallResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAppChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcAppAddress\",\"type\":\"address\"}],\"name\":\"CallerUnMatched\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"messagePayload\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgport\",\"outputs\":[{\"internalType\":\"contractIMessagePort\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcAppChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_remoteAppAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_localAppAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_remoteBridge\",\"type\":\"address\"}],\"name\":\"registerRemoteReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_remoteBridge\",\"type\":\"address\"}],\"name\":\"registerRemoteSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"remoteAppReceivers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"remoteAppSenders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"remoteMessagers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"msgportRemoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"messager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgport\",\"type\":\"address\"}],\"name\":\"setMsgPort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_appRemoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_msgportRemoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_remoteMessager\",\"type\":\"address\"}],\"name\":\"setRemoteMessager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setWhiteList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whiteList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MsgportMessagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MsgportMessagerMetaData.ABI instead.
var MsgportMessagerABI = MsgportMessagerMetaData.ABI

// MsgportMessager is an auto generated Go binding around an Ethereum contract.
type MsgportMessager struct {
	MsgportMessagerCaller     // Read-only binding to the contract
	MsgportMessagerTransactor // Write-only binding to the contract
	MsgportMessagerFilterer   // Log filterer for contract events
}

// MsgportMessagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MsgportMessagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgportMessagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MsgportMessagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgportMessagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MsgportMessagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgportMessagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MsgportMessagerSession struct {
	Contract     *MsgportMessager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MsgportMessagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MsgportMessagerCallerSession struct {
	Contract *MsgportMessagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MsgportMessagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MsgportMessagerTransactorSession struct {
	Contract     *MsgportMessagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MsgportMessagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MsgportMessagerRaw struct {
	Contract *MsgportMessager // Generic contract binding to access the raw methods on
}

// MsgportMessagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MsgportMessagerCallerRaw struct {
	Contract *MsgportMessagerCaller // Generic read-only contract binding to access the raw methods on
}

// MsgportMessagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MsgportMessagerTransactorRaw struct {
	Contract *MsgportMessagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMsgportMessager creates a new instance of MsgportMessager, bound to a specific deployed contract.
func NewMsgportMessager(address common.Address, backend bind.ContractBackend) (*MsgportMessager, error) {
	contract, err := bindMsgportMessager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MsgportMessager{MsgportMessagerCaller: MsgportMessagerCaller{contract: contract}, MsgportMessagerTransactor: MsgportMessagerTransactor{contract: contract}, MsgportMessagerFilterer: MsgportMessagerFilterer{contract: contract}}, nil
}

// NewMsgportMessagerCaller creates a new read-only instance of MsgportMessager, bound to a specific deployed contract.
func NewMsgportMessagerCaller(address common.Address, caller bind.ContractCaller) (*MsgportMessagerCaller, error) {
	contract, err := bindMsgportMessager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MsgportMessagerCaller{contract: contract}, nil
}

// NewMsgportMessagerTransactor creates a new write-only instance of MsgportMessager, bound to a specific deployed contract.
func NewMsgportMessagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MsgportMessagerTransactor, error) {
	contract, err := bindMsgportMessager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MsgportMessagerTransactor{contract: contract}, nil
}

// NewMsgportMessagerFilterer creates a new log filterer instance of MsgportMessager, bound to a specific deployed contract.
func NewMsgportMessagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MsgportMessagerFilterer, error) {
	contract, err := bindMsgportMessager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MsgportMessagerFilterer{contract: contract}, nil
}

// bindMsgportMessager binds a generic wrapper to an already deployed contract.
func bindMsgportMessager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MsgportMessagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgportMessager *MsgportMessagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgportMessager.Contract.MsgportMessagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgportMessager *MsgportMessagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgportMessager.Contract.MsgportMessagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgportMessager *MsgportMessagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgportMessager.Contract.MsgportMessagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgportMessager *MsgportMessagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgportMessager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgportMessager *MsgportMessagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgportMessager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgportMessager *MsgportMessagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgportMessager.Contract.contract.Transact(opts, method, params...)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_MsgportMessager *MsgportMessagerCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_MsgportMessager *MsgportMessagerSession) Dao() (common.Address, error) {
	return _MsgportMessager.Contract.Dao(&_MsgportMessager.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_MsgportMessager *MsgportMessagerCallerSession) Dao() (common.Address, error) {
	return _MsgportMessager.Contract.Dao(&_MsgportMessager.CallOpts)
}

// MessagePayload is a free data retrieval call binding the contract method 0x6b57ff2b.
//
// Solidity: function messagePayload(address _from, address _to, bytes _message) view returns(bytes)
func (_MsgportMessager *MsgportMessagerCaller) MessagePayload(opts *bind.CallOpts, _from common.Address, _to common.Address, _message []byte) ([]byte, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "messagePayload", _from, _to, _message)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MessagePayload is a free data retrieval call binding the contract method 0x6b57ff2b.
//
// Solidity: function messagePayload(address _from, address _to, bytes _message) view returns(bytes)
func (_MsgportMessager *MsgportMessagerSession) MessagePayload(_from common.Address, _to common.Address, _message []byte) ([]byte, error) {
	return _MsgportMessager.Contract.MessagePayload(&_MsgportMessager.CallOpts, _from, _to, _message)
}

// MessagePayload is a free data retrieval call binding the contract method 0x6b57ff2b.
//
// Solidity: function messagePayload(address _from, address _to, bytes _message) view returns(bytes)
func (_MsgportMessager *MsgportMessagerCallerSession) MessagePayload(_from common.Address, _to common.Address, _message []byte) ([]byte, error) {
	return _MsgportMessager.Contract.MessagePayload(&_MsgportMessager.CallOpts, _from, _to, _message)
}

// Msgport is a free data retrieval call binding the contract method 0x2a02eadd.
//
// Solidity: function msgport() view returns(address)
func (_MsgportMessager *MsgportMessagerCaller) Msgport(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "msgport")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Msgport is a free data retrieval call binding the contract method 0x2a02eadd.
//
// Solidity: function msgport() view returns(address)
func (_MsgportMessager *MsgportMessagerSession) Msgport() (common.Address, error) {
	return _MsgportMessager.Contract.Msgport(&_MsgportMessager.CallOpts)
}

// Msgport is a free data retrieval call binding the contract method 0x2a02eadd.
//
// Solidity: function msgport() view returns(address)
func (_MsgportMessager *MsgportMessagerCallerSession) Msgport() (common.Address, error) {
	return _MsgportMessager.Contract.Msgport(&_MsgportMessager.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_MsgportMessager *MsgportMessagerCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_MsgportMessager *MsgportMessagerSession) Operator() (common.Address, error) {
	return _MsgportMessager.Contract.Operator(&_MsgportMessager.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_MsgportMessager *MsgportMessagerCallerSession) Operator() (common.Address, error) {
	return _MsgportMessager.Contract.Operator(&_MsgportMessager.CallOpts)
}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_MsgportMessager *MsgportMessagerCaller) PendingDao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "pendingDao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_MsgportMessager *MsgportMessagerSession) PendingDao() (common.Address, error) {
	return _MsgportMessager.Contract.PendingDao(&_MsgportMessager.CallOpts)
}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_MsgportMessager *MsgportMessagerCallerSession) PendingDao() (common.Address, error) {
	return _MsgportMessager.Contract.PendingDao(&_MsgportMessager.CallOpts)
}

// RemoteAppReceivers is a free data retrieval call binding the contract method 0x5a4220ad.
//
// Solidity: function remoteAppReceivers(bytes32 ) view returns(address)
func (_MsgportMessager *MsgportMessagerCaller) RemoteAppReceivers(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "remoteAppReceivers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteAppReceivers is a free data retrieval call binding the contract method 0x5a4220ad.
//
// Solidity: function remoteAppReceivers(bytes32 ) view returns(address)
func (_MsgportMessager *MsgportMessagerSession) RemoteAppReceivers(arg0 [32]byte) (common.Address, error) {
	return _MsgportMessager.Contract.RemoteAppReceivers(&_MsgportMessager.CallOpts, arg0)
}

// RemoteAppReceivers is a free data retrieval call binding the contract method 0x5a4220ad.
//
// Solidity: function remoteAppReceivers(bytes32 ) view returns(address)
func (_MsgportMessager *MsgportMessagerCallerSession) RemoteAppReceivers(arg0 [32]byte) (common.Address, error) {
	return _MsgportMessager.Contract.RemoteAppReceivers(&_MsgportMessager.CallOpts, arg0)
}

// RemoteAppSenders is a free data retrieval call binding the contract method 0x31c3ac91.
//
// Solidity: function remoteAppSenders(bytes32 ) view returns(address)
func (_MsgportMessager *MsgportMessagerCaller) RemoteAppSenders(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "remoteAppSenders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteAppSenders is a free data retrieval call binding the contract method 0x31c3ac91.
//
// Solidity: function remoteAppSenders(bytes32 ) view returns(address)
func (_MsgportMessager *MsgportMessagerSession) RemoteAppSenders(arg0 [32]byte) (common.Address, error) {
	return _MsgportMessager.Contract.RemoteAppSenders(&_MsgportMessager.CallOpts, arg0)
}

// RemoteAppSenders is a free data retrieval call binding the contract method 0x31c3ac91.
//
// Solidity: function remoteAppSenders(bytes32 ) view returns(address)
func (_MsgportMessager *MsgportMessagerCallerSession) RemoteAppSenders(arg0 [32]byte) (common.Address, error) {
	return _MsgportMessager.Contract.RemoteAppSenders(&_MsgportMessager.CallOpts, arg0)
}

// RemoteMessagers is a free data retrieval call binding the contract method 0xced4e6e0.
//
// Solidity: function remoteMessagers(uint256 ) view returns(uint256 msgportRemoteChainId, address messager)
func (_MsgportMessager *MsgportMessagerCaller) RemoteMessagers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MsgportRemoteChainId *big.Int
	Messager             common.Address
}, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "remoteMessagers", arg0)

	outstruct := new(struct {
		MsgportRemoteChainId *big.Int
		Messager             common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MsgportRemoteChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Messager = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RemoteMessagers is a free data retrieval call binding the contract method 0xced4e6e0.
//
// Solidity: function remoteMessagers(uint256 ) view returns(uint256 msgportRemoteChainId, address messager)
func (_MsgportMessager *MsgportMessagerSession) RemoteMessagers(arg0 *big.Int) (struct {
	MsgportRemoteChainId *big.Int
	Messager             common.Address
}, error) {
	return _MsgportMessager.Contract.RemoteMessagers(&_MsgportMessager.CallOpts, arg0)
}

// RemoteMessagers is a free data retrieval call binding the contract method 0xced4e6e0.
//
// Solidity: function remoteMessagers(uint256 ) view returns(uint256 msgportRemoteChainId, address messager)
func (_MsgportMessager *MsgportMessagerCallerSession) RemoteMessagers(arg0 *big.Int) (struct {
	MsgportRemoteChainId *big.Int
	Messager             common.Address
}, error) {
	return _MsgportMessager.Contract.RemoteMessagers(&_MsgportMessager.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_MsgportMessager *MsgportMessagerCaller) WhiteList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MsgportMessager.contract.Call(opts, &out, "whiteList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_MsgportMessager *MsgportMessagerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _MsgportMessager.Contract.WhiteList(&_MsgportMessager.CallOpts, arg0)
}

// WhiteList is a free data retrieval call binding the contract method 0x372c12b1.
//
// Solidity: function whiteList(address ) view returns(bool)
func (_MsgportMessager *MsgportMessagerCallerSession) WhiteList(arg0 common.Address) (bool, error) {
	return _MsgportMessager.Contract.WhiteList(&_MsgportMessager.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MsgportMessager *MsgportMessagerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MsgportMessager *MsgportMessagerSession) AcceptOwnership() (*types.Transaction, error) {
	return _MsgportMessager.Contract.AcceptOwnership(&_MsgportMessager.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MsgportMessager.Contract.AcceptOwnership(&_MsgportMessager.TransactOpts)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x1841a477.
//
// Solidity: function receiveMessage(uint256 _srcAppChainId, address _remoteAppAddress, address _localAppAddress, bytes _message) returns()
func (_MsgportMessager *MsgportMessagerTransactor) ReceiveMessage(opts *bind.TransactOpts, _srcAppChainId *big.Int, _remoteAppAddress common.Address, _localAppAddress common.Address, _message []byte) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "receiveMessage", _srcAppChainId, _remoteAppAddress, _localAppAddress, _message)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x1841a477.
//
// Solidity: function receiveMessage(uint256 _srcAppChainId, address _remoteAppAddress, address _localAppAddress, bytes _message) returns()
func (_MsgportMessager *MsgportMessagerSession) ReceiveMessage(_srcAppChainId *big.Int, _remoteAppAddress common.Address, _localAppAddress common.Address, _message []byte) (*types.Transaction, error) {
	return _MsgportMessager.Contract.ReceiveMessage(&_MsgportMessager.TransactOpts, _srcAppChainId, _remoteAppAddress, _localAppAddress, _message)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x1841a477.
//
// Solidity: function receiveMessage(uint256 _srcAppChainId, address _remoteAppAddress, address _localAppAddress, bytes _message) returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) ReceiveMessage(_srcAppChainId *big.Int, _remoteAppAddress common.Address, _localAppAddress common.Address, _message []byte) (*types.Transaction, error) {
	return _MsgportMessager.Contract.ReceiveMessage(&_MsgportMessager.TransactOpts, _srcAppChainId, _remoteAppAddress, _localAppAddress, _message)
}

// RegisterRemoteReceiver is a paid mutator transaction binding the contract method 0xc2fdb804.
//
// Solidity: function registerRemoteReceiver(uint256 _remoteChainId, address _remoteBridge) returns()
func (_MsgportMessager *MsgportMessagerTransactor) RegisterRemoteReceiver(opts *bind.TransactOpts, _remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "registerRemoteReceiver", _remoteChainId, _remoteBridge)
}

// RegisterRemoteReceiver is a paid mutator transaction binding the contract method 0xc2fdb804.
//
// Solidity: function registerRemoteReceiver(uint256 _remoteChainId, address _remoteBridge) returns()
func (_MsgportMessager *MsgportMessagerSession) RegisterRemoteReceiver(_remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.RegisterRemoteReceiver(&_MsgportMessager.TransactOpts, _remoteChainId, _remoteBridge)
}

// RegisterRemoteReceiver is a paid mutator transaction binding the contract method 0xc2fdb804.
//
// Solidity: function registerRemoteReceiver(uint256 _remoteChainId, address _remoteBridge) returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) RegisterRemoteReceiver(_remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.RegisterRemoteReceiver(&_MsgportMessager.TransactOpts, _remoteChainId, _remoteBridge)
}

// RegisterRemoteSender is a paid mutator transaction binding the contract method 0xc07d7713.
//
// Solidity: function registerRemoteSender(uint256 _remoteChainId, address _remoteBridge) returns()
func (_MsgportMessager *MsgportMessagerTransactor) RegisterRemoteSender(opts *bind.TransactOpts, _remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "registerRemoteSender", _remoteChainId, _remoteBridge)
}

// RegisterRemoteSender is a paid mutator transaction binding the contract method 0xc07d7713.
//
// Solidity: function registerRemoteSender(uint256 _remoteChainId, address _remoteBridge) returns()
func (_MsgportMessager *MsgportMessagerSession) RegisterRemoteSender(_remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.RegisterRemoteSender(&_MsgportMessager.TransactOpts, _remoteChainId, _remoteBridge)
}

// RegisterRemoteSender is a paid mutator transaction binding the contract method 0xc07d7713.
//
// Solidity: function registerRemoteSender(uint256 _remoteChainId, address _remoteBridge) returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) RegisterRemoteSender(_remoteChainId *big.Int, _remoteBridge common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.RegisterRemoteSender(&_MsgportMessager.TransactOpts, _remoteChainId, _remoteBridge)
}

// SendMessage is a paid mutator transaction binding the contract method 0x6ea9cec9.
//
// Solidity: function sendMessage(uint256 _remoteChainId, bytes _message, bytes _params) payable returns()
func (_MsgportMessager *MsgportMessagerTransactor) SendMessage(opts *bind.TransactOpts, _remoteChainId *big.Int, _message []byte, _params []byte) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "sendMessage", _remoteChainId, _message, _params)
}

// SendMessage is a paid mutator transaction binding the contract method 0x6ea9cec9.
//
// Solidity: function sendMessage(uint256 _remoteChainId, bytes _message, bytes _params) payable returns()
func (_MsgportMessager *MsgportMessagerSession) SendMessage(_remoteChainId *big.Int, _message []byte, _params []byte) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SendMessage(&_MsgportMessager.TransactOpts, _remoteChainId, _message, _params)
}

// SendMessage is a paid mutator transaction binding the contract method 0x6ea9cec9.
//
// Solidity: function sendMessage(uint256 _remoteChainId, bytes _message, bytes _params) payable returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) SendMessage(_remoteChainId *big.Int, _message []byte, _params []byte) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SendMessage(&_MsgportMessager.TransactOpts, _remoteChainId, _message, _params)
}

// SetMsgPort is a paid mutator transaction binding the contract method 0x231c271c.
//
// Solidity: function setMsgPort(address _msgport) returns()
func (_MsgportMessager *MsgportMessagerTransactor) SetMsgPort(opts *bind.TransactOpts, _msgport common.Address) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "setMsgPort", _msgport)
}

// SetMsgPort is a paid mutator transaction binding the contract method 0x231c271c.
//
// Solidity: function setMsgPort(address _msgport) returns()
func (_MsgportMessager *MsgportMessagerSession) SetMsgPort(_msgport common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SetMsgPort(&_MsgportMessager.TransactOpts, _msgport)
}

// SetMsgPort is a paid mutator transaction binding the contract method 0x231c271c.
//
// Solidity: function setMsgPort(address _msgport) returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) SetMsgPort(_msgport common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SetMsgPort(&_MsgportMessager.TransactOpts, _msgport)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_MsgportMessager *MsgportMessagerTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_MsgportMessager *MsgportMessagerSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SetOperator(&_MsgportMessager.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SetOperator(&_MsgportMessager.TransactOpts, _operator)
}

// SetRemoteMessager is a paid mutator transaction binding the contract method 0xc59b21c8.
//
// Solidity: function setRemoteMessager(uint256 _appRemoteChainId, uint256 _msgportRemoteChainId, address _remoteMessager) returns()
func (_MsgportMessager *MsgportMessagerTransactor) SetRemoteMessager(opts *bind.TransactOpts, _appRemoteChainId *big.Int, _msgportRemoteChainId *big.Int, _remoteMessager common.Address) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "setRemoteMessager", _appRemoteChainId, _msgportRemoteChainId, _remoteMessager)
}

// SetRemoteMessager is a paid mutator transaction binding the contract method 0xc59b21c8.
//
// Solidity: function setRemoteMessager(uint256 _appRemoteChainId, uint256 _msgportRemoteChainId, address _remoteMessager) returns()
func (_MsgportMessager *MsgportMessagerSession) SetRemoteMessager(_appRemoteChainId *big.Int, _msgportRemoteChainId *big.Int, _remoteMessager common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SetRemoteMessager(&_MsgportMessager.TransactOpts, _appRemoteChainId, _msgportRemoteChainId, _remoteMessager)
}

// SetRemoteMessager is a paid mutator transaction binding the contract method 0xc59b21c8.
//
// Solidity: function setRemoteMessager(uint256 _appRemoteChainId, uint256 _msgportRemoteChainId, address _remoteMessager) returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) SetRemoteMessager(_appRemoteChainId *big.Int, _msgportRemoteChainId *big.Int, _remoteMessager common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SetRemoteMessager(&_MsgportMessager.TransactOpts, _appRemoteChainId, _msgportRemoteChainId, _remoteMessager)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _caller, bool _enable) returns()
func (_MsgportMessager *MsgportMessagerTransactor) SetWhiteList(opts *bind.TransactOpts, _caller common.Address, _enable bool) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "setWhiteList", _caller, _enable)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _caller, bool _enable) returns()
func (_MsgportMessager *MsgportMessagerSession) SetWhiteList(_caller common.Address, _enable bool) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SetWhiteList(&_MsgportMessager.TransactOpts, _caller, _enable)
}

// SetWhiteList is a paid mutator transaction binding the contract method 0x8d14e127.
//
// Solidity: function setWhiteList(address _caller, bool _enable) returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) SetWhiteList(_caller common.Address, _enable bool) (*types.Transaction, error) {
	return _MsgportMessager.Contract.SetWhiteList(&_MsgportMessager.TransactOpts, _caller, _enable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_MsgportMessager *MsgportMessagerTransactor) TransferOwnership(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _MsgportMessager.contract.Transact(opts, "transferOwnership", _dao)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_MsgportMessager *MsgportMessagerSession) TransferOwnership(_dao common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.TransferOwnership(&_MsgportMessager.TransactOpts, _dao)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_MsgportMessager *MsgportMessagerTransactorSession) TransferOwnership(_dao common.Address) (*types.Transaction, error) {
	return _MsgportMessager.Contract.TransferOwnership(&_MsgportMessager.TransactOpts, _dao)
}

// MsgportMessagerCallResultIterator is returned from FilterCallResult and is used to iterate over the raw logs and unpacked data for CallResult events raised by the MsgportMessager contract.
type MsgportMessagerCallResultIterator struct {
	Event *MsgportMessagerCallResult // Event containing the contract specifics and raw log

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
func (it *MsgportMessagerCallResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgportMessagerCallResult)
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
		it.Event = new(MsgportMessagerCallResult)
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
func (it *MsgportMessagerCallResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgportMessagerCallResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgportMessagerCallResult represents a CallResult event raised by the MsgportMessager contract.
type MsgportMessagerCallResult struct {
	SrcAppChainId *big.Int
	Result        bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCallResult is a free log retrieval operation binding the contract event 0x9d94ccd595d0b48c5a63dfa57d2feefeb060d6c47f945076accb66b1d2104422.
//
// Solidity: event CallResult(uint256 srcAppChainId, bool result)
func (_MsgportMessager *MsgportMessagerFilterer) FilterCallResult(opts *bind.FilterOpts) (*MsgportMessagerCallResultIterator, error) {

	logs, sub, err := _MsgportMessager.contract.FilterLogs(opts, "CallResult")
	if err != nil {
		return nil, err
	}
	return &MsgportMessagerCallResultIterator{contract: _MsgportMessager.contract, event: "CallResult", logs: logs, sub: sub}, nil
}

// WatchCallResult is a free log subscription operation binding the contract event 0x9d94ccd595d0b48c5a63dfa57d2feefeb060d6c47f945076accb66b1d2104422.
//
// Solidity: event CallResult(uint256 srcAppChainId, bool result)
func (_MsgportMessager *MsgportMessagerFilterer) WatchCallResult(opts *bind.WatchOpts, sink chan<- *MsgportMessagerCallResult) (event.Subscription, error) {

	logs, sub, err := _MsgportMessager.contract.WatchLogs(opts, "CallResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgportMessagerCallResult)
				if err := _MsgportMessager.contract.UnpackLog(event, "CallResult", log); err != nil {
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

// ParseCallResult is a log parse operation binding the contract event 0x9d94ccd595d0b48c5a63dfa57d2feefeb060d6c47f945076accb66b1d2104422.
//
// Solidity: event CallResult(uint256 srcAppChainId, bool result)
func (_MsgportMessager *MsgportMessagerFilterer) ParseCallResult(log types.Log) (*MsgportMessagerCallResult, error) {
	event := new(MsgportMessagerCallResult)
	if err := _MsgportMessager.contract.UnpackLog(event, "CallResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgportMessagerCallerUnMatchedIterator is returned from FilterCallerUnMatched and is used to iterate over the raw logs and unpacked data for CallerUnMatched events raised by the MsgportMessager contract.
type MsgportMessagerCallerUnMatchedIterator struct {
	Event *MsgportMessagerCallerUnMatched // Event containing the contract specifics and raw log

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
func (it *MsgportMessagerCallerUnMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MsgportMessagerCallerUnMatched)
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
		it.Event = new(MsgportMessagerCallerUnMatched)
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
func (it *MsgportMessagerCallerUnMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MsgportMessagerCallerUnMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MsgportMessagerCallerUnMatched represents a CallerUnMatched event raised by the MsgportMessager contract.
type MsgportMessagerCallerUnMatched struct {
	SrcAppChainId *big.Int
	SrcAppAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCallerUnMatched is a free log retrieval operation binding the contract event 0xad25dc355e3dabdf8324b630aec3edece57dbaa0713df1ce4e2fb51657256440.
//
// Solidity: event CallerUnMatched(uint256 srcAppChainId, address srcAppAddress)
func (_MsgportMessager *MsgportMessagerFilterer) FilterCallerUnMatched(opts *bind.FilterOpts) (*MsgportMessagerCallerUnMatchedIterator, error) {

	logs, sub, err := _MsgportMessager.contract.FilterLogs(opts, "CallerUnMatched")
	if err != nil {
		return nil, err
	}
	return &MsgportMessagerCallerUnMatchedIterator{contract: _MsgportMessager.contract, event: "CallerUnMatched", logs: logs, sub: sub}, nil
}

// WatchCallerUnMatched is a free log subscription operation binding the contract event 0xad25dc355e3dabdf8324b630aec3edece57dbaa0713df1ce4e2fb51657256440.
//
// Solidity: event CallerUnMatched(uint256 srcAppChainId, address srcAppAddress)
func (_MsgportMessager *MsgportMessagerFilterer) WatchCallerUnMatched(opts *bind.WatchOpts, sink chan<- *MsgportMessagerCallerUnMatched) (event.Subscription, error) {

	logs, sub, err := _MsgportMessager.contract.WatchLogs(opts, "CallerUnMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MsgportMessagerCallerUnMatched)
				if err := _MsgportMessager.contract.UnpackLog(event, "CallerUnMatched", log); err != nil {
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

// ParseCallerUnMatched is a log parse operation binding the contract event 0xad25dc355e3dabdf8324b630aec3edece57dbaa0713df1ce4e2fb51657256440.
//
// Solidity: event CallerUnMatched(uint256 srcAppChainId, address srcAppAddress)
func (_MsgportMessager *MsgportMessagerFilterer) ParseCallerUnMatched(log types.Log) (*MsgportMessagerCallerUnMatched, error) {
	event := new(MsgportMessagerCallerUnMatched)
	if err := _MsgportMessager.contract.UnpackLog(event, "CallerUnMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


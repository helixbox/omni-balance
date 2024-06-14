// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi_v3

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

// LnBridgeSourceV3TransferParams is an auto generated low-level Go binding around an user-defined struct.
type LnBridgeSourceV3TransferParams struct {
	RemoteChainId *big.Int
	Provider      common.Address
	SourceToken   common.Address
	TargetToken   common.Address
	TotalFee      *big.Int
	Amount        *big.Int
	Receiver      common.Address
	Timestamp     *big.Int
}

// AbiV3MetaData contains all meta data concerning the AbiV3 contract.
var AbiV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"totalFee\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLnBridgeSourceV3.TransferParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"lockAndRemoteRelease\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AbiV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiV3MetaData.ABI instead.
var AbiV3ABI = AbiV3MetaData.ABI

// AbiV3 is an auto generated Go binding around an Ethereum contract.
type AbiV3 struct {
	AbiV3Caller     // Read-only binding to the contract
	AbiV3Transactor // Write-only binding to the contract
	AbiV3Filterer   // Log filterer for contract events
}

// AbiV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type AbiV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiV3Session struct {
	Contract     *AbiV3            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiV3CallerSession struct {
	Contract *AbiV3Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiV3TransactorSession struct {
	Contract     *AbiV3Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type AbiV3Raw struct {
	Contract *AbiV3 // Generic contract binding to access the raw methods on
}

// AbiV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiV3CallerRaw struct {
	Contract *AbiV3Caller // Generic read-only contract binding to access the raw methods on
}

// AbiV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiV3TransactorRaw struct {
	Contract *AbiV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAbiV3 creates a new instance of AbiV3, bound to a specific deployed contract.
func NewAbiV3(address common.Address, backend bind.ContractBackend) (*AbiV3, error) {
	contract, err := bindAbiV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbiV3{AbiV3Caller: AbiV3Caller{contract: contract}, AbiV3Transactor: AbiV3Transactor{contract: contract}, AbiV3Filterer: AbiV3Filterer{contract: contract}}, nil
}

// NewAbiV3Caller creates a new read-only instance of AbiV3, bound to a specific deployed contract.
func NewAbiV3Caller(address common.Address, caller bind.ContractCaller) (*AbiV3Caller, error) {
	contract, err := bindAbiV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiV3Caller{contract: contract}, nil
}

// NewAbiV3Transactor creates a new write-only instance of AbiV3, bound to a specific deployed contract.
func NewAbiV3Transactor(address common.Address, transactor bind.ContractTransactor) (*AbiV3Transactor, error) {
	contract, err := bindAbiV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiV3Transactor{contract: contract}, nil
}

// NewAbiV3Filterer creates a new log filterer instance of AbiV3, bound to a specific deployed contract.
func NewAbiV3Filterer(address common.Address, filterer bind.ContractFilterer) (*AbiV3Filterer, error) {
	contract, err := bindAbiV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiV3Filterer{contract: contract}, nil
}

// bindAbiV3 binds a generic wrapper to an already deployed contract.
func bindAbiV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbiV3 *AbiV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbiV3.Contract.AbiV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbiV3 *AbiV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbiV3.Contract.AbiV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbiV3 *AbiV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbiV3.Contract.AbiV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbiV3 *AbiV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbiV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbiV3 *AbiV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbiV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbiV3 *AbiV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbiV3.Contract.contract.Transact(opts, method, params...)
}

// LockAndRemoteRelease is a paid mutator transaction binding the contract method 0x9cd13471.
//
// Solidity: function lockAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params) payable returns()
func (_AbiV3 *AbiV3Transactor) LockAndRemoteRelease(opts *bind.TransactOpts, _params LnBridgeSourceV3TransferParams) (*types.Transaction, error) {
	return _AbiV3.contract.Transact(opts, "lockAndRemoteRelease", _params)
}

// LockAndRemoteRelease is a paid mutator transaction binding the contract method 0x9cd13471.
//
// Solidity: function lockAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params) payable returns()
func (_AbiV3 *AbiV3Session) LockAndRemoteRelease(_params LnBridgeSourceV3TransferParams) (*types.Transaction, error) {
	return _AbiV3.Contract.LockAndRemoteRelease(&_AbiV3.TransactOpts, _params)
}

// LockAndRemoteRelease is a paid mutator transaction binding the contract method 0x9cd13471.
//
// Solidity: function lockAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params) payable returns()
func (_AbiV3 *AbiV3TransactorSession) LockAndRemoteRelease(_params LnBridgeSourceV3TransferParams) (*types.Transaction, error) {
	return _AbiV3.Contract.LockAndRemoteRelease(&_AbiV3.TransactOpts, _params)
}

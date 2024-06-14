// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi_v2_opposite

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

// LnOppositeBridgeSourceSnapshot is an auto generated low-level Go binding around an user-defined struct.
type LnOppositeBridgeSourceSnapshot struct {
	RemoteChainId   *big.Int
	Provider        common.Address
	SourceToken     common.Address
	TargetToken     common.Address
	TransferId      [32]byte
	TotalFee        *big.Int
	DepositedMargin *big.Int
}

// AbiV2OppositeMetaData contains all meta data concerning the AbiV2Opposite contract.
var AbiV2OppositeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"uint112\",\"name\":\"totalFee\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"depositedMargin\",\"type\":\"uint112\"}],\"internalType\":\"structLnOppositeBridgeSource.Snapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"},{\"internalType\":\"uint112\",\"name\":\"_amount\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"transferAndLockMargin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AbiV2OppositeABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiV2OppositeMetaData.ABI instead.
var AbiV2OppositeABI = AbiV2OppositeMetaData.ABI

// AbiV2Opposite is an auto generated Go binding around an Ethereum contract.
type AbiV2Opposite struct {
	AbiV2OppositeCaller     // Read-only binding to the contract
	AbiV2OppositeTransactor // Write-only binding to the contract
	AbiV2OppositeFilterer   // Log filterer for contract events
}

// AbiV2OppositeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiV2OppositeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiV2OppositeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiV2OppositeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiV2OppositeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiV2OppositeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiV2OppositeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiV2OppositeSession struct {
	Contract     *AbiV2Opposite    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiV2OppositeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiV2OppositeCallerSession struct {
	Contract *AbiV2OppositeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AbiV2OppositeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiV2OppositeTransactorSession struct {
	Contract     *AbiV2OppositeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AbiV2OppositeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiV2OppositeRaw struct {
	Contract *AbiV2Opposite // Generic contract binding to access the raw methods on
}

// AbiV2OppositeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiV2OppositeCallerRaw struct {
	Contract *AbiV2OppositeCaller // Generic read-only contract binding to access the raw methods on
}

// AbiV2OppositeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiV2OppositeTransactorRaw struct {
	Contract *AbiV2OppositeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbiV2Opposite creates a new instance of AbiV2Opposite, bound to a specific deployed contract.
func NewAbiV2Opposite(address common.Address, backend bind.ContractBackend) (*AbiV2Opposite, error) {
	contract, err := bindAbiV2Opposite(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbiV2Opposite{AbiV2OppositeCaller: AbiV2OppositeCaller{contract: contract}, AbiV2OppositeTransactor: AbiV2OppositeTransactor{contract: contract}, AbiV2OppositeFilterer: AbiV2OppositeFilterer{contract: contract}}, nil
}

// NewAbiV2OppositeCaller creates a new read-only instance of AbiV2Opposite, bound to a specific deployed contract.
func NewAbiV2OppositeCaller(address common.Address, caller bind.ContractCaller) (*AbiV2OppositeCaller, error) {
	contract, err := bindAbiV2Opposite(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiV2OppositeCaller{contract: contract}, nil
}

// NewAbiV2OppositeTransactor creates a new write-only instance of AbiV2Opposite, bound to a specific deployed contract.
func NewAbiV2OppositeTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiV2OppositeTransactor, error) {
	contract, err := bindAbiV2Opposite(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiV2OppositeTransactor{contract: contract}, nil
}

// NewAbiV2OppositeFilterer creates a new log filterer instance of AbiV2Opposite, bound to a specific deployed contract.
func NewAbiV2OppositeFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiV2OppositeFilterer, error) {
	contract, err := bindAbiV2Opposite(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiV2OppositeFilterer{contract: contract}, nil
}

// bindAbiV2Opposite binds a generic wrapper to an already deployed contract.
func bindAbiV2Opposite(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbiV2OppositeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbiV2Opposite *AbiV2OppositeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbiV2Opposite.Contract.AbiV2OppositeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbiV2Opposite *AbiV2OppositeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbiV2Opposite.Contract.AbiV2OppositeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbiV2Opposite *AbiV2OppositeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbiV2Opposite.Contract.AbiV2OppositeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbiV2Opposite *AbiV2OppositeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbiV2Opposite.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbiV2Opposite *AbiV2OppositeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbiV2Opposite.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbiV2Opposite *AbiV2OppositeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbiV2Opposite.Contract.contract.Transact(opts, method, params...)
}

// TransferAndLockMargin is a paid mutator transaction binding the contract method 0x2656c147.
//
// Solidity: function transferAndLockMargin((uint256,address,address,address,bytes32,uint112,uint112) _snapshot, uint112 _amount, address _receiver) payable returns()
func (_AbiV2Opposite *AbiV2OppositeTransactor) TransferAndLockMargin(opts *bind.TransactOpts, _snapshot LnOppositeBridgeSourceSnapshot, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _AbiV2Opposite.contract.Transact(opts, "transferAndLockMargin", _snapshot, _amount, _receiver)
}

// TransferAndLockMargin is a paid mutator transaction binding the contract method 0x2656c147.
//
// Solidity: function transferAndLockMargin((uint256,address,address,address,bytes32,uint112,uint112) _snapshot, uint112 _amount, address _receiver) payable returns()
func (_AbiV2Opposite *AbiV2OppositeSession) TransferAndLockMargin(_snapshot LnOppositeBridgeSourceSnapshot, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _AbiV2Opposite.Contract.TransferAndLockMargin(&_AbiV2Opposite.TransactOpts, _snapshot, _amount, _receiver)
}

// TransferAndLockMargin is a paid mutator transaction binding the contract method 0x2656c147.
//
// Solidity: function transferAndLockMargin((uint256,address,address,address,bytes32,uint112,uint112) _snapshot, uint112 _amount, address _receiver) payable returns()
func (_AbiV2Opposite *AbiV2OppositeTransactorSession) TransferAndLockMargin(_snapshot LnOppositeBridgeSourceSnapshot, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _AbiV2Opposite.Contract.TransferAndLockMargin(&_AbiV2Opposite.TransactOpts, _snapshot, _amount, _receiver)
}

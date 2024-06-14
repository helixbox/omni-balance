// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi_v2_default

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

// LnDefaultBridgeSourceSnapshot is an auto generated low-level Go binding around an user-defined struct.
type LnDefaultBridgeSourceSnapshot struct {
	RemoteChainId *big.Int
	Provider      common.Address
	SourceToken   common.Address
	TargetToken   common.Address
	TransferId    [32]byte
	TotalFee      *big.Int
	WithdrawNonce uint64
}

// HelixMetaData contains all meta data concerning the Helix contract.
var HelixMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"uint112\",\"name\":\"totalFee\",\"type\":\"uint112\"},{\"internalType\":\"uint64\",\"name\":\"withdrawNonce\",\"type\":\"uint64\"}],\"internalType\":\"structLnDefaultBridgeSource.Snapshot\",\"name\":\"_snapshot\",\"type\":\"tuple\"},{\"internalType\":\"uint112\",\"name\":\"_amount\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"transferAndLockMargin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// HelixABI is the input ABI used to generate the binding from.
// Deprecated: Use HelixMetaData.ABI instead.
var HelixABI = HelixMetaData.ABI

// Helix is an auto generated Go binding around an Ethereum contract.
type Helix struct {
	HelixCaller     // Read-only binding to the contract
	HelixTransactor // Write-only binding to the contract
	HelixFilterer   // Log filterer for contract events
}

// HelixCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelixCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelixTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelixTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelixFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelixFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelixSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelixSession struct {
	Contract     *Helix            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelixCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelixCallerSession struct {
	Contract *HelixCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelixTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelixTransactorSession struct {
	Contract     *HelixTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelixRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelixRaw struct {
	Contract *Helix // Generic contract binding to access the raw methods on
}

// HelixCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelixCallerRaw struct {
	Contract *HelixCaller // Generic read-only contract binding to access the raw methods on
}

// HelixTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelixTransactorRaw struct {
	Contract *HelixTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHelix creates a new instance of Helix, bound to a specific deployed contract.
func NewHelix(address common.Address, backend bind.ContractBackend) (*Helix, error) {
	contract, err := bindHelix(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Helix{HelixCaller: HelixCaller{contract: contract}, HelixTransactor: HelixTransactor{contract: contract}, HelixFilterer: HelixFilterer{contract: contract}}, nil
}

// NewHelixCaller creates a new read-only instance of Helix, bound to a specific deployed contract.
func NewHelixCaller(address common.Address, caller bind.ContractCaller) (*HelixCaller, error) {
	contract, err := bindHelix(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelixCaller{contract: contract}, nil
}

// NewHelixTransactor creates a new write-only instance of Helix, bound to a specific deployed contract.
func NewHelixTransactor(address common.Address, transactor bind.ContractTransactor) (*HelixTransactor, error) {
	contract, err := bindHelix(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelixTransactor{contract: contract}, nil
}

// NewHelixFilterer creates a new log filterer instance of Helix, bound to a specific deployed contract.
func NewHelixFilterer(address common.Address, filterer bind.ContractFilterer) (*HelixFilterer, error) {
	contract, err := bindHelix(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelixFilterer{contract: contract}, nil
}

// bindHelix binds a generic wrapper to an already deployed contract.
func bindHelix(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HelixMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helix *HelixRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helix.Contract.HelixCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helix *HelixRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helix.Contract.HelixTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helix *HelixRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helix.Contract.HelixTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Helix *HelixCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Helix.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Helix *HelixTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Helix.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Helix *HelixTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Helix.Contract.contract.Transact(opts, method, params...)
}

// TransferAndLockMargin is a paid mutator transaction binding the contract method 0x0be175a5.
//
// Solidity: function transferAndLockMargin((uint256,address,address,address,bytes32,uint112,uint64) _snapshot, uint112 _amount, address _receiver) payable returns()
func (_Helix *HelixTransactor) TransferAndLockMargin(opts *bind.TransactOpts, _snapshot LnDefaultBridgeSourceSnapshot, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Helix.contract.Transact(opts, "transferAndLockMargin", _snapshot, _amount, _receiver)
}

// TransferAndLockMargin is a paid mutator transaction binding the contract method 0x0be175a5.
//
// Solidity: function transferAndLockMargin((uint256,address,address,address,bytes32,uint112,uint64) _snapshot, uint112 _amount, address _receiver) payable returns()
func (_Helix *HelixSession) TransferAndLockMargin(_snapshot LnDefaultBridgeSourceSnapshot, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Helix.Contract.TransferAndLockMargin(&_Helix.TransactOpts, _snapshot, _amount, _receiver)
}

// TransferAndLockMargin is a paid mutator transaction binding the contract method 0x0be175a5.
//
// Solidity: function transferAndLockMargin((uint256,address,address,address,bytes32,uint112,uint64) _snapshot, uint112 _amount, address _receiver) payable returns()
func (_Helix *HelixTransactorSession) TransferAndLockMargin(_snapshot LnDefaultBridgeSourceSnapshot, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Helix.Contract.TransferAndLockMargin(&_Helix.TransactOpts, _snapshot, _amount, _receiver)
}

// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gnosis_transmuter

import (
	"bytes"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = bytes.Equal
	_ = errors.New
	_ = big.NewInt
	_ = common.Big1
	_ = types.BloomLookup
	_ = abi.ConvertType
)

// GnosisTransmuterMetaData contains all meta data concerning the GnosisTransmuter contract.
var GnosisTransmuterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"HOME_OMNIBRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC_E\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDC_ON_XDAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableTransmuter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenBridged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"rebalanceUSDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "GnosisTransmuter",
}

// GnosisTransmuter is an auto generated Go binding around an Ethereum contract.
type GnosisTransmuter struct {
	abi abi.ABI
}

// NewGnosisTransmuter creates a new instance of GnosisTransmuter.
func NewGnosisTransmuter() *GnosisTransmuter {
	parsed, err := GnosisTransmuterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &GnosisTransmuter{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *GnosisTransmuter) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackHOMEOMNIBRIDGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed324463.
//
// Solidity: function HOME_OMNIBRIDGE() view returns(address)
func (gnosisTransmuter *GnosisTransmuter) PackHOMEOMNIBRIDGE() []byte {
	enc, err := gnosisTransmuter.abi.Pack("HOME_OMNIBRIDGE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackHOMEOMNIBRIDGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed324463.
//
// Solidity: function HOME_OMNIBRIDGE() view returns(address)
func (gnosisTransmuter *GnosisTransmuter) UnpackHOMEOMNIBRIDGE(data []byte) (common.Address, error) {
	out, err := gnosisTransmuter.abi.Unpack("HOME_OMNIBRIDGE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUSDCE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0d2c5441.
//
// Solidity: function USDC_E() view returns(address)
func (gnosisTransmuter *GnosisTransmuter) PackUSDCE() []byte {
	enc, err := gnosisTransmuter.abi.Pack("USDC_E")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUSDCE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0d2c5441.
//
// Solidity: function USDC_E() view returns(address)
func (gnosisTransmuter *GnosisTransmuter) UnpackUSDCE(data []byte) (common.Address, error) {
	out, err := gnosisTransmuter.abi.Unpack("USDC_E", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUSDCONXDAI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6dfe0fda.
//
// Solidity: function USDC_ON_XDAI() view returns(address)
func (gnosisTransmuter *GnosisTransmuter) PackUSDCONXDAI() []byte {
	enc, err := gnosisTransmuter.abi.Pack("USDC_ON_XDAI")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUSDCONXDAI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6dfe0fda.
//
// Solidity: function USDC_ON_XDAI() view returns(address)
func (gnosisTransmuter *GnosisTransmuter) UnpackUSDCONXDAI(data []byte) (common.Address, error) {
	out, err := gnosisTransmuter.abi.Unpack("USDC_ON_XDAI", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDeposit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (gnosisTransmuter *GnosisTransmuter) PackDeposit(amount *big.Int) []byte {
	enc, err := gnosisTransmuter.abi.Pack("deposit", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDisableTransmuter is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xca4872cc.
//
// Solidity: function disableTransmuter() returns()
func (gnosisTransmuter *GnosisTransmuter) PackDisableTransmuter() []byte {
	enc, err := gnosisTransmuter.abi.Pack("disableTransmuter")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6aa633b6.
//
// Solidity: function isEnabled() view returns(bool)
func (gnosisTransmuter *GnosisTransmuter) PackIsEnabled() []byte {
	enc, err := gnosisTransmuter.abi.Pack("isEnabled")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsEnabled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6aa633b6.
//
// Solidity: function isEnabled() view returns(bool)
func (gnosisTransmuter *GnosisTransmuter) UnpackIsEnabled(data []byte) (bool, error) {
	out, err := gnosisTransmuter.abi.Unpack("isEnabled", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOnTokenBridged is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb7af854.
//
// Solidity: function onTokenBridged(address token, uint256 value, bytes data) returns()
func (gnosisTransmuter *GnosisTransmuter) PackOnTokenBridged(token common.Address, value *big.Int, data []byte) []byte {
	enc, err := gnosisTransmuter.abi.Pack("onTokenBridged", token, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (gnosisTransmuter *GnosisTransmuter) PackOwner() []byte {
	enc, err := gnosisTransmuter.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (gnosisTransmuter *GnosisTransmuter) UnpackOwner(data []byte) (common.Address, error) {
	out, err := gnosisTransmuter.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackRebalanceUSDC is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x90950640.
//
// Solidity: function rebalanceUSDC(address receiver) returns()
func (gnosisTransmuter *GnosisTransmuter) PackRebalanceUSDC(receiver common.Address) []byte {
	enc, err := gnosisTransmuter.abi.Pack("rebalanceUSDC", receiver)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (gnosisTransmuter *GnosisTransmuter) PackRenounceOwnership() []byte {
	enc, err := gnosisTransmuter.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (gnosisTransmuter *GnosisTransmuter) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := gnosisTransmuter.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithdraw is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (gnosisTransmuter *GnosisTransmuter) PackWithdraw(amount *big.Int) []byte {
	enc, err := gnosisTransmuter.abi.Pack("withdraw", amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// GnosisTransmuterDeposit represents a Deposit event raised by the GnosisTransmuter contract.
type GnosisTransmuterDeposit struct {
	Depositor common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const GnosisTransmuterDepositEventName = "Deposit"

// ContractEventName returns the user-defined event name.
func (GnosisTransmuterDeposit) ContractEventName() string {
	return GnosisTransmuterDepositEventName
}

// UnpackDepositEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Deposit(address indexed depositor, uint256 indexed amount)
func (gnosisTransmuter *GnosisTransmuter) UnpackDepositEvent(log *types.Log) (*GnosisTransmuterDeposit, error) {
	event := "Deposit"
	if log.Topics[0] != gnosisTransmuter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisTransmuterDeposit)
	if len(log.Data) > 0 {
		if err := gnosisTransmuter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisTransmuter.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GnosisTransmuterOwnershipTransferred represents a OwnershipTransferred event raised by the GnosisTransmuter contract.
type GnosisTransmuterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const GnosisTransmuterOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (GnosisTransmuterOwnershipTransferred) ContractEventName() string {
	return GnosisTransmuterOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (gnosisTransmuter *GnosisTransmuter) UnpackOwnershipTransferredEvent(log *types.Log) (*GnosisTransmuterOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != gnosisTransmuter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisTransmuterOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := gnosisTransmuter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisTransmuter.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

// GnosisTransmuterWithdraw represents a Withdraw event raised by the GnosisTransmuter contract.
type GnosisTransmuterWithdraw struct {
	Depositor common.Address
	Amount    *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const GnosisTransmuterWithdrawEventName = "Withdraw"

// ContractEventName returns the user-defined event name.
func (GnosisTransmuterWithdraw) ContractEventName() string {
	return GnosisTransmuterWithdrawEventName
}

// UnpackWithdrawEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Withdraw(address indexed depositor, uint256 indexed amount)
func (gnosisTransmuter *GnosisTransmuter) UnpackWithdrawEvent(log *types.Log) (*GnosisTransmuterWithdraw, error) {
	event := "Withdraw"
	if log.Topics[0] != gnosisTransmuter.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisTransmuterWithdraw)
	if len(log.Data) > 0 {
		if err := gnosisTransmuter.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisTransmuter.abi.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}
	out.Raw = log
	return out, nil
}

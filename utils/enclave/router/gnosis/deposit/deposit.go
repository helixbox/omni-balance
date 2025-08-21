// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gnosis_deposit

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

// GnosisDepositMetaData contains all meta data concerning the GnosisDeposit contract.
var GnosisDepositMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_suffix\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"DailyLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"ExecutionDailyLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"FailedMessageFixed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"}],\"name\":\"NewTokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"TokensBridged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageId\",\"type\":\"bytes32\"}],\"name\":\"TokensBridgingInitiated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeContract\",\"outputs\":[{\"internalType\":\"contractIAMB\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeToken\",\"type\":\"address\"}],\"name\":\"bridgedTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"claimTokensFromTokenContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"dailyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"deployAndHandleBridgedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"deployAndHandleBridgedTokensAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"disableInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"executionDailyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"executionMaxPerTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"fixFailedMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"fixMediatorBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeInterfacesVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"major\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minor\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"patch\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeMode\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"_data\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"handleBridgedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"handleBridgedTokensAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"handleNativeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"handleNativeTokensAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mediatorContract\",\"type\":\"address\"},{\"internalType\":\"uint256[3]\",\"name\":\"_dailyLimitMaxPerTxMinPerTxArray\",\"type\":\"uint256[3]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_executionDailyLimitExecutionMaxPerTxArray\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"_requestGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenFactory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_impl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCashThreshold\",\"type\":\"uint256\"}],\"name\":\"initializeInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"interestImplementation\",\"outputs\":[{\"internalType\":\"contractIInterestImplementation\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"invest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isBridgedTokenDeployAcknowledged\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isRegisteredAsNativeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isTokenRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxAvailablePerTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"maxPerTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"mediatorBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mediatorContractOnOtherSide\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"messageFixed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_interestImplementation\",\"type\":\"address\"}],\"name\":\"migrateTo_3_3_0\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"minCashThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"minPerTx\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgedToken\",\"type\":\"address\"}],\"name\":\"nativeTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC677\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"relayTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC677\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"relayTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC677\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"relayTokensAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"requestFailedMessageFix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeContract\",\"type\":\"address\"}],\"name\":\"setBridgeContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgedToken\",\"type\":\"address\"}],\"name\":\"setCustomTokenAddressPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"setDailyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"setExecutionDailyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerTx\",\"type\":\"uint256\"}],\"name\":\"setExecutionMaxPerTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxPerTx\",\"type\":\"uint256\"}],\"name\":\"setMaxPerTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mediatorContract\",\"type\":\"address\"}],\"name\":\"setMediatorContractOnOtherSide\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCashThreshold\",\"type\":\"uint256\"}],\"name\":\"setMinCashThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minPerTx\",\"type\":\"uint256\"}],\"name\":\"setMinPerTx\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"setRequestGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenFactory\",\"type\":\"address\"}],\"name\":\"setTokenFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenFactory\",\"outputs\":[{\"internalType\":\"contractTokenFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_day\",\"type\":\"uint256\"}],\"name\":\"totalExecutedPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_day\",\"type\":\"uint256\"}],\"name\":\"totalSpentPerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withinExecutionLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withinLimit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "GnosisDeposit",
}

// GnosisDeposit is an auto generated Go binding around an Ethereum contract.
type GnosisDeposit struct {
	abi abi.ABI
}

// NewGnosisDeposit creates a new instance of GnosisDeposit.
func NewGnosisDeposit() *GnosisDeposit {
	parsed, err := GnosisDepositMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &GnosisDeposit{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *GnosisDeposit) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(string _suffix) returns()
func (gnosisDeposit *GnosisDeposit) PackConstructor(_suffix string) []byte {
	enc, err := gnosisDeposit.abi.Pack("", _suffix)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBridgeContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (gnosisDeposit *GnosisDeposit) PackBridgeContract() []byte {
	enc, err := gnosisDeposit.abi.Pack("bridgeContract")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBridgeContract is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (gnosisDeposit *GnosisDeposit) UnpackBridgeContract(data []byte) (common.Address, error) {
	out, err := gnosisDeposit.abi.Unpack("bridgeContract", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackBridgedTokenAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2d70061f.
//
// Solidity: function bridgedTokenAddress(address _nativeToken) view returns(address)
func (gnosisDeposit *GnosisDeposit) PackBridgedTokenAddress(nativeToken common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("bridgedTokenAddress", nativeToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackBridgedTokenAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2d70061f.
//
// Solidity: function bridgedTokenAddress(address _nativeToken) view returns(address)
func (gnosisDeposit *GnosisDeposit) UnpackBridgedTokenAddress(data []byte) (common.Address, error) {
	out, err := gnosisDeposit.abi.Unpack("bridgedTokenAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackClaimTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x69ffa08a.
//
// Solidity: function claimTokens(address _token, address _to) returns()
func (gnosisDeposit *GnosisDeposit) PackClaimTokens(token common.Address, to common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("claimTokens", token, to)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackClaimTokensFromTokenContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x64696f97.
//
// Solidity: function claimTokensFromTokenContract(address _bridgedToken, address _token, address _to) returns()
func (gnosisDeposit *GnosisDeposit) PackClaimTokensFromTokenContract(bridgedToken common.Address, token common.Address, to common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("claimTokensFromTokenContract", bridgedToken, token, to)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDailyLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3f51415.
//
// Solidity: function dailyLimit(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackDailyLimit(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("dailyLimit", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDailyLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf3f51415.
//
// Solidity: function dailyLimit(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackDailyLimit(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("dailyLimit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDeployAndHandleBridgedTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2ae87cdd.
//
// Solidity: function deployAndHandleBridgedTokens(address _token, string _name, string _symbol, uint8 _decimals, address _recipient, uint256 _value) returns()
func (gnosisDeposit *GnosisDeposit) PackDeployAndHandleBridgedTokens(token common.Address, name string, symbol string, decimals uint8, recipient common.Address, value *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("deployAndHandleBridgedTokens", token, name, symbol, decimals, recipient, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDeployAndHandleBridgedTokensAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd522cfd7.
//
// Solidity: function deployAndHandleBridgedTokensAndCall(address _token, string _name, string _symbol, uint8 _decimals, address _recipient, uint256 _value, bytes _data) returns()
func (gnosisDeposit *GnosisDeposit) PackDeployAndHandleBridgedTokensAndCall(token common.Address, name string, symbol string, decimals uint8, recipient common.Address, value *big.Int, data []byte) []byte {
	enc, err := gnosisDeposit.abi.Pack("deployAndHandleBridgedTokensAndCall", token, name, symbol, decimals, recipient, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDisableInterest is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf50dace6.
//
// Solidity: function disableInterest(address _token) returns()
func (gnosisDeposit *GnosisDeposit) PackDisableInterest(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("disableInterest", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExecutionDailyLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x40f8dd86.
//
// Solidity: function executionDailyLimit(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackExecutionDailyLimit(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("executionDailyLimit", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecutionDailyLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x40f8dd86.
//
// Solidity: function executionDailyLimit(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackExecutionDailyLimit(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("executionDailyLimit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackExecutionMaxPerTx is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16ef1913.
//
// Solidity: function executionMaxPerTx(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackExecutionMaxPerTx(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("executionMaxPerTx", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackExecutionMaxPerTx is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x16ef1913.
//
// Solidity: function executionMaxPerTx(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackExecutionMaxPerTx(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("executionMaxPerTx", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackFixFailedMessage is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0950d515.
//
// Solidity: function fixFailedMessage(bytes32 _messageId) returns()
func (gnosisDeposit *GnosisDeposit) PackFixFailedMessage(messageId [32]byte) []byte {
	enc, err := gnosisDeposit.abi.Pack("fixFailedMessage", messageId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackFixMediatorBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd0342acd.
//
// Solidity: function fixMediatorBalance(address _token, address _receiver) returns()
func (gnosisDeposit *GnosisDeposit) PackFixMediatorBalance(token common.Address, receiver common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("fixMediatorBalance", token, receiver)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackGetBridgeInterfacesVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9cb7595a.
//
// Solidity: function getBridgeInterfacesVersion() pure returns(uint64 major, uint64 minor, uint64 patch)
func (gnosisDeposit *GnosisDeposit) PackGetBridgeInterfacesVersion() []byte {
	enc, err := gnosisDeposit.abi.Pack("getBridgeInterfacesVersion")
	if err != nil {
		panic(err)
	}
	return enc
}

// GetBridgeInterfacesVersionOutput serves as a container for the return parameters of contract
// method GetBridgeInterfacesVersion.
type GetBridgeInterfacesVersionOutput struct {
	Major uint64
	Minor uint64
	Patch uint64
}

// UnpackGetBridgeInterfacesVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9cb7595a.
//
// Solidity: function getBridgeInterfacesVersion() pure returns(uint64 major, uint64 minor, uint64 patch)
func (gnosisDeposit *GnosisDeposit) UnpackGetBridgeInterfacesVersion(data []byte) (GetBridgeInterfacesVersionOutput, error) {
	out, err := gnosisDeposit.abi.Unpack("getBridgeInterfacesVersion", data)
	outstruct := new(GetBridgeInterfacesVersionOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Major = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Minor = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Patch = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	return *outstruct, err

}

// PackGetBridgeMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x437764df.
//
// Solidity: function getBridgeMode() pure returns(bytes4 _data)
func (gnosisDeposit *GnosisDeposit) PackGetBridgeMode() []byte {
	enc, err := gnosisDeposit.abi.Pack("getBridgeMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetBridgeMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x437764df.
//
// Solidity: function getBridgeMode() pure returns(bytes4 _data)
func (gnosisDeposit *GnosisDeposit) UnpackGetBridgeMode(data []byte) ([4]byte, error) {
	out, err := gnosisDeposit.abi.Unpack("getBridgeMode", data)
	if err != nil {
		return *new([4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	return out0, err
}

// PackGetCurrentDay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackGetCurrentDay() []byte {
	enc, err := gnosisDeposit.abi.Pack("getCurrentDay")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGetCurrentDay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3e6968b6.
//
// Solidity: function getCurrentDay() view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackGetCurrentDay(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("getCurrentDay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackHandleBridgedTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x125e4cfb.
//
// Solidity: function handleBridgedTokens(address _token, address _recipient, uint256 _value) returns()
func (gnosisDeposit *GnosisDeposit) PackHandleBridgedTokens(token common.Address, recipient common.Address, value *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("handleBridgedTokens", token, recipient, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHandleBridgedTokensAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc5345761.
//
// Solidity: function handleBridgedTokensAndCall(address _token, address _recipient, uint256 _value, bytes _data) returns()
func (gnosisDeposit *GnosisDeposit) PackHandleBridgedTokensAndCall(token common.Address, recipient common.Address, value *big.Int, data []byte) []byte {
	enc, err := gnosisDeposit.abi.Pack("handleBridgedTokensAndCall", token, recipient, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHandleNativeTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x272255bb.
//
// Solidity: function handleNativeTokens(address _token, address _recipient, uint256 _value) returns()
func (gnosisDeposit *GnosisDeposit) PackHandleNativeTokens(token common.Address, recipient common.Address, value *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("handleNativeTokens", token, recipient, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackHandleNativeTokensAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x867f7a4d.
//
// Solidity: function handleNativeTokensAndCall(address _token, address _recipient, uint256 _value, bytes _data) returns()
func (gnosisDeposit *GnosisDeposit) PackHandleNativeTokensAndCall(token common.Address, recipient common.Address, value *big.Int, data []byte) []byte {
	enc, err := gnosisDeposit.abi.Pack("handleNativeTokensAndCall", token, recipient, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2c3500a6.
//
// Solidity: function initialize(address _bridgeContract, address _mediatorContract, uint256[3] _dailyLimitMaxPerTxMinPerTxArray, uint256[2] _executionDailyLimitExecutionMaxPerTxArray, uint256 _requestGasLimit, address _owner, address _tokenFactory) returns(bool)
func (gnosisDeposit *GnosisDeposit) PackInitialize(bridgeContract common.Address, mediatorContract common.Address, dailyLimitMaxPerTxMinPerTxArray [3]*big.Int, executionDailyLimitExecutionMaxPerTxArray [2]*big.Int, requestGasLimit *big.Int, owner common.Address, tokenFactory common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("initialize", bridgeContract, mediatorContract, dailyLimitMaxPerTxMinPerTxArray, executionDailyLimitExecutionMaxPerTxArray, requestGasLimit, owner, tokenFactory)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackInitialize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2c3500a6.
//
// Solidity: function initialize(address _bridgeContract, address _mediatorContract, uint256[3] _dailyLimitMaxPerTxMinPerTxArray, uint256[2] _executionDailyLimitExecutionMaxPerTxArray, uint256 _requestGasLimit, address _owner, address _tokenFactory) returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackInitialize(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("initialize", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackInitializeInterest is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab4f5dc5.
//
// Solidity: function initializeInterest(address _token, address _impl, uint256 _minCashThreshold) returns()
func (gnosisDeposit *GnosisDeposit) PackInitializeInterest(token common.Address, impl common.Address, minCashThreshold *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("initializeInterest", token, impl, minCashThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInterestImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85df73bd.
//
// Solidity: function interestImplementation(address _token) view returns(address)
func (gnosisDeposit *GnosisDeposit) PackInterestImplementation(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("interestImplementation", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackInterestImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85df73bd.
//
// Solidity: function interestImplementation(address _token) view returns(address)
func (gnosisDeposit *GnosisDeposit) UnpackInterestImplementation(data []byte) (common.Address, error) {
	out, err := gnosisDeposit.abi.Unpack("interestImplementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInvest is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x03f9c793.
//
// Solidity: function invest(address _token) returns()
func (gnosisDeposit *GnosisDeposit) PackInvest(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("invest", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackIsBridgedTokenDeployAcknowledged is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xae813e9f.
//
// Solidity: function isBridgedTokenDeployAcknowledged(address _token) view returns(bool)
func (gnosisDeposit *GnosisDeposit) PackIsBridgedTokenDeployAcknowledged(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("isBridgedTokenDeployAcknowledged", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsBridgedTokenDeployAcknowledged is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xae813e9f.
//
// Solidity: function isBridgedTokenDeployAcknowledged(address _token) view returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackIsBridgedTokenDeployAcknowledged(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("isBridgedTokenDeployAcknowledged", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsInitialized is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (gnosisDeposit *GnosisDeposit) PackIsInitialized() []byte {
	enc, err := gnosisDeposit.abi.Pack("isInitialized")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsInitialized is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackIsInitialized(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("isInitialized", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsRegisteredAsNativeToken is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2173d43.
//
// Solidity: function isRegisteredAsNativeToken(address _token) view returns(bool)
func (gnosisDeposit *GnosisDeposit) PackIsRegisteredAsNativeToken(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("isRegisteredAsNativeToken", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsRegisteredAsNativeToken is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc2173d43.
//
// Solidity: function isRegisteredAsNativeToken(address _token) view returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackIsRegisteredAsNativeToken(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("isRegisteredAsNativeToken", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackIsTokenRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x26aa101f.
//
// Solidity: function isTokenRegistered(address _token) view returns(bool)
func (gnosisDeposit *GnosisDeposit) PackIsTokenRegistered(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("isTokenRegistered", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackIsTokenRegistered is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x26aa101f.
//
// Solidity: function isTokenRegistered(address _token) view returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackIsTokenRegistered(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("isTokenRegistered", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMaxAvailablePerTx is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7610722f.
//
// Solidity: function maxAvailablePerTx(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackMaxAvailablePerTx(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("maxAvailablePerTx", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxAvailablePerTx is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7610722f.
//
// Solidity: function maxAvailablePerTx(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackMaxAvailablePerTx(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("maxAvailablePerTx", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMaxPerTx is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x032f693f.
//
// Solidity: function maxPerTx(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackMaxPerTx(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("maxPerTx", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMaxPerTx is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x032f693f.
//
// Solidity: function maxPerTx(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackMaxPerTx(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("maxPerTx", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMediatorBalance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x194153d3.
//
// Solidity: function mediatorBalance(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackMediatorBalance(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("mediatorBalance", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMediatorBalance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x194153d3.
//
// Solidity: function mediatorBalance(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackMediatorBalance(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("mediatorBalance", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMediatorContractOnOtherSide is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x871c0760.
//
// Solidity: function mediatorContractOnOtherSide() view returns(address)
func (gnosisDeposit *GnosisDeposit) PackMediatorContractOnOtherSide() []byte {
	enc, err := gnosisDeposit.abi.Pack("mediatorContractOnOtherSide")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMediatorContractOnOtherSide is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x871c0760.
//
// Solidity: function mediatorContractOnOtherSide() view returns(address)
func (gnosisDeposit *GnosisDeposit) UnpackMediatorContractOnOtherSide(data []byte) (common.Address, error) {
	out, err := gnosisDeposit.abi.Unpack("mediatorContractOnOtherSide", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackMessageFixed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x59339982.
//
// Solidity: function messageFixed(bytes32 _messageId) view returns(bool)
func (gnosisDeposit *GnosisDeposit) PackMessageFixed(messageId [32]byte) []byte {
	enc, err := gnosisDeposit.abi.Pack("messageFixed", messageId)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMessageFixed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x59339982.
//
// Solidity: function messageFixed(bytes32 _messageId) view returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackMessageFixed(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("messageFixed", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackMigrateTo330 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd814b1d7.
//
// Solidity: function migrateTo_3_3_0(address _tokenFactory, address _interestImplementation) returns()
func (gnosisDeposit *GnosisDeposit) PackMigrateTo330(tokenFactory common.Address, interestImplementation common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("migrateTo_3_3_0", tokenFactory, interestImplementation)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMinCashThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5726ff30.
//
// Solidity: function minCashThreshold(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackMinCashThreshold(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("minCashThreshold", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMinCashThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5726ff30.
//
// Solidity: function minCashThreshold(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackMinCashThreshold(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("minCashThreshold", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackMinPerTx is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4b1c243.
//
// Solidity: function minPerTx(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackMinPerTx(token common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("minPerTx", token)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMinPerTx is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa4b1c243.
//
// Solidity: function minPerTx(address _token) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackMinPerTx(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("minPerTx", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackNativeTokenAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x61c04f84.
//
// Solidity: function nativeTokenAddress(address _bridgedToken) view returns(address)
func (gnosisDeposit *GnosisDeposit) PackNativeTokenAddress(bridgedToken common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("nativeTokenAddress", bridgedToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNativeTokenAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x61c04f84.
//
// Solidity: function nativeTokenAddress(address _bridgedToken) view returns(address)
func (gnosisDeposit *GnosisDeposit) UnpackNativeTokenAddress(data []byte) (common.Address, error) {
	out, err := gnosisDeposit.abi.Unpack("nativeTokenAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackOnTokenTransfer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _from, uint256 _value, bytes _data) returns(bool)
func (gnosisDeposit *GnosisDeposit) PackOnTokenTransfer(from common.Address, value *big.Int, data []byte) []byte {
	enc, err := gnosisDeposit.abi.Pack("onTokenTransfer", from, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOnTokenTransfer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address _from, uint256 _value, bytes _data) returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackOnTokenTransfer(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("onTokenTransfer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (gnosisDeposit *GnosisDeposit) PackOwner() []byte {
	enc, err := gnosisDeposit.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (gnosisDeposit *GnosisDeposit) UnpackOwner(data []byte) (common.Address, error) {
	out, err := gnosisDeposit.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackRelayTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01e4f53a.
//
// Solidity: function relayTokens(address token, uint256 _value) returns()
func (gnosisDeposit *GnosisDeposit) PackRelayTokens(token common.Address, value *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("relayTokens", token, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelayTokens0 is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad58bdd1.
//
// Solidity: function relayTokens(address token, address _receiver, uint256 _value) returns()
func (gnosisDeposit *GnosisDeposit) PackRelayTokens0(token common.Address, receiver common.Address, value *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("relayTokens0", token, receiver, value)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelayTokensAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7405481.
//
// Solidity: function relayTokensAndCall(address token, address _receiver, uint256 _value, bytes _data) returns()
func (gnosisDeposit *GnosisDeposit) PackRelayTokensAndCall(token common.Address, receiver common.Address, value *big.Int, data []byte) []byte {
	enc, err := gnosisDeposit.abi.Pack("relayTokensAndCall", token, receiver, value, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRequestFailedMessageFix is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a4a4395.
//
// Solidity: function requestFailedMessageFix(bytes32 _messageId) returns()
func (gnosisDeposit *GnosisDeposit) PackRequestFailedMessageFix(messageId [32]byte) []byte {
	enc, err := gnosisDeposit.abi.Pack("requestFailedMessageFix", messageId)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRequestGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbe3b625b.
//
// Solidity: function requestGasLimit() view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackRequestGasLimit() []byte {
	enc, err := gnosisDeposit.abi.Pack("requestGasLimit")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRequestGasLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbe3b625b.
//
// Solidity: function requestGasLimit() view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackRequestGasLimit(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("requestGasLimit", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackSetBridgeContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b26cf66.
//
// Solidity: function setBridgeContract(address _bridgeContract) returns()
func (gnosisDeposit *GnosisDeposit) PackSetBridgeContract(bridgeContract common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("setBridgeContract", bridgeContract)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetCustomTokenAddressPair is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0b71a4a7.
//
// Solidity: function setCustomTokenAddressPair(address _nativeToken, address _bridgedToken) returns()
func (gnosisDeposit *GnosisDeposit) PackSetCustomTokenAddressPair(nativeToken common.Address, bridgedToken common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("setCustomTokenAddressPair", nativeToken, bridgedToken)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetDailyLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2803212f.
//
// Solidity: function setDailyLimit(address _token, uint256 _dailyLimit) returns()
func (gnosisDeposit *GnosisDeposit) PackSetDailyLimit(token common.Address, dailyLimit *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("setDailyLimit", token, dailyLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetExecutionDailyLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7837cf91.
//
// Solidity: function setExecutionDailyLimit(address _token, uint256 _dailyLimit) returns()
func (gnosisDeposit *GnosisDeposit) PackSetExecutionDailyLimit(token common.Address, dailyLimit *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("setExecutionDailyLimit", token, dailyLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetExecutionMaxPerTx is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01fcc1d3.
//
// Solidity: function setExecutionMaxPerTx(address _token, uint256 _maxPerTx) returns()
func (gnosisDeposit *GnosisDeposit) PackSetExecutionMaxPerTx(token common.Address, maxPerTx *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("setExecutionMaxPerTx", token, maxPerTx)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMaxPerTx is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb6fff8c.
//
// Solidity: function setMaxPerTx(address _token, uint256 _maxPerTx) returns()
func (gnosisDeposit *GnosisDeposit) PackSetMaxPerTx(token common.Address, maxPerTx *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("setMaxPerTx", token, maxPerTx)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMediatorContractOnOtherSide is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6e5d6bea.
//
// Solidity: function setMediatorContractOnOtherSide(address _mediatorContract) returns()
func (gnosisDeposit *GnosisDeposit) PackSetMediatorContractOnOtherSide(mediatorContract common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("setMediatorContractOnOtherSide", mediatorContract)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinCashThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa4b4b233.
//
// Solidity: function setMinCashThreshold(address _token, uint256 _minCashThreshold) returns()
func (gnosisDeposit *GnosisDeposit) PackSetMinCashThreshold(token common.Address, minCashThreshold *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("setMinCashThreshold", token, minCashThreshold)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetMinPerTx is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec47de2a.
//
// Solidity: function setMinPerTx(address _token, uint256 _minPerTx) returns()
func (gnosisDeposit *GnosisDeposit) PackSetMinPerTx(token common.Address, minPerTx *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("setMinPerTx", token, minPerTx)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetRequestGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf3b83791.
//
// Solidity: function setRequestGasLimit(uint256 _gasLimit) returns()
func (gnosisDeposit *GnosisDeposit) PackSetRequestGasLimit(gasLimit *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("setRequestGasLimit", gasLimit)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetTokenFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2f73a9f8.
//
// Solidity: function setTokenFactory(address _tokenFactory) returns()
func (gnosisDeposit *GnosisDeposit) PackSetTokenFactory(tokenFactory common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("setTokenFactory", tokenFactory)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTokenFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (gnosisDeposit *GnosisDeposit) PackTokenFactory() []byte {
	enc, err := gnosisDeposit.abi.Pack("tokenFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe77772fe.
//
// Solidity: function tokenFactory() view returns(address)
func (gnosisDeposit *GnosisDeposit) UnpackTokenFactory(data []byte) (common.Address, error) {
	out, err := gnosisDeposit.abi.Unpack("tokenFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackTotalExecutedPerDay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2c54fe8.
//
// Solidity: function totalExecutedPerDay(address _token, uint256 _day) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackTotalExecutedPerDay(token common.Address, day *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("totalExecutedPerDay", token, day)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalExecutedPerDay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2c54fe8.
//
// Solidity: function totalExecutedPerDay(address _token, uint256 _day) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackTotalExecutedPerDay(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("totalExecutedPerDay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTotalSpentPerDay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xab3a25d9.
//
// Solidity: function totalSpentPerDay(address _token, uint256 _day) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) PackTotalSpentPerDay(token common.Address, day *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("totalSpentPerDay", token, day)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTotalSpentPerDay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xab3a25d9.
//
// Solidity: function totalSpentPerDay(address _token, uint256 _day) view returns(uint256)
func (gnosisDeposit *GnosisDeposit) UnpackTotalSpentPerDay(data []byte) (*big.Int, error) {
	out, err := gnosisDeposit.abi.Unpack("totalSpentPerDay", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (gnosisDeposit *GnosisDeposit) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := gnosisDeposit.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackWithinExecutionLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a50bc87.
//
// Solidity: function withinExecutionLimit(address _token, uint256 _amount) view returns(bool)
func (gnosisDeposit *GnosisDeposit) PackWithinExecutionLimit(token common.Address, amount *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("withinExecutionLimit", token, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackWithinExecutionLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a50bc87.
//
// Solidity: function withinExecutionLimit(address _token, uint256 _amount) view returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackWithinExecutionLimit(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("withinExecutionLimit", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackWithinLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x10775238.
//
// Solidity: function withinLimit(address _token, uint256 _amount) view returns(bool)
func (gnosisDeposit *GnosisDeposit) PackWithinLimit(token common.Address, amount *big.Int) []byte {
	enc, err := gnosisDeposit.abi.Pack("withinLimit", token, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackWithinLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x10775238.
//
// Solidity: function withinLimit(address _token, uint256 _amount) view returns(bool)
func (gnosisDeposit *GnosisDeposit) UnpackWithinLimit(data []byte) (bool, error) {
	out, err := gnosisDeposit.abi.Unpack("withinLimit", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// GnosisDepositDailyLimitChanged represents a DailyLimitChanged event raised by the GnosisDeposit contract.
type GnosisDepositDailyLimitChanged struct {
	Token    common.Address
	NewLimit *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const GnosisDepositDailyLimitChangedEventName = "DailyLimitChanged"

// ContractEventName returns the user-defined event name.
func (GnosisDepositDailyLimitChanged) ContractEventName() string {
	return GnosisDepositDailyLimitChangedEventName
}

// UnpackDailyLimitChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DailyLimitChanged(address indexed token, uint256 newLimit)
func (gnosisDeposit *GnosisDeposit) UnpackDailyLimitChangedEvent(log *types.Log) (*GnosisDepositDailyLimitChanged, error) {
	event := "DailyLimitChanged"
	if log.Topics[0] != gnosisDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisDepositDailyLimitChanged)
	if len(log.Data) > 0 {
		if err := gnosisDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisDeposit.abi.Events[event].Inputs {
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

// GnosisDepositExecutionDailyLimitChanged represents a ExecutionDailyLimitChanged event raised by the GnosisDeposit contract.
type GnosisDepositExecutionDailyLimitChanged struct {
	Token    common.Address
	NewLimit *big.Int
	Raw      *types.Log // Blockchain specific contextual infos
}

const GnosisDepositExecutionDailyLimitChangedEventName = "ExecutionDailyLimitChanged"

// ContractEventName returns the user-defined event name.
func (GnosisDepositExecutionDailyLimitChanged) ContractEventName() string {
	return GnosisDepositExecutionDailyLimitChangedEventName
}

// UnpackExecutionDailyLimitChangedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ExecutionDailyLimitChanged(address indexed token, uint256 newLimit)
func (gnosisDeposit *GnosisDeposit) UnpackExecutionDailyLimitChangedEvent(log *types.Log) (*GnosisDepositExecutionDailyLimitChanged, error) {
	event := "ExecutionDailyLimitChanged"
	if log.Topics[0] != gnosisDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisDepositExecutionDailyLimitChanged)
	if len(log.Data) > 0 {
		if err := gnosisDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisDeposit.abi.Events[event].Inputs {
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

// GnosisDepositFailedMessageFixed represents a FailedMessageFixed event raised by the GnosisDeposit contract.
type GnosisDepositFailedMessageFixed struct {
	MessageId [32]byte
	Token     common.Address
	Recipient common.Address
	Value     *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const GnosisDepositFailedMessageFixedEventName = "FailedMessageFixed"

// ContractEventName returns the user-defined event name.
func (GnosisDepositFailedMessageFixed) ContractEventName() string {
	return GnosisDepositFailedMessageFixedEventName
}

// UnpackFailedMessageFixedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FailedMessageFixed(bytes32 indexed messageId, address token, address recipient, uint256 value)
func (gnosisDeposit *GnosisDeposit) UnpackFailedMessageFixedEvent(log *types.Log) (*GnosisDepositFailedMessageFixed, error) {
	event := "FailedMessageFixed"
	if log.Topics[0] != gnosisDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisDepositFailedMessageFixed)
	if len(log.Data) > 0 {
		if err := gnosisDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisDeposit.abi.Events[event].Inputs {
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

// GnosisDepositNewTokenRegistered represents a NewTokenRegistered event raised by the GnosisDeposit contract.
type GnosisDepositNewTokenRegistered struct {
	NativeToken  common.Address
	BridgedToken common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const GnosisDepositNewTokenRegisteredEventName = "NewTokenRegistered"

// ContractEventName returns the user-defined event name.
func (GnosisDepositNewTokenRegistered) ContractEventName() string {
	return GnosisDepositNewTokenRegisteredEventName
}

// UnpackNewTokenRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NewTokenRegistered(address indexed nativeToken, address indexed bridgedToken)
func (gnosisDeposit *GnosisDeposit) UnpackNewTokenRegisteredEvent(log *types.Log) (*GnosisDepositNewTokenRegistered, error) {
	event := "NewTokenRegistered"
	if log.Topics[0] != gnosisDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisDepositNewTokenRegistered)
	if len(log.Data) > 0 {
		if err := gnosisDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisDeposit.abi.Events[event].Inputs {
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

// GnosisDepositOwnershipTransferred represents a OwnershipTransferred event raised by the GnosisDeposit contract.
type GnosisDepositOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const GnosisDepositOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (GnosisDepositOwnershipTransferred) ContractEventName() string {
	return GnosisDepositOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address previousOwner, address newOwner)
func (gnosisDeposit *GnosisDeposit) UnpackOwnershipTransferredEvent(log *types.Log) (*GnosisDepositOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != gnosisDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisDepositOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := gnosisDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisDeposit.abi.Events[event].Inputs {
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

// GnosisDepositTokensBridged represents a TokensBridged event raised by the GnosisDeposit contract.
type GnosisDepositTokensBridged struct {
	Token     common.Address
	Recipient common.Address
	Value     *big.Int
	MessageId [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const GnosisDepositTokensBridgedEventName = "TokensBridged"

// ContractEventName returns the user-defined event name.
func (GnosisDepositTokensBridged) ContractEventName() string {
	return GnosisDepositTokensBridgedEventName
}

// UnpackTokensBridgedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokensBridged(address indexed token, address indexed recipient, uint256 value, bytes32 indexed messageId)
func (gnosisDeposit *GnosisDeposit) UnpackTokensBridgedEvent(log *types.Log) (*GnosisDepositTokensBridged, error) {
	event := "TokensBridged"
	if log.Topics[0] != gnosisDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisDepositTokensBridged)
	if len(log.Data) > 0 {
		if err := gnosisDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisDeposit.abi.Events[event].Inputs {
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

// GnosisDepositTokensBridgingInitiated represents a TokensBridgingInitiated event raised by the GnosisDeposit contract.
type GnosisDepositTokensBridgingInitiated struct {
	Token     common.Address
	Sender    common.Address
	Value     *big.Int
	MessageId [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const GnosisDepositTokensBridgingInitiatedEventName = "TokensBridgingInitiated"

// ContractEventName returns the user-defined event name.
func (GnosisDepositTokensBridgingInitiated) ContractEventName() string {
	return GnosisDepositTokensBridgingInitiatedEventName
}

// UnpackTokensBridgingInitiatedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TokensBridgingInitiated(address indexed token, address indexed sender, uint256 value, bytes32 indexed messageId)
func (gnosisDeposit *GnosisDeposit) UnpackTokensBridgingInitiatedEvent(log *types.Log) (*GnosisDepositTokensBridgingInitiated, error) {
	event := "TokensBridgingInitiated"
	if log.Topics[0] != gnosisDeposit.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisDepositTokensBridgingInitiated)
	if len(log.Data) > 0 {
		if err := gnosisDeposit.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisDeposit.abi.Events[event].Inputs {
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

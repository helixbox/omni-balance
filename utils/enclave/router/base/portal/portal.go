// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package base_portal

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

// TypesOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockhash          [32]byte
}

// TypesWithdrawalTransaction is an auto generated low-level Go binding around an user-defined struct.
type TypesWithdrawalTransaction struct {
	Nonce    *big.Int
	Sender   common.Address
	Target   common.Address
	Value    *big.Int
	GasLimit *big.Int
	Data     []byte
}

// BasePortalMetaData contains all meta data concerning the BasePortal contract.
var BasePortalMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proofMaturityDelaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_disputeGameFinalityDelaySeconds\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Blacklisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContentLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasEstimation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataRemainder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDisputeGame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGameType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidHeader\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LargeCalldata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LegacyGame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonReentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposalNotValidated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SmallGasLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unproven\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIDisputeGame\",\"name\":\"disputeGame\",\"type\":\"address\"}],\"name\":\"DisputeGameBlacklisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"GameType\",\"name\":\"newGameType\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"Timestamp\",\"name\":\"updatedAt\",\"type\":\"uint64\"}],\"name\":\"RespectedGameTypeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"opaqueData\",\"type\":\"bytes\"}],\"name\":\"TransactionDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"WithdrawalProven\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"withdrawalHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proofSubmitter\",\"type\":\"address\"}],\"name\":\"WithdrawalProvenExtension1\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIDisputeGame\",\"name\":\"_disputeGame\",\"type\":\"address\"}],\"name\":\"blacklistDisputeGame\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_withdrawalHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_proofSubmitter\",\"type\":\"address\"}],\"name\":\"checkWithdrawal\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"_isCreation\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"depositTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDisputeGame\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"disputeGameBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeGameFactory\",\"outputs\":[{\"internalType\":\"contractIDisputeGameFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeGameFinalityDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.WithdrawalTransaction\",\"name\":\"_tx\",\"type\":\"tuple\"}],\"name\":\"finalizeWithdrawalTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.WithdrawalTransaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_proofSubmitter\",\"type\":\"address\"}],\"name\":\"finalizeWithdrawalTransactionExternalProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"finalizedWithdrawals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDisputeGameFactory\",\"name\":\"_disputeGameFactory\",\"type\":\"address\"},{\"internalType\":\"contractISystemConfig\",\"name\":\"_systemConfig\",\"type\":\"address\"},{\"internalType\":\"contractISuperchainConfig\",\"name\":\"_superchainConfig\",\"type\":\"address\"},{\"internalType\":\"GameType\",\"name\":\"_initialRespectedGameType\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_byteCount\",\"type\":\"uint64\"}],\"name\":\"minimumGasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_withdrawalHash\",\"type\":\"bytes32\"}],\"name\":\"numProofSubmitters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"params\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"prevBaseFee\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"prevBoughtGas\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"prevBlockNum\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proofMaturityDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proofSubmitters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.WithdrawalTransaction\",\"name\":\"_tx\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_disputeGameIndex\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"latestBlockhash\",\"type\":\"bytes32\"}],\"internalType\":\"structTypes.OutputRootProof\",\"name\":\"_outputRootProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_withdrawalProof\",\"type\":\"bytes[]\"}],\"name\":\"proveWithdrawalTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"provenWithdrawals\",\"outputs\":[{\"internalType\":\"contractIDisputeGame\",\"name\":\"disputeGameProxy\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"respectedGameType\",\"outputs\":[{\"internalType\":\"GameType\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"respectedGameTypeUpdatedAt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"_gameType\",\"type\":\"uint32\"}],\"name\":\"setRespectedGameType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"superchainConfig\",\"outputs\":[{\"internalType\":\"contractISuperchainConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemConfig\",\"outputs\":[{\"internalType\":\"contractISystemConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	ID:  "BasePortal",
}

// BasePortal is an auto generated Go binding around an Ethereum contract.
type BasePortal struct {
	abi abi.ABI
}

// NewBasePortal creates a new instance of BasePortal.
func NewBasePortal() *BasePortal {
	parsed, err := BasePortalMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &BasePortal{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *BasePortal) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(uint256 _proofMaturityDelaySeconds, uint256 _disputeGameFinalityDelaySeconds) returns()
func (basePortal *BasePortal) PackConstructor(_proofMaturityDelaySeconds *big.Int, _disputeGameFinalityDelaySeconds *big.Int) []byte {
	enc, err := basePortal.abi.Pack("", _proofMaturityDelaySeconds, _disputeGameFinalityDelaySeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackBlacklistDisputeGame is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7d6be8dc.
//
// Solidity: function blacklistDisputeGame(address _disputeGame) returns()
func (basePortal *BasePortal) PackBlacklistDisputeGame(disputeGame common.Address) []byte {
	enc, err := basePortal.abi.Pack("blacklistDisputeGame", disputeGame)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCheckWithdrawal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x71c1566e.
//
// Solidity: function checkWithdrawal(bytes32 _withdrawalHash, address _proofSubmitter) view returns()
func (basePortal *BasePortal) PackCheckWithdrawal(withdrawalHash [32]byte, proofSubmitter common.Address) []byte {
	enc, err := basePortal.abi.Pack("checkWithdrawal", withdrawalHash, proofSubmitter)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDepositTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (basePortal *BasePortal) PackDepositTransaction(to common.Address, value *big.Int, gasLimit uint64, isCreation bool, data []byte) []byte {
	enc, err := basePortal.abi.Pack("depositTransaction", to, value, gasLimit, isCreation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackDisputeGameBlacklist is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x45884d32.
//
// Solidity: function disputeGameBlacklist(address ) view returns(bool)
func (basePortal *BasePortal) PackDisputeGameBlacklist(arg0 common.Address) []byte {
	enc, err := basePortal.abi.Pack("disputeGameBlacklist", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDisputeGameBlacklist is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x45884d32.
//
// Solidity: function disputeGameBlacklist(address ) view returns(bool)
func (basePortal *BasePortal) UnpackDisputeGameBlacklist(data []byte) (bool, error) {
	out, err := basePortal.abi.Unpack("disputeGameBlacklist", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackDisputeGameFactory is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (basePortal *BasePortal) PackDisputeGameFactory() []byte {
	enc, err := basePortal.abi.Pack("disputeGameFactory")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDisputeGameFactory is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (basePortal *BasePortal) UnpackDisputeGameFactory(data []byte) (common.Address, error) {
	out, err := basePortal.abi.Unpack("disputeGameFactory", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackDisputeGameFinalityDelaySeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x952b2797.
//
// Solidity: function disputeGameFinalityDelaySeconds() view returns(uint256)
func (basePortal *BasePortal) PackDisputeGameFinalityDelaySeconds() []byte {
	enc, err := basePortal.abi.Pack("disputeGameFinalityDelaySeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDisputeGameFinalityDelaySeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x952b2797.
//
// Solidity: function disputeGameFinalityDelaySeconds() view returns(uint256)
func (basePortal *BasePortal) UnpackDisputeGameFinalityDelaySeconds(data []byte) (*big.Int, error) {
	out, err := basePortal.abi.Unpack("disputeGameFinalityDelaySeconds", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackDonateETH is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (basePortal *BasePortal) PackDonateETH() []byte {
	enc, err := basePortal.abi.Pack("donateETH")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackFinalizeWithdrawalTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (basePortal *BasePortal) PackFinalizeWithdrawalTransaction(tx TypesWithdrawalTransaction) []byte {
	enc, err := basePortal.abi.Pack("finalizeWithdrawalTransaction", tx)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackFinalizeWithdrawalTransactionExternalProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x43ca1c50.
//
// Solidity: function finalizeWithdrawalTransactionExternalProof((uint256,address,address,uint256,uint256,bytes) _tx, address _proofSubmitter) returns()
func (basePortal *BasePortal) PackFinalizeWithdrawalTransactionExternalProof(tx TypesWithdrawalTransaction, proofSubmitter common.Address) []byte {
	enc, err := basePortal.abi.Pack("finalizeWithdrawalTransactionExternalProof", tx, proofSubmitter)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackFinalizedWithdrawals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (basePortal *BasePortal) PackFinalizedWithdrawals(arg0 [32]byte) []byte {
	enc, err := basePortal.abi.Pack("finalizedWithdrawals", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFinalizedWithdrawals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (basePortal *BasePortal) UnpackFinalizedWithdrawals(data []byte) (bool, error) {
	out, err := basePortal.abi.Unpack("finalizedWithdrawals", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackGuardian is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (basePortal *BasePortal) PackGuardian() []byte {
	enc, err := basePortal.abi.Pack("guardian")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackGuardian is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (basePortal *BasePortal) UnpackGuardian(data []byte) (common.Address, error) {
	out, err := basePortal.abi.Unpack("guardian", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e819e54.
//
// Solidity: function initialize(address _disputeGameFactory, address _systemConfig, address _superchainConfig, uint32 _initialRespectedGameType) returns()
func (basePortal *BasePortal) PackInitialize(disputeGameFactory common.Address, systemConfig common.Address, superchainConfig common.Address, initialRespectedGameType uint32) []byte {
	enc, err := basePortal.abi.Pack("initialize", disputeGameFactory, systemConfig, superchainConfig, initialRespectedGameType)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackL2Sender is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (basePortal *BasePortal) PackL2Sender() []byte {
	enc, err := basePortal.abi.Pack("l2Sender")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackL2Sender is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (basePortal *BasePortal) UnpackL2Sender(data []byte) (common.Address, error) {
	out, err := basePortal.abi.Unpack("l2Sender", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackMinimumGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (basePortal *BasePortal) PackMinimumGasLimit(byteCount uint64) []byte {
	enc, err := basePortal.abi.Pack("minimumGasLimit", byteCount)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackMinimumGasLimit is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (basePortal *BasePortal) UnpackMinimumGasLimit(data []byte) (uint64, error) {
	out, err := basePortal.abi.Unpack("minimumGasLimit", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackNumProofSubmitters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x513747ab.
//
// Solidity: function numProofSubmitters(bytes32 _withdrawalHash) view returns(uint256)
func (basePortal *BasePortal) PackNumProofSubmitters(withdrawalHash [32]byte) []byte {
	enc, err := basePortal.abi.Pack("numProofSubmitters", withdrawalHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackNumProofSubmitters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x513747ab.
//
// Solidity: function numProofSubmitters(bytes32 _withdrawalHash) view returns(uint256)
func (basePortal *BasePortal) UnpackNumProofSubmitters(data []byte) (*big.Int, error) {
	out, err := basePortal.abi.Unpack("numProofSubmitters", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (basePortal *BasePortal) PackParams() []byte {
	enc, err := basePortal.abi.Pack("params")
	if err != nil {
		panic(err)
	}
	return enc
}

// ParamsOutput serves as a container for the return parameters of contract
// method Params.
type ParamsOutput struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}

// UnpackParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (basePortal *BasePortal) UnpackParams(data []byte) (ParamsOutput, error) {
	out, err := basePortal.abi.Unpack("params", data)
	outstruct := new(ParamsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.PrevBaseFee = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.PrevBoughtGas = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PrevBlockNum = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	return *outstruct, err

}

// PackPaused is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (basePortal *BasePortal) PackPaused() []byte {
	enc, err := basePortal.abi.Pack("paused")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackPaused is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (basePortal *BasePortal) UnpackPaused(data []byte) (bool, error) {
	out, err := basePortal.abi.Unpack("paused", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, err
}

// PackProofMaturityDelaySeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf653a5c.
//
// Solidity: function proofMaturityDelaySeconds() view returns(uint256)
func (basePortal *BasePortal) PackProofMaturityDelaySeconds() []byte {
	enc, err := basePortal.abi.Pack("proofMaturityDelaySeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProofMaturityDelaySeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf653a5c.
//
// Solidity: function proofMaturityDelaySeconds() view returns(uint256)
func (basePortal *BasePortal) UnpackProofMaturityDelaySeconds(data []byte) (*big.Int, error) {
	out, err := basePortal.abi.Unpack("proofMaturityDelaySeconds", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, err
}

// PackProofSubmitters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3860f48.
//
// Solidity: function proofSubmitters(bytes32 , uint256 ) view returns(address)
func (basePortal *BasePortal) PackProofSubmitters(arg0 [32]byte, arg1 *big.Int) []byte {
	enc, err := basePortal.abi.Pack("proofSubmitters", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackProofSubmitters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3860f48.
//
// Solidity: function proofSubmitters(bytes32 , uint256 ) view returns(address)
func (basePortal *BasePortal) UnpackProofSubmitters(data []byte) (common.Address, error) {
	out, err := basePortal.abi.Unpack("proofSubmitters", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackProveWithdrawalTransaction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4870496f.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _disputeGameIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (basePortal *BasePortal) PackProveWithdrawalTransaction(tx TypesWithdrawalTransaction, disputeGameIndex *big.Int, outputRootProof TypesOutputRootProof, withdrawalProof [][]byte) []byte {
	enc, err := basePortal.abi.Pack("proveWithdrawalTransaction", tx, disputeGameIndex, outputRootProof, withdrawalProof)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackProvenWithdrawals is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb2c727e.
//
// Solidity: function provenWithdrawals(bytes32 , address ) view returns(address disputeGameProxy, uint64 timestamp)
func (basePortal *BasePortal) PackProvenWithdrawals(arg0 [32]byte, arg1 common.Address) []byte {
	enc, err := basePortal.abi.Pack("provenWithdrawals", arg0, arg1)
	if err != nil {
		panic(err)
	}
	return enc
}

// ProvenWithdrawalsOutput serves as a container for the return parameters of contract
// method ProvenWithdrawals.
type ProvenWithdrawalsOutput struct {
	DisputeGameProxy common.Address
	Timestamp        uint64
}

// UnpackProvenWithdrawals is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbb2c727e.
//
// Solidity: function provenWithdrawals(bytes32 , address ) view returns(address disputeGameProxy, uint64 timestamp)
func (basePortal *BasePortal) UnpackProvenWithdrawals(data []byte) (ProvenWithdrawalsOutput, error) {
	out, err := basePortal.abi.Unpack("provenWithdrawals", data)
	outstruct := new(ProvenWithdrawalsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.DisputeGameProxy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, err

}

// PackRespectedGameType is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c9f397c.
//
// Solidity: function respectedGameType() view returns(uint32)
func (basePortal *BasePortal) PackRespectedGameType() []byte {
	enc, err := basePortal.abi.Pack("respectedGameType")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRespectedGameType is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3c9f397c.
//
// Solidity: function respectedGameType() view returns(uint32)
func (basePortal *BasePortal) UnpackRespectedGameType(data []byte) (uint32, error) {
	out, err := basePortal.abi.Unpack("respectedGameType", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, err
}

// PackRespectedGameTypeUpdatedAt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4fd0434c.
//
// Solidity: function respectedGameTypeUpdatedAt() view returns(uint64)
func (basePortal *BasePortal) PackRespectedGameTypeUpdatedAt() []byte {
	enc, err := basePortal.abi.Pack("respectedGameTypeUpdatedAt")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackRespectedGameTypeUpdatedAt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4fd0434c.
//
// Solidity: function respectedGameTypeUpdatedAt() view returns(uint64)
func (basePortal *BasePortal) UnpackRespectedGameTypeUpdatedAt(data []byte) (uint64, error) {
	out, err := basePortal.abi.Unpack("respectedGameTypeUpdatedAt", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, err
}

// PackSetRespectedGameType is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fc48504.
//
// Solidity: function setRespectedGameType(uint32 _gameType) returns()
func (basePortal *BasePortal) PackSetRespectedGameType(gameType uint32) []byte {
	enc, err := basePortal.abi.Pack("setRespectedGameType", gameType)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSuperchainConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (basePortal *BasePortal) PackSuperchainConfig() []byte {
	enc, err := basePortal.abi.Pack("superchainConfig")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSuperchainConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (basePortal *BasePortal) UnpackSuperchainConfig(data []byte) (common.Address, error) {
	out, err := basePortal.abi.Unpack("superchainConfig", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackSystemConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (basePortal *BasePortal) PackSystemConfig() []byte {
	enc, err := basePortal.abi.Pack("systemConfig")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackSystemConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (basePortal *BasePortal) UnpackSystemConfig(data []byte) (common.Address, error) {
	out, err := basePortal.abi.Unpack("systemConfig", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackVersion is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (basePortal *BasePortal) PackVersion() []byte {
	enc, err := basePortal.abi.Pack("version")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackVersion is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (basePortal *BasePortal) UnpackVersion(data []byte) (string, error) {
	out, err := basePortal.abi.Unpack("version", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, err
}

// BasePortalDisputeGameBlacklisted represents a DisputeGameBlacklisted event raised by the BasePortal contract.
type BasePortalDisputeGameBlacklisted struct {
	DisputeGame common.Address
	Raw         *types.Log // Blockchain specific contextual infos
}

const BasePortalDisputeGameBlacklistedEventName = "DisputeGameBlacklisted"

// ContractEventName returns the user-defined event name.
func (BasePortalDisputeGameBlacklisted) ContractEventName() string {
	return BasePortalDisputeGameBlacklistedEventName
}

// UnpackDisputeGameBlacklistedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DisputeGameBlacklisted(address indexed disputeGame)
func (basePortal *BasePortal) UnpackDisputeGameBlacklistedEvent(log *types.Log) (*BasePortalDisputeGameBlacklisted, error) {
	event := "DisputeGameBlacklisted"
	if log.Topics[0] != basePortal.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BasePortalDisputeGameBlacklisted)
	if len(log.Data) > 0 {
		if err := basePortal.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range basePortal.abi.Events[event].Inputs {
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

// BasePortalInitialized represents a Initialized event raised by the BasePortal contract.
type BasePortalInitialized struct {
	Version uint8
	Raw     *types.Log // Blockchain specific contextual infos
}

const BasePortalInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (BasePortalInitialized) ContractEventName() string {
	return BasePortalInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint8 version)
func (basePortal *BasePortal) UnpackInitializedEvent(log *types.Log) (*BasePortalInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != basePortal.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BasePortalInitialized)
	if len(log.Data) > 0 {
		if err := basePortal.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range basePortal.abi.Events[event].Inputs {
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

// BasePortalRespectedGameTypeSet represents a RespectedGameTypeSet event raised by the BasePortal contract.
type BasePortalRespectedGameTypeSet struct {
	NewGameType uint32
	UpdatedAt   uint64
	Raw         *types.Log // Blockchain specific contextual infos
}

const BasePortalRespectedGameTypeSetEventName = "RespectedGameTypeSet"

// ContractEventName returns the user-defined event name.
func (BasePortalRespectedGameTypeSet) ContractEventName() string {
	return BasePortalRespectedGameTypeSetEventName
}

// UnpackRespectedGameTypeSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RespectedGameTypeSet(uint32 indexed newGameType, uint64 indexed updatedAt)
func (basePortal *BasePortal) UnpackRespectedGameTypeSetEvent(log *types.Log) (*BasePortalRespectedGameTypeSet, error) {
	event := "RespectedGameTypeSet"
	if log.Topics[0] != basePortal.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BasePortalRespectedGameTypeSet)
	if len(log.Data) > 0 {
		if err := basePortal.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range basePortal.abi.Events[event].Inputs {
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

// BasePortalTransactionDeposited represents a TransactionDeposited event raised by the BasePortal contract.
type BasePortalTransactionDeposited struct {
	From       common.Address
	To         common.Address
	Version    *big.Int
	OpaqueData []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const BasePortalTransactionDepositedEventName = "TransactionDeposited"

// ContractEventName returns the user-defined event name.
func (BasePortalTransactionDeposited) ContractEventName() string {
	return BasePortalTransactionDepositedEventName
}

// UnpackTransactionDepositedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (basePortal *BasePortal) UnpackTransactionDepositedEvent(log *types.Log) (*BasePortalTransactionDeposited, error) {
	event := "TransactionDeposited"
	if log.Topics[0] != basePortal.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BasePortalTransactionDeposited)
	if len(log.Data) > 0 {
		if err := basePortal.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range basePortal.abi.Events[event].Inputs {
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

// BasePortalWithdrawalFinalized represents a WithdrawalFinalized event raised by the BasePortal contract.
type BasePortalWithdrawalFinalized struct {
	WithdrawalHash [32]byte
	Success        bool
	Raw            *types.Log // Blockchain specific contextual infos
}

const BasePortalWithdrawalFinalizedEventName = "WithdrawalFinalized"

// ContractEventName returns the user-defined event name.
func (BasePortalWithdrawalFinalized) ContractEventName() string {
	return BasePortalWithdrawalFinalizedEventName
}

// UnpackWithdrawalFinalizedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (basePortal *BasePortal) UnpackWithdrawalFinalizedEvent(log *types.Log) (*BasePortalWithdrawalFinalized, error) {
	event := "WithdrawalFinalized"
	if log.Topics[0] != basePortal.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BasePortalWithdrawalFinalized)
	if len(log.Data) > 0 {
		if err := basePortal.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range basePortal.abi.Events[event].Inputs {
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

// BasePortalWithdrawalProven represents a WithdrawalProven event raised by the BasePortal contract.
type BasePortalWithdrawalProven struct {
	WithdrawalHash [32]byte
	From           common.Address
	To             common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const BasePortalWithdrawalProvenEventName = "WithdrawalProven"

// ContractEventName returns the user-defined event name.
func (BasePortalWithdrawalProven) ContractEventName() string {
	return BasePortalWithdrawalProvenEventName
}

// UnpackWithdrawalProvenEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (basePortal *BasePortal) UnpackWithdrawalProvenEvent(log *types.Log) (*BasePortalWithdrawalProven, error) {
	event := "WithdrawalProven"
	if log.Topics[0] != basePortal.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BasePortalWithdrawalProven)
	if len(log.Data) > 0 {
		if err := basePortal.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range basePortal.abi.Events[event].Inputs {
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

// BasePortalWithdrawalProvenExtension1 represents a WithdrawalProvenExtension1 event raised by the BasePortal contract.
type BasePortalWithdrawalProvenExtension1 struct {
	WithdrawalHash [32]byte
	ProofSubmitter common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const BasePortalWithdrawalProvenExtension1EventName = "WithdrawalProvenExtension1"

// ContractEventName returns the user-defined event name.
func (BasePortalWithdrawalProvenExtension1) ContractEventName() string {
	return BasePortalWithdrawalProvenExtension1EventName
}

// UnpackWithdrawalProvenExtension1Event is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event WithdrawalProvenExtension1(bytes32 indexed withdrawalHash, address indexed proofSubmitter)
func (basePortal *BasePortal) UnpackWithdrawalProvenExtension1Event(log *types.Log) (*BasePortalWithdrawalProvenExtension1, error) {
	event := "WithdrawalProvenExtension1"
	if log.Topics[0] != basePortal.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(BasePortalWithdrawalProvenExtension1)
	if len(log.Data) > 0 {
		if err := basePortal.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range basePortal.abi.Events[event].Inputs {
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

// UnpackError attempts to decode the provided error data using user-defined
// error definitions.
func (basePortal *BasePortal) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], basePortal.abi.Errors["AlreadyFinalized"].ID.Bytes()[:4]) {
		return basePortal.UnpackAlreadyFinalizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["BadTarget"].ID.Bytes()[:4]) {
		return basePortal.UnpackBadTargetError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["Blacklisted"].ID.Bytes()[:4]) {
		return basePortal.UnpackBlacklistedError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["CallPaused"].ID.Bytes()[:4]) {
		return basePortal.UnpackCallPausedError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["ContentLengthMismatch"].ID.Bytes()[:4]) {
		return basePortal.UnpackContentLengthMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["EmptyItem"].ID.Bytes()[:4]) {
		return basePortal.UnpackEmptyItemError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["GasEstimation"].ID.Bytes()[:4]) {
		return basePortal.UnpackGasEstimationError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["InvalidDataRemainder"].ID.Bytes()[:4]) {
		return basePortal.UnpackInvalidDataRemainderError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["InvalidDisputeGame"].ID.Bytes()[:4]) {
		return basePortal.UnpackInvalidDisputeGameError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["InvalidGameType"].ID.Bytes()[:4]) {
		return basePortal.UnpackInvalidGameTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["InvalidHeader"].ID.Bytes()[:4]) {
		return basePortal.UnpackInvalidHeaderError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["InvalidMerkleProof"].ID.Bytes()[:4]) {
		return basePortal.UnpackInvalidMerkleProofError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["InvalidProof"].ID.Bytes()[:4]) {
		return basePortal.UnpackInvalidProofError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["LargeCalldata"].ID.Bytes()[:4]) {
		return basePortal.UnpackLargeCalldataError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["LegacyGame"].ID.Bytes()[:4]) {
		return basePortal.UnpackLegacyGameError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["NonReentrant"].ID.Bytes()[:4]) {
		return basePortal.UnpackNonReentrantError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["OutOfGas"].ID.Bytes()[:4]) {
		return basePortal.UnpackOutOfGasError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["ProposalNotValidated"].ID.Bytes()[:4]) {
		return basePortal.UnpackProposalNotValidatedError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["SmallGasLimit"].ID.Bytes()[:4]) {
		return basePortal.UnpackSmallGasLimitError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["Unauthorized"].ID.Bytes()[:4]) {
		return basePortal.UnpackUnauthorizedError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["UnexpectedList"].ID.Bytes()[:4]) {
		return basePortal.UnpackUnexpectedListError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["UnexpectedString"].ID.Bytes()[:4]) {
		return basePortal.UnpackUnexpectedStringError(raw[4:])
	}
	if bytes.Equal(raw[:4], basePortal.abi.Errors["Unproven"].ID.Bytes()[:4]) {
		return basePortal.UnpackUnprovenError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// BasePortalAlreadyFinalized represents a AlreadyFinalized error raised by the BasePortal contract.
type BasePortalAlreadyFinalized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyFinalized()
func BasePortalAlreadyFinalizedErrorID() common.Hash {
	return common.HexToHash("0x475a25352a73f4a08da552a02be7a21ad2f80f3f9491d94257fb6a042cfac983")
}

// UnpackAlreadyFinalizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyFinalized()
func (basePortal *BasePortal) UnpackAlreadyFinalizedError(raw []byte) (*BasePortalAlreadyFinalized, error) {
	out := new(BasePortalAlreadyFinalized)
	if err := basePortal.abi.UnpackIntoInterface(out, "AlreadyFinalized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalBadTarget represents a BadTarget error raised by the BasePortal contract.
type BasePortalBadTarget struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BadTarget()
func BasePortalBadTargetErrorID() common.Hash {
	return common.HexToHash("0x13496fdad3875138129e92bf0d7a120d9f7f119281644e3e5ea665ccbbe21488")
}

// UnpackBadTargetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BadTarget()
func (basePortal *BasePortal) UnpackBadTargetError(raw []byte) (*BasePortalBadTarget, error) {
	out := new(BasePortalBadTarget)
	if err := basePortal.abi.UnpackIntoInterface(out, "BadTarget", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalBlacklisted represents a Blacklisted error raised by the BasePortal contract.
type BasePortalBlacklisted struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Blacklisted()
func BasePortalBlacklistedErrorID() common.Hash {
	return common.HexToHash("0x09550c770d1e7568680ee03aa94c6a5fca20c34f72b0bd46af877d58210a302f")
}

// UnpackBlacklistedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Blacklisted()
func (basePortal *BasePortal) UnpackBlacklistedError(raw []byte) (*BasePortalBlacklisted, error) {
	out := new(BasePortalBlacklisted)
	if err := basePortal.abi.UnpackIntoInterface(out, "Blacklisted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalCallPaused represents a CallPaused error raised by the BasePortal contract.
type BasePortalCallPaused struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CallPaused()
func BasePortalCallPausedErrorID() common.Hash {
	return common.HexToHash("0xf480973e0c9de46b32d662d48e82e8a3d28e25a06bf4401edda2e286235a3eb9")
}

// UnpackCallPausedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CallPaused()
func (basePortal *BasePortal) UnpackCallPausedError(raw []byte) (*BasePortalCallPaused, error) {
	out := new(BasePortalCallPaused)
	if err := basePortal.abi.UnpackIntoInterface(out, "CallPaused", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalContentLengthMismatch represents a ContentLengthMismatch error raised by the BasePortal contract.
type BasePortalContentLengthMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ContentLengthMismatch()
func BasePortalContentLengthMismatchErrorID() common.Hash {
	return common.HexToHash("0x66c944854c74f126ddb347ede5ddccb741d5d7926bed34960c8d068de5a8f475")
}

// UnpackContentLengthMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ContentLengthMismatch()
func (basePortal *BasePortal) UnpackContentLengthMismatchError(raw []byte) (*BasePortalContentLengthMismatch, error) {
	out := new(BasePortalContentLengthMismatch)
	if err := basePortal.abi.UnpackIntoInterface(out, "ContentLengthMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalEmptyItem represents a EmptyItem error raised by the BasePortal contract.
type BasePortalEmptyItem struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyItem()
func BasePortalEmptyItemErrorID() common.Hash {
	return common.HexToHash("0x5ab458fbe9490fbea962b45039c8271fd72a0f01d48093712043b25dc6ba708f")
}

// UnpackEmptyItemError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyItem()
func (basePortal *BasePortal) UnpackEmptyItemError(raw []byte) (*BasePortalEmptyItem, error) {
	out := new(BasePortalEmptyItem)
	if err := basePortal.abi.UnpackIntoInterface(out, "EmptyItem", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalGasEstimation represents a GasEstimation error raised by the BasePortal contract.
type BasePortalGasEstimation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GasEstimation()
func BasePortalGasEstimationErrorID() common.Hash {
	return common.HexToHash("0xeeae4ed317fad2887a8e15d44b7d5fba14a5f2673acb2ee5081bbedc688a8e2e")
}

// UnpackGasEstimationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GasEstimation()
func (basePortal *BasePortal) UnpackGasEstimationError(raw []byte) (*BasePortalGasEstimation, error) {
	out := new(BasePortalGasEstimation)
	if err := basePortal.abi.UnpackIntoInterface(out, "GasEstimation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalInvalidDataRemainder represents a InvalidDataRemainder error raised by the BasePortal contract.
type BasePortalInvalidDataRemainder struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDataRemainder()
func BasePortalInvalidDataRemainderErrorID() common.Hash {
	return common.HexToHash("0x5c5537b8753217199634d3b78fb6e68bd912e435605abb57ae70f055c98106de")
}

// UnpackInvalidDataRemainderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDataRemainder()
func (basePortal *BasePortal) UnpackInvalidDataRemainderError(raw []byte) (*BasePortalInvalidDataRemainder, error) {
	out := new(BasePortalInvalidDataRemainder)
	if err := basePortal.abi.UnpackIntoInterface(out, "InvalidDataRemainder", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalInvalidDisputeGame represents a InvalidDisputeGame error raised by the BasePortal contract.
type BasePortalInvalidDisputeGame struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidDisputeGame()
func BasePortalInvalidDisputeGameErrorID() common.Hash {
	return common.HexToHash("0xd3573474548064dfaddea84e04d97470a02d608bdcb9143d8877ba297061e536")
}

// UnpackInvalidDisputeGameError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidDisputeGame()
func (basePortal *BasePortal) UnpackInvalidDisputeGameError(raw []byte) (*BasePortalInvalidDisputeGame, error) {
	out := new(BasePortalInvalidDisputeGame)
	if err := basePortal.abi.UnpackIntoInterface(out, "InvalidDisputeGame", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalInvalidGameType represents a InvalidGameType error raised by the BasePortal contract.
type BasePortalInvalidGameType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGameType()
func BasePortalInvalidGameTypeErrorID() common.Hash {
	return common.HexToHash("0x27a10cc2008fefb3a25d677d8f9363901337a5b58afec35aea1299b7ef260472")
}

// UnpackInvalidGameTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGameType()
func (basePortal *BasePortal) UnpackInvalidGameTypeError(raw []byte) (*BasePortalInvalidGameType, error) {
	out := new(BasePortalInvalidGameType)
	if err := basePortal.abi.UnpackIntoInterface(out, "InvalidGameType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalInvalidHeader represents a InvalidHeader error raised by the BasePortal contract.
type BasePortalInvalidHeader struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidHeader()
func BasePortalInvalidHeaderErrorID() common.Hash {
	return common.HexToHash("0xbabb01ddbf35fa5cb4266dbe78b38e673d97e7c3b147273565c15eb8980ded33")
}

// UnpackInvalidHeaderError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidHeader()
func (basePortal *BasePortal) UnpackInvalidHeaderError(raw []byte) (*BasePortalInvalidHeader, error) {
	out := new(BasePortalInvalidHeader)
	if err := basePortal.abi.UnpackIntoInterface(out, "InvalidHeader", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalInvalidMerkleProof represents a InvalidMerkleProof error raised by the BasePortal contract.
type BasePortalInvalidMerkleProof struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidMerkleProof()
func BasePortalInvalidMerkleProofErrorID() common.Hash {
	return common.HexToHash("0xb05e92facfa0fd6ba9338977017107202232768a12b21141a91a36a56212ad1e")
}

// UnpackInvalidMerkleProofError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidMerkleProof()
func (basePortal *BasePortal) UnpackInvalidMerkleProofError(raw []byte) (*BasePortalInvalidMerkleProof, error) {
	out := new(BasePortalInvalidMerkleProof)
	if err := basePortal.abi.UnpackIntoInterface(out, "InvalidMerkleProof", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalInvalidProof represents a InvalidProof error raised by the BasePortal contract.
type BasePortalInvalidProof struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidProof()
func BasePortalInvalidProofErrorID() common.Hash {
	return common.HexToHash("0x09bde339c6b182be216ee7ef8ccff6338c6ef7993445216112ae575c5438fd27")
}

// UnpackInvalidProofError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidProof()
func (basePortal *BasePortal) UnpackInvalidProofError(raw []byte) (*BasePortalInvalidProof, error) {
	out := new(BasePortalInvalidProof)
	if err := basePortal.abi.UnpackIntoInterface(out, "InvalidProof", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalLargeCalldata represents a LargeCalldata error raised by the BasePortal contract.
type BasePortalLargeCalldata struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LargeCalldata()
func BasePortalLargeCalldataErrorID() common.Hash {
	return common.HexToHash("0x73052b0f64faca17fa96ffc1fa21ba7ee5b6c418c2f3ffaeb9af08b3c534092f")
}

// UnpackLargeCalldataError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LargeCalldata()
func (basePortal *BasePortal) UnpackLargeCalldataError(raw []byte) (*BasePortalLargeCalldata, error) {
	out := new(BasePortalLargeCalldata)
	if err := basePortal.abi.UnpackIntoInterface(out, "LargeCalldata", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalLegacyGame represents a LegacyGame error raised by the BasePortal contract.
type BasePortalLegacyGame struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LegacyGame()
func BasePortalLegacyGameErrorID() common.Hash {
	return common.HexToHash("0xd502c9a620330678ef0ae23173dba173810726fb7a9e43a38c1d9192ffb84694")
}

// UnpackLegacyGameError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LegacyGame()
func (basePortal *BasePortal) UnpackLegacyGameError(raw []byte) (*BasePortalLegacyGame, error) {
	out := new(BasePortalLegacyGame)
	if err := basePortal.abi.UnpackIntoInterface(out, "LegacyGame", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalNonReentrant represents a NonReentrant error raised by the BasePortal contract.
type BasePortalNonReentrant struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NonReentrant()
func BasePortalNonReentrantErrorID() common.Hash {
	return common.HexToHash("0x9396d156a2826722287cac55c202086fc9790a2ba575119c1f37b3bce209f2fa")
}

// UnpackNonReentrantError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NonReentrant()
func (basePortal *BasePortal) UnpackNonReentrantError(raw []byte) (*BasePortalNonReentrant, error) {
	out := new(BasePortalNonReentrant)
	if err := basePortal.abi.UnpackIntoInterface(out, "NonReentrant", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalOutOfGas represents a OutOfGas error raised by the BasePortal contract.
type BasePortalOutOfGas struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OutOfGas()
func BasePortalOutOfGasErrorID() common.Hash {
	return common.HexToHash("0x77ebef4d460014cfc45ab055292b3a8f4a7c9e9ab42973e79fdaff6b27233461")
}

// UnpackOutOfGasError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OutOfGas()
func (basePortal *BasePortal) UnpackOutOfGasError(raw []byte) (*BasePortalOutOfGas, error) {
	out := new(BasePortalOutOfGas)
	if err := basePortal.abi.UnpackIntoInterface(out, "OutOfGas", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalProposalNotValidated represents a ProposalNotValidated error raised by the BasePortal contract.
type BasePortalProposalNotValidated struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ProposalNotValidated()
func BasePortalProposalNotValidatedErrorID() common.Hash {
	return common.HexToHash("0xa080a3c4475df694083ceef07de99f7d96f1e61e8c3f6b16d23c308e28f4948b")
}

// UnpackProposalNotValidatedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ProposalNotValidated()
func (basePortal *BasePortal) UnpackProposalNotValidatedError(raw []byte) (*BasePortalProposalNotValidated, error) {
	out := new(BasePortalProposalNotValidated)
	if err := basePortal.abi.UnpackIntoInterface(out, "ProposalNotValidated", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalSmallGasLimit represents a SmallGasLimit error raised by the BasePortal contract.
type BasePortalSmallGasLimit struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SmallGasLimit()
func BasePortalSmallGasLimitErrorID() common.Hash {
	return common.HexToHash("0x4929b808dec65c5ed411ea9b10187a0d4d141812282a141b020d83e8c7494eb7")
}

// UnpackSmallGasLimitError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SmallGasLimit()
func (basePortal *BasePortal) UnpackSmallGasLimitError(raw []byte) (*BasePortalSmallGasLimit, error) {
	out := new(BasePortalSmallGasLimit)
	if err := basePortal.abi.UnpackIntoInterface(out, "SmallGasLimit", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalUnauthorized represents a Unauthorized error raised by the BasePortal contract.
type BasePortalUnauthorized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Unauthorized()
func BasePortalUnauthorizedErrorID() common.Hash {
	return common.HexToHash("0x82b4290015f7ec7256ca2a6247d3c2a89c4865c0e791456df195f40ad0a81367")
}

// UnpackUnauthorizedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Unauthorized()
func (basePortal *BasePortal) UnpackUnauthorizedError(raw []byte) (*BasePortalUnauthorized, error) {
	out := new(BasePortalUnauthorized)
	if err := basePortal.abi.UnpackIntoInterface(out, "Unauthorized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalUnexpectedList represents a UnexpectedList error raised by the BasePortal contract.
type BasePortalUnexpectedList struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnexpectedList()
func BasePortalUnexpectedListErrorID() common.Hash {
	return common.HexToHash("0x1ff9b2e4a9eba439988cd6b8e020496ddcf6d6cefbd8ccaa0c6006330e0183fd")
}

// UnpackUnexpectedListError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnexpectedList()
func (basePortal *BasePortal) UnpackUnexpectedListError(raw []byte) (*BasePortalUnexpectedList, error) {
	out := new(BasePortalUnexpectedList)
	if err := basePortal.abi.UnpackIntoInterface(out, "UnexpectedList", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalUnexpectedString represents a UnexpectedString error raised by the BasePortal contract.
type BasePortalUnexpectedString struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnexpectedString()
func BasePortalUnexpectedStringErrorID() common.Hash {
	return common.HexToHash("0x4b9c6abeb08b7a3e5561c35d6edd32e38a8bf05c863166482044d76d816eaad0")
}

// UnpackUnexpectedStringError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnexpectedString()
func (basePortal *BasePortal) UnpackUnexpectedStringError(raw []byte) (*BasePortalUnexpectedString, error) {
	out := new(BasePortalUnexpectedString)
	if err := basePortal.abi.UnpackIntoInterface(out, "UnexpectedString", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// BasePortalUnproven represents a Unproven error raised by the BasePortal contract.
type BasePortalUnproven struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Unproven()
func BasePortalUnprovenErrorID() common.Hash {
	return common.HexToHash("0x94efd49b02cb0e171f0d2992bdbe2f2be71d2e81a763f3a235aa23112f75598d")
}

// UnpackUnprovenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Unproven()
func (basePortal *BasePortal) UnpackUnprovenError(raw []byte) (*BasePortalUnproven, error) {
	out := new(BasePortalUnproven)
	if err := basePortal.abi.UnpackIntoInterface(out, "Unproven", raw); err != nil {
		return nil, err
	}
	return out, nil
}

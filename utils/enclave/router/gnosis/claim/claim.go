// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gnosis_claim

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

// GnosisClaimMetaData contains all meta data concerning the GnosisClaim contract.
var GnosisClaimMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ClaimUsdsNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FOREIGN_AMB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FOREIGN_OMNIBRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FOREIGN_XDAIBRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USDS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH_OMNIBRIDGE_ROUTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"executeSignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"executeSignaturesUSDS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverLockedFund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"relayTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"safeExecuteSignaturesWithAutoGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_route\",\"type\":\"address\"}],\"name\":\"setRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenRoutes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "GnosisClaim",
}

// GnosisClaim is an auto generated Go binding around an Ethereum contract.
type GnosisClaim struct {
	abi abi.ABI
}

// NewGnosisClaim creates a new instance of GnosisClaim.
func NewGnosisClaim() *GnosisClaim {
	parsed, err := GnosisClaimMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &GnosisClaim{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *GnosisClaim) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackDAI is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0bab4c4.
//
// Solidity: function DAI() view returns(address)
func (gnosisClaim *GnosisClaim) PackDAI() []byte {
	enc, err := gnosisClaim.abi.Pack("DAI")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackDAI is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe0bab4c4.
//
// Solidity: function DAI() view returns(address)
func (gnosisClaim *GnosisClaim) UnpackDAI(data []byte) (common.Address, error) {
	out, err := gnosisClaim.abi.Unpack("DAI", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackFOREIGNAMB is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3f02fb7.
//
// Solidity: function FOREIGN_AMB() view returns(address)
func (gnosisClaim *GnosisClaim) PackFOREIGNAMB() []byte {
	enc, err := gnosisClaim.abi.Pack("FOREIGN_AMB")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFOREIGNAMB is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3f02fb7.
//
// Solidity: function FOREIGN_AMB() view returns(address)
func (gnosisClaim *GnosisClaim) UnpackFOREIGNAMB(data []byte) (common.Address, error) {
	out, err := gnosisClaim.abi.Unpack("FOREIGN_AMB", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackFOREIGNOMNIBRIDGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x044abec9.
//
// Solidity: function FOREIGN_OMNIBRIDGE() view returns(address)
func (gnosisClaim *GnosisClaim) PackFOREIGNOMNIBRIDGE() []byte {
	enc, err := gnosisClaim.abi.Pack("FOREIGN_OMNIBRIDGE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFOREIGNOMNIBRIDGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x044abec9.
//
// Solidity: function FOREIGN_OMNIBRIDGE() view returns(address)
func (gnosisClaim *GnosisClaim) UnpackFOREIGNOMNIBRIDGE(data []byte) (common.Address, error) {
	out, err := gnosisClaim.abi.Unpack("FOREIGN_OMNIBRIDGE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackFOREIGNXDAIBRIDGE is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe08e1a70.
//
// Solidity: function FOREIGN_XDAIBRIDGE() view returns(address)
func (gnosisClaim *GnosisClaim) PackFOREIGNXDAIBRIDGE() []byte {
	enc, err := gnosisClaim.abi.Pack("FOREIGN_XDAIBRIDGE")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackFOREIGNXDAIBRIDGE is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe08e1a70.
//
// Solidity: function FOREIGN_XDAIBRIDGE() view returns(address)
func (gnosisClaim *GnosisClaim) UnpackFOREIGNXDAIBRIDGE(data []byte) (common.Address, error) {
	out, err := gnosisClaim.abi.Unpack("FOREIGN_XDAIBRIDGE", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackUSDS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc8ef95ae.
//
// Solidity: function USDS() view returns(address)
func (gnosisClaim *GnosisClaim) PackUSDS() []byte {
	enc, err := gnosisClaim.abi.Pack("USDS")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackUSDS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc8ef95ae.
//
// Solidity: function USDS() view returns(address)
func (gnosisClaim *GnosisClaim) UnpackUSDS(data []byte) (common.Address, error) {
	out, err := gnosisClaim.abi.Unpack("USDS", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackWETHOMNIBRIDGEROUTER is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c254601.
//
// Solidity: function WETH_OMNIBRIDGE_ROUTER() view returns(address)
func (gnosisClaim *GnosisClaim) PackWETHOMNIBRIDGEROUTER() []byte {
	enc, err := gnosisClaim.abi.Pack("WETH_OMNIBRIDGE_ROUTER")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackWETHOMNIBRIDGEROUTER is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4c254601.
//
// Solidity: function WETH_OMNIBRIDGE_ROUTER() view returns(address)
func (gnosisClaim *GnosisClaim) UnpackWETHOMNIBRIDGEROUTER(data []byte) (common.Address, error) {
	out, err := gnosisClaim.abi.Unpack("WETH_OMNIBRIDGE_ROUTER", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackExecuteSignatures is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3f7658fd.
//
// Solidity: function executeSignatures(bytes message, bytes signatures) returns()
func (gnosisClaim *GnosisClaim) PackExecuteSignatures(message []byte, signatures []byte) []byte {
	enc, err := gnosisClaim.abi.Pack("executeSignatures", message, signatures)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackExecuteSignaturesUSDS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xae6862cd.
//
// Solidity: function executeSignaturesUSDS(bytes message, bytes signatures) returns()
func (gnosisClaim *GnosisClaim) PackExecuteSignaturesUSDS(message []byte, signatures []byte) []byte {
	enc, err := gnosisClaim.abi.Pack("executeSignaturesUSDS", message, signatures)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (gnosisClaim *GnosisClaim) PackInitialize(owner common.Address) []byte {
	enc, err := gnosisClaim.abi.Pack("initialize", owner)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackOwner is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (gnosisClaim *GnosisClaim) PackOwner() []byte {
	enc, err := gnosisClaim.abi.Pack("owner")
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackOwner is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (gnosisClaim *GnosisClaim) UnpackOwner(data []byte) (common.Address, error) {
	out, err := gnosisClaim.abi.Unpack("owner", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackRecoverLockedFund is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x423fd7dc.
//
// Solidity: function recoverLockedFund(address token, address recipient, uint256 amount) returns()
func (gnosisClaim *GnosisClaim) PackRecoverLockedFund(token common.Address, recipient common.Address, amount *big.Int) []byte {
	enc, err := gnosisClaim.abi.Pack("recoverLockedFund", token, recipient, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRelayTokens is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad58bdd1.
//
// Solidity: function relayTokens(address _token, address _receiver, uint256 _amount) payable returns()
func (gnosisClaim *GnosisClaim) PackRelayTokens(token common.Address, receiver common.Address, amount *big.Int) []byte {
	enc, err := gnosisClaim.abi.Pack("relayTokens", token, receiver, amount)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackRenounceOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (gnosisClaim *GnosisClaim) PackRenounceOwnership() []byte {
	enc, err := gnosisClaim.abi.Pack("renounceOwnership")
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSafeExecuteSignaturesWithAutoGasLimit is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23caab49.
//
// Solidity: function safeExecuteSignaturesWithAutoGasLimit(bytes message, bytes signatures) returns()
func (gnosisClaim *GnosisClaim) PackSafeExecuteSignaturesWithAutoGasLimit(message []byte, signatures []byte) []byte {
	enc, err := gnosisClaim.abi.Pack("safeExecuteSignaturesWithAutoGasLimit", message, signatures)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackSetRoute is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0505e94d.
//
// Solidity: function setRoute(address _token, address _route) returns()
func (gnosisClaim *GnosisClaim) PackSetRoute(token common.Address, route common.Address) []byte {
	enc, err := gnosisClaim.abi.Pack("setRoute", token, route)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackTokenRoutes is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x489f945c.
//
// Solidity: function tokenRoutes(address ) view returns(address)
func (gnosisClaim *GnosisClaim) PackTokenRoutes(arg0 common.Address) []byte {
	enc, err := gnosisClaim.abi.Pack("tokenRoutes", arg0)
	if err != nil {
		panic(err)
	}
	return enc
}

// UnpackTokenRoutes is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x489f945c.
//
// Solidity: function tokenRoutes(address ) view returns(address)
func (gnosisClaim *GnosisClaim) UnpackTokenRoutes(data []byte) (common.Address, error) {
	out, err := gnosisClaim.abi.Unpack("tokenRoutes", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, err
}

// PackTransferOwnership is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (gnosisClaim *GnosisClaim) PackTransferOwnership(newOwner common.Address) []byte {
	enc, err := gnosisClaim.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		panic(err)
	}
	return enc
}

// GnosisClaimInitialized represents a Initialized event raised by the GnosisClaim contract.
type GnosisClaimInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const GnosisClaimInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (GnosisClaimInitialized) ContractEventName() string {
	return GnosisClaimInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (gnosisClaim *GnosisClaim) UnpackInitializedEvent(log *types.Log) (*GnosisClaimInitialized, error) {
	event := "Initialized"
	if log.Topics[0] != gnosisClaim.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisClaimInitialized)
	if len(log.Data) > 0 {
		if err := gnosisClaim.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisClaim.abi.Events[event].Inputs {
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

// GnosisClaimOwnershipTransferred represents a OwnershipTransferred event raised by the GnosisClaim contract.
type GnosisClaimOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           *types.Log // Blockchain specific contextual infos
}

const GnosisClaimOwnershipTransferredEventName = "OwnershipTransferred"

// ContractEventName returns the user-defined event name.
func (GnosisClaimOwnershipTransferred) ContractEventName() string {
	return GnosisClaimOwnershipTransferredEventName
}

// UnpackOwnershipTransferredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (gnosisClaim *GnosisClaim) UnpackOwnershipTransferredEvent(log *types.Log) (*GnosisClaimOwnershipTransferred, error) {
	event := "OwnershipTransferred"
	if log.Topics[0] != gnosisClaim.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(GnosisClaimOwnershipTransferred)
	if len(log.Data) > 0 {
		if err := gnosisClaim.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range gnosisClaim.abi.Events[event].Inputs {
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
func (gnosisClaim *GnosisClaim) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], gnosisClaim.abi.Errors["ClaimUsdsNotSupported"].ID.Bytes()[:4]) {
		return gnosisClaim.UnpackClaimUsdsNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], gnosisClaim.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return gnosisClaim.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], gnosisClaim.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return gnosisClaim.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], gnosisClaim.abi.Errors["OwnableInvalidOwner"].ID.Bytes()[:4]) {
		return gnosisClaim.UnpackOwnableInvalidOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], gnosisClaim.abi.Errors["OwnableUnauthorizedAccount"].ID.Bytes()[:4]) {
		return gnosisClaim.UnpackOwnableUnauthorizedAccountError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// GnosisClaimClaimUsdsNotSupported represents a ClaimUsdsNotSupported error raised by the GnosisClaim contract.
type GnosisClaimClaimUsdsNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ClaimUsdsNotSupported()
func GnosisClaimClaimUsdsNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x662554fc43b5d87090a5f9e8b365ca35213d23ae7082671886b760dc510ee8da")
}

// UnpackClaimUsdsNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ClaimUsdsNotSupported()
func (gnosisClaim *GnosisClaim) UnpackClaimUsdsNotSupportedError(raw []byte) (*GnosisClaimClaimUsdsNotSupported, error) {
	out := new(GnosisClaimClaimUsdsNotSupported)
	if err := gnosisClaim.abi.UnpackIntoInterface(out, "ClaimUsdsNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GnosisClaimInvalidInitialization represents a InvalidInitialization error raised by the GnosisClaim contract.
type GnosisClaimInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func GnosisClaimInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (gnosisClaim *GnosisClaim) UnpackInvalidInitializationError(raw []byte) (*GnosisClaimInvalidInitialization, error) {
	out := new(GnosisClaimInvalidInitialization)
	if err := gnosisClaim.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GnosisClaimNotInitializing represents a NotInitializing error raised by the GnosisClaim contract.
type GnosisClaimNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func GnosisClaimNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (gnosisClaim *GnosisClaim) UnpackNotInitializingError(raw []byte) (*GnosisClaimNotInitializing, error) {
	out := new(GnosisClaimNotInitializing)
	if err := gnosisClaim.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GnosisClaimOwnableInvalidOwner represents a OwnableInvalidOwner error raised by the GnosisClaim contract.
type GnosisClaimOwnableInvalidOwner struct {
	Owner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableInvalidOwner(address owner)
func GnosisClaimOwnableInvalidOwnerErrorID() common.Hash {
	return common.HexToHash("0x1e4fbdf7f3ef8bcaa855599e3abf48b232380f183f08f6f813d9ffa5bd585188")
}

// UnpackOwnableInvalidOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableInvalidOwner(address owner)
func (gnosisClaim *GnosisClaim) UnpackOwnableInvalidOwnerError(raw []byte) (*GnosisClaimOwnableInvalidOwner, error) {
	out := new(GnosisClaimOwnableInvalidOwner)
	if err := gnosisClaim.abi.UnpackIntoInterface(out, "OwnableInvalidOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// GnosisClaimOwnableUnauthorizedAccount represents a OwnableUnauthorizedAccount error raised by the GnosisClaim contract.
type GnosisClaimOwnableUnauthorizedAccount struct {
	Account common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func GnosisClaimOwnableUnauthorizedAccountErrorID() common.Hash {
	return common.HexToHash("0x118cdaa7a341953d1887a2245fd6665d741c67c8c50581daa59e1d03373fa188")
}

// UnpackOwnableUnauthorizedAccountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OwnableUnauthorizedAccount(address account)
func (gnosisClaim *GnosisClaim) UnpackOwnableUnauthorizedAccountError(raw []byte) (*GnosisClaimOwnableUnauthorizedAccount, error) {
	out := new(GnosisClaimOwnableUnauthorizedAccount)
	if err := gnosisClaim.abi.UnpackIntoInterface(out, "OwnableUnauthorizedAccount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

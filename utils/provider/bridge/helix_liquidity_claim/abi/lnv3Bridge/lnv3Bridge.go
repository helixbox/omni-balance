// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lnv3Bridge

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

// LnBridgeSourceV3TokenConfigure is an auto generated low-level Go binding around an user-defined struct.
type LnBridgeSourceV3TokenConfigure struct {
	ProtocolFee    *big.Int
	Penalty        *big.Int
	SourceDecimals uint8
	TargetDecimals uint8
}

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

// LnBridgeTargetV3RelayParams is an auto generated low-level Go binding around an user-defined struct.
type LnBridgeTargetV3RelayParams struct {
	RemoteChainId *big.Int
	Provider      common.Address
	SourceToken   common.Address
	TargetToken   common.Address
	SourceAmount  *big.Int
	TargetAmount  *big.Int
	Receiver      common.Address
	Timestamp     *big.Int
}

// Lnv3BridgeMetaData contains all meta data concerning the Lnv3Bridge contract.
var Lnv3BridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenInfoKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"FeeIncomeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"transferIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"LnProviderPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"baseFee\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidityfeeRate\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"transferLimit\",\"type\":\"uint112\"}],\"name\":\"LnProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedPanaltyReserve\",\"type\":\"uint256\"}],\"name\":\"PenaltyReserveUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"}],\"name\":\"SlashRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenInfoKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"protocolFee\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"penalty\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"sourceDecimals\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"targetDecimals\",\"type\":\"uint112\"}],\"name\":\"TokenInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"totalFee\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLnBridgeSourceV3.TransferParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"targetAmount\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"fee\",\"type\":\"uint112\"}],\"name\":\"TokenLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"protocolFee\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"penalty\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"TransferFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"slashAmount\",\"type\":\"uint112\"}],\"name\":\"TransferSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LIQUIDITY_FEE_RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_STATUS_LOCKED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_STATUS_SLASHED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_STATUS_WITHDRAWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOCK_TIME_DISTANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TRANSFER_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASH_EXPIRE_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_tokenInfoKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimProtocolFeeIncome\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"deleteTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositPenaltyReserve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fillTransfers\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"}],\"name\":\"getProviderKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getProviderStateKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"}],\"name\":\"getTokenKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"totalFee\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLnBridgeSourceV3.TransferParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"uint112\",\"name\":\"_remoteAmount\",\"type\":\"uint112\"}],\"name\":\"getTransferId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dao\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"totalFee\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"amount\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLnBridgeSourceV3.TransferParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"lockAndRemoteRelease\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"lockInfos\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"amountWithFeeAndPenalty\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"tokenIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"messagers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sendService\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiveService\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"penaltyReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingDao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"}],\"name\":\"providerPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"}],\"name\":\"providerUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"_baseFee\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"_liquidityFeeRate\",\"type\":\"uint16\"},{\"internalType\":\"uint112\",\"name\":\"_transferLimit\",\"type\":\"uint112\"}],\"name\":\"registerLnProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"_protocolFee\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_penalty\",\"type\":\"uint112\"},{\"internalType\":\"uint8\",\"name\":\"_sourceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_targetDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_index\",\"type\":\"uint32\"}],\"name\":\"registerTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"sourceAmount\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"targetAmount\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLnBridgeTargetV3.RelayParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_expectedTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_relayBySelf\",\"type\":\"bool\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"sourceAmount\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"targetAmount\",\"type\":\"uint112\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structLnBridgeTargetV3.RelayParams\",\"name\":\"_params\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_expectedTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_feePrepaid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extParams\",\"type\":\"bytes\"}],\"name\":\"requestSlashAndRemoteRelease\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_transferIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extParams\",\"type\":\"bytes\"}],\"name\":\"requestWithdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_extParams\",\"type\":\"bytes\"}],\"name\":\"retrySlash\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_remoteBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_service\",\"type\":\"address\"}],\"name\":\"setReceiveService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_remoteBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_service\",\"type\":\"address\"}],\"name\":\"setSendService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_transferId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_lnProvider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"slashInfos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"srcProviders\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"baseFee\",\"type\":\"uint112\"},{\"internalType\":\"uint16\",\"name\":\"liquidityFeeRate\",\"type\":\"uint16\"},{\"internalType\":\"uint112\",\"name\":\"transferLimit\",\"type\":\"uint112\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"tokenIndexer\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"tokenInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint112\",\"name\":\"protocolFee\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"penalty\",\"type\":\"uint112\"},{\"internalType\":\"uint8\",\"name\":\"sourceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"targetDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structLnBridgeSourceV3.TokenConfigure\",\"name\":\"config\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"targetToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeeIncome\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"_amount\",\"type\":\"uint112\"}],\"name\":\"totalFee\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetToken\",\"type\":\"address\"},{\"internalType\":\"uint112\",\"name\":\"_protocolFee\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_penalty\",\"type\":\"uint112\"},{\"internalType\":\"uint8\",\"name\":\"_sourceDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_targetDecimals\",\"type\":\"uint8\"}],\"name\":\"updateTokenInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_transferIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"}],\"name\":\"withdrawLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sourceToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPenaltyReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Lnv3BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use Lnv3BridgeMetaData.ABI instead.
var Lnv3BridgeABI = Lnv3BridgeMetaData.ABI

// Lnv3Bridge is an auto generated Go binding around an Ethereum contract.
type Lnv3Bridge struct {
	Lnv3BridgeCaller     // Read-only binding to the contract
	Lnv3BridgeTransactor // Write-only binding to the contract
	Lnv3BridgeFilterer   // Log filterer for contract events
}

// Lnv3BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type Lnv3BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lnv3BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Lnv3BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lnv3BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Lnv3BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Lnv3BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Lnv3BridgeSession struct {
	Contract     *Lnv3Bridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Lnv3BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Lnv3BridgeCallerSession struct {
	Contract *Lnv3BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// Lnv3BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Lnv3BridgeTransactorSession struct {
	Contract     *Lnv3BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// Lnv3BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type Lnv3BridgeRaw struct {
	Contract *Lnv3Bridge // Generic contract binding to access the raw methods on
}

// Lnv3BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Lnv3BridgeCallerRaw struct {
	Contract *Lnv3BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// Lnv3BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Lnv3BridgeTransactorRaw struct {
	Contract *Lnv3BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLnv3Bridge creates a new instance of Lnv3Bridge, bound to a specific deployed contract.
func NewLnv3Bridge(address common.Address, backend bind.ContractBackend) (*Lnv3Bridge, error) {
	contract, err := bindLnv3Bridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lnv3Bridge{Lnv3BridgeCaller: Lnv3BridgeCaller{contract: contract}, Lnv3BridgeTransactor: Lnv3BridgeTransactor{contract: contract}, Lnv3BridgeFilterer: Lnv3BridgeFilterer{contract: contract}}, nil
}

// NewLnv3BridgeCaller creates a new read-only instance of Lnv3Bridge, bound to a specific deployed contract.
func NewLnv3BridgeCaller(address common.Address, caller bind.ContractCaller) (*Lnv3BridgeCaller, error) {
	contract, err := bindLnv3Bridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeCaller{contract: contract}, nil
}

// NewLnv3BridgeTransactor creates a new write-only instance of Lnv3Bridge, bound to a specific deployed contract.
func NewLnv3BridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*Lnv3BridgeTransactor, error) {
	contract, err := bindLnv3Bridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeTransactor{contract: contract}, nil
}

// NewLnv3BridgeFilterer creates a new log filterer instance of Lnv3Bridge, bound to a specific deployed contract.
func NewLnv3BridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*Lnv3BridgeFilterer, error) {
	contract, err := bindLnv3Bridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeFilterer{contract: contract}, nil
}

// bindLnv3Bridge binds a generic wrapper to an already deployed contract.
func bindLnv3Bridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Lnv3BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lnv3Bridge *Lnv3BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lnv3Bridge.Contract.Lnv3BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lnv3Bridge *Lnv3BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Lnv3BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lnv3Bridge *Lnv3BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Lnv3BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lnv3Bridge *Lnv3BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lnv3Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lnv3Bridge *Lnv3BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lnv3Bridge *Lnv3BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.contract.Transact(opts, method, params...)
}

// LIQUIDITYFEERATEBASE is a free data retrieval call binding the contract method 0xf99bb152.
//
// Solidity: function LIQUIDITY_FEE_RATE_BASE() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCaller) LIQUIDITYFEERATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "LIQUIDITY_FEE_RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LIQUIDITYFEERATEBASE is a free data retrieval call binding the contract method 0xf99bb152.
//
// Solidity: function LIQUIDITY_FEE_RATE_BASE() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeSession) LIQUIDITYFEERATEBASE() (*big.Int, error) {
	return _Lnv3Bridge.Contract.LIQUIDITYFEERATEBASE(&_Lnv3Bridge.CallOpts)
}

// LIQUIDITYFEERATEBASE is a free data retrieval call binding the contract method 0xf99bb152.
//
// Solidity: function LIQUIDITY_FEE_RATE_BASE() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) LIQUIDITYFEERATEBASE() (*big.Int, error) {
	return _Lnv3Bridge.Contract.LIQUIDITYFEERATEBASE(&_Lnv3Bridge.CallOpts)
}

// LOCKSTATUSLOCKED is a free data retrieval call binding the contract method 0x9a168fa8.
//
// Solidity: function LOCK_STATUS_LOCKED() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeCaller) LOCKSTATUSLOCKED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "LOCK_STATUS_LOCKED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LOCKSTATUSLOCKED is a free data retrieval call binding the contract method 0x9a168fa8.
//
// Solidity: function LOCK_STATUS_LOCKED() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeSession) LOCKSTATUSLOCKED() (uint8, error) {
	return _Lnv3Bridge.Contract.LOCKSTATUSLOCKED(&_Lnv3Bridge.CallOpts)
}

// LOCKSTATUSLOCKED is a free data retrieval call binding the contract method 0x9a168fa8.
//
// Solidity: function LOCK_STATUS_LOCKED() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) LOCKSTATUSLOCKED() (uint8, error) {
	return _Lnv3Bridge.Contract.LOCKSTATUSLOCKED(&_Lnv3Bridge.CallOpts)
}

// LOCKSTATUSSLASHED is a free data retrieval call binding the contract method 0xe5f221be.
//
// Solidity: function LOCK_STATUS_SLASHED() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeCaller) LOCKSTATUSSLASHED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "LOCK_STATUS_SLASHED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LOCKSTATUSSLASHED is a free data retrieval call binding the contract method 0xe5f221be.
//
// Solidity: function LOCK_STATUS_SLASHED() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeSession) LOCKSTATUSSLASHED() (uint8, error) {
	return _Lnv3Bridge.Contract.LOCKSTATUSSLASHED(&_Lnv3Bridge.CallOpts)
}

// LOCKSTATUSSLASHED is a free data retrieval call binding the contract method 0xe5f221be.
//
// Solidity: function LOCK_STATUS_SLASHED() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) LOCKSTATUSSLASHED() (uint8, error) {
	return _Lnv3Bridge.Contract.LOCKSTATUSSLASHED(&_Lnv3Bridge.CallOpts)
}

// LOCKSTATUSWITHDRAWN is a free data retrieval call binding the contract method 0x429912cf.
//
// Solidity: function LOCK_STATUS_WITHDRAWN() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeCaller) LOCKSTATUSWITHDRAWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "LOCK_STATUS_WITHDRAWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LOCKSTATUSWITHDRAWN is a free data retrieval call binding the contract method 0x429912cf.
//
// Solidity: function LOCK_STATUS_WITHDRAWN() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeSession) LOCKSTATUSWITHDRAWN() (uint8, error) {
	return _Lnv3Bridge.Contract.LOCKSTATUSWITHDRAWN(&_Lnv3Bridge.CallOpts)
}

// LOCKSTATUSWITHDRAWN is a free data retrieval call binding the contract method 0x429912cf.
//
// Solidity: function LOCK_STATUS_WITHDRAWN() view returns(uint8)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) LOCKSTATUSWITHDRAWN() (uint8, error) {
	return _Lnv3Bridge.Contract.LOCKSTATUSWITHDRAWN(&_Lnv3Bridge.CallOpts)
}

// LOCKTIMEDISTANCE is a free data retrieval call binding the contract method 0x2300f1c9.
//
// Solidity: function LOCK_TIME_DISTANCE() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCaller) LOCKTIMEDISTANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "LOCK_TIME_DISTANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKTIMEDISTANCE is a free data retrieval call binding the contract method 0x2300f1c9.
//
// Solidity: function LOCK_TIME_DISTANCE() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeSession) LOCKTIMEDISTANCE() (*big.Int, error) {
	return _Lnv3Bridge.Contract.LOCKTIMEDISTANCE(&_Lnv3Bridge.CallOpts)
}

// LOCKTIMEDISTANCE is a free data retrieval call binding the contract method 0x2300f1c9.
//
// Solidity: function LOCK_TIME_DISTANCE() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) LOCKTIMEDISTANCE() (*big.Int, error) {
	return _Lnv3Bridge.Contract.LOCKTIMEDISTANCE(&_Lnv3Bridge.CallOpts)
}

// MAXTRANSFERAMOUNT is a free data retrieval call binding the contract method 0xfc1b3113.
//
// Solidity: function MAX_TRANSFER_AMOUNT() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCaller) MAXTRANSFERAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "MAX_TRANSFER_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTRANSFERAMOUNT is a free data retrieval call binding the contract method 0xfc1b3113.
//
// Solidity: function MAX_TRANSFER_AMOUNT() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeSession) MAXTRANSFERAMOUNT() (*big.Int, error) {
	return _Lnv3Bridge.Contract.MAXTRANSFERAMOUNT(&_Lnv3Bridge.CallOpts)
}

// MAXTRANSFERAMOUNT is a free data retrieval call binding the contract method 0xfc1b3113.
//
// Solidity: function MAX_TRANSFER_AMOUNT() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) MAXTRANSFERAMOUNT() (*big.Int, error) {
	return _Lnv3Bridge.Contract.MAXTRANSFERAMOUNT(&_Lnv3Bridge.CallOpts)
}

// SLASHEXPIRETIME is a free data retrieval call binding the contract method 0x0373e9ef.
//
// Solidity: function SLASH_EXPIRE_TIME() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCaller) SLASHEXPIRETIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "SLASH_EXPIRE_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLASHEXPIRETIME is a free data retrieval call binding the contract method 0x0373e9ef.
//
// Solidity: function SLASH_EXPIRE_TIME() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeSession) SLASHEXPIRETIME() (*big.Int, error) {
	return _Lnv3Bridge.Contract.SLASHEXPIRETIME(&_Lnv3Bridge.CallOpts)
}

// SLASHEXPIRETIME is a free data retrieval call binding the contract method 0x0373e9ef.
//
// Solidity: function SLASH_EXPIRE_TIME() view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) SLASHEXPIRETIME() (*big.Int, error) {
	return _Lnv3Bridge.Contract.SLASHEXPIRETIME(&_Lnv3Bridge.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeSession) Dao() (common.Address, error) {
	return _Lnv3Bridge.Contract.Dao(&_Lnv3Bridge.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) Dao() (common.Address, error) {
	return _Lnv3Bridge.Contract.Dao(&_Lnv3Bridge.CallOpts)
}

// FillTransfers is a free data retrieval call binding the contract method 0xc08fc262.
//
// Solidity: function fillTransfers(bytes32 ) view returns(uint64 timestamp, address provider)
func (_Lnv3Bridge *Lnv3BridgeCaller) FillTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timestamp uint64
	Provider  common.Address
}, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "fillTransfers", arg0)

	outstruct := new(struct {
		Timestamp uint64
		Provider  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Provider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// FillTransfers is a free data retrieval call binding the contract method 0xc08fc262.
//
// Solidity: function fillTransfers(bytes32 ) view returns(uint64 timestamp, address provider)
func (_Lnv3Bridge *Lnv3BridgeSession) FillTransfers(arg0 [32]byte) (struct {
	Timestamp uint64
	Provider  common.Address
}, error) {
	return _Lnv3Bridge.Contract.FillTransfers(&_Lnv3Bridge.CallOpts, arg0)
}

// FillTransfers is a free data retrieval call binding the contract method 0xc08fc262.
//
// Solidity: function fillTransfers(bytes32 ) view returns(uint64 timestamp, address provider)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) FillTransfers(arg0 [32]byte) (struct {
	Timestamp uint64
	Provider  common.Address
}, error) {
	return _Lnv3Bridge.Contract.FillTransfers(&_Lnv3Bridge.CallOpts, arg0)
}

// GetProviderKey is a free data retrieval call binding the contract method 0xb7fbfcc3.
//
// Solidity: function getProviderKey(uint256 _remoteChainId, address _provider, address _sourceToken, address _targetToken) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCaller) GetProviderKey(opts *bind.CallOpts, _remoteChainId *big.Int, _provider common.Address, _sourceToken common.Address, _targetToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "getProviderKey", _remoteChainId, _provider, _sourceToken, _targetToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetProviderKey is a free data retrieval call binding the contract method 0xb7fbfcc3.
//
// Solidity: function getProviderKey(uint256 _remoteChainId, address _provider, address _sourceToken, address _targetToken) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeSession) GetProviderKey(_remoteChainId *big.Int, _provider common.Address, _sourceToken common.Address, _targetToken common.Address) ([32]byte, error) {
	return _Lnv3Bridge.Contract.GetProviderKey(&_Lnv3Bridge.CallOpts, _remoteChainId, _provider, _sourceToken, _targetToken)
}

// GetProviderKey is a free data retrieval call binding the contract method 0xb7fbfcc3.
//
// Solidity: function getProviderKey(uint256 _remoteChainId, address _provider, address _sourceToken, address _targetToken) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) GetProviderKey(_remoteChainId *big.Int, _provider common.Address, _sourceToken common.Address, _targetToken common.Address) ([32]byte, error) {
	return _Lnv3Bridge.Contract.GetProviderKey(&_Lnv3Bridge.CallOpts, _remoteChainId, _provider, _sourceToken, _targetToken)
}

// GetProviderStateKey is a free data retrieval call binding the contract method 0x2e1ab85f.
//
// Solidity: function getProviderStateKey(address _sourceToken, address provider) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCaller) GetProviderStateKey(opts *bind.CallOpts, _sourceToken common.Address, provider common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "getProviderStateKey", _sourceToken, provider)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetProviderStateKey is a free data retrieval call binding the contract method 0x2e1ab85f.
//
// Solidity: function getProviderStateKey(address _sourceToken, address provider) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeSession) GetProviderStateKey(_sourceToken common.Address, provider common.Address) ([32]byte, error) {
	return _Lnv3Bridge.Contract.GetProviderStateKey(&_Lnv3Bridge.CallOpts, _sourceToken, provider)
}

// GetProviderStateKey is a free data retrieval call binding the contract method 0x2e1ab85f.
//
// Solidity: function getProviderStateKey(address _sourceToken, address provider) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) GetProviderStateKey(_sourceToken common.Address, provider common.Address) ([32]byte, error) {
	return _Lnv3Bridge.Contract.GetProviderStateKey(&_Lnv3Bridge.CallOpts, _sourceToken, provider)
}

// GetTokenKey is a free data retrieval call binding the contract method 0x926d1dff.
//
// Solidity: function getTokenKey(uint256 _remoteChainId, address _sourceToken, address _targetToken) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCaller) GetTokenKey(opts *bind.CallOpts, _remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "getTokenKey", _remoteChainId, _sourceToken, _targetToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTokenKey is a free data retrieval call binding the contract method 0x926d1dff.
//
// Solidity: function getTokenKey(uint256 _remoteChainId, address _sourceToken, address _targetToken) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeSession) GetTokenKey(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) ([32]byte, error) {
	return _Lnv3Bridge.Contract.GetTokenKey(&_Lnv3Bridge.CallOpts, _remoteChainId, _sourceToken, _targetToken)
}

// GetTokenKey is a free data retrieval call binding the contract method 0x926d1dff.
//
// Solidity: function getTokenKey(uint256 _remoteChainId, address _sourceToken, address _targetToken) pure returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) GetTokenKey(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) ([32]byte, error) {
	return _Lnv3Bridge.Contract.GetTokenKey(&_Lnv3Bridge.CallOpts, _remoteChainId, _sourceToken, _targetToken)
}

// GetTransferId is a free data retrieval call binding the contract method 0x0a838aec.
//
// Solidity: function getTransferId((uint256,address,address,address,uint112,uint112,address,uint256) _params, uint112 _remoteAmount) view returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCaller) GetTransferId(opts *bind.CallOpts, _params LnBridgeSourceV3TransferParams, _remoteAmount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "getTransferId", _params, _remoteAmount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTransferId is a free data retrieval call binding the contract method 0x0a838aec.
//
// Solidity: function getTransferId((uint256,address,address,address,uint112,uint112,address,uint256) _params, uint112 _remoteAmount) view returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeSession) GetTransferId(_params LnBridgeSourceV3TransferParams, _remoteAmount *big.Int) ([32]byte, error) {
	return _Lnv3Bridge.Contract.GetTransferId(&_Lnv3Bridge.CallOpts, _params, _remoteAmount)
}

// GetTransferId is a free data retrieval call binding the contract method 0x0a838aec.
//
// Solidity: function getTransferId((uint256,address,address,address,uint112,uint112,address,uint256) _params, uint112 _remoteAmount) view returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) GetTransferId(_params LnBridgeSourceV3TransferParams, _remoteAmount *big.Int) ([32]byte, error) {
	return _Lnv3Bridge.Contract.GetTransferId(&_Lnv3Bridge.CallOpts, _params, _remoteAmount)
}

// LockInfos is a free data retrieval call binding the contract method 0xf499b5c3.
//
// Solidity: function lockInfos(bytes32 ) view returns(uint112 amountWithFeeAndPenalty, uint32 tokenIndex, uint8 status)
func (_Lnv3Bridge *Lnv3BridgeCaller) LockInfos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	AmountWithFeeAndPenalty *big.Int
	TokenIndex              uint32
	Status                  uint8
}, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "lockInfos", arg0)

	outstruct := new(struct {
		AmountWithFeeAndPenalty *big.Int
		TokenIndex              uint32
		Status                  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountWithFeeAndPenalty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenIndex = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// LockInfos is a free data retrieval call binding the contract method 0xf499b5c3.
//
// Solidity: function lockInfos(bytes32 ) view returns(uint112 amountWithFeeAndPenalty, uint32 tokenIndex, uint8 status)
func (_Lnv3Bridge *Lnv3BridgeSession) LockInfos(arg0 [32]byte) (struct {
	AmountWithFeeAndPenalty *big.Int
	TokenIndex              uint32
	Status                  uint8
}, error) {
	return _Lnv3Bridge.Contract.LockInfos(&_Lnv3Bridge.CallOpts, arg0)
}

// LockInfos is a free data retrieval call binding the contract method 0xf499b5c3.
//
// Solidity: function lockInfos(bytes32 ) view returns(uint112 amountWithFeeAndPenalty, uint32 tokenIndex, uint8 status)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) LockInfos(arg0 [32]byte) (struct {
	AmountWithFeeAndPenalty *big.Int
	TokenIndex              uint32
	Status                  uint8
}, error) {
	return _Lnv3Bridge.Contract.LockInfos(&_Lnv3Bridge.CallOpts, arg0)
}

// Messagers is a free data retrieval call binding the contract method 0x99a1900f.
//
// Solidity: function messagers(uint256 ) view returns(address sendService, address receiveService)
func (_Lnv3Bridge *Lnv3BridgeCaller) Messagers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SendService    common.Address
	ReceiveService common.Address
}, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "messagers", arg0)

	outstruct := new(struct {
		SendService    common.Address
		ReceiveService common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SendService = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ReceiveService = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Messagers is a free data retrieval call binding the contract method 0x99a1900f.
//
// Solidity: function messagers(uint256 ) view returns(address sendService, address receiveService)
func (_Lnv3Bridge *Lnv3BridgeSession) Messagers(arg0 *big.Int) (struct {
	SendService    common.Address
	ReceiveService common.Address
}, error) {
	return _Lnv3Bridge.Contract.Messagers(&_Lnv3Bridge.CallOpts, arg0)
}

// Messagers is a free data retrieval call binding the contract method 0x99a1900f.
//
// Solidity: function messagers(uint256 ) view returns(address sendService, address receiveService)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) Messagers(arg0 *big.Int) (struct {
	SendService    common.Address
	ReceiveService common.Address
}, error) {
	return _Lnv3Bridge.Contract.Messagers(&_Lnv3Bridge.CallOpts, arg0)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeSession) Operator() (common.Address, error) {
	return _Lnv3Bridge.Contract.Operator(&_Lnv3Bridge.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) Operator() (common.Address, error) {
	return _Lnv3Bridge.Contract.Operator(&_Lnv3Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lnv3Bridge *Lnv3BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lnv3Bridge *Lnv3BridgeSession) Paused() (bool, error) {
	return _Lnv3Bridge.Contract.Paused(&_Lnv3Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) Paused() (bool, error) {
	return _Lnv3Bridge.Contract.Paused(&_Lnv3Bridge.CallOpts)
}

// PenaltyReserves is a free data retrieval call binding the contract method 0xd9f81ded.
//
// Solidity: function penaltyReserves(bytes32 ) view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCaller) PenaltyReserves(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "penaltyReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyReserves is a free data retrieval call binding the contract method 0xd9f81ded.
//
// Solidity: function penaltyReserves(bytes32 ) view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeSession) PenaltyReserves(arg0 [32]byte) (*big.Int, error) {
	return _Lnv3Bridge.Contract.PenaltyReserves(&_Lnv3Bridge.CallOpts, arg0)
}

// PenaltyReserves is a free data retrieval call binding the contract method 0xd9f81ded.
//
// Solidity: function penaltyReserves(bytes32 ) view returns(uint256)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) PenaltyReserves(arg0 [32]byte) (*big.Int, error) {
	return _Lnv3Bridge.Contract.PenaltyReserves(&_Lnv3Bridge.CallOpts, arg0)
}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeCaller) PendingDao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "pendingDao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeSession) PendingDao() (common.Address, error) {
	return _Lnv3Bridge.Contract.PendingDao(&_Lnv3Bridge.CallOpts)
}

// PendingDao is a free data retrieval call binding the contract method 0x67af1bdf.
//
// Solidity: function pendingDao() view returns(address)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) PendingDao() (common.Address, error) {
	return _Lnv3Bridge.Contract.PendingDao(&_Lnv3Bridge.CallOpts)
}

// SlashInfos is a free data retrieval call binding the contract method 0x453d1340.
//
// Solidity: function slashInfos(bytes32 ) view returns(uint256 remoteChainId, address slasher)
func (_Lnv3Bridge *Lnv3BridgeCaller) SlashInfos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	RemoteChainId *big.Int
	Slasher       common.Address
}, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "slashInfos", arg0)

	outstruct := new(struct {
		RemoteChainId *big.Int
		Slasher       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RemoteChainId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slasher = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SlashInfos is a free data retrieval call binding the contract method 0x453d1340.
//
// Solidity: function slashInfos(bytes32 ) view returns(uint256 remoteChainId, address slasher)
func (_Lnv3Bridge *Lnv3BridgeSession) SlashInfos(arg0 [32]byte) (struct {
	RemoteChainId *big.Int
	Slasher       common.Address
}, error) {
	return _Lnv3Bridge.Contract.SlashInfos(&_Lnv3Bridge.CallOpts, arg0)
}

// SlashInfos is a free data retrieval call binding the contract method 0x453d1340.
//
// Solidity: function slashInfos(bytes32 ) view returns(uint256 remoteChainId, address slasher)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) SlashInfos(arg0 [32]byte) (struct {
	RemoteChainId *big.Int
	Slasher       common.Address
}, error) {
	return _Lnv3Bridge.Contract.SlashInfos(&_Lnv3Bridge.CallOpts, arg0)
}

// SrcProviders is a free data retrieval call binding the contract method 0xd4377f1d.
//
// Solidity: function srcProviders(bytes32 ) view returns(uint112 baseFee, uint16 liquidityFeeRate, uint112 transferLimit, bool pause)
func (_Lnv3Bridge *Lnv3BridgeCaller) SrcProviders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	BaseFee          *big.Int
	LiquidityFeeRate uint16
	TransferLimit    *big.Int
	Pause            bool
}, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "srcProviders", arg0)

	outstruct := new(struct {
		BaseFee          *big.Int
		LiquidityFeeRate uint16
		TransferLimit    *big.Int
		Pause            bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidityFeeRate = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.TransferLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Pause = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// SrcProviders is a free data retrieval call binding the contract method 0xd4377f1d.
//
// Solidity: function srcProviders(bytes32 ) view returns(uint112 baseFee, uint16 liquidityFeeRate, uint112 transferLimit, bool pause)
func (_Lnv3Bridge *Lnv3BridgeSession) SrcProviders(arg0 [32]byte) (struct {
	BaseFee          *big.Int
	LiquidityFeeRate uint16
	TransferLimit    *big.Int
	Pause            bool
}, error) {
	return _Lnv3Bridge.Contract.SrcProviders(&_Lnv3Bridge.CallOpts, arg0)
}

// SrcProviders is a free data retrieval call binding the contract method 0xd4377f1d.
//
// Solidity: function srcProviders(bytes32 ) view returns(uint112 baseFee, uint16 liquidityFeeRate, uint112 transferLimit, bool pause)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) SrcProviders(arg0 [32]byte) (struct {
	BaseFee          *big.Int
	LiquidityFeeRate uint16
	TransferLimit    *big.Int
	Pause            bool
}, error) {
	return _Lnv3Bridge.Contract.SrcProviders(&_Lnv3Bridge.CallOpts, arg0)
}

// TokenIndexer is a free data retrieval call binding the contract method 0x89e46c81.
//
// Solidity: function tokenIndexer(uint32 ) view returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCaller) TokenIndexer(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "tokenIndexer", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenIndexer is a free data retrieval call binding the contract method 0x89e46c81.
//
// Solidity: function tokenIndexer(uint32 ) view returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeSession) TokenIndexer(arg0 uint32) ([32]byte, error) {
	return _Lnv3Bridge.Contract.TokenIndexer(&_Lnv3Bridge.CallOpts, arg0)
}

// TokenIndexer is a free data retrieval call binding the contract method 0x89e46c81.
//
// Solidity: function tokenIndexer(uint32 ) view returns(bytes32)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) TokenIndexer(arg0 uint32) ([32]byte, error) {
	return _Lnv3Bridge.Contract.TokenIndexer(&_Lnv3Bridge.CallOpts, arg0)
}

// TokenInfos is a free data retrieval call binding the contract method 0x90f1f17b.
//
// Solidity: function tokenInfos(bytes32 ) view returns((uint112,uint112,uint8,uint8) config, uint32 index, address sourceToken, address targetToken, uint256 protocolFeeIncome)
func (_Lnv3Bridge *Lnv3BridgeCaller) TokenInfos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Config            LnBridgeSourceV3TokenConfigure
	Index             uint32
	SourceToken       common.Address
	TargetToken       common.Address
	ProtocolFeeIncome *big.Int
}, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "tokenInfos", arg0)

	outstruct := new(struct {
		Config LnBridgeSourceV3TokenConfigure
		Index  uint32
		SourceToken       common.Address
		TargetToken       common.Address
		ProtocolFeeIncome *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Config = *abi.ConvertType(out[0], new(LnBridgeSourceV3TokenConfigure)).(*LnBridgeSourceV3TokenConfigure)
	outstruct.Index = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.SourceToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TargetToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ProtocolFeeIncome = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenInfos is a free data retrieval call binding the contract method 0x90f1f17b.
//
// Solidity: function tokenInfos(bytes32 ) view returns((uint112,uint112,uint8,uint8) config, uint32 index, address sourceToken, address targetToken, uint256 protocolFeeIncome)
func (_Lnv3Bridge *Lnv3BridgeSession) TokenInfos(arg0 [32]byte) (struct {
	Config            LnBridgeSourceV3TokenConfigure
	Index             uint32
	SourceToken       common.Address
	TargetToken       common.Address
	ProtocolFeeIncome *big.Int
}, error) {
	return _Lnv3Bridge.Contract.TokenInfos(&_Lnv3Bridge.CallOpts, arg0)
}

// TokenInfos is a free data retrieval call binding the contract method 0x90f1f17b.
//
// Solidity: function tokenInfos(bytes32 ) view returns((uint112,uint112,uint8,uint8) config, uint32 index, address sourceToken, address targetToken, uint256 protocolFeeIncome)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) TokenInfos(arg0 [32]byte) (struct {
	Config            LnBridgeSourceV3TokenConfigure
	Index             uint32
	SourceToken       common.Address
	TargetToken       common.Address
	ProtocolFeeIncome *big.Int
}, error) {
	return _Lnv3Bridge.Contract.TokenInfos(&_Lnv3Bridge.CallOpts, arg0)
}

// TotalFee is a free data retrieval call binding the contract method 0xc3ac5e40.
//
// Solidity: function totalFee(uint256 _remoteChainId, address _provider, address _sourceToken, address _targetToken, uint112 _amount) view returns(uint112)
func (_Lnv3Bridge *Lnv3BridgeCaller) TotalFee(opts *bind.CallOpts, _remoteChainId *big.Int, _provider common.Address, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lnv3Bridge.contract.Call(opts, &out, "totalFee", _remoteChainId, _provider, _sourceToken, _targetToken, _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalFee is a free data retrieval call binding the contract method 0xc3ac5e40.
//
// Solidity: function totalFee(uint256 _remoteChainId, address _provider, address _sourceToken, address _targetToken, uint112 _amount) view returns(uint112)
func (_Lnv3Bridge *Lnv3BridgeSession) TotalFee(_remoteChainId *big.Int, _provider common.Address, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, error) {
	return _Lnv3Bridge.Contract.TotalFee(&_Lnv3Bridge.CallOpts, _remoteChainId, _provider, _sourceToken, _targetToken, _amount)
}

// TotalFee is a free data retrieval call binding the contract method 0xc3ac5e40.
//
// Solidity: function totalFee(uint256 _remoteChainId, address _provider, address _sourceToken, address _targetToken, uint112 _amount) view returns(uint112)
func (_Lnv3Bridge *Lnv3BridgeCallerSession) TotalFee(_remoteChainId *big.Int, _provider common.Address, _sourceToken common.Address, _targetToken common.Address, _amount *big.Int) (*big.Int, error) {
	return _Lnv3Bridge.Contract.TotalFee(&_Lnv3Bridge.CallOpts, _remoteChainId, _provider, _sourceToken, _targetToken, _amount)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Lnv3Bridge *Lnv3BridgeSession) AcceptOwnership() (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.AcceptOwnership(&_Lnv3Bridge.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.AcceptOwnership(&_Lnv3Bridge.TransactOpts)
}

// ClaimProtocolFeeIncome is a paid mutator transaction binding the contract method 0x8dcbec8e.
//
// Solidity: function claimProtocolFeeIncome(bytes32 _tokenInfoKey, uint256 _amount, address _receiver) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) ClaimProtocolFeeIncome(opts *bind.TransactOpts, _tokenInfoKey [32]byte, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "claimProtocolFeeIncome", _tokenInfoKey, _amount, _receiver)
}

// ClaimProtocolFeeIncome is a paid mutator transaction binding the contract method 0x8dcbec8e.
//
// Solidity: function claimProtocolFeeIncome(bytes32 _tokenInfoKey, uint256 _amount, address _receiver) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) ClaimProtocolFeeIncome(_tokenInfoKey [32]byte, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.ClaimProtocolFeeIncome(&_Lnv3Bridge.TransactOpts, _tokenInfoKey, _amount, _receiver)
}

// ClaimProtocolFeeIncome is a paid mutator transaction binding the contract method 0x8dcbec8e.
//
// Solidity: function claimProtocolFeeIncome(bytes32 _tokenInfoKey, uint256 _amount, address _receiver) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) ClaimProtocolFeeIncome(_tokenInfoKey [32]byte, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.ClaimProtocolFeeIncome(&_Lnv3Bridge.TransactOpts, _tokenInfoKey, _amount, _receiver)
}

// DeleteTokenInfo is a paid mutator transaction binding the contract method 0xbc15fba5.
//
// Solidity: function deleteTokenInfo(bytes32 key) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) DeleteTokenInfo(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "deleteTokenInfo", key)
}

// DeleteTokenInfo is a paid mutator transaction binding the contract method 0xbc15fba5.
//
// Solidity: function deleteTokenInfo(bytes32 key) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) DeleteTokenInfo(key [32]byte) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.DeleteTokenInfo(&_Lnv3Bridge.TransactOpts, key)
}

// DeleteTokenInfo is a paid mutator transaction binding the contract method 0xbc15fba5.
//
// Solidity: function deleteTokenInfo(bytes32 key) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) DeleteTokenInfo(key [32]byte) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.DeleteTokenInfo(&_Lnv3Bridge.TransactOpts, key)
}

// DepositPenaltyReserve is a paid mutator transaction binding the contract method 0x2fd5da4f.
//
// Solidity: function depositPenaltyReserve(address _sourceToken, uint256 _amount) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) DepositPenaltyReserve(opts *bind.TransactOpts, _sourceToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "depositPenaltyReserve", _sourceToken, _amount)
}

// DepositPenaltyReserve is a paid mutator transaction binding the contract method 0x2fd5da4f.
//
// Solidity: function depositPenaltyReserve(address _sourceToken, uint256 _amount) payable returns()
func (_Lnv3Bridge *Lnv3BridgeSession) DepositPenaltyReserve(_sourceToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.DepositPenaltyReserve(&_Lnv3Bridge.TransactOpts, _sourceToken, _amount)
}

// DepositPenaltyReserve is a paid mutator transaction binding the contract method 0x2fd5da4f.
//
// Solidity: function depositPenaltyReserve(address _sourceToken, uint256 _amount) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) DepositPenaltyReserve(_sourceToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.DepositPenaltyReserve(&_Lnv3Bridge.TransactOpts, _sourceToken, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address dao) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) Initialize(opts *bind.TransactOpts, dao common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "initialize", dao)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address dao) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) Initialize(dao common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Initialize(&_Lnv3Bridge.TransactOpts, dao)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address dao) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) Initialize(dao common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Initialize(&_Lnv3Bridge.TransactOpts, dao)
}

// LockAndRemoteRelease is a paid mutator transaction binding the contract method 0x9cd13471.
//
// Solidity: function lockAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) LockAndRemoteRelease(opts *bind.TransactOpts, _params LnBridgeSourceV3TransferParams) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "lockAndRemoteRelease", _params)
}

// LockAndRemoteRelease is a paid mutator transaction binding the contract method 0x9cd13471.
//
// Solidity: function lockAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params) payable returns()
func (_Lnv3Bridge *Lnv3BridgeSession) LockAndRemoteRelease(_params LnBridgeSourceV3TransferParams) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.LockAndRemoteRelease(&_Lnv3Bridge.TransactOpts, _params)
}

// LockAndRemoteRelease is a paid mutator transaction binding the contract method 0x9cd13471.
//
// Solidity: function lockAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) LockAndRemoteRelease(_params LnBridgeSourceV3TransferParams) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.LockAndRemoteRelease(&_Lnv3Bridge.TransactOpts, _params)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lnv3Bridge *Lnv3BridgeSession) Pause() (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Pause(&_Lnv3Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Pause(&_Lnv3Bridge.TransactOpts)
}

// ProviderPause is a paid mutator transaction binding the contract method 0x484763e6.
//
// Solidity: function providerPause(uint256 _remoteChainId, address _sourceToken, address _targetToken) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) ProviderPause(opts *bind.TransactOpts, _remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "providerPause", _remoteChainId, _sourceToken, _targetToken)
}

// ProviderPause is a paid mutator transaction binding the contract method 0x484763e6.
//
// Solidity: function providerPause(uint256 _remoteChainId, address _sourceToken, address _targetToken) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) ProviderPause(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.ProviderPause(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken)
}

// ProviderPause is a paid mutator transaction binding the contract method 0x484763e6.
//
// Solidity: function providerPause(uint256 _remoteChainId, address _sourceToken, address _targetToken) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) ProviderPause(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.ProviderPause(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken)
}

// ProviderUnpause is a paid mutator transaction binding the contract method 0xf67caca4.
//
// Solidity: function providerUnpause(uint256 _remoteChainId, address _sourceToken, address _targetToken) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) ProviderUnpause(opts *bind.TransactOpts, _remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "providerUnpause", _remoteChainId, _sourceToken, _targetToken)
}

// ProviderUnpause is a paid mutator transaction binding the contract method 0xf67caca4.
//
// Solidity: function providerUnpause(uint256 _remoteChainId, address _sourceToken, address _targetToken) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) ProviderUnpause(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.ProviderUnpause(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken)
}

// ProviderUnpause is a paid mutator transaction binding the contract method 0xf67caca4.
//
// Solidity: function providerUnpause(uint256 _remoteChainId, address _sourceToken, address _targetToken) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) ProviderUnpause(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.ProviderUnpause(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken)
}

// RegisterLnProvider is a paid mutator transaction binding the contract method 0x48faf229.
//
// Solidity: function registerLnProvider(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _baseFee, uint16 _liquidityFeeRate, uint112 _transferLimit) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) RegisterLnProvider(opts *bind.TransactOpts, _remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _baseFee *big.Int, _liquidityFeeRate uint16, _transferLimit *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "registerLnProvider", _remoteChainId, _sourceToken, _targetToken, _baseFee, _liquidityFeeRate, _transferLimit)
}

// RegisterLnProvider is a paid mutator transaction binding the contract method 0x48faf229.
//
// Solidity: function registerLnProvider(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _baseFee, uint16 _liquidityFeeRate, uint112 _transferLimit) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) RegisterLnProvider(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _baseFee *big.Int, _liquidityFeeRate uint16, _transferLimit *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RegisterLnProvider(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken, _baseFee, _liquidityFeeRate, _transferLimit)
}

// RegisterLnProvider is a paid mutator transaction binding the contract method 0x48faf229.
//
// Solidity: function registerLnProvider(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _baseFee, uint16 _liquidityFeeRate, uint112 _transferLimit) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) RegisterLnProvider(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _baseFee *big.Int, _liquidityFeeRate uint16, _transferLimit *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RegisterLnProvider(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken, _baseFee, _liquidityFeeRate, _transferLimit)
}

// RegisterTokenInfo is a paid mutator transaction binding the contract method 0x293a3c8a.
//
// Solidity: function registerTokenInfo(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _protocolFee, uint112 _penalty, uint8 _sourceDecimals, uint8 _targetDecimals, uint32 _index) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) RegisterTokenInfo(opts *bind.TransactOpts, _remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _protocolFee *big.Int, _penalty *big.Int, _sourceDecimals uint8, _targetDecimals uint8, _index uint32) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "registerTokenInfo", _remoteChainId, _sourceToken, _targetToken, _protocolFee, _penalty, _sourceDecimals, _targetDecimals, _index)
}

// RegisterTokenInfo is a paid mutator transaction binding the contract method 0x293a3c8a.
//
// Solidity: function registerTokenInfo(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _protocolFee, uint112 _penalty, uint8 _sourceDecimals, uint8 _targetDecimals, uint32 _index) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) RegisterTokenInfo(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _protocolFee *big.Int, _penalty *big.Int, _sourceDecimals uint8, _targetDecimals uint8, _index uint32) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RegisterTokenInfo(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken, _protocolFee, _penalty, _sourceDecimals, _targetDecimals, _index)
}

// RegisterTokenInfo is a paid mutator transaction binding the contract method 0x293a3c8a.
//
// Solidity: function registerTokenInfo(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _protocolFee, uint112 _penalty, uint8 _sourceDecimals, uint8 _targetDecimals, uint32 _index) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) RegisterTokenInfo(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _protocolFee *big.Int, _penalty *big.Int, _sourceDecimals uint8, _targetDecimals uint8, _index uint32) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RegisterTokenInfo(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken, _protocolFee, _penalty, _sourceDecimals, _targetDecimals, _index)
}

// Relay is a paid mutator transaction binding the contract method 0x23744da5.
//
// Solidity: function relay((uint256,address,address,address,uint112,uint112,address,uint256) _params, bytes32 _expectedTransferId, bool _relayBySelf) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) Relay(opts *bind.TransactOpts, _params LnBridgeTargetV3RelayParams, _expectedTransferId [32]byte, _relayBySelf bool) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "relay", _params, _expectedTransferId, _relayBySelf)
}

// Relay is a paid mutator transaction binding the contract method 0x23744da5.
//
// Solidity: function relay((uint256,address,address,address,uint112,uint112,address,uint256) _params, bytes32 _expectedTransferId, bool _relayBySelf) payable returns()
func (_Lnv3Bridge *Lnv3BridgeSession) Relay(_params LnBridgeTargetV3RelayParams, _expectedTransferId [32]byte, _relayBySelf bool) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Relay(&_Lnv3Bridge.TransactOpts, _params, _expectedTransferId, _relayBySelf)
}

// Relay is a paid mutator transaction binding the contract method 0x23744da5.
//
// Solidity: function relay((uint256,address,address,address,uint112,uint112,address,uint256) _params, bytes32 _expectedTransferId, bool _relayBySelf) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) Relay(_params LnBridgeTargetV3RelayParams, _expectedTransferId [32]byte, _relayBySelf bool) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Relay(&_Lnv3Bridge.TransactOpts, _params, _expectedTransferId, _relayBySelf)
}

// RequestSlashAndRemoteRelease is a paid mutator transaction binding the contract method 0x58b8ff52.
//
// Solidity: function requestSlashAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params, bytes32 _expectedTransferId, uint256 _feePrepaid, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) RequestSlashAndRemoteRelease(opts *bind.TransactOpts, _params LnBridgeTargetV3RelayParams, _expectedTransferId [32]byte, _feePrepaid *big.Int, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "requestSlashAndRemoteRelease", _params, _expectedTransferId, _feePrepaid, _extParams)
}

// RequestSlashAndRemoteRelease is a paid mutator transaction binding the contract method 0x58b8ff52.
//
// Solidity: function requestSlashAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params, bytes32 _expectedTransferId, uint256 _feePrepaid, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeSession) RequestSlashAndRemoteRelease(_params LnBridgeTargetV3RelayParams, _expectedTransferId [32]byte, _feePrepaid *big.Int, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RequestSlashAndRemoteRelease(&_Lnv3Bridge.TransactOpts, _params, _expectedTransferId, _feePrepaid, _extParams)
}

// RequestSlashAndRemoteRelease is a paid mutator transaction binding the contract method 0x58b8ff52.
//
// Solidity: function requestSlashAndRemoteRelease((uint256,address,address,address,uint112,uint112,address,uint256) _params, bytes32 _expectedTransferId, uint256 _feePrepaid, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) RequestSlashAndRemoteRelease(_params LnBridgeTargetV3RelayParams, _expectedTransferId [32]byte, _feePrepaid *big.Int, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RequestSlashAndRemoteRelease(&_Lnv3Bridge.TransactOpts, _params, _expectedTransferId, _feePrepaid, _extParams)
}

// RequestWithdrawLiquidity is a paid mutator transaction binding the contract method 0x29f085f4.
//
// Solidity: function requestWithdrawLiquidity(uint256 _remoteChainId, bytes32[] _transferIds, address _provider, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) RequestWithdrawLiquidity(opts *bind.TransactOpts, _remoteChainId *big.Int, _transferIds [][32]byte, _provider common.Address, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "requestWithdrawLiquidity", _remoteChainId, _transferIds, _provider, _extParams)
}

// RequestWithdrawLiquidity is a paid mutator transaction binding the contract method 0x29f085f4.
//
// Solidity: function requestWithdrawLiquidity(uint256 _remoteChainId, bytes32[] _transferIds, address _provider, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeSession) RequestWithdrawLiquidity(_remoteChainId *big.Int, _transferIds [][32]byte, _provider common.Address, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RequestWithdrawLiquidity(&_Lnv3Bridge.TransactOpts, _remoteChainId, _transferIds, _provider, _extParams)
}

// RequestWithdrawLiquidity is a paid mutator transaction binding the contract method 0x29f085f4.
//
// Solidity: function requestWithdrawLiquidity(uint256 _remoteChainId, bytes32[] _transferIds, address _provider, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) RequestWithdrawLiquidity(_remoteChainId *big.Int, _transferIds [][32]byte, _provider common.Address, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RequestWithdrawLiquidity(&_Lnv3Bridge.TransactOpts, _remoteChainId, _transferIds, _provider, _extParams)
}

// RetrySlash is a paid mutator transaction binding the contract method 0x8d6a62e9.
//
// Solidity: function retrySlash(bytes32 transferId, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) RetrySlash(opts *bind.TransactOpts, transferId [32]byte, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "retrySlash", transferId, _extParams)
}

// RetrySlash is a paid mutator transaction binding the contract method 0x8d6a62e9.
//
// Solidity: function retrySlash(bytes32 transferId, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeSession) RetrySlash(transferId [32]byte, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RetrySlash(&_Lnv3Bridge.TransactOpts, transferId, _extParams)
}

// RetrySlash is a paid mutator transaction binding the contract method 0x8d6a62e9.
//
// Solidity: function retrySlash(bytes32 transferId, bytes _extParams) payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) RetrySlash(transferId [32]byte, _extParams []byte) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.RetrySlash(&_Lnv3Bridge.TransactOpts, transferId, _extParams)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) SetOperator(opts *bind.TransactOpts, _operator common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "setOperator", _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.SetOperator(&_Lnv3Bridge.TransactOpts, _operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _operator) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) SetOperator(_operator common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.SetOperator(&_Lnv3Bridge.TransactOpts, _operator)
}

// SetReceiveService is a paid mutator transaction binding the contract method 0x8c4cff20.
//
// Solidity: function setReceiveService(uint256 _remoteChainId, address _remoteBridge, address _service) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) SetReceiveService(opts *bind.TransactOpts, _remoteChainId *big.Int, _remoteBridge common.Address, _service common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "setReceiveService", _remoteChainId, _remoteBridge, _service)
}

// SetReceiveService is a paid mutator transaction binding the contract method 0x8c4cff20.
//
// Solidity: function setReceiveService(uint256 _remoteChainId, address _remoteBridge, address _service) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) SetReceiveService(_remoteChainId *big.Int, _remoteBridge common.Address, _service common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.SetReceiveService(&_Lnv3Bridge.TransactOpts, _remoteChainId, _remoteBridge, _service)
}

// SetReceiveService is a paid mutator transaction binding the contract method 0x8c4cff20.
//
// Solidity: function setReceiveService(uint256 _remoteChainId, address _remoteBridge, address _service) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) SetReceiveService(_remoteChainId *big.Int, _remoteBridge common.Address, _service common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.SetReceiveService(&_Lnv3Bridge.TransactOpts, _remoteChainId, _remoteBridge, _service)
}

// SetSendService is a paid mutator transaction binding the contract method 0x2fe5718b.
//
// Solidity: function setSendService(uint256 _remoteChainId, address _remoteBridge, address _service) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) SetSendService(opts *bind.TransactOpts, _remoteChainId *big.Int, _remoteBridge common.Address, _service common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "setSendService", _remoteChainId, _remoteBridge, _service)
}

// SetSendService is a paid mutator transaction binding the contract method 0x2fe5718b.
//
// Solidity: function setSendService(uint256 _remoteChainId, address _remoteBridge, address _service) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) SetSendService(_remoteChainId *big.Int, _remoteBridge common.Address, _service common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.SetSendService(&_Lnv3Bridge.TransactOpts, _remoteChainId, _remoteBridge, _service)
}

// SetSendService is a paid mutator transaction binding the contract method 0x2fe5718b.
//
// Solidity: function setSendService(uint256 _remoteChainId, address _remoteBridge, address _service) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) SetSendService(_remoteChainId *big.Int, _remoteBridge common.Address, _service common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.SetSendService(&_Lnv3Bridge.TransactOpts, _remoteChainId, _remoteBridge, _service)
}

// Slash is a paid mutator transaction binding the contract method 0xc0474fba.
//
// Solidity: function slash(uint256 _remoteChainId, bytes32 _transferId, address _lnProvider, address _slasher) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) Slash(opts *bind.TransactOpts, _remoteChainId *big.Int, _transferId [32]byte, _lnProvider common.Address, _slasher common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "slash", _remoteChainId, _transferId, _lnProvider, _slasher)
}

// Slash is a paid mutator transaction binding the contract method 0xc0474fba.
//
// Solidity: function slash(uint256 _remoteChainId, bytes32 _transferId, address _lnProvider, address _slasher) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) Slash(_remoteChainId *big.Int, _transferId [32]byte, _lnProvider common.Address, _slasher common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Slash(&_Lnv3Bridge.TransactOpts, _remoteChainId, _transferId, _lnProvider, _slasher)
}

// Slash is a paid mutator transaction binding the contract method 0xc0474fba.
//
// Solidity: function slash(uint256 _remoteChainId, bytes32 _transferId, address _lnProvider, address _slasher) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) Slash(_remoteChainId *big.Int, _transferId [32]byte, _lnProvider common.Address, _slasher common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Slash(&_Lnv3Bridge.TransactOpts, _remoteChainId, _transferId, _lnProvider, _slasher)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) TransferOwnership(opts *bind.TransactOpts, _dao common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "transferOwnership", _dao)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) TransferOwnership(_dao common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.TransferOwnership(&_Lnv3Bridge.TransactOpts, _dao)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _dao) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) TransferOwnership(_dao common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.TransferOwnership(&_Lnv3Bridge.TransactOpts, _dao)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lnv3Bridge *Lnv3BridgeSession) Unpause() (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Unpause(&_Lnv3Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Unpause(&_Lnv3Bridge.TransactOpts)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x88396e34.
//
// Solidity: function updateTokenInfo(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _protocolFee, uint112 _penalty, uint8 _sourceDecimals, uint8 _targetDecimals) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) UpdateTokenInfo(opts *bind.TransactOpts, _remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _protocolFee *big.Int, _penalty *big.Int, _sourceDecimals uint8, _targetDecimals uint8) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "updateTokenInfo", _remoteChainId, _sourceToken, _targetToken, _protocolFee, _penalty, _sourceDecimals, _targetDecimals)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x88396e34.
//
// Solidity: function updateTokenInfo(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _protocolFee, uint112 _penalty, uint8 _sourceDecimals, uint8 _targetDecimals) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) UpdateTokenInfo(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _protocolFee *big.Int, _penalty *big.Int, _sourceDecimals uint8, _targetDecimals uint8) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.UpdateTokenInfo(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken, _protocolFee, _penalty, _sourceDecimals, _targetDecimals)
}

// UpdateTokenInfo is a paid mutator transaction binding the contract method 0x88396e34.
//
// Solidity: function updateTokenInfo(uint256 _remoteChainId, address _sourceToken, address _targetToken, uint112 _protocolFee, uint112 _penalty, uint8 _sourceDecimals, uint8 _targetDecimals) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) UpdateTokenInfo(_remoteChainId *big.Int, _sourceToken common.Address, _targetToken common.Address, _protocolFee *big.Int, _penalty *big.Int, _sourceDecimals uint8, _targetDecimals uint8) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.UpdateTokenInfo(&_Lnv3Bridge.TransactOpts, _remoteChainId, _sourceToken, _targetToken, _protocolFee, _penalty, _sourceDecimals, _targetDecimals)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x7425b8b5.
//
// Solidity: function withdrawLiquidity(bytes32[] _transferIds, uint256 _remoteChainId, address _provider) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) WithdrawLiquidity(opts *bind.TransactOpts, _transferIds [][32]byte, _remoteChainId *big.Int, _provider common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "withdrawLiquidity", _transferIds, _remoteChainId, _provider)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x7425b8b5.
//
// Solidity: function withdrawLiquidity(bytes32[] _transferIds, uint256 _remoteChainId, address _provider) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) WithdrawLiquidity(_transferIds [][32]byte, _remoteChainId *big.Int, _provider common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.WithdrawLiquidity(&_Lnv3Bridge.TransactOpts, _transferIds, _remoteChainId, _provider)
}

// WithdrawLiquidity is a paid mutator transaction binding the contract method 0x7425b8b5.
//
// Solidity: function withdrawLiquidity(bytes32[] _transferIds, uint256 _remoteChainId, address _provider) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) WithdrawLiquidity(_transferIds [][32]byte, _remoteChainId *big.Int, _provider common.Address) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.WithdrawLiquidity(&_Lnv3Bridge.TransactOpts, _transferIds, _remoteChainId, _provider)
}

// WithdrawPenaltyReserve is a paid mutator transaction binding the contract method 0xd6a57514.
//
// Solidity: function withdrawPenaltyReserve(address _sourceToken, uint256 _amount) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) WithdrawPenaltyReserve(opts *bind.TransactOpts, _sourceToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.Transact(opts, "withdrawPenaltyReserve", _sourceToken, _amount)
}

// WithdrawPenaltyReserve is a paid mutator transaction binding the contract method 0xd6a57514.
//
// Solidity: function withdrawPenaltyReserve(address _sourceToken, uint256 _amount) returns()
func (_Lnv3Bridge *Lnv3BridgeSession) WithdrawPenaltyReserve(_sourceToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.WithdrawPenaltyReserve(&_Lnv3Bridge.TransactOpts, _sourceToken, _amount)
}

// WithdrawPenaltyReserve is a paid mutator transaction binding the contract method 0xd6a57514.
//
// Solidity: function withdrawPenaltyReserve(address _sourceToken, uint256 _amount) returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) WithdrawPenaltyReserve(_sourceToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.WithdrawPenaltyReserve(&_Lnv3Bridge.TransactOpts, _sourceToken, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lnv3Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lnv3Bridge *Lnv3BridgeSession) Receive() (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Receive(&_Lnv3Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Lnv3Bridge *Lnv3BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Lnv3Bridge.Contract.Receive(&_Lnv3Bridge.TransactOpts)
}

// Lnv3BridgeFeeIncomeClaimedIterator is returned from FilterFeeIncomeClaimed and is used to iterate over the raw logs and unpacked data for FeeIncomeClaimed events raised by the Lnv3Bridge contract.
type Lnv3BridgeFeeIncomeClaimedIterator struct {
	Event *Lnv3BridgeFeeIncomeClaimed // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeFeeIncomeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeFeeIncomeClaimed)
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
		it.Event = new(Lnv3BridgeFeeIncomeClaimed)
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
func (it *Lnv3BridgeFeeIncomeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeFeeIncomeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeFeeIncomeClaimed represents a FeeIncomeClaimed event raised by the Lnv3Bridge contract.
type Lnv3BridgeFeeIncomeClaimed struct {
	TokenInfoKey [32]byte
	Amount       *big.Int
	Receiver     common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeIncomeClaimed is a free log retrieval operation binding the contract event 0x175d2cfc854b7a48ec725e3d265a97656e5ea13c906c7dcdd112686f7c508386.
//
// Solidity: event FeeIncomeClaimed(bytes32 tokenInfoKey, uint256 amount, address receiver)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterFeeIncomeClaimed(opts *bind.FilterOpts) (*Lnv3BridgeFeeIncomeClaimedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "FeeIncomeClaimed")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeFeeIncomeClaimedIterator{contract: _Lnv3Bridge.contract, event: "FeeIncomeClaimed", logs: logs, sub: sub}, nil
}

// WatchFeeIncomeClaimed is a free log subscription operation binding the contract event 0x175d2cfc854b7a48ec725e3d265a97656e5ea13c906c7dcdd112686f7c508386.
//
// Solidity: event FeeIncomeClaimed(bytes32 tokenInfoKey, uint256 amount, address receiver)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchFeeIncomeClaimed(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeFeeIncomeClaimed) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "FeeIncomeClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeFeeIncomeClaimed)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "FeeIncomeClaimed", log); err != nil {
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

// ParseFeeIncomeClaimed is a log parse operation binding the contract event 0x175d2cfc854b7a48ec725e3d265a97656e5ea13c906c7dcdd112686f7c508386.
//
// Solidity: event FeeIncomeClaimed(bytes32 tokenInfoKey, uint256 amount, address receiver)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseFeeIncomeClaimed(log types.Log) (*Lnv3BridgeFeeIncomeClaimed, error) {
	event := new(Lnv3BridgeFeeIncomeClaimed)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "FeeIncomeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Lnv3Bridge contract.
type Lnv3BridgeInitializedIterator struct {
	Event *Lnv3BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeInitialized)
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
		it.Event = new(Lnv3BridgeInitialized)
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
func (it *Lnv3BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeInitialized represents a Initialized event raised by the Lnv3Bridge contract.
type Lnv3BridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*Lnv3BridgeInitializedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeInitializedIterator{contract: _Lnv3Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeInitialized)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseInitialized(log types.Log) (*Lnv3BridgeInitialized, error) {
	event := new(Lnv3BridgeInitialized)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeLiquidityWithdrawnIterator is returned from FilterLiquidityWithdrawn and is used to iterate over the raw logs and unpacked data for LiquidityWithdrawn events raised by the Lnv3Bridge contract.
type Lnv3BridgeLiquidityWithdrawnIterator struct {
	Event *Lnv3BridgeLiquidityWithdrawn // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeLiquidityWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeLiquidityWithdrawn)
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
		it.Event = new(Lnv3BridgeLiquidityWithdrawn)
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
func (it *Lnv3BridgeLiquidityWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeLiquidityWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeLiquidityWithdrawn represents a LiquidityWithdrawn event raised by the Lnv3Bridge contract.
type Lnv3BridgeLiquidityWithdrawn struct {
	TransferIds [][32]byte
	Provider    common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLiquidityWithdrawn is a free log retrieval operation binding the contract event 0xbee375a5295c9c5af94066364af2e5ed7767c74915a74c5bce99b21ba4bdc1ed.
//
// Solidity: event LiquidityWithdrawn(bytes32[] transferIds, address provider, uint256 amount)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterLiquidityWithdrawn(opts *bind.FilterOpts) (*Lnv3BridgeLiquidityWithdrawnIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "LiquidityWithdrawn")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeLiquidityWithdrawnIterator{contract: _Lnv3Bridge.contract, event: "LiquidityWithdrawn", logs: logs, sub: sub}, nil
}

// WatchLiquidityWithdrawn is a free log subscription operation binding the contract event 0xbee375a5295c9c5af94066364af2e5ed7767c74915a74c5bce99b21ba4bdc1ed.
//
// Solidity: event LiquidityWithdrawn(bytes32[] transferIds, address provider, uint256 amount)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchLiquidityWithdrawn(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeLiquidityWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "LiquidityWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeLiquidityWithdrawn)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "LiquidityWithdrawn", log); err != nil {
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

// ParseLiquidityWithdrawn is a log parse operation binding the contract event 0xbee375a5295c9c5af94066364af2e5ed7767c74915a74c5bce99b21ba4bdc1ed.
//
// Solidity: event LiquidityWithdrawn(bytes32[] transferIds, address provider, uint256 amount)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseLiquidityWithdrawn(log types.Log) (*Lnv3BridgeLiquidityWithdrawn, error) {
	event := new(Lnv3BridgeLiquidityWithdrawn)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "LiquidityWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeLnProviderPausedIterator is returned from FilterLnProviderPaused and is used to iterate over the raw logs and unpacked data for LnProviderPaused events raised by the Lnv3Bridge contract.
type Lnv3BridgeLnProviderPausedIterator struct {
	Event *Lnv3BridgeLnProviderPaused // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeLnProviderPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeLnProviderPaused)
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
		it.Event = new(Lnv3BridgeLnProviderPaused)
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
func (it *Lnv3BridgeLnProviderPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeLnProviderPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeLnProviderPaused represents a LnProviderPaused event raised by the Lnv3Bridge contract.
type Lnv3BridgeLnProviderPaused struct {
	Provider      common.Address
	RemoteChainId *big.Int
	SourceToken   common.Address
	TargetToken   common.Address
	Paused        bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLnProviderPaused is a free log retrieval operation binding the contract event 0x357026723683cd411410642e5879fb1965366fd656c8ca7448d84fc3a082b42e.
//
// Solidity: event LnProviderPaused(address provider, uint256 remoteChainId, address sourceToken, address targetToken, bool paused)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterLnProviderPaused(opts *bind.FilterOpts) (*Lnv3BridgeLnProviderPausedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "LnProviderPaused")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeLnProviderPausedIterator{contract: _Lnv3Bridge.contract, event: "LnProviderPaused", logs: logs, sub: sub}, nil
}

// WatchLnProviderPaused is a free log subscription operation binding the contract event 0x357026723683cd411410642e5879fb1965366fd656c8ca7448d84fc3a082b42e.
//
// Solidity: event LnProviderPaused(address provider, uint256 remoteChainId, address sourceToken, address targetToken, bool paused)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchLnProviderPaused(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeLnProviderPaused) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "LnProviderPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeLnProviderPaused)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "LnProviderPaused", log); err != nil {
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

// ParseLnProviderPaused is a log parse operation binding the contract event 0x357026723683cd411410642e5879fb1965366fd656c8ca7448d84fc3a082b42e.
//
// Solidity: event LnProviderPaused(address provider, uint256 remoteChainId, address sourceToken, address targetToken, bool paused)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseLnProviderPaused(log types.Log) (*Lnv3BridgeLnProviderPaused, error) {
	event := new(Lnv3BridgeLnProviderPaused)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "LnProviderPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeLnProviderUpdatedIterator is returned from FilterLnProviderUpdated and is used to iterate over the raw logs and unpacked data for LnProviderUpdated events raised by the Lnv3Bridge contract.
type Lnv3BridgeLnProviderUpdatedIterator struct {
	Event *Lnv3BridgeLnProviderUpdated // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeLnProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeLnProviderUpdated)
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
		it.Event = new(Lnv3BridgeLnProviderUpdated)
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
func (it *Lnv3BridgeLnProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeLnProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeLnProviderUpdated represents a LnProviderUpdated event raised by the Lnv3Bridge contract.
type Lnv3BridgeLnProviderUpdated struct {
	RemoteChainId    *big.Int
	Provider         common.Address
	SourceToken      common.Address
	TargetToken      common.Address
	BaseFee          *big.Int
	LiquidityfeeRate uint16
	TransferLimit    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLnProviderUpdated is a free log retrieval operation binding the contract event 0x89d09e625a1f04abb654b8315b57f9acefd0bfd504d79d1d4bf610f83a358ca8.
//
// Solidity: event LnProviderUpdated(uint256 remoteChainId, address provider, address sourceToken, address targetToken, uint112 baseFee, uint16 liquidityfeeRate, uint112 transferLimit)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterLnProviderUpdated(opts *bind.FilterOpts) (*Lnv3BridgeLnProviderUpdatedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "LnProviderUpdated")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeLnProviderUpdatedIterator{contract: _Lnv3Bridge.contract, event: "LnProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchLnProviderUpdated is a free log subscription operation binding the contract event 0x89d09e625a1f04abb654b8315b57f9acefd0bfd504d79d1d4bf610f83a358ca8.
//
// Solidity: event LnProviderUpdated(uint256 remoteChainId, address provider, address sourceToken, address targetToken, uint112 baseFee, uint16 liquidityfeeRate, uint112 transferLimit)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchLnProviderUpdated(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeLnProviderUpdated) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "LnProviderUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeLnProviderUpdated)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "LnProviderUpdated", log); err != nil {
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

// ParseLnProviderUpdated is a log parse operation binding the contract event 0x89d09e625a1f04abb654b8315b57f9acefd0bfd504d79d1d4bf610f83a358ca8.
//
// Solidity: event LnProviderUpdated(uint256 remoteChainId, address provider, address sourceToken, address targetToken, uint112 baseFee, uint16 liquidityfeeRate, uint112 transferLimit)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseLnProviderUpdated(log types.Log) (*Lnv3BridgeLnProviderUpdated, error) {
	event := new(Lnv3BridgeLnProviderUpdated)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "LnProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Lnv3Bridge contract.
type Lnv3BridgePausedIterator struct {
	Event *Lnv3BridgePaused // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgePaused)
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
		it.Event = new(Lnv3BridgePaused)
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
func (it *Lnv3BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgePaused represents a Paused event raised by the Lnv3Bridge contract.
type Lnv3BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*Lnv3BridgePausedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgePausedIterator{contract: _Lnv3Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *Lnv3BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgePaused)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParsePaused(log types.Log) (*Lnv3BridgePaused, error) {
	event := new(Lnv3BridgePaused)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgePenaltyReserveUpdatedIterator is returned from FilterPenaltyReserveUpdated and is used to iterate over the raw logs and unpacked data for PenaltyReserveUpdated events raised by the Lnv3Bridge contract.
type Lnv3BridgePenaltyReserveUpdatedIterator struct {
	Event *Lnv3BridgePenaltyReserveUpdated // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgePenaltyReserveUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgePenaltyReserveUpdated)
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
		it.Event = new(Lnv3BridgePenaltyReserveUpdated)
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
func (it *Lnv3BridgePenaltyReserveUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgePenaltyReserveUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgePenaltyReserveUpdated represents a PenaltyReserveUpdated event raised by the Lnv3Bridge contract.
type Lnv3BridgePenaltyReserveUpdated struct {
	Provider              common.Address
	SourceToken           common.Address
	UpdatedPanaltyReserve *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPenaltyReserveUpdated is a free log retrieval operation binding the contract event 0xcb96d44bdc94c9da498125fd745b97752bbc57eb3274a9535977e34da277e32d.
//
// Solidity: event PenaltyReserveUpdated(address provider, address sourceToken, uint256 updatedPanaltyReserve)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterPenaltyReserveUpdated(opts *bind.FilterOpts) (*Lnv3BridgePenaltyReserveUpdatedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "PenaltyReserveUpdated")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgePenaltyReserveUpdatedIterator{contract: _Lnv3Bridge.contract, event: "PenaltyReserveUpdated", logs: logs, sub: sub}, nil
}

// WatchPenaltyReserveUpdated is a free log subscription operation binding the contract event 0xcb96d44bdc94c9da498125fd745b97752bbc57eb3274a9535977e34da277e32d.
//
// Solidity: event PenaltyReserveUpdated(address provider, address sourceToken, uint256 updatedPanaltyReserve)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchPenaltyReserveUpdated(opts *bind.WatchOpts, sink chan<- *Lnv3BridgePenaltyReserveUpdated) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "PenaltyReserveUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgePenaltyReserveUpdated)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "PenaltyReserveUpdated", log); err != nil {
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

// ParsePenaltyReserveUpdated is a log parse operation binding the contract event 0xcb96d44bdc94c9da498125fd745b97752bbc57eb3274a9535977e34da277e32d.
//
// Solidity: event PenaltyReserveUpdated(address provider, address sourceToken, uint256 updatedPanaltyReserve)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParsePenaltyReserveUpdated(log types.Log) (*Lnv3BridgePenaltyReserveUpdated, error) {
	event := new(Lnv3BridgePenaltyReserveUpdated)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "PenaltyReserveUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeSlashRequestIterator is returned from FilterSlashRequest and is used to iterate over the raw logs and unpacked data for SlashRequest events raised by the Lnv3Bridge contract.
type Lnv3BridgeSlashRequestIterator struct {
	Event *Lnv3BridgeSlashRequest // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeSlashRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeSlashRequest)
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
		it.Event = new(Lnv3BridgeSlashRequest)
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
func (it *Lnv3BridgeSlashRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeSlashRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeSlashRequest represents a SlashRequest event raised by the Lnv3Bridge contract.
type Lnv3BridgeSlashRequest struct {
	TransferId    [32]byte
	RemoteChainId *big.Int
	Provider      common.Address
	SourceToken   common.Address
	TargetToken   common.Address
	Slasher       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSlashRequest is a free log retrieval operation binding the contract event 0x88db78985c3a67e1d3ed326247cef7c51a31f73e4428674250bffc6029f3e49d.
//
// Solidity: event SlashRequest(bytes32 transferId, uint256 remoteChainId, address provider, address sourceToken, address targetToken, address slasher)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterSlashRequest(opts *bind.FilterOpts) (*Lnv3BridgeSlashRequestIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "SlashRequest")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeSlashRequestIterator{contract: _Lnv3Bridge.contract, event: "SlashRequest", logs: logs, sub: sub}, nil
}

// WatchSlashRequest is a free log subscription operation binding the contract event 0x88db78985c3a67e1d3ed326247cef7c51a31f73e4428674250bffc6029f3e49d.
//
// Solidity: event SlashRequest(bytes32 transferId, uint256 remoteChainId, address provider, address sourceToken, address targetToken, address slasher)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchSlashRequest(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeSlashRequest) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "SlashRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeSlashRequest)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "SlashRequest", log); err != nil {
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

// ParseSlashRequest is a log parse operation binding the contract event 0x88db78985c3a67e1d3ed326247cef7c51a31f73e4428674250bffc6029f3e49d.
//
// Solidity: event SlashRequest(bytes32 transferId, uint256 remoteChainId, address provider, address sourceToken, address targetToken, address slasher)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseSlashRequest(log types.Log) (*Lnv3BridgeSlashRequest, error) {
	event := new(Lnv3BridgeSlashRequest)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "SlashRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeTokenInfoUpdatedIterator is returned from FilterTokenInfoUpdated and is used to iterate over the raw logs and unpacked data for TokenInfoUpdated events raised by the Lnv3Bridge contract.
type Lnv3BridgeTokenInfoUpdatedIterator struct {
	Event *Lnv3BridgeTokenInfoUpdated // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeTokenInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeTokenInfoUpdated)
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
		it.Event = new(Lnv3BridgeTokenInfoUpdated)
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
func (it *Lnv3BridgeTokenInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeTokenInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeTokenInfoUpdated represents a TokenInfoUpdated event raised by the Lnv3Bridge contract.
type Lnv3BridgeTokenInfoUpdated struct {
	TokenInfoKey   [32]byte
	ProtocolFee    *big.Int
	Penalty        *big.Int
	SourceDecimals *big.Int
	TargetDecimals *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTokenInfoUpdated is a free log retrieval operation binding the contract event 0x1e5cf0e3e832b0e553adfb29acd7d687c2a4d2f67de5a201ed6d2313c14a240f.
//
// Solidity: event TokenInfoUpdated(bytes32 tokenInfoKey, uint112 protocolFee, uint112 penalty, uint112 sourceDecimals, uint112 targetDecimals)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterTokenInfoUpdated(opts *bind.FilterOpts) (*Lnv3BridgeTokenInfoUpdatedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "TokenInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeTokenInfoUpdatedIterator{contract: _Lnv3Bridge.contract, event: "TokenInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenInfoUpdated is a free log subscription operation binding the contract event 0x1e5cf0e3e832b0e553adfb29acd7d687c2a4d2f67de5a201ed6d2313c14a240f.
//
// Solidity: event TokenInfoUpdated(bytes32 tokenInfoKey, uint112 protocolFee, uint112 penalty, uint112 sourceDecimals, uint112 targetDecimals)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchTokenInfoUpdated(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeTokenInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "TokenInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeTokenInfoUpdated)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "TokenInfoUpdated", log); err != nil {
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

// ParseTokenInfoUpdated is a log parse operation binding the contract event 0x1e5cf0e3e832b0e553adfb29acd7d687c2a4d2f67de5a201ed6d2313c14a240f.
//
// Solidity: event TokenInfoUpdated(bytes32 tokenInfoKey, uint112 protocolFee, uint112 penalty, uint112 sourceDecimals, uint112 targetDecimals)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseTokenInfoUpdated(log types.Log) (*Lnv3BridgeTokenInfoUpdated, error) {
	event := new(Lnv3BridgeTokenInfoUpdated)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "TokenInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeTokenLockedIterator is returned from FilterTokenLocked and is used to iterate over the raw logs and unpacked data for TokenLocked events raised by the Lnv3Bridge contract.
type Lnv3BridgeTokenLockedIterator struct {
	Event *Lnv3BridgeTokenLocked // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeTokenLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeTokenLocked)
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
		it.Event = new(Lnv3BridgeTokenLocked)
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
func (it *Lnv3BridgeTokenLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeTokenLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeTokenLocked represents a TokenLocked event raised by the Lnv3Bridge contract.
type Lnv3BridgeTokenLocked struct {
	Params     LnBridgeSourceV3TransferParams
	TransferId [32]byte
	TargetAmount *big.Int
	Fee          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenLocked is a free log retrieval operation binding the contract event 0x83c35d6c5591e4558a0e2999947de65e31db92eb6ff31371ede54f03264efa0d.
//
// Solidity: event TokenLocked((uint256,address,address,address,uint112,uint112,address,uint256) params, bytes32 transferId, uint112 targetAmount, uint112 fee)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterTokenLocked(opts *bind.FilterOpts) (*Lnv3BridgeTokenLockedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "TokenLocked")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeTokenLockedIterator{contract: _Lnv3Bridge.contract, event: "TokenLocked", logs: logs, sub: sub}, nil
}

// WatchTokenLocked is a free log subscription operation binding the contract event 0x83c35d6c5591e4558a0e2999947de65e31db92eb6ff31371ede54f03264efa0d.
//
// Solidity: event TokenLocked((uint256,address,address,address,uint112,uint112,address,uint256) params, bytes32 transferId, uint112 targetAmount, uint112 fee)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchTokenLocked(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeTokenLocked) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "TokenLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeTokenLocked)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "TokenLocked", log); err != nil {
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

// ParseTokenLocked is a log parse operation binding the contract event 0x83c35d6c5591e4558a0e2999947de65e31db92eb6ff31371ede54f03264efa0d.
//
// Solidity: event TokenLocked((uint256,address,address,address,uint112,uint112,address,uint256) params, bytes32 transferId, uint112 targetAmount, uint112 fee)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseTokenLocked(log types.Log) (*Lnv3BridgeTokenLocked, error) {
	event := new(Lnv3BridgeTokenLocked)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "TokenLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the Lnv3Bridge contract.
type Lnv3BridgeTokenRegisteredIterator struct {
	Event *Lnv3BridgeTokenRegistered // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeTokenRegistered)
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
		it.Event = new(Lnv3BridgeTokenRegistered)
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
func (it *Lnv3BridgeTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeTokenRegistered represents a TokenRegistered event raised by the Lnv3Bridge contract.
type Lnv3BridgeTokenRegistered struct {
	Key           [32]byte
	RemoteChainId *big.Int
	SourceToken   common.Address
	TargetToken   common.Address
	ProtocolFee   *big.Int
	Penalty       *big.Int
	Index         uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xc88686a66177c5149684c415fd3fff315c45023438ca1737750c300f46412c8b.
//
// Solidity: event TokenRegistered(bytes32 key, uint256 remoteChainId, address sourceToken, address targetToken, uint112 protocolFee, uint112 penalty, uint32 index)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterTokenRegistered(opts *bind.FilterOpts) (*Lnv3BridgeTokenRegisteredIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "TokenRegistered")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeTokenRegisteredIterator{contract: _Lnv3Bridge.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xc88686a66177c5149684c415fd3fff315c45023438ca1737750c300f46412c8b.
//
// Solidity: event TokenRegistered(bytes32 key, uint256 remoteChainId, address sourceToken, address targetToken, uint112 protocolFee, uint112 penalty, uint32 index)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeTokenRegistered) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "TokenRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeTokenRegistered)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// ParseTokenRegistered is a log parse operation binding the contract event 0xc88686a66177c5149684c415fd3fff315c45023438ca1737750c300f46412c8b.
//
// Solidity: event TokenRegistered(bytes32 key, uint256 remoteChainId, address sourceToken, address targetToken, uint112 protocolFee, uint112 penalty, uint32 index)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseTokenRegistered(log types.Log) (*Lnv3BridgeTokenRegistered, error) {
	event := new(Lnv3BridgeTokenRegistered)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeTransferFilledIterator is returned from FilterTransferFilled and is used to iterate over the raw logs and unpacked data for TransferFilled events raised by the Lnv3Bridge contract.
type Lnv3BridgeTransferFilledIterator struct {
	Event *Lnv3BridgeTransferFilled // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeTransferFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeTransferFilled)
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
		it.Event = new(Lnv3BridgeTransferFilled)
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
func (it *Lnv3BridgeTransferFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeTransferFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeTransferFilled represents a TransferFilled event raised by the Lnv3Bridge contract.
type Lnv3BridgeTransferFilled struct {
	TransferId [32]byte
	Provider   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferFilled is a free log retrieval operation binding the contract event 0x5f3f39cdc66af8643701985897f6bea88129dec0331a8e0f739916223c0d398c.
//
// Solidity: event TransferFilled(bytes32 transferId, address provider)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterTransferFilled(opts *bind.FilterOpts) (*Lnv3BridgeTransferFilledIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "TransferFilled")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeTransferFilledIterator{contract: _Lnv3Bridge.contract, event: "TransferFilled", logs: logs, sub: sub}, nil
}

// WatchTransferFilled is a free log subscription operation binding the contract event 0x5f3f39cdc66af8643701985897f6bea88129dec0331a8e0f739916223c0d398c.
//
// Solidity: event TransferFilled(bytes32 transferId, address provider)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchTransferFilled(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeTransferFilled) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "TransferFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeTransferFilled)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "TransferFilled", log); err != nil {
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

// ParseTransferFilled is a log parse operation binding the contract event 0x5f3f39cdc66af8643701985897f6bea88129dec0331a8e0f739916223c0d398c.
//
// Solidity: event TransferFilled(bytes32 transferId, address provider)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseTransferFilled(log types.Log) (*Lnv3BridgeTransferFilled, error) {
	event := new(Lnv3BridgeTransferFilled)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "TransferFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeTransferSlashedIterator is returned from FilterTransferSlashed and is used to iterate over the raw logs and unpacked data for TransferSlashed events raised by the Lnv3Bridge contract.
type Lnv3BridgeTransferSlashedIterator struct {
	Event *Lnv3BridgeTransferSlashed // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeTransferSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeTransferSlashed)
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
		it.Event = new(Lnv3BridgeTransferSlashed)
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
func (it *Lnv3BridgeTransferSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeTransferSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeTransferSlashed represents a TransferSlashed event raised by the Lnv3Bridge contract.
type Lnv3BridgeTransferSlashed struct {
	TransferId  [32]byte
	Provider    common.Address
	Slasher     common.Address
	SlashAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferSlashed is a free log retrieval operation binding the contract event 0x8093e14577907e78b2fbe5796d0b0b6f3e012cdd858bd4f48904f16f5e7cfa7b.
//
// Solidity: event TransferSlashed(bytes32 transferId, address provider, address slasher, uint112 slashAmount)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterTransferSlashed(opts *bind.FilterOpts) (*Lnv3BridgeTransferSlashedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "TransferSlashed")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeTransferSlashedIterator{contract: _Lnv3Bridge.contract, event: "TransferSlashed", logs: logs, sub: sub}, nil
}

// WatchTransferSlashed is a free log subscription operation binding the contract event 0x8093e14577907e78b2fbe5796d0b0b6f3e012cdd858bd4f48904f16f5e7cfa7b.
//
// Solidity: event TransferSlashed(bytes32 transferId, address provider, address slasher, uint112 slashAmount)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchTransferSlashed(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeTransferSlashed) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "TransferSlashed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeTransferSlashed)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "TransferSlashed", log); err != nil {
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

// ParseTransferSlashed is a log parse operation binding the contract event 0x8093e14577907e78b2fbe5796d0b0b6f3e012cdd858bd4f48904f16f5e7cfa7b.
//
// Solidity: event TransferSlashed(bytes32 transferId, address provider, address slasher, uint112 slashAmount)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseTransferSlashed(log types.Log) (*Lnv3BridgeTransferSlashed, error) {
	event := new(Lnv3BridgeTransferSlashed)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "TransferSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Lnv3BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Lnv3Bridge contract.
type Lnv3BridgeUnpausedIterator struct {
	Event *Lnv3BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *Lnv3BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Lnv3BridgeUnpaused)
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
		it.Event = new(Lnv3BridgeUnpaused)
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
func (it *Lnv3BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Lnv3BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Lnv3BridgeUnpaused represents a Unpaused event raised by the Lnv3Bridge contract.
type Lnv3BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lnv3Bridge *Lnv3BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*Lnv3BridgeUnpausedIterator, error) {

	logs, sub, err := _Lnv3Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &Lnv3BridgeUnpausedIterator{contract: _Lnv3Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lnv3Bridge *Lnv3BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *Lnv3BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Lnv3Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Lnv3BridgeUnpaused)
				if err := _Lnv3Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Lnv3Bridge *Lnv3BridgeFilterer) ParseUnpaused(log types.Log) (*Lnv3BridgeUnpaused, error) {
	event := new(Lnv3BridgeUnpaused)
	if err := _Lnv3Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}


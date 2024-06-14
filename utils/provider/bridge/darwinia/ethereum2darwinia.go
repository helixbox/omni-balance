package darwinia

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"omni-balance/utils/constant"
	"strings"
)

var (
	ethereum2darwinia = map[string]Params{
		"RING": Params{
			contractAddress:  common.HexToAddress("0xc29dCb1F12a1618262eF9FBA673b77140adc02D6"),
			TokenAddress:     common.HexToAddress("0x9469d013805bffb7d3debe5e7839237e535ec483"),
			originalToken:    common.HexToAddress("0xE7578598Aac020abFB918f33A20faD5B71d670b4"),
			remoteAppAddress: common.HexToAddress("0xDc0C760c0fB4672D06088515F6446a71Df0c64C1"),
			localAppAddress:  common.HexToAddress("0x2B496f19A420C02490dB859fefeCCD71eDc2c046"),
			sourceMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			targetMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			recipient:        common.HexToAddress("0x4CA75992d2750BEC270731A72DfDedE6b9E71cC7"),
			sourceChainId:    1,
			targetChainId:    46,
			extData:          "000000000000000000000000092e19c46c9daab7824393f1cd9c22f5bea1356000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000014{sender}000000000000000000000000",
		},
		"KTON": Params{
			contractAddress:  common.Address{}, // must replace to sender address
			TokenAddress:     common.HexToAddress("0x9469d013805bffb7d3debe5e7839237e535ec483"),
			originalToken:    common.HexToAddress("0x0000000000000000000000000000000000000402"),
			remoteAppAddress: common.HexToAddress("0xDc0C760c0fB4672D06088515F6446a71Df0c64C1"),
			localAppAddress:  common.HexToAddress("0x2B496f19A420C02490dB859fefeCCD71eDc2c046"),
			sourceMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			targetMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			recipient:        common.HexToAddress("0x4CA75992d2750BEC270731A72DfDedE6b9E71cC7"),
			xTokenAddress:    common.HexToAddress("0x9f284e1337a815fe77d2ff4ae46544645b20c5ff"),
			sourceChainId:    1,
			targetChainId:    46,
			extData:          "000000000000000000000000{sender}00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000",
		},
	}
)

func Ethereum2darwinia(ctx context.Context, args SwapParams) (tx *types.LegacyTx, err error) {

	var (
		wallet     = args.Sender
		realWallet = wallet.GetAddress(true)
		tokenConf  = ethereum2darwinia[strings.ToUpper(args.TokenName)]
	)

	if tokenConf.sourceChainId == 0 {
		return nil, errors.New("token not support")
	}
	if !constant.IsTestCtx(ctx) {
		err := approve(ctx, tokenConf.sourceChainId, tokenConf.TokenAddress, tokenConf.remoteAppAddress,
			wallet, args.Amount.Mul(decimal.NewFromInt(2)), args.Client)
		if err != nil {
			return nil, errors.Wrap(err, "approve")
		}
	}
	extData := common.Hex2Bytes(ReplaceExtData(tokenConf.extData, realWallet.Hex()))
	contractAddress := tokenConf.contractAddress
	if strings.EqualFold(contractAddress.Hex(), constant.ZeroAddress.Hex()) {
		contractAddress = realWallet
	}
	data, err := Unlock(
		tokenConf.sourceChainId,
		tokenConf.originalToken,
		contractAddress,
		tokenConf.recipient,
		realWallet,
		args.Amount,
		args.Nonce,
		extData)
	if err != nil {
		return nil, errors.Wrap(err, "Unlock")
	}

	data, err = ReceiveMessage(tokenConf.sourceChainId, tokenConf.remoteAppAddress, tokenConf.localAppAddress,
		common.Bytes2Hex(data))
	if err != nil {
		return nil, errors.Wrap(err, "ReceiveMessage")
	}

	fee, param, gas, err := FetchMsglineFeeAndParams(ctx, tokenConf.sourceChainId, tokenConf.targetChainId,
		tokenConf.sourceMessager, tokenConf.targetMessager, realWallet, common.Bytes2Hex(data))
	if err != nil {
		return nil, errors.Wrap(err, "FetchMsglineFeeAndParams")
	}

	if strings.EqualFold(tokenConf.contractAddress.Hex(), constant.ZeroAddress.Hex()) {
		contractAddress = tokenConf.remoteAppAddress
		data, err = BurnAndXUnlock(tokenConf.xTokenAddress, tokenConf.recipient, realWallet, args.Amount,
			args.Nonce, extData, param)
	} else {
		data, err = XTokenBurnAndXUnlock(tokenConf.recipient, realWallet, args.Amount, args.Nonce, extData, param)

	}

	if err != nil {
		return nil, errors.Wrap(err, "XTokenBurnAndXUnlock")
	}
	// remoteAppAddress
	return &types.LegacyTx{
		To:    &contractAddress,
		Value: fee.BigInt(),
		Data:  data,
		Gas:   uint64(gas.IntPart()),
	}, nil
}

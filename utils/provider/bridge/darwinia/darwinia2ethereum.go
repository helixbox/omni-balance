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
	darwinia2ethereum = map[string]Params{
		"RING": Params{
			contractAddress:  common.HexToAddress("0x092e19c46c9daab7824393f1cd9c22f5bea13560"),
			originalToken:    common.HexToAddress("0xE7578598Aac020abFB918f33A20faD5B71d670b4"),
			recipient:        common.HexToAddress("0x4CA75992d2750BEC270731A72DfDedE6b9E71cC7"),
			remoteAppAddress: common.HexToAddress("0x2B496f19A420C02490dB859fefeCCD71eDc2c046"),
			localAppAddress:  common.HexToAddress("0xDc0C760c0fB4672D06088515F6446a71Df0c64C1"),
			sourceMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			targetMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			sourceChainId:    46,
			targetChainId:    1,
			extData:          "000000000000000000000000c29dcb1f12a1618262ef9fba673b77140adc02d600000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000014{sender}000000000000000000000000",
		},
		"KTON": Params{
			TokenAddress:     common.HexToAddress("0x0000000000000000000000000000000000000402"),
			contractAddress:  common.Address{}, // must replace to sender address
			originalToken:    common.HexToAddress("0x0000000000000000000000000000000000000402"),
			remoteAppAddress: common.HexToAddress("0x2B496f19A420C02490dB859fefeCCD71eDc2c046"),
			localAppAddress:  common.HexToAddress("0xDc0C760c0fB4672D06088515F6446a71Df0c64C1"),
			sourceMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			targetMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			recipient:        common.HexToAddress("0x4CA75992d2750BEC270731A72DfDedE6b9E71cC7"),
			sourceChainId:    46,
			targetChainId:    1,
			extData:          "000000000000000000000000{sender}00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000",
		},
	}
)

func Darwinia2ethereum(ctx context.Context, args SwapParams) (tx *types.LegacyTx, err error) {
	tokenConf := darwinia2ethereum[strings.ToUpper(args.TokenName)]
	if tokenConf.sourceChainId == 0 {
		return nil, errors.Errorf("not support token %s", args.TokenName)
	}
	var (
		data          []byte
		isNativeToken = strings.EqualFold(args.TokenName, "RING")
		wallet        = args.Sender
		realWallet    = wallet.GetAddress(true)
	)
	if !isNativeToken && !constant.IsTestCtx(ctx) {
		if err := approve(ctx, tokenConf.sourceChainId, tokenConf.TokenAddress, tokenConf.localAppAddress,
			wallet, args.Amount.Mul(decimal.NewFromInt(2)), args.Client); err != nil {
			return nil, errors.Wrap(err, "approve")
		}
	}

	contractAddress := tokenConf.contractAddress
	if strings.EqualFold(contractAddress.Hex(), constant.ZeroAddress.Hex()) {
		contractAddress = realWallet
	}
	extData := common.Hex2Bytes(ReplaceExtData(tokenConf.extData, realWallet.Hex()))
	data, err = Issue(
		tokenConf.sourceChainId, tokenConf.originalToken, contractAddress, tokenConf.recipient,
		realWallet, args.Amount, args.Nonce, extData)
	if err != nil {
		return nil, errors.Wrap(err, "issue")
	}

	data, err = ReceiveMessage(tokenConf.sourceChainId, tokenConf.remoteAppAddress, tokenConf.localAppAddress,
		common.Bytes2Hex(data))
	if err != nil {
		return nil, errors.Wrap(err, "receiveMessage")
	}

	fee, param, gas, err := FetchMsglineFeeAndParams(ctx, tokenConf.sourceChainId, tokenConf.targetChainId,
		tokenConf.sourceMessager, tokenConf.targetMessager, realWallet, common.Bytes2Hex(data))
	if err != nil {
		return nil, errors.Wrap(err, "fetchMsglineFeeAndParams")
	}

	if isNativeToken {
		data, err = WTokenLockAndXIssue(tokenConf.targetChainId, tokenConf.recipient, realWallet,
			args.Amount, args.Nonce, common.Bytes2Hex(extData), common.Bytes2Hex(param))
		if err != nil {
			return nil, errors.Wrap(err, "wTokenLockAndXIssue")
		}
	}
	if !isNativeToken {
		data, err = XTokenLockAndXIssue(tokenConf.targetChainId, tokenConf.originalToken, tokenConf.recipient,
			realWallet, args.Amount, args.Nonce, common.Bytes2Hex(extData), common.Bytes2Hex(param))
		if err != nil {
			return nil, errors.Wrap(err, "xTokenLockAndXIssue")
		}
	}
	if strings.EqualFold(contractAddress.Hex(), realWallet.Hex()) {
		contractAddress = tokenConf.remoteAppAddress
	}
	a := &types.LegacyTx{
		To:    &contractAddress,
		Value: args.Amount.Add(fee).BigInt(),
		Data:  data,
		Gas:   uint64(gas.IntPart()),
	}
	if !isNativeToken {
		a.Value = fee.BigInt()
	}
	return a, nil
}

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
	crab2darwinia = map[string]Params{
		"RING": {
			contractAddress:  common.HexToAddress("0xf6372ab2d35B32156A19F2d2F23FA6dDeFBE58bd"),
			originalToken:    common.HexToAddress("0xE7578598Aac020abFB918f33A20faD5B71d670b4"),
			remoteAppAddress: common.HexToAddress("0xf6372ab2d35B32156A19F2d2F23FA6dDeFBE58bd"),
			localAppAddress:  common.HexToAddress("0xa64D1c284280b22f921E7B2A55040C7bbfD4d9d0"),
			sourceMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			targetMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			xTokenAddress:    common.HexToAddress("0x273131F7CB50ac002BDd08cA721988731F7e1092"),
			recipient:        common.HexToAddress("0xA8d0E9a45249Ec839C397fa0F371f5F64eCAB7F7"),
			sourceChainId:    44,
			targetChainId:    46,
		},
		"CRAB": {
			contractAddress:  common.HexToAddress("0x004D0dE211BC148c3Ce696C51Cbc85BD421727E9"),
			originalToken:    common.HexToAddress("0x2D2b97EA380b0185e9fDF8271d1AFB5d2Bf18329"),
			remoteAppAddress: common.HexToAddress("0xa64D1c284280b22f921E7B2A55040C7bbfD4d9d0"),
			localAppAddress:  common.HexToAddress("0xf6372ab2d35B32156A19F2d2F23FA6dDeFBE58bd"),
			sourceMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			targetMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			sourceChainId:    44,
			targetChainId:    46,
		},
	}
)

func Crab2Darwinia(ctx context.Context, args SwapParams) (tx *types.LegacyTx, err error) {
	conf := crab2darwinia[strings.ToUpper(args.TokenName)]
	if conf.sourceChainId == 0 {
		return nil, errors.Errorf("not support token %s", args.TokenName)
	}
	var (
		data          []byte
		isNativeToken = strings.EqualFold(args.TokenName, "CRAB")
		wallet        = args.Sender
		realWallet    = wallet.GetAddress(true)
	)
	if !isNativeToken {
		if !constant.IsTestCtx(ctx) {

			if err := approve(ctx, conf.sourceChainId, conf.TokenAddress, conf.localAppAddress, wallet,
				args.Amount.Mul(decimal.NewFromInt(2)), args.Client); err != nil {

				return nil, errors.Wrap(err, "approve")
			}
		}

		data, err = Unlock(
			conf.sourceChainId,
			conf.originalToken,
			realWallet,
			conf.recipient,
			realWallet,
			args.Amount,
			args.Nonce,
			realWallet.Bytes(),
		)
		if err != nil {
			return nil, errors.Wrap(err, "unlock")
		}
	}
	if isNativeToken {
		data, err = Issue(
			conf.sourceChainId,
			conf.originalToken,
			conf.contractAddress,
			realWallet,
			realWallet,
			args.Amount,
			args.Nonce,
			[]byte(""))
		if err != nil {
			return nil, errors.Wrap(err, "issue")
		}
	}
	data, err = ReceiveMessage(conf.sourceChainId, conf.remoteAppAddress, conf.localAppAddress, common.Bytes2Hex(data))
	if err != nil {
		return nil, errors.Wrap(err, "receiveMessage")
	}
	fee, param, gas, err := FetchMsglineFeeAndParams(ctx, conf.sourceChainId, conf.targetChainId, conf.sourceMessager,
		conf.targetMessager, wallet.GetAddress(true), common.Bytes2Hex(data))
	if err != nil {
		return nil, errors.Wrap(err, "fetchMsglineFeeAndParams")
	}
	if ctx.Value(constant.FeeTestKeyInCtx) != nil { // for test
		fee = ctx.Value(constant.FeeTestKeyInCtx).(decimal.Decimal)
	}
	if isNativeToken {
		data, err = WTokenLockAndXIssue(conf.targetChainId,
			wallet.GetAddress(true), wallet.GetAddress(true), args.Amount, args.Nonce, "", common.Bytes2Hex(param))
		if err != nil {
			return nil, errors.Wrap(err, "wTokenLockAndXIssue")
		}
	}
	if !isNativeToken {
		data, err = BurnAndXUnlock(conf.xTokenAddress, conf.recipient, realWallet, args.Amount, args.Nonce, realWallet.Bytes(), param)
		if err != nil {
			return nil, errors.Wrap(err, "burnAndXUnlock")
		}
	}

	a := &types.LegacyTx{
		To:    &conf.contractAddress,
		Value: args.Amount.Add(fee).BigInt(),
		Data:  data,
		Gas:   uint64(gas.IntPart()),
	}
	if !isNativeToken {
		a.Value = fee.BigInt()
	}
	return a, nil
}

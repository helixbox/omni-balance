package darwinia

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"omni-balance/utils"
	"omni-balance/utils/constant"
	"strings"
)

var (
	darwinia2Crab = map[string]Params{
		"RING": {
			contractAddress:  common.HexToAddress("0xA8d0E9a45249Ec839C397fa0F371f5F64eCAB7F7"),
			originalToken:    common.HexToAddress("0xE7578598Aac020abFB918f33A20faD5B71d670b4"),
			remoteAppAddress: common.HexToAddress("0xa64D1c284280b22f921E7B2A55040C7bbfD4d9d0"),
			localAppAddress:  common.HexToAddress("0xf6372ab2d35B32156A19F2d2F23FA6dDeFBE58bd"),
			sourceMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			targetMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			sourceChainId:    46,
			targetChainId:    44,
		},
		"CRAB": {
			TokenAddress:     common.HexToAddress("0x656567eb75b765fc320783cc6edd86bd854b2305"),
			contractAddress:  common.HexToAddress(""),
			originalToken:    common.HexToAddress("0x2D2b97EA380b0185e9fDF8271d1AFB5d2Bf18329"),
			remoteAppAddress: common.HexToAddress("0xf6372ab2d35B32156A19F2d2F23FA6dDeFBE58bd"),
			localAppAddress:  common.HexToAddress("0xa64D1c284280b22f921E7B2A55040C7bbfD4d9d0"),
			sourceMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			targetMessager:   common.HexToAddress("0x65Be094765731F394bc6d9DF53bDF3376F1Fc8B0"),
			xTokenAddress:    common.HexToAddress("0x656567Eb75b765FC320783cc6EDd86bD854b2305"),
			recipient:        common.HexToAddress("0x004D0dE211BC148c3Ce696C51Cbc85BD421727E9"),
			sourceChainId:    46,
			targetChainId:    44,
		},
	}
)

func Darwinia2Crab(ctx context.Context, args SwapParams) (tx *types.LegacyTx, err error) {
	tokenConf := darwinia2Crab[strings.ToUpper(args.TokenName)]
	if tokenConf.sourceChainId == 0 {
		return nil, errors.Errorf("not support token %s", args.TokenName)
	}
	var (
		data          []byte
		isNativeToken = utils.InArray(args.TokenName, []string{"RING"})
		wallet        = args.Sender
		realWallet    = wallet.GetAddress(true)
	)
	if !isNativeToken {
		if !constant.IsTestCtx(ctx) {
			if err := approve(ctx, tokenConf.sourceChainId, tokenConf.TokenAddress,
				tokenConf.remoteAppAddress, wallet, args.Amount.Mul(decimal.NewFromInt(2)), args.Client); err != nil {
				return nil, errors.Wrap(err, "approve")
			}
		}
		data, err = Unlock(
			tokenConf.sourceChainId,
			tokenConf.originalToken,
			realWallet,
			tokenConf.recipient,
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
			tokenConf.sourceChainId, tokenConf.originalToken, tokenConf.contractAddress,
			realWallet, realWallet, args.Amount, args.Nonce, []byte(""))
		if err != nil {
			return nil, errors.Wrap(err, "issue")
		}
	}

	data, err = ReceiveMessage(tokenConf.sourceChainId, tokenConf.remoteAppAddress, tokenConf.localAppAddress, common.Bytes2Hex(data))
	if err != nil {
		return nil, errors.Wrap(err, "receiveMessage")
	}
	fee, param, gas, err := FetchMsglineFeeAndParams(ctx, tokenConf.sourceChainId, tokenConf.targetChainId,
		tokenConf.sourceMessager, tokenConf.targetMessager, realWallet, common.Bytes2Hex(data))
	if err != nil {
		return nil, errors.Wrap(err, "fetchMsglineFeeAndParams")
	}
	if isNativeToken {
		data, err = WTokenLockAndXIssue(tokenConf.targetChainId, realWallet, realWallet, args.Amount,
			args.Nonce, "", common.Bytes2Hex(param))
		if err != nil {
			return nil, errors.Wrap(err, "wTokenLockAndXIssue")
		}
	}
	if !isNativeToken {
		data, err = BurnAndXUnlock(tokenConf.xTokenAddress, tokenConf.recipient, realWallet, args.Amount,
			args.Nonce, realWallet.Bytes(), param)
		if err != nil {
			return nil, errors.Wrap(err, "burnAndXUnlock")
		}
	}
	a := &types.LegacyTx{
		To:    &tokenConf.contractAddress,
		Value: args.Amount.Add(fee).BigInt(),
		Data:  data,
		Gas:   uint64(gas.IntPart()),
	}
	if !isNativeToken {
		a.Value = fee.BigInt()
		a.To = &tokenConf.remoteAppAddress
	}
	return a, nil
}

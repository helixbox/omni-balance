package helix_liquidity_claim

import (
	"context"
	"fmt"
	"math/big"
	"omni-balance/utils/chains"
	"omni-balance/utils/provider/bridge/darwinia"
	"omni-balance/utils/provider/bridge/helix_liquidity_claim/abi/layerzeroMessager"
	"omni-balance/utils/provider/bridge/helix_liquidity_claim/abi/msgportMessager"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	messagerInstance = map[string]Messager{
		"msgline":   MessagePortMessager{},
		"layerzero": LayerzeroMessager{},
	}
)

type LayerzeroMessager struct {
}

func (l LayerzeroMessager) EncodePayload(ctx context.Context, args EncodePayloadArgs) (payload []byte, err error) {
	abiArgs := &abi.Arguments{
		{Type: chains.MostNewAbiType("address", "address", nil)},
		{Type: chains.MostNewAbiType("address", "address", nil)},
		{Type: chains.MostNewAbiType("bytes", "bytes", nil)},
	}
	return abiArgs.Pack(args.LocalAppAddress, args.RemoteAppAddress, args.Message)
}

func (l LayerzeroMessager) Params(ctx context.Context, args MessagerArgs) (params MessagerParams, err error) {
	c, err := layerzeroMessager.NewLayerzeroMessagerCaller(args.ContractAddress, args.Client)
	if err != nil {
		return MessagerParams{}, errors.Wrap(err, "new layerzero messager caller")
	}
	fmt.Println(args.ContractAddress.String(), args.ToChainId)
	fee, err := c.Fee(&bind.CallOpts{}, big.NewInt(int64(args.ToChainId)), args.Payload)
	if err != nil {
		return MessagerParams{}, errors.Wrap(err, "get fee")
	}
	return MessagerParams{
		Fee:       decimal.NewFromBigInt(fee.NativeFee, 0),
		ExtParams: args.Refunder.Bytes(),
	}, nil
}

type MessagePortMessager struct {
}

func (m MessagePortMessager) EncodePayload(ctx context.Context, args EncodePayloadArgs) (payload []byte, err error) {
	abi, err := msgportMessager.MsgportMessagerMetaData.GetAbi()
	if err != nil {
		return nil, errors.Wrap(err, "get abi")
	}
	return abi.Pack("receiveMessage", big.NewInt(int64(args.FromChainId)), args.LocalAppAddress, args.RemoteAppAddress, args.Message)
}

func (m MessagePortMessager) Params(ctx context.Context, args MessagerArgs) (params MessagerParams, err error) {
	fee, extParams, _, err := darwinia.FetchMsglineFeeAndParams(ctx, int64(args.FromChainId), int64(args.ToChainId),
		args.ContractAddress, args.RemoteMessager, args.Refunder, common.Bytes2Hex(args.Payload))
	if err != nil {
		return MessagerParams{}, errors.Wrap(err, "fetch msgline fee and params")
	}
	return MessagerParams{
		Fee:       fee,
		ExtParams: extParams,
	}, nil
}

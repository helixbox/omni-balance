package utils

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
	"github.com/uniswapv3-go/uniswapv3-universal-router-decoder-go/command"
	"math/big"
	"omni-balance/utils"
	"omni-balance/utils/chains"
	"omni-balance/utils/provider/dex/uniswap/abi/v3UniversalRouter"
	"strings"
)

type Paths []Path

type Path struct {
	TokenIn  common.Address
	Fee      int64
	TokenOut common.Address
}

func (p Paths) ToV2() []common.Address {
	var result []common.Address
	for index, path := range p {
		result = append(result, path.TokenIn)
		if index == len(p)-1 { // last need to append tokenOut
			result = append(result, path.TokenOut)
			continue
		}
	}
	return result
}

func (p Paths) ToBytes(fnName byte) []byte {
	var result = bytes.NewBuffer(nil)
	switch fnName {
	case command.V3_SWAP_EXACT_OUT, command.V3_SWAP_EXACT_IN:
		for index, item := range p {
			tokenIn := item.TokenIn.Hex()
			tokenOut := item.TokenOut.Hex()
			fee := fmt.Sprintf("%06x", item.Fee)
			if fnName == command.V3_SWAP_EXACT_OUT {
				tokenIn = p[len(p)-index-1].TokenOut.Hex()
				tokenOut = p[len(p)-index-1].TokenIn.Hex()
				fee = fmt.Sprintf("%06x", p[len(p)-index-1].Fee)
			}
			result.Write(common.Hex2Bytes(tokenIn[2:]))
			result.Write(common.Hex2Bytes(fee))
			if index == len(p)-1 {
				result.Write(common.Hex2Bytes(tokenOut[2:]))
				continue
			}
		}

	default:
		return nil
	}
	return result.Bytes()
}

type Execute struct {
	commands []command.Command
	errs     []error
}

func (e *Execute) appendCommand(cmd byte, types []string, values ...interface{}) *Execute {
	input, err := AbiPack(
		types,
		values...,
	)
	if err != nil {
		e.errs = append(e.errs, errors.WithStack(err))
		return e
	}
	e.commands = append(e.commands, command.Command{
		Command: cmd,
		Input:   input,
	})
	return e
}

func (e *Execute) lastRevert(isCanRevert ...bool) *Execute {
	if len(e.commands) == 0 {
		return e
	}
	if len(isCanRevert) > 0 && isCanRevert[0] {
		e.commands[len(e.commands)-1] = command.Command{
			Command: utils.SetBit(e.commands[len(e.commands)-1].Command, 7, true),
			Input:   e.commands[len(e.commands)-1].Input,
		}
	}
	return e
}

// UnwrapWEth see https://docs.uniswap.org/contracts/universal-router/technical-reference#unwrap_weth
func (e *Execute) UnwrapWEth(ETHRecipient common.Address, minETHAmount *big.Int, isCanRevert ...bool) *Execute {
	return e.appendCommand(command.UNWRAP_WETH, []string{"address", "uint256"}, ETHRecipient, minETHAmount).
		lastRevert(isCanRevert...)
}

// WrapEth see https://docs.uniswap.org/contracts/universal-router/technical-reference#wrap_eth
func (e *Execute) WrapEth(WETHRecipient common.Address, wrapEthAmount *big.Int, isCanRevert ...bool) *Execute {
	return e.appendCommand(command.WRAP_ETH, []string{"address", "uint256"}, WETHRecipient, wrapEthAmount).
		lastRevert(isCanRevert...)
}

// Permit2Permit see https://docs.uniswap.org/contracts/universal-router/technical-reference#permit2_permit
func (e *Execute) Permit2Permit(permitSingle PermitSingle, signFn func(msg []byte) (sig []byte, err error),
	isCanRevert ...bool) *Execute {

	typedData := apitypes.TypedData{
		Types: apitypes.Types{
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
			"PermitDetails": []apitypes.Type{
				{Name: "token", Type: "address"},
				{Name: "amount", Type: "uint160"},
				{Name: "expiration", Type: "uint48"},
				{Name: "nonce", Type: "uint48"},
			},
			"PermitSingle": []apitypes.Type{
				{Name: "details", Type: "PermitDetails"},
				{Name: "spender", Type: "address"},
				{Name: "sigDeadline", Type: "uint256"},
			},
		},
		PrimaryType: "PermitSingle",
		Domain: apitypes.TypedDataDomain{
			Name:              "Permit2",
			ChainId:           permitSingle.ChainId,
			VerifyingContract: permitSingle.VerifyingContract.Hex(),
		},
		Message: apitypes.TypedDataMessage{
			"details": map[string]interface{}{
				"token":      permitSingle.Details.Token.Hex(),
				"amount":     permitSingle.Details.Amount,
				"expiration": permitSingle.Details.Expiration,
				"nonce":      permitSingle.Details.Nonce,
			},
			"spender":     permitSingle.Spender.Hex(),
			"sigDeadline": permitSingle.SigDeadline,
		},
	}
	sig, err := chains.SignTypedData(typedData, signFn)
	if err != nil {
		e.errs = append(e.errs, errors.Wrap(err, "SignTypedData error"))
		return e
	}
	var args = abi.Arguments{
		{
			Name: "PermitSingle",
			Type: chains.MostNewAbiType("tuple", "", []abi.ArgumentMarshaling{
				{Name: "details", Type: "tuple", Components: []abi.ArgumentMarshaling{
					{Name: "token", Type: "address"},
					{Name: "amount", Type: "uint160"},
					{Name: "expiration", Type: "uint48"},
					{Name: "nonce", Type: "uint48"},
				}},
				{Name: "spender", Type: "address"},
				{Name: "sigDeadline", Type: "uint256"},
			}),
		},
		{
			Name: "signature",
			Type: chains.MostNewAbiType("bytes", "", []abi.ArgumentMarshaling{}),
		},
	}
	type PermitSingle struct {
		Details     PermitDetails  `json:"details"`
		Spender     common.Address `json:"spender"`
		SigDeadline *big.Int       `json:"sigDeadline"`
	}
	permit := &PermitSingle{
		Details:     permitSingle.Details,
		Spender:     permitSingle.Spender,
		SigDeadline: permitSingle.SigDeadline,
	}
	input, err := args.Pack(permit, sig)
	if err != nil {
		e.errs = append(e.errs, errors.Wrap(err, "Pack"))
		return e
	}
	e.commands = append(e.commands, command.Command{
		Command: command.PERMIT2_PERMIT,
		Input:   input,
	})
	return e.lastRevert(isCanRevert...)
}

// V2SwapExactOut see https://docs.uniswap.org/contracts/universal-router/technical-reference#v2_swap_exact_out
func (e *Execute) V2SwapExactOut(recipient common.Address, tokenOutAmount, tokenInMaxAmount *big.Int, path Paths,
	isSenderFromPermit2 bool, isCanRevert ...bool) *Execute {

	return e.appendCommand(command.V2_SWAP_EXACT_OUT, []string{"address", "uint256", "uint256", "address[]", "bool"},
		recipient, tokenOutAmount, tokenInMaxAmount, path.ToV2(), isSenderFromPermit2).
		lastRevert(isCanRevert...)
}

// V2SwapExactIn see https://docs.uniswap.org/contracts/universal-router/technical-reference#v2_swap_exact_in
func (e *Execute) V2SwapExactIn(recipient common.Address, tokenInAmount, tokenInMinAmount *big.Int, path Paths,
	isSenderFromPermit2 bool, isCanRevert ...bool) *Execute {
	return e.appendCommand(command.V2_SWAP_EXACT_IN, []string{"address", "uint256", "uint256", "bytes", "bool"},
		recipient, tokenInAmount, tokenInMinAmount, path.ToV2(), isSenderFromPermit2).
		lastRevert(isCanRevert...)
}

// PayPortion see https://docs.uniswap.org/contracts/universal-router/technical-reference#pay_portion
func (e *Execute) PayPortion(token, recipient common.Address, amount *big.Int, isCanRevert ...bool) *Execute {
	return e.appendCommand(command.PAY_PORTION, []string{"address", "address", "uint256"}, token, recipient, amount).
		lastRevert(isCanRevert...)
}

// Transfer see https://docs.uniswap.org/contracts/universal-router/technical-reference#transfer
func (e *Execute) Transfer(token, recipient common.Address, amount *big.Int, isCanRevert ...bool) *Execute {
	return e.appendCommand(command.TRANSFER, []string{"address", "address", "uint256"}, token, recipient, amount).
		lastRevert(isCanRevert...)
}

// Sweep see https://docs.uniswap.org/contracts/universal-router/technical-reference#sweep
func (e *Execute) Sweep(token, recipient common.Address, tokenMinAmount *big.Int, isCanRevert ...bool) *Execute {
	return e.appendCommand(command.SWEEP, []string{"address", "address", "uint256"}, token, recipient, tokenMinAmount).
		lastRevert(isCanRevert...)
}

// Permit2TransferFrom see https://docs.uniswap.org/contracts/universal-router/technical-reference#permit2_transfer_from
func (e *Execute) Permit2TransferFrom(fromPermit2Token, recipient common.Address, amount *big.Int,
	isCanRevert ...bool) *Execute {

	return e.appendCommand(command.PERMIT2_TRANSFER_FROM,
		[]string{"address", "address", "uint256"}, fromPermit2Token, recipient, amount).lastRevert(isCanRevert...)
}

// V3SwapExactOut see https://docs.uniswap.org/contracts/universal-router/technical-reference#v3_swap_exact_out
func (e *Execute) V3SwapExactOut(recipient common.Address, tokenOutAmountWei, tokenInMaxAmountWei *big.Int, path Paths,
	isSenderFromPermit2 bool, isCanRevert ...bool) *Execute {

	return e.appendCommand(command.V3_SWAP_EXACT_OUT, []string{"address", "uint256", "uint256", "bytes", "bool"},
		recipient, tokenOutAmountWei, tokenInMaxAmountWei, path.ToBytes(command.V3_SWAP_EXACT_OUT), isSenderFromPermit2).
		lastRevert(isCanRevert...)
}

// V3SwapExactIn see https://docs.uniswap.org/contracts/universal-router/technical-reference#v3_swap_exact_in
func (e *Execute) V3SwapExactIn(recipient common.Address, tokenInAmountWei, tokenOutMinAmountWei *big.Int, path Paths,
	isSenderFromPermit2 bool, isCanRevert ...bool) *Execute {

	return e.appendCommand(command.V3_SWAP_EXACT_IN, []string{"address", "uint256", "uint256", "bytes", "bool"},
		recipient, tokenInAmountWei, tokenOutMinAmountWei, path.ToBytes(command.V3_SWAP_EXACT_IN), isSenderFromPermit2).
		lastRevert(isCanRevert...)
}

func (e *Execute) Error() error {
	var errs []string
	for _, err := range e.errs {
		errs = append(errs, err.Error())
	}
	if len(errs) == 0 {
		return nil
	}
	return fmt.Errorf(strings.Join(errs, "\n"))
}

func (e *Execute) Build(deadline *big.Int) ([]byte, error) {
	if err := e.Error(); err != nil {
		return nil, errors.Wrap(err, "processing commands")
	}
	var (
		// 2 unused bytes, reserved for future use. Leaving these 2 bits as 0 will save gas,
		// but any value passed into the contract will be ignored.
		// Later versions of the UniversalRouter will likely expand the 5 bits used for command to use at least 1 of these bits.
		commands = make([]byte, 0)
		inputs   [][]byte
	)
	for _, v := range e.commands {
		commands = append(commands, v.Command)
		inputs = append(inputs, utils.ZFillByte(v.Input, 64))
	}
	abiObj, err := v3UniversalRouter.V3UniversalRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return abiObj.Pack("execute0", commands, inputs, deadline)
}

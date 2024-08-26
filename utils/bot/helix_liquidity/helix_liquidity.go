package helix_liquidity

import (
	"context"
	"omni-balance/utils/bot"
	"omni-balance/utils/provider"
	"omni-balance/utils/provider/bridge/helix_liquidity_claim"

	log "omni-balance/utils/logging"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

func init() {
	bot.Register(HelixLiquidity{}.Name(), HelixLiquidity{})
}

type HelixLiquidity struct {
}

func (b HelixLiquidity) Balance(ctx context.Context, args bot.Params) (decimal.Decimal, error) {
	hlc, err := helix_liquidity_claim.New(args.Conf)
	if err != nil {
		return decimal.Zero, err
	}
	claim := hlc.(helix_liquidity_claim.Claim)
	records, err := claim.ListNeedWithdrawRecords(ctx, args.Info.Wallet.GetAddress(), args.Info.Chain, args.Info.TokenName)
	if err != nil {
		return decimal.Zero, err
	}
	token := args.Conf.GetTokenInfoOnChain(args.Info.TokenName, args.Info.Chain)
	balance, err := args.Info.Wallet.GetExternalBalance(ctx, common.HexToAddress(token.ContractAddress), token.Decimals, args.Client)
	if err != nil {
		return decimal.Zero, err
	}
	var (
		total = balance
	)
	for _, v := range records {
		total = total.Add(v.TotalAmount)
	}
	for _, v := range debtImpl {
		balance, err := v.BalanceOf(ctx, DebtParams{
			Address: args.Info.Wallet.GetAddress(),
			Chain:   args.Info.Chain,
			Client:  args.Client,
			Token:   token.Name,
		})
		if err != nil {
			return decimal.Zero, errors.Wrapf(err, "get %s balance of %s error", token.Name, v.Name())
		}
		total = total.Add(balance)
	}
	return total, nil
}

func (h HelixLiquidity) Check(ctx context.Context, args bot.Params) ([]bot.Task, bot.ProcessType, error) {

	threshold := args.Conf.GetTokenThreshold(args.Info.Wallet.GetAddress().Hex(), args.Info.TokenName, args.Info.Chain)
	purchaseAmount := args.Conf.GetTokenPurchaseAmount(args.Info.Wallet.GetAddress().Hex(), args.Info.TokenName, args.Info.Chain)
	total, err := h.Balance(ctx, args)
	if err != nil {
		return nil, bot.Parallel, err
	}
	log.Debugf("wallet %s token %s on chain %s total balance is %s, threshold is %s", args.Info.Wallet.GetAddress(), args.Info.TokenName, args.Info.Chain, total, threshold)
	if total.GreaterThan(threshold) {
		return nil, bot.Parallel, nil
	}

	if total.Add(purchaseAmount).LessThanOrEqual(threshold) {
		newAmount := threshold.Add(threshold.Mul(decimal.RequireFromString("0.01")))
		log.Infof("The %s %s on %s current balance is %s, amount in config is %s, balance(%s) + amount(%s) <= threshold(%s), so set amount to %s",
			args.Info.Wallet.GetAddress(), args.Info.TokenName, args.Info.Chain, total, purchaseAmount, total, purchaseAmount, threshold, newAmount)
		purchaseAmount = newAmount
	}
	return []bot.Task{
		{
			Wallet:            args.Info.Wallet.GetAddress().Hex(),
			TokenOutName:      args.Info.TokenName,
			TokenOutChainName: args.Info.Chain,
			Amount:            purchaseAmount,
			Status:            provider.TxStatusPending,
			CurrentBalance:    total,
		},
	}, bot.Queue, nil
}

func (h HelixLiquidity) Name() string {
	return "helix_liquidity"
}

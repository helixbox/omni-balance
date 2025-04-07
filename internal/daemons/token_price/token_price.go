package token_price

import (
	"context"
	"omni-balance/internal/daemons"
	"omni-balance/internal/db"
	"omni-balance/internal/models"
	"omni-balance/utils"
	"omni-balance/utils/configs"
	log "omni-balance/utils/logging"
	"omni-balance/utils/token_price"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

func init() {
	daemons.RegisterIntervalTask(daemons.Task{
		Name:            "getTokenPriceInUsdt",
		Description:     "Responsible for obtaining the token price, denominated in USDT.",
		TaskFunc:        Run,
		DefaultInterval: time.Minute * 3,
		RunOnStart:      true,
	})
}

// Run coordinates the token price aggregation process from multiple providers.
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - conf: Configuration containing source tokens to monitor
//
// Returns:
//   - error: Aggregation errors including:
//   - Provider-specific price fetch failures
//   - Database save failures
//
// The function:
// 1. Initializes synchronization primitives for concurrent execution
// 2. Iterates through all registered price providers
// 3. Fetches prices concurrently from each provider
// 4. Validates and aggregates price data
// 5. Persists validated prices to the database
func Run(ctx context.Context, conf configs.Config) error {
	var (
		w          sync.WaitGroup
		m          sync.Mutex
		tokenPrice []models.TokenPrice
	)
	providers := token_price.ListTokenPriceProviders()
	for _, provider := range providers {
		w.Add(1)
		go func(provider token_price.TokenPrice) {
			defer utils.Recover()
			defer w.Done()
			result, err := provider.GetTokenPriceInUSDT(ctx, conf.SourceTokens...)
			if err != nil {
				log.Warnf("%s get token price error: %s", provider.Name(), err)
				return
			}
			m.Lock()
			defer m.Unlock()
			for _, r := range result {
				if r.Price.LessThan(decimal.Zero) {
					log.Warnf("%s get %s token price is less than 0", provider.Name(), r.TokeName)
					continue
				}
				tokenPrice = append(tokenPrice, models.TokenPrice{
					TokenName: r.TokeName,
					Price:     r.Price,
					Source:    provider.Name(),
				})
			}

		}(provider)
	}
	w.Wait()
	return models.SaveTokenPrice(ctx, db.DB().Begin(), tokenPrice)
}

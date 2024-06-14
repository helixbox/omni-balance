package constant

import "context"

type ContextKey string

const (
	// TestPrivateKey and TestWalletAddress  is for testnet, nothing to do with mainnet
	TestPrivateKey               = "2f8f3049314beb2adffeac27ce9ea799b254cc5035673c861b1c0a09539b3de3" // gitleaks:allow
	TestWalletAddress            = "0x43Ef13E84D9992d1461a1f90CAc4653658CEA4FD"                       // gitleaks:allow
	testKey           ContextKey = "is_test"
)

func WithTestCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, testKey, true)
}

func IsTestCtx(ctx context.Context) bool {
	return ctx.Value(testKey) == true
}

package base

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPollReceiveTxHash(t *testing.T) {
	depositTxHash := "0x6b1115f3ba8f76b42e483d1c0eb1300f18f750dd5404644c6dd20bd7b21771f6"
	receiveTxHash := "0x4545b85ab81e87c12ab9d27a6342779b7c604270cfcc6c9988ccada61999b0db"

	evAddr := "0x9003d8731df107aA5E3FEADdFC165787b910Ff1e"

	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	result, err := WaitForChildTransactionReceipt(ctx, depositTxHash, evAddr)
	require.NoError(t, err)
	require.Equal(t, receiveTxHash, result)
}

func TestGetProve(t *testing.T) {
	txHash := "0xf34da0e5596f64b6cfe218bb9b4576008c00d4577a2e34066431dbf1517289eb"
	trader := "0xd1fc331dbf956e21da5c2d89caa2f98c80317d33"

	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	prove, err := getProve(ctx, txHash, trader)
	require.NoError(t, err)
	require.NotEmpty(t, prove)
}

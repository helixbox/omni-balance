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
	txHash := "0xfde3e860dd7c22c87aafd2aa171c0dc761aa5e9cfaf0faa238fc33c26e456d3b"
	trader := "0x9003d8731df107aA5E3FEADdFC165787b910Ff1e"

	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	prove, err := getProve(ctx, txHash, trader)
	require.NoError(t, err)
	require.NotEmpty(t, prove)
}

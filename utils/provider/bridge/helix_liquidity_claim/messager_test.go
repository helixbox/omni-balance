package helix_liquidity_claim

import (
	"context"
	"omni-balance/utils/chains"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestLayerzeroMessager_EncodePayload(t *testing.T) {
	data, err := new(LayerzeroMessager).EncodePayload(context.Background(), EncodePayloadArgs{
		LocalAppAddress:  common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		RemoteAppAddress: common.HexToAddress("0xbA5D580B18b6436411562981e02c8A9aA1776D10"),
		Message:          common.Hex2Bytes("7425b8b500000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000089000000000000000000000000000000000bb6a011db294ce3f3423f00eac4959e0000000000000000000000000000000000000000000000000000000000000002d84be70f3b8d54cc248da2aae1d73f2536b1e4c02fa53697991e284fd88bb7378bf1db785ff2091cf02cd03c399d441d46289729c074a99c06f7e90cfd3f77ad"),
	})
	assert.NoError(t, err)
	expected := "000000000000000000000000ba5d580b18b6436411562981e02c8a9aa1776d10000000000000000000000000ba5d580b18b6436411562981e02c8a9aa1776d10000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c47425b8b500000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000089000000000000000000000000000000000bb6a011db294ce3f3423f00eac4959e0000000000000000000000000000000000000000000000000000000000000002d84be70f3b8d54cc248da2aae1d73f2536b1e4c02fa53697991e284fd88bb7378bf1db785ff2091cf02cd03c399d441d46289729c074a99c06f7e90cfd3f77ad00000000000000000000000000000000000000000000000000000000"
	assert.Equal(t, expected, common.Bytes2Hex(data))
}

func TestLayerzeroMessager_Params(t *testing.T) {
	client, err := chains.NewTryClient(context.TODO(), []string{
		"https://polygon-rpc.com",
	})
	assert.NoError(t, err)
	params, err := new(LayerzeroMessager).Params(context.Background(), MessagerArgs{
		ContractAddress: common.HexToAddress("0x463D1730a8527CA58d48EF70C7460B9920346567"),
		FromChainId:     137,
		ToChainId:       42161,
		RemoteMessager:  common.HexToAddress("0x509354A4ebf98aCC7a65d2264694A65a2938cac9"),
		Payload:         common.Hex2Bytes("000000000000000000000000ba5d580b18b6436411562981e02c8a9aa1776d10000000000000000000000000ba5d580b18b6436411562981e02c8a9aa1776d10000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c47425b8b500000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000089000000000000000000000000000000000bb6a011db294ce3f3423f00eac4959e0000000000000000000000000000000000000000000000000000000000000002d84be70f3b8d54cc248da2aae1d73f2536b1e4c02fa53697991e284fd88bb7378bf1db785ff2091cf02cd03c399d441d46289729c074a99c06f7e90cfd3f77ad00000000000000000000000000000000000000000000000000000000"),
		Refunder:        common.HexToAddress("0x000000000Bb6a011dB294ce3F3423f00EAc4959e"),
		Client:          client,
	})
	assert.NoError(t, err)
	assert.Equal(t, "0x000000000Bb6a011dB294ce3F3423f00EAc4959e", common.BytesToAddress(params.ExtParams).Hex())
}

func TestMessagePortMessager_EncodePayload(t *testing.T) {
	type args struct {
		ctx  context.Context
		args EncodePayloadArgs
	}
	tests := []struct {
		name        string
		m           MessagePortMessager
		args        args
		wantPayload []byte
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MessagePortMessager{}
			gotPayload, err := m.EncodePayload(tt.args.ctx, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessagePortMessager.EncodePayload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPayload, tt.wantPayload) {
				t.Errorf("MessagePortMessager.EncodePayload() = %v, want %v", gotPayload, tt.wantPayload)
			}
		})
	}
}

func TestMessagePortMessager_Params(t *testing.T) {
	type args struct {
		ctx  context.Context
		args MessagerArgs
	}
	tests := []struct {
		name       string
		m          MessagePortMessager
		args       args
		wantParams MessagerParams
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MessagePortMessager{}
			gotParams, err := m.Params(tt.args.ctx, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("MessagePortMessager.Params() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotParams, tt.wantParams) {
				t.Errorf("MessagePortMessager.Params() = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}

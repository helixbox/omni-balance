package enclave

import (
	"encoding/json"

	base_portal "omni-balance/utils/enclave/router/base/portal"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type BaseClaimRequest struct {
	BaseClaim BaseClaim `json:"baseClaim"`
}

func (a BaseClaimRequest) GetRequestType() string {
	return "base_claim"
}

// function finalizeWithdrawalTransaction(WithdrawalTransaction memory tx);
type BaseClaim struct {
	Tx   WithdrawalTransaction `json:"tx"`
	Meta Meta                  `json:"meta"`
}

func (c *Client) SignBaseClaim(tx *types.Transaction, chainID int64) (*types.Transaction, error) {
	req, err := buildClaimRequest(tx.Data(), tx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return c.signRequest(req, tx, chainID)
}

func buildClaimRequest(input []byte, tx *types.Transaction) (BaseClaimRequest, error) {
	if len(input) < 4 {
		return BaseClaimRequest{}, errors.New("data too short")
	}

	portalAbi, err := base_portal.BasePortalMetaData.ParseABI()
	if err != nil {
		return BaseClaimRequest{}, errors.WithStack(err)
	}

	args, err := portalAbi.Methods["finalizeWithdrawalTransactionExternalProof"].Inputs.Unpack(input[4:])
	if err != nil {
		return BaseClaimRequest{}, errors.Wrap(err, "unpack")
	}

	if len(args) != 2 {
		return BaseClaimRequest{}, errors.New("invalid number of args")
	}

	b, _ := json.Marshal(args[0])
	var txw base_portal.TypesWithdrawalTransaction
	json.Unmarshal(b, &txw)

	prove := BaseClaim{
		Tx: WithdrawalTransaction{
			Nonce:    txw.Nonce.String(),
			Sender:   txw.Sender,
			Target:   txw.Target,
			Value:    txw.Value.String(),
			GasLimit: txw.GasLimit.String(),
			Data:     common.Bytes2Hex(txw.Data[:]),
		},
		Meta: Meta{
			Nonce:                tx.Nonce(),
			GasLimit:             tx.Gas(),
			MaxFeePerGas:         tx.GasFeeCap(),
			MaxPriorityFeePerGas: tx.GasTipCap(),
		},
	}

	return BaseClaimRequest{BaseClaim: prove}, nil
}

package chains

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/pkg/errors"
)

// SignTypedData - Sign typed data
func SignTypedData(typedData apitypes.TypedData, signFn func(msg []byte) (sig []byte, err error)) (sig []byte, err error) {
	hash, err := EncodeForSigning(typedData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to encode typed data")
	}
	sig, err = signFn(hash.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign typed data")
	}
	sig[64] += 27
	return
}

// EncodeForSigning - Encoding the typed data
func EncodeForSigning(typedData apitypes.TypedData) (hash common.Hash, err error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return hash, errors.Wrap(err, "failed to hash domain")
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return hash, errors.Wrap(err, "failed to hash typed data")
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash = common.BytesToHash(crypto.Keccak256(rawData))
	return
}

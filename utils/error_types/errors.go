package error_types

import "errors"

var (
	ErrNotFound                     = errors.New("not found")
	ErrInsufficientLiquidity        = errors.New("insufficient liquidity")
	ErrNativeTokenInsufficient      = errors.New("native token insufficient")
	ErrPrivateKeyEmpty              = errors.New("private key empty")
	ErrPrivateKeyNotMatch           = errors.New("private key not match")
	ErrUnsupportedMultiSign         = errors.New("unsupported multi sign")
	ErrInsufficientBalance          = errors.New("insufficient balance")
	ErrUnsupportedTokenAndChain     = errors.New("unsupported token and chain")
	ErrUnsupportedWalletType        = errors.New("unsupported wallet type")
	ErrUnsupportedActions           = errors.New("unsupported actions")
	ErrWalletLocked                 = errors.New("wallet locked")
	ErrNoProvider                   = errors.New("no available provider was found")
	ErrEnclaveNotSupportNativeToken = errors.New("enclave not support native token")
)

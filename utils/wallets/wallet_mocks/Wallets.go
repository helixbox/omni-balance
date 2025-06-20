// Code generated by mockery v2.43.2. DO NOT EDIT.

package wallet_mocks

import (
	context "context"

	common "github.com/ethereum/go-ethereum/common"

	decimal "github.com/shopspring/decimal"

	mock "github.com/stretchr/testify/mock"

	simulated "github.com/ethereum/go-ethereum/ethclient/simulated"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Wallets is an autogenerated mock type for the Wallets type
type Wallets struct {
	mock.Mock
}


func (s *Wallets) Name(_ context.Context) string {
	return "safe"
}

// CheckFullAccess provides a mock function with given fields: ctx
func (_m *Wallets) CheckFullAccess(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CheckFullAccess")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAddress provides a mock function with given fields: isReal
func (_m *Wallets) GetAddress(isReal ...bool) common.Address {
	_va := make([]interface{}, len(isReal))
	for _i := range isReal {
		_va[_i] = isReal[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetAddress")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func(...bool) common.Address); ok {
		r0 = rf(isReal...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// GetBalance provides a mock function with given fields: ctx, tokenAddress, decimals, client
func (_m *Wallets) GetBalance(ctx context.Context, tokenAddress common.Address, decimals int32, client simulated.Client) (decimal.Decimal, error) {
	ret := _m.Called(ctx, tokenAddress, decimals, client)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 decimal.Decimal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, int32, simulated.Client) (decimal.Decimal, error)); ok {
		return rf(ctx, tokenAddress, decimals, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, int32, simulated.Client) decimal.Decimal); ok {
		r0 = rf(ctx, tokenAddress, decimals, client)
	} else {
		r0 = ret.Get(0).(decimal.Decimal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, int32, simulated.Client) error); ok {
		r1 = rf(ctx, tokenAddress, decimals, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExternalBalance provides a mock function with given fields: ctx, tokenAddress, decimals, client
func (_m *Wallets) GetExternalBalance(ctx context.Context, tokenAddress common.Address, decimals int32, client simulated.Client) (decimal.Decimal, error) {
	ret := _m.Called(ctx, tokenAddress, decimals, client)

	if len(ret) == 0 {
		panic("no return value specified for GetExternalBalance")
	}

	var r0 decimal.Decimal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, int32, simulated.Client) (decimal.Decimal, error)); ok {
		return rf(ctx, tokenAddress, decimals, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, int32, simulated.Client) decimal.Decimal); ok {
		r0 = rf(ctx, tokenAddress, decimals, client)
	} else {
		r0 = ret.Get(0).(decimal.Decimal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, int32, simulated.Client) error); ok {
		r1 = rf(ctx, tokenAddress, decimals, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNonce provides a mock function with given fields: ctx, client
func (_m *Wallets) GetNonce(ctx context.Context, client simulated.Client) (uint64, error) {
	ret := _m.Called(ctx, client)

	if len(ret) == 0 {
		panic("no return value specified for GetNonce")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, simulated.Client) (uint64, error)); ok {
		return rf(ctx, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, simulated.Client) uint64); ok {
		r0 = rf(ctx, client)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, simulated.Client) error); ok {
		r1 = rf(ctx, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRealHash provides a mock function with given fields: ctx, txHash, client
func (_m *Wallets) GetRealHash(ctx context.Context, txHash common.Hash, client simulated.Client) (common.Hash, error) {
	ret := _m.Called(ctx, txHash, client)

	if len(ret) == 0 {
		panic("no return value specified for GetRealHash")
	}

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, simulated.Client) (common.Hash, error)); ok {
		return rf(ctx, txHash, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, simulated.Client) common.Hash); ok {
		r0 = rf(ctx, txHash, client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, simulated.Client) error); ok {
		r1 = rf(ctx, txHash, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsDifferentAddress provides a mock function with given fields:
func (_m *Wallets) IsDifferentAddress() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsDifferentAddress")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsSupportEip712 provides a mock function with given fields:
func (_m *Wallets) IsSupportEip712() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSupportEip712")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MarshalJSON provides a mock function with given fields:
func (_m *Wallets) MarshalJSON() ([]byte, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MarshalJSON")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]byte, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendTransaction provides a mock function with given fields: ctx, tx, client
func (_m *Wallets) SendTransaction(ctx context.Context, tx *types.DynamicFeeTx, client simulated.Client) (common.Hash, error) {
	ret := _m.Called(ctx, tx, client)

	if len(ret) == 0 {
		panic("no return value specified for SendTransaction")
	}

	var r0 common.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.DynamicFeeTx, simulated.Client) (common.Hash, error)); ok {
		return rf(ctx, tx, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.DynamicFeeTx, simulated.Client) common.Hash); ok {
		r0 = rf(ctx, tx, client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.DynamicFeeTx, simulated.Client) error); ok {
		r1 = rf(ctx, tx, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignRawMessage provides a mock function with given fields: msg
func (_m *Wallets) SignRawMessage(msg []byte) ([]byte, error) {
	ret := _m.Called(msg)

	if len(ret) == 0 {
		panic("no return value specified for SignRawMessage")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(msg)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitTransaction provides a mock function with given fields: ctx, txHash, client
func (_m *Wallets) WaitTransaction(ctx context.Context, txHash common.Hash, client simulated.Client) error {
	ret := _m.Called(ctx, txHash, client)

	if len(ret) == 0 {
		panic("no return value specified for WaitTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, simulated.Client) error); ok {
		r0 = rf(ctx, txHash, client)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWallets creates a new instance of Wallets. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWallets(t interface {
	mock.TestingT
	Cleanup(func())
}) *Wallets {
	mock := &Wallets{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

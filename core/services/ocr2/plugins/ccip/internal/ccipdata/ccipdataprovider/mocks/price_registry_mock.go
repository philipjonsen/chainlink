// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	ccip "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// PriceRegistry is an autogenerated mock type for the PriceRegistry type
type PriceRegistry struct {
	mock.Mock
}

type PriceRegistry_Expecter struct {
	mock *mock.Mock
}

func (_m *PriceRegistry) EXPECT() *PriceRegistry_Expecter {
	return &PriceRegistry_Expecter{mock: &_m.Mock}
}

// NewPriceRegistryReader provides a mock function with given fields: ctx, addr
func (_m *PriceRegistry) NewPriceRegistryReader(ctx context.Context, addr ccip.Address) (ccip.PriceRegistryReader, error) {
	ret := _m.Called(ctx, addr)

	if len(ret) == 0 {
		panic("no return value specified for NewPriceRegistryReader")
	}

	var r0 ccip.PriceRegistryReader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ccip.Address) (ccip.PriceRegistryReader, error)); ok {
		return rf(ctx, addr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ccip.Address) ccip.PriceRegistryReader); ok {
		r0 = rf(ctx, addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ccip.PriceRegistryReader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ccip.Address) error); ok {
		r1 = rf(ctx, addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PriceRegistry_NewPriceRegistryReader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewPriceRegistryReader'
type PriceRegistry_NewPriceRegistryReader_Call struct {
	*mock.Call
}

// NewPriceRegistryReader is a helper method to define mock.On call
//   - ctx context.Context
//   - addr ccip.Address
func (_e *PriceRegistry_Expecter) NewPriceRegistryReader(ctx interface{}, addr interface{}) *PriceRegistry_NewPriceRegistryReader_Call {
	return &PriceRegistry_NewPriceRegistryReader_Call{Call: _e.mock.On("NewPriceRegistryReader", ctx, addr)}
}

func (_c *PriceRegistry_NewPriceRegistryReader_Call) Run(run func(ctx context.Context, addr ccip.Address)) *PriceRegistry_NewPriceRegistryReader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ccip.Address))
	})
	return _c
}

func (_c *PriceRegistry_NewPriceRegistryReader_Call) Return(_a0 ccip.PriceRegistryReader, _a1 error) *PriceRegistry_NewPriceRegistryReader_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PriceRegistry_NewPriceRegistryReader_Call) RunAndReturn(run func(context.Context, ccip.Address) (ccip.PriceRegistryReader, error)) *PriceRegistry_NewPriceRegistryReader_Call {
	_c.Call.Return(run)
	return _c
}

// NewPriceRegistry creates a new instance of PriceRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPriceRegistry(t interface {
	mock.TestingT
	Cleanup(func())
}) *PriceRegistry {
	mock := &PriceRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "homework/internal/pack/domain"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Calc provides a mock function with given fields: _a0, _a1
func (_m *Service) Calc(_a0 context.Context, _a1 domain.PacksCalcRequest) (*domain.CalculatePacksResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Calc")
	}

	var r0 *domain.CalculatePacksResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PacksCalcRequest) (*domain.CalculatePacksResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PacksCalcRequest) *domain.CalculatePacksResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.CalculatePacksResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PacksCalcRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

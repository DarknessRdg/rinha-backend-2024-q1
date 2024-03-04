// Code generated by mockery v2.42.0. DO NOT EDIT.

package service

import (
	dto "github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/dto"
	mock "github.com/stretchr/testify/mock"
)

// MockIAccountService is an autogenerated mock type for the IAccountService type
type MockIAccountService struct {
	mock.Mock
}

type MockIAccountService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockIAccountService) EXPECT() *MockIAccountService_Expecter {
	return &MockIAccountService_Expecter{mock: &_m.Mock}
}

// CreditOrDebit provides a mock function with given fields: accountId, amountCents, operation
func (_m *MockIAccountService) CreditOrDebit(accountId int, amountCents int, operation string) (dto.AccountDto, error) {
	ret := _m.Called(accountId, amountCents, operation)

	if len(ret) == 0 {
		panic("no return value specified for CreditOrDebit")
	}

	var r0 dto.AccountDto
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string) (dto.AccountDto, error)); ok {
		return rf(accountId, amountCents, operation)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) dto.AccountDto); ok {
		r0 = rf(accountId, amountCents, operation)
	} else {
		r0 = ret.Get(0).(dto.AccountDto)
	}

	if rf, ok := ret.Get(1).(func(int, int, string) error); ok {
		r1 = rf(accountId, amountCents, operation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockIAccountService_CreditOrDebit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreditOrDebit'
type MockIAccountService_CreditOrDebit_Call struct {
	*mock.Call
}

// CreditOrDebit is a helper method to define mock.On call
//   - accountId int
//   - amountCents int
//   - operation string
func (_e *MockIAccountService_Expecter) CreditOrDebit(accountId interface{}, amountCents interface{}, operation interface{}) *MockIAccountService_CreditOrDebit_Call {
	return &MockIAccountService_CreditOrDebit_Call{Call: _e.mock.On("CreditOrDebit", accountId, amountCents, operation)}
}

func (_c *MockIAccountService_CreditOrDebit_Call) Run(run func(accountId int, amountCents int, operation string)) *MockIAccountService_CreditOrDebit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].(string))
	})
	return _c
}

func (_c *MockIAccountService_CreditOrDebit_Call) Return(_a0 dto.AccountDto, _a1 error) *MockIAccountService_CreditOrDebit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockIAccountService_CreditOrDebit_Call) RunAndReturn(run func(int, int, string) (dto.AccountDto, error)) *MockIAccountService_CreditOrDebit_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockIAccountService creates a new instance of MockIAccountService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIAccountService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIAccountService {
	mock := &MockIAccountService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

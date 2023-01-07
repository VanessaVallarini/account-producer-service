// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	models "account-producer-service/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IAccountRepository is an autogenerated mock type for the IAccountRepository type
type IAccountRepository struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: ctx, a
func (_m *IAccountRepository) GetByEmail(ctx context.Context, a models.AccountRequestByEmail) (*models.Account, error) {
	ret := _m.Called(ctx, a)

	var r0 *models.Account
	if rf, ok := ret.Get(0).(func(context.Context, models.AccountRequestByEmail) *models.Account); ok {
		r0 = rf(ctx, a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.AccountRequestByEmail) error); ok {
		r1 = rf(ctx, a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *IAccountRepository) List(ctx context.Context) ([]models.Account, error) {
	ret := _m.Called(ctx)

	var r0 []models.Account
	if rf, ok := ret.Get(0).(func(context.Context) []models.Account); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIAccountRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewIAccountRepository creates a new instance of IAccountRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAccountRepository(t mockConstructorTestingTNewIAccountRepository) *IAccountRepository {
	mock := &IAccountRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
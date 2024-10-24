// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	"github.com/onattech/invest/models"
	mock "github.com/stretchr/testify/mock"
)

// SignupService is an autogenerated mock type for the SignupService type
type SignupService struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, user
func (_m *SignupService) Create(c context.Context, user *models.User) error {
	ret := _m.Called(c, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAccessToken provides a mock function with given fields: user, secret, expiry
func (_m *SignupService) CreateAccessToken(user *models.User, secret string, expiry int) (string, error) {
	ret := _m.Called(user, secret, expiry)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.User, string, int) string); ok {
		r0 = rf(user, secret, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.User, string, int) error); ok {
		r1 = rf(user, secret, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: user, secret, expiry
func (_m *SignupService) CreateRefreshToken(user *models.User, secret string, expiry int) (string, error) {
	ret := _m.Called(user, secret, expiry)

	var r0 string
	if rf, ok := ret.Get(0).(func(*models.User, string, int) string); ok {
		r0 = rf(user, secret, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.User, string, int) error); ok {
		r1 = rf(user, secret, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: c, email
func (_m *SignupService) GetUserByEmail(c context.Context, email string) (models.User, error) {
	ret := _m.Called(c, email)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(c, email)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSignupService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSignupService creates a new instance of SignupService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSignupService(t mockConstructorTestingTNewSignupService) *SignupService {
	mock := &SignupService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

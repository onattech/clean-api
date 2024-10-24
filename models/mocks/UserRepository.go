// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	"github.com/onattech/invest/models"
	mock "github.com/stretchr/testify/mock"
)

// UserStore is an autogenerated mock type for the UserStore type
type UserStore struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, user
func (_m *UserStore) Create(c context.Context, user *models.User) error {
	ret := _m.Called(c, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: c
func (_m *UserStore) Fetch(c context.Context) ([]models.User, error) {
	ret := _m.Called(c)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context) []models.User); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: c, email
func (_m *UserStore) GetByEmail(c context.Context, email string) (models.User, error) {
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

// GetByID provides a mock function with given fields: c, id
func (_m *UserStore) GetByID(c context.Context, id string) (models.User, error) {
	ret := _m.Called(c, id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserStore creates a new instance of UserStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserStore(t mockConstructorTestingTNewUserStore) *UserStore {
	mock := &UserStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	"github.com/onattech/invest/models"
	mock "github.com/stretchr/testify/mock"
)

// TaskService is an autogenerated mock type for the TaskService type
type TaskService struct {
	mock.Mock
}

// Create provides a mock function with given fields: c, task
func (_m *TaskService) Create(c context.Context, task *models.Task) error {
	ret := _m.Called(c, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Task) error); ok {
		r0 = rf(c, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchByUserID provides a mock function with given fields: c, userID
func (_m *TaskService) FetchByUserID(c context.Context, userID string) ([]models.Task, error) {
	ret := _m.Called(c, userID)

	var r0 []models.Task
	if rf, ok := ret.Get(0).(func(context.Context, string) []models.Task); ok {
		r0 = rf(c, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTaskService interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskService creates a new instance of TaskService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskService(t mockConstructorTestingTNewTaskService) *TaskService {
	mock := &TaskService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
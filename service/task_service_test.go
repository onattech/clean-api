package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/onattech/invest/models"
	"github.com/onattech/invest/models/mocks"
	"github.com/onattech/invest/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchByUserID(t *testing.T) {
	mockTaskStore := new(mocks.TaskStore)
	userUUID := uuid.New()
	userID := userUUID

	t.Run("success", func(t *testing.T) {
		mockTask := models.Task{
			ID:     uuid.New(),
			Title:  "Test Title",
			UserID: userUUID, // Using uuid.UUID for the user ID
		}

		mockListTask := make([]models.Task, 0)
		mockListTask = append(mockListTask, mockTask)

		// Mock the FetchByUserID method call
		mockTaskStore.On("FetchByUserID", mock.Anything, userID).Return(mockListTask, nil).Once()

		// Create TaskService
		u := service.NewTaskService(mockTaskStore, time.Second*2)

		// Call the method
		list, err := u.FetchByUserID(context.Background(), userID)

		// Assertions
		assert.NoError(t, err)
		assert.NotNil(t, list)
		assert.Len(t, list, len(mockListTask))

		// Verify that expectations were met
		mockTaskStore.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Mock the FetchByUserID method call returning an error
		mockTaskStore.On("FetchByUserID", mock.Anything, userID).Return(nil, errors.New("Unexpected")).Once()

		// Create TaskService
		u := service.NewTaskService(mockTaskStore, time.Second*2)

		// Call the method
		list, err := u.FetchByUserID(context.Background(), userID)

		// Assertions
		assert.Error(t, err)
		assert.Nil(t, list)

		// Verify that expectations were met
		mockTaskStore.AssertExpectations(t)
	})
}

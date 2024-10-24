package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/onattech/invest/models"
)

type taskService struct {
	taskStore      models.TaskStore
	contextTimeout time.Duration
}

func NewTaskService(taskStore models.TaskStore, timeout time.Duration) models.TaskService {
	return &taskService{
		taskStore:      taskStore,
		contextTimeout: timeout,
	}
}

func (service *taskService) Create(c context.Context, task *models.Task) error {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()
	return service.taskStore.Create(ctx, task)
}

func (service *taskService) FetchByUserID(c context.Context, userID uuid.UUID) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()
	return service.taskStore.FetchByUserID(ctx, userID)
}

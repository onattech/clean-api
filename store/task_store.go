package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/onattech/invest/models"
	"gorm.io/gorm"
)

type taskStore struct {
	db *gorm.DB
}

func NewTaskStore(db *gorm.DB) models.TaskStore {
	return &taskStore{
		db: db,
	}
}

func (ts *taskStore) Create(ctx context.Context, task *models.Task) error {
	// Create the task in the database
	return ts.db.WithContext(ctx).Create(task).Error
}

func (ts *taskStore) FetchByUserID(ctx context.Context, userID uuid.UUID) ([]models.Task, error) {
	var tasks []models.Task

	// Find all tasks for the given user ID
	err := ts.db.WithContext(ctx).Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

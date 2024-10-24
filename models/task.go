package models

import (
	"context"

	"github.com/google/uuid"
)

type Task struct {
	ID     uuid.UUID `json:"-"`                                     // Use UUID for Task ID
	Title  string    `form:"title" binding:"required" json:"title"` // Task title with required validation
	UserID uuid.UUID `json:"-"`                                     // Use UUID for UserID
}

type TaskService interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID uuid.UUID) ([]Task, error)
}

type TaskStore interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID uuid.UUID) ([]Task, error)
}

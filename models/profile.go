package models

import (
	"context"

	"github.com/google/uuid"
)

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ProfileService interface {
	GetProfileByID(c context.Context, userID uuid.UUID) (*Profile, error)
}

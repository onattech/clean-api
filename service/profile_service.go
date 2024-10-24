package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/onattech/invest/models"
)

type profileService struct {
	userStore      models.UserStore
	contextTimeout time.Duration
}

func NewProfileService(userStore models.UserStore, timeout time.Duration) models.ProfileService {
	return &profileService{
		userStore:      userStore,
		contextTimeout: timeout,
	}
}

func (service *profileService) GetProfileByID(c context.Context, userID uuid.UUID) (*models.Profile, error) {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()

	user, err := service.userStore.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &models.Profile{Name: user.Name, Email: user.Email}, nil
}

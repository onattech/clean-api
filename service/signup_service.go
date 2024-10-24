package service

import (
	"context"
	"time"

	"github.com/onattech/invest/models"
	"github.com/onattech/invest/utils/tokenutil"
)

type signupService struct {
	userStore      models.UserStore
	contextTimeout time.Duration
}

func NewSignupService(userStore models.UserStore, timeout time.Duration) models.SignupService {
	return &signupService{
		userStore:      userStore,
		contextTimeout: timeout,
	}
}

func (service *signupService) Create(c context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()
	return service.userStore.Create(ctx, user)
}

func (service *signupService) GetUserByEmail(c context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()
	return service.userStore.GetByEmail(ctx, email)
}

func (service *signupService) CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (service *signupService) CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

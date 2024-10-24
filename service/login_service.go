package service

import (
	"context"
	"time"

	"github.com/onattech/invest/models"
	"github.com/onattech/invest/utils/tokenutil"
)

type loginService struct {
	userStore      models.UserStore
	contextTimeout time.Duration
}

func NewLoginService(userStore models.UserStore, timeout time.Duration) models.LoginService {
	return &loginService{
		userStore:      userStore,
		contextTimeout: timeout,
	}
}

func (service *loginService) GetUserByEmail(c context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()
	return service.userStore.GetByEmail(ctx, email)
}

func (service *loginService) CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (service *loginService) CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

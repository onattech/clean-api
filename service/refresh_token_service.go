package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/onattech/invest/models"
	"github.com/onattech/invest/utils/tokenutil"
)

type refreshTokenService struct {
	userStore      models.UserStore
	contextTimeout time.Duration
}

func NewRefreshTokenService(userStore models.UserStore, timeout time.Duration) models.RefreshTokenService {
	return &refreshTokenService{
		userStore:      userStore,
		contextTimeout: timeout,
	}
}

func (service *refreshTokenService) GetUserByID(c context.Context, id uuid.UUID) (models.User, error) {
	ctx, cancel := context.WithTimeout(c, service.contextTimeout)
	defer cancel()
	return service.userStore.GetByID(ctx, id)
}

func (service *refreshTokenService) CreateAccessToken(user *models.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (service *refreshTokenService) CreateRefreshToken(user *models.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (service *refreshTokenService) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}

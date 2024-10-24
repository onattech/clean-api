package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onattech/invest/config"
	"github.com/onattech/invest/models"
)

type RefreshTokenHandler struct {
	RefreshTokenService models.RefreshTokenService
	Env                 *config.Env
}

func (rth *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var request models.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := rth.RefreshTokenService.ExtractIDFromToken(request.RefreshToken, rth.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "User not found"})
		return
	}

	// Convert the string id to uuid.UUID
	userUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid user ID format"})
		return
	}

	user, err := rth.RefreshTokenService.GetUserByID(c, userUUID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := rth.RefreshTokenService.CreateAccessToken(&user, rth.Env.AccessTokenSecret, rth.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rth.RefreshTokenService.CreateRefreshToken(&user, rth.Env.RefreshTokenSecret, rth.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := models.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}

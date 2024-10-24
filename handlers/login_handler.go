package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/onattech/invest/config"
	"github.com/onattech/invest/models"
)

type LoginHandler struct {
	LoginService models.LoginService
	Env          *config.Env
}

func (lh *LoginHandler) Login(c *gin.Context) {
	var request models.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lh.LoginService.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lh.LoginService.CreateAccessToken(&user, lh.Env.AccessTokenSecret, lh.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lh.LoginService.CreateRefreshToken(&user, lh.Env.RefreshTokenSecret, lh.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}

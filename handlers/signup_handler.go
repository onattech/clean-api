package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onattech/invest/config"
	"github.com/onattech/invest/models"
	"golang.org/x/crypto/bcrypt"
)

type SignupHandler struct {
	SignupService models.SignupService
	Env           *config.Env
}

// Signup handles user registration
func (sh *SignupHandler) Signup(c *gin.Context) {
	var request models.SignupRequest

	// Bind the incoming request body to the SignupRequest struct
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Check if the user already exists by email
	_, err = sh.SignupService.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, models.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	// Hash the password
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Update the request password with the hashed password
	request.Password = string(encryptedPassword)

	// Create the new user with a UUID for the ID
	user := models.User{
		ID:       uuid.New(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	// Attempt to create the user in the database
	err = sh.SignupService.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Generate access and refresh tokens for the new user
	accessToken, err := sh.SignupService.CreateAccessToken(&user, sh.Env.AccessTokenSecret, sh.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := sh.SignupService.CreateRefreshToken(&user, sh.Env.RefreshTokenSecret, sh.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Prepare the signup response with tokens
	signupResponse := models.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	// Return the signup response
	c.JSON(http.StatusOK, signupResponse)
}

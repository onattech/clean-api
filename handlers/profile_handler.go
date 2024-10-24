package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onattech/invest/models"
)

type ProfileHandler struct {
	ProfileService models.ProfileService
}

func (ph *ProfileHandler) Fetch(c *gin.Context) {
	// Fetch userID as string
	userIDStr := c.GetString("x-user-id")

	// Convert userID string to uuid.UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid user ID format"})
		return
	}

	// Pass uuid.UUID type to GetProfileByID
	profile, err := ph.ProfileService.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

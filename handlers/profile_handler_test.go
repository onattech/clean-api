package handlers_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onattech/invest/handlers"
	"github.com/onattech/invest/models"
	"github.com/onattech/invest/models/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// setUserID is a middleware that sets the x-user-id header for the request
func setUserID(userID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("x-user-id", userID)
		c.Next()
	}
}

func TestFetch(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		// Create a mock profile
		mockProfile := &models.Profile{
			Name:  "Test Name",
			Email: "test@gmail.com",
		}

		// Generate a UUID for the userID (as would be used in PostgreSQL)
		userID := uuid.New().String()

		// Create a mock ProfileService
		mockProfileService := new(mocks.ProfileService)

		// Mock the GetProfileByID function to return the mockProfile
		mockProfileService.On("GetProfileByID", mock.Anything, userID).Return(mockProfile, nil)

		// Initialize Gin
		gin := gin.Default()

		// Create a test recorder to record HTTP responses
		rec := httptest.NewRecorder()

		// Create a ProfileHandler with the mock ProfileService
		pc := &handlers.ProfileHandler{
			ProfileService: mockProfileService,
		}

		// Set up the middleware and route
		gin.Use(setUserID(userID))
		gin.GET("/profile", pc.Fetch)

		// Marshal the mockProfile into JSON format
		body, err := json.Marshal(mockProfile)
		assert.NoError(t, err)

		bodyString := string(body)

		// Create a new GET request
		req := httptest.NewRequest(http.MethodGet, "/profile", nil)

		// Serve the request
		gin.ServeHTTP(rec, req)

		// Assert that the status code is 200 OK
		assert.Equal(t, http.StatusOK, rec.Code)

		// Assert that the response body matches the expected JSON body
		assert.Equal(t, bodyString, rec.Body.String())

		// Ensure all expectations were met
		mockProfileService.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Generate a UUID for the userID
		userID := uuid.New().String()

		// Create a mock ProfileService
		mockProfileService := new(mocks.ProfileService)

		// Create a custom error to simulate a failure
		customErr := errors.New("Unexpected")

		// Mock the GetProfileByID function to return an error
		mockProfileService.On("GetProfileByID", mock.Anything, userID).Return(nil, customErr)

		// Initialize Gin
		gin := gin.Default()

		// Create a test recorder to record HTTP responses
		rec := httptest.NewRecorder()

		// Create a ProfileHandler with the mock ProfileService
		pc := &handlers.ProfileHandler{
			ProfileService: mockProfileService,
		}

		// Set up the middleware and route
		gin.Use(setUserID(userID))
		gin.GET("/profile", pc.Fetch)

		// Marshal the error response into JSON format
		body, err := json.Marshal(models.ErrorResponse{Message: customErr.Error()})
		assert.NoError(t, err)

		bodyString := string(body)

		// Create a new GET request
		req := httptest.NewRequest(http.MethodGet, "/profile", nil)

		// Serve the request
		gin.ServeHTTP(rec, req)

		// Assert that the status code is 500 Internal Server Error
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		// Assert that the response body matches the expected JSON error response
		assert.Equal(t, bodyString, rec.Body.String())

		// Ensure all expectations were met
		mockProfileService.AssertExpectations(t)
	})
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/onattech/invest/models"
)

type TaskHandler struct {
	TaskService models.TaskService
}

// Create handles the creation of a new task
func (th *TaskHandler) Create(c *gin.Context) {
	var task models.Task

	// Bind the incoming request body to the Task struct
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Extract the user ID from the context
	userID := c.GetString("x-user-id")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	// Assign a new UUID to the task and associate it with the user
	task.ID = uuid.New()
	task.UserID = userUUID

	// Call the use case to create the task in the database
	err = th.TaskService.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Return a success message upon successful creation
	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Task created successfully",
	})
}

// Fetch retrieves tasks associated with the authenticated user
func (th *TaskHandler) Fetch(c *gin.Context) {
	// Extract the user ID from the context
	userID := c.GetString("x-user-id")

	// Parse the user ID into a UUID format
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "Invalid user ID"})
		return
	}

	// Fetch tasks by the user ID
	tasks, err := th.TaskService.FetchByUserID(c, userUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}

	// Return the fetched tasks
	c.JSON(http.StatusOK, tasks)
}

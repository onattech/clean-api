package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onattech/invest/config"
	"github.com/onattech/invest/handlers"
	"github.com/onattech/invest/service"
	"github.com/onattech/invest/store"
	"gorm.io/gorm"
)

func RegisterTaskRouter(env *config.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	store := store.NewTaskStore(db)
	handler := &handlers.TaskHandler{
		TaskService: service.NewTaskService(store, timeout),
	}
	group.GET("/task", handler.Fetch)
	group.POST("/task", handler.Create)
}

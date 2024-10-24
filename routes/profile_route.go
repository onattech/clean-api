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

func RegisterProfileRouter(env *config.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	store := store.NewUserStore(db)
	handler := &handlers.ProfileHandler{
		ProfileService: service.NewProfileService(store, timeout),
	}
	group.GET("/profile", handler.Fetch)
}

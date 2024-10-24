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

func RegisterRefreshTokenRouter(env *config.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	store := store.NewUserStore(db)
	handler := &handlers.RefreshTokenHandler{
		RefreshTokenService: service.NewRefreshTokenService(store, timeout),
		Env:                 env,
	}
	group.POST("/refresh", handler.RefreshToken)
}

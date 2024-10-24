package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onattech/invest/config"
	"github.com/onattech/invest/middleware"
	"gorm.io/gorm"
)

// RegisterRoutes sets up both public and protected routes for the application
func RegisterRoutes(env *config.Env, timeout time.Duration, db *gorm.DB, router *gin.Engine) {
	// Set up public routes
	publicRouter := router.Group("")
	registerPublicRoutes(env, timeout, db, publicRouter)

	// Set up protected routes with JWT authentication middleware
	protectedRouter := router.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	registerProtectedRoutes(env, timeout, db, protectedRouter)
}

// registerPublicRoutes registers all public APIs
func registerPublicRoutes(env *config.Env, timeout time.Duration, db *gorm.DB, router *gin.RouterGroup) {
	RegisterSignupRouter(env, timeout, db, router)
	RegisterLoginRouter(env, timeout, db, router)
	RegisterRefreshTokenRouter(env, timeout, db, router)
}

// registerProtectedRoutes registers all protected APIs
func registerProtectedRoutes(env *config.Env, timeout time.Duration, db *gorm.DB, router *gin.RouterGroup) {
	RegisterProfileRouter(env, timeout, db, router)
	RegisterTaskRouter(env, timeout, db, router)
}

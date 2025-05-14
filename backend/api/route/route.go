package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(cfg *config.RuntimeConfig, db *gorm.DB, cache *redis.Client, objStore *s3.Client, g *gin.Engine) {
	// Statics
	g.Static("/uploads", "./static/uploads")

	// Probes
	probeRouter := g.Group("")
	NewReadinessRouter(probeRouter)
	NewLivenessRouter(probeRouter)

	// Public APIs
	publicRouter := g.Group("")
	NewSignupRouter(cfg, db, cache, publicRouter)
	NewSigninRouter(cfg, db, cache, publicRouter)
	NewRefreshTokenRouter(cfg, db, publicRouter)
	NewSignoutRouter(cfg, cache, publicRouter)
	NewForgotPasswordRouter(cfg, db, cache, publicRouter)

	// Protected APIs
	protectedRouter := g.Group("v1")
	userGroup := protectedRouter.Group("user")
	NewResetPasswordRouter(cfg, db, userGroup)
	NewUploadAvatarRouter(cfg, db, objStore, userGroup)
	NewUpdateUserRouter(cfg, db, cache, userGroup)
	NewGetUserRouter(cfg, db, userGroup)
}

func NewGetUserRouter(cfg *config.RuntimeConfig, db *gorm.DB, group *gin.RouterGroup) {
	
}

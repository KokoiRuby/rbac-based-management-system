package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(cfg *config.RuntimeConfig, db *gorm.DB, cache *redis.Client, g *gin.Engine) {
	// Statics
	g.Static("/uploads", "./static/uploads")

	// Probes
	probeRouter := g.Group("")
	NewReadinessRouter(probeRouter)
	NewLivenessRouter(probeRouter)

	// Public APIs
	publicRouter := g.Group("")
	_ = publicRouter
	NewSignupRouter(cfg, db, publicRouter)
	NewSigninRouter(cfg, db, publicRouter)
	NewRefreshTokenRouter(cfg, db, publicRouter)
	//NewSignoutRouter(cfg, db, publicRouter)

	// Protected APIs
	protectedRouter := g.Group("v1")
	_ = protectedRouter
}

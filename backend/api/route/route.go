package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(cfg *config.RuntimeConfig, db *gorm.DB, g *gin.Engine) {
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
	//NewLoginRouter(env, timeout, db, publicRouter)
	//NewRefreshTokenRouter(env, timeout, db, publicRouter)

	// Protected APIs
	protectedRouter := g.Group("v1")
	_ = protectedRouter
}

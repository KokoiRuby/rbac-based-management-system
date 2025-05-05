package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/handler"
	"github.com/gin-gonic/gin"
)

func NewLivenessRouter(group *gin.RouterGroup) {
	h := handler.LivenessHandler{}
	group.GET("/liveness", h.Probe)
}

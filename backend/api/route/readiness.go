package route

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/handler"
	"github.com/gin-gonic/gin"
)

func NewReadinessRouter(group *gin.RouterGroup) {
	h := handler.ReadinessHandler{}
	group.GET("/readiness", h.Probe)
}

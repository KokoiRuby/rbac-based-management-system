package handler

import (
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type SignoutHandler struct {
	SignoutService service.SignoutService
	RuntimeConfig  *config.RuntimeConfig
}

func (handler *SignoutHandler) Signout(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		zap.S().Error("failed to get claims from context")
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signout.")
		return
	}

	email := claims.(*utils.CustomClaims).Email
	expireAt := claims.(*utils.CustomClaims).ExpiresAt
	key := fmt.Sprintf("signout_%s", email)

	flag, err := handler.SignoutService.IsSignedOut(c, key)
	if err != nil {
		zap.S().Errorf("failed to check if key exists: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signout.")
		return
	}
	if flag {
		utils.FailWithMsg(c, http.StatusConflict, "Already signed out.")
		return
	}

	ttl := expireAt.Sub(time.Now())
	_, err = handler.SignoutService.SetKeyWithTTLToCache(c, key, "", ttl)
	if err != nil {
		zap.S().Errorf("failed to set key to cache: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signout.")
		return
	}

	utils.OKWithMsg(c, http.StatusOK, "Signout successfully")
	return
}

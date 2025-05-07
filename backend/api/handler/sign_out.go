package handler

import (
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/api/middleware"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
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
	req := middleware.GetBind[model.SignoutRequest](c)

	fmt.Println(req.AccessToken)

	expireAt, err := handler.SignoutService.ExtractExpireAtFromToken(req.AccessToken, handler.RuntimeConfig.JWT)
	if err != nil {
		zap.S().Errorf("failed to extract expireAt from token: %v", err)
		utils.FailWithMsg(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	key := fmt.Sprintf("signout_%s", req.AccessToken)

	flag, err := handler.SignoutService.IsSignedOut(c, key)
	if err != nil {
		zap.S().Errorf("failed to check if key exists: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signout.")
		return
	}
	if flag {
		utils.FailWithMsg(c, http.StatusUnauthorized, "Already signed out.")
		return
	}

	ttl := expireAt.Sub(time.Now())
	_, err = handler.SignoutService.SetKeyWithTTLToCache(c, key, "", ttl)
	if err != nil {
		zap.S().Errorf("failed to set key to cache: %v", err)
		utils.FailWithMsg(c, http.StatusInternalServerError, "Failed to signout.")
		return
	}

	utils.FailWithMsg(c, http.StatusOK, "Signout successfully")
	return
}

package middleware

import (
	"context"
	"fmt"
	"github.com/KokoiRuby/rbac-based-management-system/backend/global"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func AuthNMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		utils.FailWithMsg(c, http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}

	tokenString := parts[1]
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		zap.S().Errorf("failed to parse token: %v", err)
		utils.FailWithMsg(c, http.StatusUnauthorized, "Invalid token")
		c.Abort()
		return
	}

	// Is logged out?
	key := fmt.Sprintf("signout_%s", claims.Email)
	_, err = global.Redis.Exists(context.Background(), key).Result()
	if err != nil {
		utils.OKWithMsg(c, http.StatusFound, "Already signed out.")
		c.Abort()
		return
	}

	// Propagate in context
	c.Set("claims", claims)
	c.Next()
}

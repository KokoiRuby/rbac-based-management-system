package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
)

type RefreshTokenService interface {
	ExtractIDFromToken(requestToken string, cfg runtime.JWT) (uint, error)
	GetUserByID(c context.Context, id uint) (model.User, error)
	CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error)
	CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error)
}

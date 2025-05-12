package service

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"time"
)

type forgotPasswordService struct {
	rdb            service.UserRDB
	cache          service.RedisCache
	contextTimeout time.Duration
}

func NewForgotPasswordService(rdb service.UserRDB, cache service.RedisCache, timeout time.Duration) service.ForgotPasswordService {
	return &forgotPasswordService{
		rdb:            rdb,
		cache:          cache,
		contextTimeout: timeout,
	}
}

func (s forgotPasswordService) GetUserByEmail(c *gin.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByEmail(ctx, email)
}

func (s forgotPasswordService) CreateConfirmToken(req *model.ForgotPasswordRequest, cfg runtime.JWT) (confirmToken string, err error) {
	return utils.CreateForgotPasswordConfirmToken(req, cfg)
}

func (s forgotPasswordService) SendConfirmEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error {
	return utils.SendEmail(msg, cfg)
}

func (s forgotPasswordService) Confirm(token string) (string, error) {
	return utils.ExtractEmailFromToken(token)
}

func (s forgotPasswordService) SetKeyWithTTLToCache(c *gin.Context, key string, value string, ttl time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.SetKeyWithTTL(ctx, key, value, ttl)
}

func (s forgotPasswordService) DelKeyFromCache(c *gin.Context, key string) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.DelKey(ctx, key)
}

func (s forgotPasswordService) IsKeyExist(c *gin.Context, key string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.IsKeyExist(ctx, key)
}

func (s forgotPasswordService) UpdateUser(c *gin.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.Update(ctx, user)
}

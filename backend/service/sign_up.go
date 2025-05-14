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

type signupService struct {
	rdb            service.UserRDB
	cache          service.RedisCache
	contextTimeout time.Duration
}

func NewSignupService(rdb service.UserRDB, cache service.RedisCache, timeout time.Duration) service.SignupService {
	return &signupService{
		rdb:            rdb,
		cache:          cache,
		contextTimeout: timeout,
	}
}

func (s signupService) GetUserByEmail(c *gin.Context, email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByEmail(ctx, email)
}

func (s signupService) CreateConfirmToken(req *model.SignupConfirmRequest, cfg runtime.JWT) (confirmToken string, err error) {
	return utils.CreateSignupConfirmToken(req, cfg)
}

func (s signupService) SendConfirmEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error {
	return utils.SendEmail(msg, cfg)
}

func (s signupService) Confirm(token string) (*model.SignupRequest, error) {
	return utils.ExtractSignupRequestFromToken(token)
}

func (s signupService) CreateAccessToken(user *model.User, cfg runtime.JWT) (accessToken string, err error) {
	return utils.CreateAccessToken(user, cfg)
}

func (s signupService) CreateUser(c *gin.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.Create(ctx, user)
}

func (s signupService) CreateRefreshToken(user *model.User, cfg runtime.JWT) (refreshToken string, err error) {
	return utils.CreateRefreshToken(user, cfg)
}

func (s signupService) SetKeyWithTTLToCache(c *gin.Context, key string, value string, ttl time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.SetKeyWithTTL(ctx, key, value, ttl)
}

func (s signupService) DelKeyFromCache(c *gin.Context, key string) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.DelKey(ctx, key)
}

func (s signupService) IsKeyExist(c *gin.Context, key string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.cache.IsKeyExist(ctx, key)
}

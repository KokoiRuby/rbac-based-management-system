package service

import (
	"context"
	"errors"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/model"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/KokoiRuby/rbac-based-management-system/backend/infra/repository/query"
	"github.com/KokoiRuby/rbac-based-management-system/backend/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type updateUserService struct {
	rdb            service.UserRDB
	cache          service.RedisCache
	contextTimeout time.Duration
}

func NewUpdateUserService(rdb service.UserRDB, cache service.RedisCache, contextTimeout time.Duration) service.UpdateUserService {
	return &updateUserService{
		rdb:            rdb,
		cache:          cache,
		contextTimeout: contextTimeout,
	}
}

func (s updateUserService) GetByID(c *gin.Context, id uint) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.GetByID(ctx, id)
}

func (s updateUserService) ValidateUserNameUniqueness(c *gin.Context, user *model.User, username string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()

	conds := []gen.Condition{
		query.User.Username.Eq(username),
		query.User.ID.Neq(user.ID),
	}

	_, found, err := s.rdb.GetByCond(ctx, conds...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return !found, nil
		}
		return found, err
	}

	return found, nil
}

func (s updateUserService) ValidateEmailUniqueness(c *gin.Context, user *model.User, email string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()

	conds := []gen.Condition{
		query.User.Email.Eq(email),
		query.User.ID.Neq(user.ID),
	}

	_, found, err := s.rdb.GetByCond(ctx, conds...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return !found, nil
		}
		return found, err
	}

	return found, nil
}

func (s updateUserService) UpdateUser(c *gin.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout*time.Second)
	defer cancel()
	return s.rdb.Update(ctx, user)
}

func (s updateUserService) CreateConfirmToken(req *model.UserUpdateConfirmRequest, cfg runtime.JWT) (confirmToken string, err error) {
	return utils.CreateUserUpdateConfirmToken(req, cfg)
}

func (s updateUserService) SendConfirmEmail(msg *gomail.Message, cfg runtime.SMTPConfig) error {
	return utils.SendEmail(msg, cfg)
}

func (s updateUserService) Confirm(token string) (*model.UserUpdateConfirmRequest, error) {
	return utils.ExtractUserUpdateRequestFromToken(token)
}

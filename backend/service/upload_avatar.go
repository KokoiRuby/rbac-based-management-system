package service

import (
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"time"
)

type uploadAvatarService struct {
	contextTimeout time.Duration
}

func NewUploadAvatarService(timeout time.Duration) service.UploadAvatarService {
	return &uploadAvatarService{
		contextTimeout: timeout,
	}
}

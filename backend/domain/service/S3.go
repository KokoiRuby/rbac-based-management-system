package service

import (
	"context"
	"mime/multipart"
)

type AvatarObjectStorage interface {
	Put(c context.Context, bucket string, key string, body multipart.File) error
}

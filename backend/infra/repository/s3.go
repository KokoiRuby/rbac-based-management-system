package repository

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/domain/service"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"mime/multipart"
)

type avatarS3 struct {
	userRDB
	client *s3.Client
}

func NewAvatarRepository(client *s3.Client) service.AvatarObjectStorage {
	return &avatarS3{
		client: client,
	}
}

func (a *avatarS3) Put(c context.Context, bucket string, key string, body multipart.File) error {
	_, err := a.client.PutObject(c, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   body,
	})
	if err != nil {
		return err
	}
	return nil
}

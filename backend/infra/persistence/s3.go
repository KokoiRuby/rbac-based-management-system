package persistence

import (
	"context"
	"github.com/KokoiRuby/rbac-based-management-system/backend/config/runtime"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.uber.org/zap"
)

func NewS3Client(ctx context.Context, cfg runtime.AWS) *s3.Client {
	// TODO: security practice required
	creds := credentials.NewStaticCredentialsProvider(cfg.S3.KeyID, cfg.S3.AccessKey, "")

	s3cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.S3.Region),
		config.WithCredentialsProvider(creds),
	)
	if err != nil {
		zap.S().Fatalf("unable to load S3 config: %v", err)
	}

	client := s3.NewFromConfig(s3cfg)

	return client
}

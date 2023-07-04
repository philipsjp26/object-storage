package strategies

import (
	"bytes"
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorageStrategy struct {
	config *MinioConfig
	client *minio.Client
}

func NewMinioStorageStrategy() (*MinioStorageStrategy, error) {
	config := GetMinioConfig()
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretKey, ""),
		Secure: config.useSSL,
	})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &MinioStorageStrategy{
		config: config,
		client: client,
	}, nil

}

func (s *MinioStorageStrategy) StoreObject(ctx context.Context, objectName string, file []byte) (string, error) {
	info, err := s.client.PutObject(ctx, s.config.BucketName, objectName, bytes.NewReader(file), int64(len(file)), minio.PutObjectOptions{})
	if err != nil {

		return "", err
	}
	return info.Bucket, nil
}

func (s *MinioStorageStrategy) GetObject(ctx context.Context, objectName string) (string, error) {
	return "", nil
}

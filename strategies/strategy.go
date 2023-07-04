package strategies

import "context"

type StorageStrategy interface {
	StoreObject(ctx context.Context, object string, file []byte) (string, error)
	GetObject(ctx context.Context, object string) (string, error)
}

type MinioConfig struct {
	Endpoint    string
	AccessKeyID string
	SecretKey   string
	BucketName  string
	useSSL      bool
}

func GetMinioConfig() *MinioConfig {
	return &MinioConfig{
		Endpoint:    "localhost:9000",
		AccessKeyID: "adminadmin99",
		SecretKey:   "InoueKNepjo9M6gBY73WjKIxlSmh1csEGOozeZns",
		BucketName:  "demo-development",
		useSSL:      false,
	}
}

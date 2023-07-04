package strategies

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type GCSStrategy struct {
	client *storage.Client
}

func NewGCSStorageStrategy(creds string) (*GCSStrategy, error) {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(creds))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return &GCSStrategy{
		client: client,
	}, nil
}

func (g *GCSStrategy) StoreObject(ctx context.Context, objectName string, file []byte) (string, error) {
	writer := g.client.Bucket("privy-acceleration").Object(objectName).NewWriter(ctx)
	defer writer.Close()

	attributes, err := g.client.Bucket("privy-acceleration").Object(objectName).Attrs(ctx)
	if err != nil {
		log.Fatalf("Error store object : %v", err.Error())
		return "", err
	}
	return attributes.MediaLink, nil
}

func (g *GCSStrategy) GetObject(ctx context.Context, objectName string) (string, error) {
	u := g.client.Bucket("privy-acceleration").Object(objectName)
	attributes, _ := u.Attrs(ctx)
	return attributes.MediaLink, nil
}

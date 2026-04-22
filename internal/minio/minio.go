package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIO struct {
	client *minio.Client
}

func NewMinIO() (*MinIO, error) {
	client, err := minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4(os.Getenv("MINIO_ROOT_USER"), os.Getenv("MINIO_ROOT_PASSWORD"), ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("cant create minIO client")
	}

	return &MinIO{client: client}, nil
}

func (m *MinIO) Put(ctx context.Context, bucket, name string, data []byte, contentType string) error {
	_, err := m.client.PutObject(
		ctx,
		bucket,
		name,
		bytes.NewReader(data),
		int64(len(data)),
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	return err
}

func (m *MinIO) Get(ctx context.Context, bucket, name string) ([]byte, string, error) {
	obj, err := m.client.GetObject(ctx, bucket, name, minio.GetObjectOptions{})
	if err != nil {
		return nil, "", err
	}
	defer obj.Close()

	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, "", err
	}

	stat, err := obj.Stat()
	if err != nil {
		return nil, "", err
	}

	return data, stat.ContentType, nil
}

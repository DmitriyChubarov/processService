package entity

import "context"

type MinIo interface {
	Put(ctx context.Context, bucket, name string, data []byte, contentType string) error
	Get(ctx context.Context, bucket, name string) ([]byte, string, error)
}

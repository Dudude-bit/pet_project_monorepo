package queue

import "context"

type UserStorageInterface interface {
}

type StorageInterface interface {
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
}

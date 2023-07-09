package queue

import "context"

type UserStorageInterface interface {
}

type StorageInterface interface {
	Pint(ctx context.Context) error
	Close(ctx context.Context) error
}

package database

import "context"

type UserStorageInterface interface {
	CreateUser(ctx context.Context, dto *User) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
}

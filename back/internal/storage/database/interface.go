package database

import "context"

type UserStorageInterface interface {
	CreateUser(ctx context.Context, dto *CreateUserDTO) (*CreateUserReturn, error)
	GetUser(ctx context.Context, dto *GetUserDTO) (*GetUserReturn, error)
	GetUserByUsername(ctx context.Context, dto *GetUserByUsernameDTO) (*GetUserByUsernameReturn, error)
}

type VideoStorageInterface interface {
	CreateVideo(ctx context.Context, dto *CreateVideoDTO) (*CreateVideoReturn, error)
	GetVideo(ctx context.Context, dto *GetVideoDTO) (*GetVideoReturn, error)
}

type DonationStorageInterface interface {
	CreateDonation(ctx context.Context, dto *CreateDonationDTO) (*CreateDonationReturn, error)
	UpdateDonationStatus(ctx context.Context, dto *UpdateDonationStatusDTO) (*UpdateDonationStatusReturn, error)
}

type StorageInterface interface {
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
}

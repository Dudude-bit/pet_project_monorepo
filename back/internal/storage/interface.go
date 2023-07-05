package storage

import "context"

type UserStorageInterface interface {
	CreateUser(ctx context.Context, dto *CreateUserDTO) (*CreateUserReturn, error)
	GetUser(ctx context.Context, dto *GetUserDTO) (*GetUserReturn, error)
}

type VideoStorage interface {
	CreateVideo(ctx context.Context, dto *CreateVideoDTO) (*CreateVideoReturn, error)
	GetVideo(ctx context.Context, dto *GetVideoDTO) (*GetVideoReturn, error)
}

type DonationStorage interface {
	CreateDonation(ctx context.Context, dto *CreateDonationDTO) (*CreateDonationReturn, error)
	UpdateDonationStatus(ctx context.Context, dto *UpdateDonationStatusDTO) (*UpdateDonationStatusReturn, error)
}

type DatabaseInterface interface {
	UserStorageInterface
	VideoStorage
	DonationStorage
}

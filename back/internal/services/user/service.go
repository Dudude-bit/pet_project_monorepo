package user

import (
	"context"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database"
)

type Service struct {
	storage database.UserStorageInterface
}

func NewUserService(storage database.UserStorageInterface) *Service {
	return &Service{storage: storage}
}

func (s *Service) RegisterUser(ctx context.Context, dto *RegisterUserDTO) (*RegisterUserReturn, error) {
	return nil, nil
}

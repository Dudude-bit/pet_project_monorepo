package user

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database"
	"github.com/Dudude-bit/pet_project_monorepo/back/internal/utils"
)

type Service struct {
	storage      database.UserStorageInterface
	JWTSecretKey string
}

func NewService(storageInstance database.UserStorageInterface, JWTSecretKey string) *Service {
	return &Service{storage: storageInstance, JWTSecretKey: JWTSecretKey}
}

func (s *Service) RegisterUser(ctx context.Context, dto *RegisterUserDTO) (*RegisterUserReturn, error) {
	hashedPassword, hashPasswordErr := utils.HashPassword(dto.Password)
	if hashPasswordErr != nil {
		return nil, hashPasswordErr
	}

	user, createUserErr := s.storage.CreateUser(ctx, &database.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: hashedPassword,
	})
	if createUserErr != nil {
		return nil, createUserErr
	}

	return &RegisterUserReturn{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *Service) LoginUser(ctx context.Context, dto *LoginUserDTO) (*LoginUserReturn, error) {
	user, getUserByUsernameErr := s.storage.GetUserByUsername(ctx, dto.Username)

	if getUserByUsernameErr != nil {
		return nil, getUserByUsernameErr
	}

	if !utils.CheckPasswordHash(dto.Password, user.Password) {
		return nil, fmt.Errorf("user not found: %s", dto.Username)
	}

	payload := jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, signErr := token.SignedString([]byte(s.JWTSecretKey))
	if signErr != nil {
		return nil, signErr
	}

	return &LoginUserReturn{AccessToken: signedToken}, nil
}

func (s *Service) GetUser(ctx context.Context, dto *GetUserDTO) (*GetUserReturn, error) {
	user, getUserErr := s.storage.GetUser(ctx, dto.Id)
	if getUserErr != nil {
		return nil, getUserErr
	}

	return &GetUserReturn{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

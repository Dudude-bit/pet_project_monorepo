package user

import "context"

type ServiceInterface interface {
	RegisterUser(ctx context.Context, dto *RegisterUserDTO) (*RegisterUserReturn, error)
	LoginUser(ctx context.Context, dto *LoginUserDTO) (*LoginUserReturn, error)
	GetUser(ctx context.Context, dto *GetUserDTO) (*GetUserReturn, error)
}

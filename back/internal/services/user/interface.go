package user

import "context"

type ServiceInterface interface {
	RegisterUser(ctx context.Context, dto *RegisterUserDTO) (*RegisterUserReturn, error)
}

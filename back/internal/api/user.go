package api

import (
	"context"
	"github.com/Dudude-bit/pet_project_monorepo/back/internal/services/user"
)

func (ae *Server) RegisterUser(ctx context.Context, request RegisterUserRequestObject) (RegisterUserResponseObject, error) {

	// TODO add validation

	registeredUser, registerUserErr := ae.UserService.RegisterUser(ctx, &user.RegisterUserDTO{
		Username: request.Body.Username,
		Email:    request.Body.Email,
		Password: request.Body.Password,
	})

	if registerUserErr != nil {
		return nil, registerUserErr
	}

	return RegisterUser200JSONResponse{
		Data: &User{
			Email:    registeredUser.Email,
			Username: registeredUser.Username,
		},
	}, registerUserErr
}

func (ae *Server) LoginUser(ctx context.Context, request LoginUserRequestObject) (LoginUserResponseObject, error) {

	// TODO add validation

	loginData, loginUserErr := ae.UserService.LoginUser(ctx, &user.LoginUserDTO{
		Username: request.Body.Username,
		Password: request.Body.Password,
	})

	if loginUserErr != nil {
		return nil, loginUserErr
	}

	return LoginUser200JSONResponse{
		Data: &Authorization{
			AccessToken: loginData.AccessToken,
		},
	}, nil
}

func (ae *Server) UserMe(ctx context.Context, _ UserMeRequestObject) (UserMeResponseObject, error) {

	userId := ctx.Value(JWTUserContextKey).(string)

	userResult, getUserErr := ae.UserService.GetUser(ctx, &user.GetUserDTO{Id: userId})
	if getUserErr != nil {
		return nil, getUserErr
	}

	return UserMe200JSONResponse{
		Data: &User{
			Email:    userResult.Email,
			Username: userResult.Username,
		},
	}, nil
}

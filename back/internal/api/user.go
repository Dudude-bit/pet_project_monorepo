package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/services/user"
)

func (ae *Server) RegisterUser(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	b, readErr := io.ReadAll(request.Body)
	defer request.Body.Close()

	if readErr != nil {
		sendErrorResponse(response, http.StatusBadRequest, readErr)
		return
	}

	var data RegisterUserJSONRequestBody
	unmarshalErr := json.Unmarshal(b, &data)
	if unmarshalErr != nil {
		sendErrorResponse(response, http.StatusInternalServerError, readErr)
		return
	}

	if validationErr := data.Validate(ctx); validationErr != nil {
		sendErrorResponse(response, http.StatusBadRequest, validationErr)
		return
	}

	registeredUser, registerUserErr := ae.UserService.RegisterUser(ctx, &user.RegisterUserDTO{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	})

	if registerUserErr != nil {
		sendErrorResponse(response, http.StatusInternalServerError, registerUserErr)
		return
	}

	if sendErr := sendSuccessResponse(response, http.StatusCreated, User{
		Email:    registeredUser.Email,
		Username: registeredUser.Username,
	}); sendErr != nil {
		sendErrorResponse(response, http.StatusInternalServerError, sendErr)
		return
	}
}

func (ae *Server) LoginUser(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	b, readErr := io.ReadAll(request.Body)
	defer request.Body.Close()

	if readErr != nil {
		sendErrorResponse(response, http.StatusBadRequest, readErr)
		return
	}

	var data LoginUserJSONRequestBody
	unmarshalErr := json.Unmarshal(b, &data)
	if unmarshalErr != nil {
		sendErrorResponse(response, http.StatusInternalServerError, unmarshalErr)
		return
	}

	if validationErr := data.Validate(ctx); validationErr != nil {
		sendErrorResponse(response, http.StatusBadRequest, validationErr)
		return
	}

	loginData, loginUserErr := ae.UserService.LoginUser(ctx, &user.LoginUserDTO{
		Username: data.Username,
		Password: data.Password,
	})

	if loginUserErr != nil {
		sendErrorResponse(response, http.StatusInternalServerError, loginUserErr)
		return
	}

	if sendErr := sendSuccessResponse(response, http.StatusCreated, Authorization{
		AccessToken: loginData.AccessToken,
	}); sendErr != nil {
		sendErrorResponse(response, http.StatusInternalServerError, sendErr)
		return
	}
}

func (ae *Server) UserMe(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	userId, ok := request.Context().Value(userIdKey).(string)
	if !ok {
		sendErrorResponse(response, http.StatusInternalServerError, fmt.Errorf("something went wrong"))
		return
	}

	user, getUserErr := ae.UserService.GetUser(ctx, &user.GetUserDTO{Id: userId})
	if getUserErr != nil {
		sendErrorResponse(response, http.StatusInternalServerError, getUserErr)
		return
	}

	if sendErr := sendSuccessResponse(response, http.StatusCreated, User{
		Username: user.Username,
		Email:    user.Email,
	}); sendErr != nil {
		sendErrorResponse(response, http.StatusInternalServerError, sendErr)
		return
	}
}

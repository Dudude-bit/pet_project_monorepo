package api

import (
	"context"

	"github.com/go-playground/validator/v10"
)

func (params *RegisterUserJSONRequestBody) Validate(ctx context.Context) error {
	validatorInstance := validator.New()

	return validatorInstance.StructCtx(ctx, params)
}

func (params *LoginUserJSONRequestBody) Validate(ctx context.Context) error {
	validatorInstance := validator.New()

	return validatorInstance.StructCtx(ctx, params)
}

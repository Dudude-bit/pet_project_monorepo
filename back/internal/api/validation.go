package api

import (
	"context"

	"github.com/go-playground/validator/v10"
)

func (params *RegisterUser) Validate(ctx context.Context) error {
	validatorInstance := validator.New()

	return validatorInstance.StructCtx(ctx, params)
}

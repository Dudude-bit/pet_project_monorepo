package api

import (
	"time"

	userService "github.com/Dudude-bit/pet_project_monorepo/back/internal/services/user"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	UserService *userService.Service `json:"storage"`
}

type ServerParams struct {
	BaseURL      string               `json:"base_url"`
	ReadTimeout  time.Duration        `json:"read_timeout"`
	WriteTimeout time.Duration        `json:"write_timeout"`
	JWTSecretKey string               `json:"jwt_secret_key"`
	UserService  *userService.Service `json:"user_service"`
}

func NewServer(params *ServerParams) *fiber.App {
	server := Server{
		UserService: params.UserService,
	}

	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: false,
		ReadTimeout:   params.ReadTimeout,
		WriteTimeout:  params.WriteTimeout,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// todo add custom response code

			return ctx.JSON(struct {
				Error string `json:"error"`
			}{
				Error: err.Error(),
			})
		},
		DisableStartupMessage: true,
	})
	// справа налево
	ssi := NewStrictHandler(&server, []StrictMiddlewareFunc{WithJWTMiddleware([]byte(params.JWTSecretKey)), RequestIDMiddleware})
	RegisterHandlersWithOptions(app, ssi, FiberServerOptions{
		BaseURL: params.BaseURL,
	})
	return app
}

package api

import (
	"time"

	userService "github.com/Dudude-bit/pet_project_monorepo/back/internal/services/user"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	UserService userService.ServiceInterface `json:"storage"`
}

type ServerParams struct {
	BaseURL      string                       `json:"base_url"`
	ReadTimeout  time.Duration                `json:"read_timeout"`
	WriteTimeout time.Duration                `json:"write_timeout"`
	JWTSecretKey string                       `json:"jwt_secret_key"`
	UserService  userService.ServiceInterface `json:"user_service"`
}

func NewServer(params *ServerParams) *fiber.App {
	server := Server{
		UserService: params.UserService,
	}

	jwtMiddleware := jwtware.New(jwtware.Config{
		Filter: func(ctx *fiber.Ctx) bool {
			return ctx.Context().UserValue(JWTUserContextKey) == nil
		},
		SuccessHandler: nil,
		ErrorHandler:   nil,
		SigningKey:     jwtware.SigningKey{Key: params.JWTSecretKey},
		ContextKey:     JWTUserContextKey,
		Claims:         nil,
		TokenLookup:    "header:Authorization",
		AuthScheme:     "Bearer",
	})

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

	ssi := NewStrictHandler(&server, nil)
	RegisterHandlersWithOptions(app, ssi, FiberServerOptions{
		BaseURL:     params.BaseURL,
		Middlewares: []MiddlewareFunc{jwtMiddleware},
	})
	return app
}

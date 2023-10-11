package api

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func RequestIDMiddleware(f StrictHandlerFunc, _ string) StrictHandlerFunc {
	return func(ctx *fiber.Ctx, args interface{}) (interface{}, error) {
		ctx.Context().SetUserValue(RequestIdContextKey, uuid.New().String())
		return f(ctx, args)
	}
}

func WithJWTMiddleware(secret []byte) func(f StrictHandlerFunc, operationID string) StrictHandlerFunc {
	return func(f StrictHandlerFunc, operationID string) StrictHandlerFunc {
		return func(ctx *fiber.Ctx, args interface{}) (interface{}, error) {
			if ctx.Context().UserValue(JWTAuthScopes) == nil {
				return f(ctx, args)
			}
			token, err := jwt.Parse(ctx.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("unknown signing method: %s", token.Method.Alg())
				}
				return secret, nil
			})
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, errors.New("invalid token")
			}

			userCtx := context.WithValue(ctx.UserContext(), JWTUserContextKey, claims["sub"])
			ctx.SetUserContext(userCtx)
			return f(ctx, args)
		}
	}
}

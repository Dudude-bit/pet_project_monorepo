package api

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func WithUserMiddleware(JWTSecretKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Context().Value(JWTAuthScopes) == nil {
				next.ServeHTTP(w, r)
				return
			}
			tokenBearerSlice, ok := r.Header["Authorization"]
			if !ok {
				sendErrorResponse(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
				return
			}
			if len(tokenBearerSlice) != 1 {
				sendErrorResponse(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
				return
			}
			tokenBearer := tokenBearerSlice[0]
			tokenJWT := strings.Replace(tokenBearer, "Bearer ", "", -1)

			claims := &jwt.MapClaims{}
			tkn, parseErr := jwt.ParseWithClaims(tokenJWT, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(JWTSecretKey), nil
			})
			if parseErr != nil {
				sendErrorResponse(w, http.StatusUnauthorized, parseErr)
				return
			}

			sub, getSubErr := tkn.Claims.GetSubject()
			if getSubErr != nil {
				sendErrorResponse(w, http.StatusUnauthorized, getSubErr)
				return
			}
			ctx := context.WithValue(r.Context(), userIdKey, sub)
			newRequest := r.WithContext(ctx)
			next.ServeHTTP(w, newRequest)
		})
	}
}

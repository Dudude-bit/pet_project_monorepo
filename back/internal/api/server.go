package api

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	searchService "github.com/Dudude-bit/pet_project_monorepo/back/internal/services/search"
	userService "github.com/Dudude-bit/pet_project_monorepo/back/internal/services/user"
)

type Server struct {
	UserService userService.ServiceInterface   `json:"storage"`
	Search      searchService.ServiceInterface `json:"search"`
}

type ServerParams struct {
	BaseURL           string                         `json:"base_url"`
	ServerAddress     string                         `json:"server_address"`
	ReadHeaderTimeout time.Duration                  `json:"read_header_timeout"`
	ReadTimeout       time.Duration                  `json:"read_timeout"`
	UserService       userService.ServiceInterface   `json:"user_service"`
	SearchService     searchService.ServiceInterface `json:"search_service"`
}

func NewServer(params *ServerParams) (*http.Server, error) {
	mux := chi.NewRouter()

	serverEnv := &Server{
		UserService: params.UserService,
		Search:      params.SearchService,
	}
	mux.With(middleware.SetHeader("Content-Type", "text/json")).
		Route(params.BaseURL, func(r chi.Router) {
			HandlerWithOptions(serverEnv, ChiServerOptions{
				BaseURL:    params.BaseURL,
				BaseRouter: r,
			})
		})

	return &http.Server{
		Addr:              params.ServerAddress,
		Handler:           mux,
		ReadHeaderTimeout: params.ReadHeaderTimeout,
		ReadTimeout:       params.ReadTimeout,
	}, nil
}

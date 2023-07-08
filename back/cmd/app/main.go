package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/api"
	services "github.com/Dudude-bit/pet_project_monorepo/back/internal/services/user"
	storage "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database/edgedb"
	utils "github.com/Dudude-bit/pet_project_monorepo/back/internal/utils"
)

func main() {
	ctx := context.Background()

	cfg, loadCfgErr := utils.LoadConfig()
	if loadCfgErr != nil {
		logrus.WithError(loadCfgErr).Fatal("cant load cfg")
	}
	utils.ConfigureLogger()

	storageInstance, newStorageErr := storage.NewStorage(ctx, cfg.Storage)
	if newStorageErr != nil {
		logrus.WithError(newStorageErr).Fatal("cant create storage")
	}

	userService := services.NewUserService(storageInstance)

	server, newServerErr := api.NewServer(&api.ServerParams{
		BaseURL:           cfg.BaseURL,
		ServerAddress:     cfg.ServerAddress,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		UserService:       userService,
		SearchService:     nil,
	})
	if newServerErr != nil {
		logrus.WithError(newServerErr).Fatal("cant create server")
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.WithError(err).Fatal("error on listen")
		return
	}
}

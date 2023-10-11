package main

import (
	"context"
	"flag"
	"github.com/sirupsen/logrus"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/api"
	userService "github.com/Dudude-bit/pet_project_monorepo/back/internal/services/user"
	edgedbStorage "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database/surrealdb"
	"github.com/Dudude-bit/pet_project_monorepo/back/internal/utils"
)

func main() {
	_ = context.Background()

	configPath := flag.String("c", "./config-docker-compose.yaml", "Path to config")
	flag.Parse()

	cfg, loadCfgErr := utils.LoadConfig(*configPath)
	if loadCfgErr != nil {
		logrus.WithError(loadCfgErr).Fatal("cant load cfg")
	}
	utils.ConfigureLogger()

	storageInstance, newStorageErr := edgedbStorage.NewStorage(cfg.Storage)
	if newStorageErr != nil {
		logrus.WithError(newStorageErr).Fatal("cant create storage")
	}

	userServiceInstance := userService.NewService(storageInstance, cfg.JWTSecretKey)

	server := api.NewServer(&api.ServerParams{
		BaseURL:      cfg.BaseURL,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
		JWTSecretKey: cfg.JWTSecretKey,
		UserService:  userServiceInstance,
	})

	logrus.Info("starting server")

	if err := server.Listen(cfg.ServerAddress); err != nil {
		logrus.WithError(err).Fatal("error on listen")
		return
	}
}

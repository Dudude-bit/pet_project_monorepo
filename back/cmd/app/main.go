package main

import (
	"context"
	"flag"

	"github.com/sirupsen/logrus"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/api"
	searchService "github.com/Dudude-bit/pet_project_monorepo/back/internal/services/search"
	userService "github.com/Dudude-bit/pet_project_monorepo/back/internal/services/user"
	edgedbStorage "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database/edgedb"
	searchStorage "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/search/meilisearch"
	"github.com/Dudude-bit/pet_project_monorepo/back/internal/utils"
)

func main() {
	ctx := context.Background()

	configPath := flag.String("c", "./config.yaml", "Path to config")
	flag.Parse()

	cfg, loadCfgErr := utils.LoadConfig(*configPath)
	if loadCfgErr != nil {
		logrus.WithError(loadCfgErr).Fatal("cant load cfg")
	}
	utils.ConfigureLogger()

	storageInstance, newStorageErr := edgedbStorage.NewStorage(ctx, cfg.Storage)
	if newStorageErr != nil {
		logrus.WithError(newStorageErr).Fatal("cant create storage")
	}

	//queueInstance, newQueueErr := rabbitmq.NewStorage(ctx, cfg.QueueStorage)
	//if newQueueErr != nil {
	//	logrus.WithError(newQueueErr).Fatal("cant create queue")
	//}

	searchInstance, newSearchInstanceErr := searchStorage.NewStorage(ctx, cfg.SearchStorage)
	if newSearchInstanceErr != nil {
		logrus.WithError(newSearchInstanceErr).Fatal("cant create search instance")
	}

	userServiceInstance := userService.NewService(storageInstance, cfg.JWTSecretKey, nil)
	searchServiceInstance := searchService.NewService(searchInstance)

	server, newServerErr := api.NewServer(&api.ServerParams{
		BaseURL:           cfg.BaseURL,
		ServerAddress:     cfg.ServerAddress,
		JWTSecretKey:      cfg.JWTSecretKey,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		UserService:       userServiceInstance,
		SearchService:     searchServiceInstance,
	})
	if newServerErr != nil {
		logrus.WithError(newServerErr).Fatal("cant create server")
	}

	logrus.Info("starting server")
	if err := server.ListenAndServe(); err != nil {
		logrus.WithError(err).Fatal("error on listen")
		return
	}
}

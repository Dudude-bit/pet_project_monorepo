package meilisearch

import (
	"context"

	meilisearchLib "github.com/meilisearch/meilisearch-go"
)

type Config struct {
	DSN       string `mapstructure:"SEARCH_DSN"`
	MasterKey string `mapstructure:"SEARCH_MASTER_KEY"`
	Timeout   int    `mapstructure:"SEARCH_TIMEOUT"`
}

type Storage struct {
	conn *meilisearchLib.Client
}

func NewStorage(ctx context.Context, params *Config) (*Storage, error) {
	conn := meilisearchLib.NewClient(
		meilisearchLib.ClientConfig{
			Host:    params.DSN,
			APIKey:  params.MasterKey,
			Timeout: 0,
		},
	)

	return &Storage{conn: conn}, nil
}

func (s *Storage) Close(ctx context.Context) error {
	return nil
}

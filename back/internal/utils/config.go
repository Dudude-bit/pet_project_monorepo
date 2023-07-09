package utils

import (
	"time"

	"github.com/spf13/viper"

	storage "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database/edgedb"
	search "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/search/meilisearch"
)

type Config struct {
	BaseURL           string        `mapstructure:"BASE_URL"`
	ServerAddress     string        `mapstructure:"SERVER_ADDRESS"`
	ReadHeaderTimeout time.Duration `mapstructure:"READ_HEADER_TIMEOUT"`
	ReadTimeout       time.Duration `mapstructure:"READ_TIMEOUT"`
	Storage           *storage.Config
	SearchStorage     *search.Config
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigType("env")
	viper.AutomaticEnv()

	unmarshalErr := viper.Unmarshal(cfg)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return cfg, nil
}

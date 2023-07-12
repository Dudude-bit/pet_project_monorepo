package utils

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"

	storage "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database/edgedb"
	queue "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/queue/rabbitmq"
	search "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/search/meilisearch"
)

type Config struct {
	BaseURL           string          `yaml:"base_url"`
	ServerAddress     string          `yaml:"server_address"`
	ReadHeaderTimeout time.Duration   `yaml:"read_header_timeout"`
	ReadTimeout       time.Duration   `yaml:"read_timeout"`
	JWTSecretKey      string          `yaml:"jwt_secret_key"`
	Storage           *storage.Config `yaml:"db"`
	QueueStorage      *queue.Config   `yaml:"queue"`
	SearchStorage     *search.Config  `yaml:"search"`
} // TODO move this structure

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

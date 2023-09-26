package utils

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"

	storage "github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database/edgedb"
)

type Config struct {
	BaseURL       string          `yaml:"base_url"`
	ServerAddress string          `yaml:"server_address"`
	ReadTimeout   time.Duration   `yaml:"read_timeout"`
	WriteTimeout  time.Duration   `yaml:"write_timeout"`
	JWTSecretKey  string          `yaml:"jwt_secret_key"`
	Storage       *storage.Config `yaml:"db"`
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

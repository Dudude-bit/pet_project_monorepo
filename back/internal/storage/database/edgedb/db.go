package edgedb

import (
	"context"
	"time"

	edgedbLib "github.com/edgedb/edgedb-go"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database"
)

type Config struct {
	Host              string        `mapstructure:"DB_HOST"`
	Port              int           `mapstructure:"DB_PORT"`
	Database          string        `mapstructure:"DB_NAME"`
	User              string        `mapstructure:"DB_USER"`
	Password          string        `mapstructure:"DB_PASSWORD"`
	ConnectionTimeout time.Duration `mapstructure:"DB_CONNECTION_TIMEOUT"`
}

type Storage struct {
	conn *edgedbLib.Client
}

func NewStorage(ctx context.Context, params *Config) (*Storage, error) {
	conn, createClientErr := edgedbLib.CreateClient(ctx, edgedbLib.Options{
		Host:           params.Host,
		Port:           params.Port,
		Database:       params.Database,
		User:           params.User,
		Password:       edgedbLib.NewOptionalStr(params.Password),
		ConnectTimeout: params.ConnectionTimeout,
	})
	if createClientErr != nil {
		return nil, createClientErr
	}

	return &Storage{conn: conn}, nil
}

func (s *Storage) CreateUser(ctx context.Context, dto *database.CreateUserDTO) (*database.CreateUserReturn, error) {
	return nil, nil
}

func (s *Storage) GetUser(ctx context.Context, dto *database.GetUserDTO) (*database.GetUserReturn, error) {
	//TODO implement me
	panic("implement me")
}

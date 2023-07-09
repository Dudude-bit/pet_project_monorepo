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

func (s *Storage) Close(ctx context.Context) error {
	return s.conn.Close()
}

func (s *Storage) CreateUser(ctx context.Context, dto *database.CreateUserDTO) (*database.CreateUserReturn, error) {
	res, insertErr := userCreate(ctx, s.conn, dto.Username, dto.Password, dto.Email)
	if insertErr != nil {
		return nil, insertErr
	}

	return &database.CreateUserReturn{
		Username: res.username,
		Email:    res.email,
	}, nil
}

func (s *Storage) GetUser(ctx context.Context, dto *database.GetUserDTO) (*database.GetUserReturn, error) {

	edgedbUUID := edgedbLib.UUID{}
	unmarshalErr := edgedbUUID.UnmarshalText([]byte(dto.Id))
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	res, selectErr := userGet(ctx, s.conn, edgedbUUID)
	if selectErr != nil {
		return nil, selectErr
	}

	return &database.GetUserReturn{
		Username: res.username,
		Email:    res.email,
	}, nil
}

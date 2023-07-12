package edgedb

import (
	"context"
	"time"

	edgedbLib "github.com/edgedb/edgedb-go"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database"
)

type TLSConfig struct {
	CAFile       string                    `yaml:"ca_file"`
	SecurityMode edgedbLib.TLSSecurityMode `yaml:"security_mode"`
}

type Config struct {
	Host              string        `yaml:"host"`
	Port              int           `yaml:"port"`
	Database          string        `yaml:"name"`
	User              string        `yaml:"user"`
	Password          string        `yaml:"password"`
	TLSOptions        TLSConfig     `yaml:"tls_options"`
	ConnectionTimeout time.Duration `yaml:"connection_timeout"`
}

type Storage struct {
	conn *edgedbLib.Client
}

func NewStorage(ctx context.Context, params *Config) (*Storage, error) {
	conn, createClientErr := edgedbLib.CreateClient(ctx, edgedbLib.Options{
		Host:     params.Host,
		Port:     params.Port,
		Database: params.Database,
		User:     params.User,
		Password: edgedbLib.NewOptionalStr(params.Password),
		TLSOptions: edgedbLib.TLSOptions{
			CAFile:       params.TLSOptions.CAFile,
			SecurityMode: params.TLSOptions.SecurityMode,
		},
		ConnectTimeout: params.ConnectionTimeout,
	})
	if createClientErr != nil {
		return nil, createClientErr
	}

	storage := &Storage{conn: conn}

	if pingErr := storage.Ping(ctx); pingErr != nil {
		return nil, pingErr
	}

	return storage, nil
}

func (s *Storage) Ping(ctx context.Context) error {
	pingQuery := "select 'ping';"

	return s.conn.QuerySingle(ctx, pingQuery, &edgedbLib.OptionalStr{})
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

func (s *Storage) GetUserByUsername(ctx context.Context, dto *database.GetUserByUsernameDTO) (*database.GetUserByUsernameReturn, error) {
	res, selectErr := userGetByUsername(ctx, s.conn, dto.Username)
	if selectErr != nil {
		return nil, selectErr
	}

	return &database.GetUserByUsernameReturn{
		Id:       res.id.String(),
		Username: res.username,
		Email:    res.email,
		Password: res.password,
	}, nil
}

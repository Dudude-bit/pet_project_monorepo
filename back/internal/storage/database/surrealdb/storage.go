package surrealdb

import (
	"context"
	"github.com/google/uuid"
	"time"

	"github.com/surrealdb/surrealdb.go"

	"github.com/Dudude-bit/pet_project_monorepo/back/internal/storage/database"
)

const (
	thing = "user"
)

type Config struct {
	URL               string        `json:"url" yaml:"url"`
	Username          string        `json:"username" yaml:"username"`
	Password          string        `json:"password" yaml:"password"`
	Namespace         string        `json:"namespace" yaml:"namespace"`
	Database          string        `json:"database" yaml:"database"`
	ConnectionTimeout time.Duration `json:"connection_timeout" yaml:"connection_timeout"`
}

type Storage struct {
	conn *surrealdb.DB
}

func NewStorage(params *Config) (*Storage, error) {
	conn, newErr := surrealdb.New(params.URL, surrealdb.WithTimeout(params.ConnectionTimeout))
	if newErr != nil {
		return nil, newErr
	}
	if _, signingErr := conn.Signin(map[string]string{
		"user": params.Username,
		"pass": params.Password,
	}); signingErr != nil {
		return nil, signingErr
	}
	if _, useErr := conn.Use(params.Namespace, params.Database); useErr != nil {
		return nil, useErr
	}

	storage := &Storage{conn: conn}

	return storage, nil
}

func (s *Storage) Close() {
	s.conn.Close()
}

func (s *Storage) CreateUser(_ context.Context, dto *database.User) (*database.User, error) {
	users, createErr := surrealdb.SmartUnmarshal[[]*database.User](s.conn.Create(thing, &database.User{
		Id:       uuid.New().String(),
		Username: dto.Username,
		Email:    dto.Email,
		Password: dto.Password,
	}))
	if createErr != nil {
		return nil, createErr
	}
	return users[0], nil
}

func (s *Storage) GetUser(_ context.Context, id string) (*database.User, error) {
	query := `
	SELECT username, mail, password from $thing WHERE id = $id
	`
	users, queryErr := surrealdb.SmartUnmarshal[[]*database.User](s.conn.Query(query, map[string]string{
		"thing": thing,
		"id":    id,
	}))
	if queryErr != nil {
		return nil, queryErr
	}
	return users[0], nil
}

func (s *Storage) GetUserByUsername(_ context.Context, username string) (*database.User, error) {
	query := `
	SELECT username, mail, password from $thing
	`
	users, queryErr := surrealdb.SmartUnmarshal[[]*database.User](s.conn.Query(query, map[string]string{
		"thing":    thing,
		"username": username,
	}))
	if queryErr != nil {
		return nil, queryErr
	}
	if len(users) != 1 {
		return nil, database.ErrNoSuchUser
	}
	return users[0], nil
}

package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	DSN string `mapstructure:"QUEUE_DSN"`
}

type Storage struct {
	conn *amqp.Connection
}

func NewStorage(ctx context.Context, params *Config) (*Storage, error) {
	conn, dialErr := amqp.Dial(params.DSN)
	if dialErr != nil {
		return nil, dialErr
	}
	return &Storage{conn: conn}, nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.conn.Close()
}

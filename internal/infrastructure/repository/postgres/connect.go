package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/ryoeuyo/test_go/internal/config"
)

func Connect(ctx context.Context, cfg config.Database) (*pgx.Conn, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		cfg.Host, cfg.Port, cfg.Name, cfg.User, cfg.Password,
	)
	connectConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	conn, err := pgx.ConnectConfig(ctx, connectConfig)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}

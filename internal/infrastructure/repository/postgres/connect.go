package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/ryoeuyo/test_go/internal/config"
)

func MustConnect(ctx context.Context, cfg config.Database) *pgx.Conn {
	connString := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s",
		cfg.Host, cfg.Port, cfg.Name, cfg.User, cfg.Password,
	)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		panic("failed connect to database")
	}

	if err := conn.Ping(ctx); err != nil {
		panic("failed ping to database")
	}

	return conn
}

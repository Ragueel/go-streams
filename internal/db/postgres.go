package db

import (
	"context"
	"github.com/jackc/pgx"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
)

func CreatePostgresConnection(ctx context.Context, logger *zap.Logger, lc fx.Lifecycle) *pgx.Conn {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	if err != nil {
		panic(err)
	}

	lc.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		return conn.Close()
	}})

	return conn
}

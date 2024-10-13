package postgres

import (
	"context"
	"fmt"
	"github.com/Tinddd28/TestTask/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	songTable   = "songs"
	limitSong   = 10
	versesTable = "verses"
)

func NewPostgres(cfg config.Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}
	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}

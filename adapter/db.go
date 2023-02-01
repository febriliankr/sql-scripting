package adapter

import (
	"fmt"
	"time"

	"github.com/febriliankr/go-sql-scripting/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.Config) (*sqlx.DB, error) {
	pg, err := sqlx.Connect(cfg.DatabaseDriver, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	pg.SetConnMaxLifetime(time.Minute * 3)
	pg.SetMaxOpenConns(100)
	pg.SetMaxIdleConns(10)

	return pg, nil
}

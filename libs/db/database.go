package db

import (
	"fmt"
	"log"

	"github.com/ecomz/backend/libs/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQLServerConnectionManager struct {
	db *sqlx.DB
}

func NewConnectionManager(cfg config.DatabaseConfig) (*SQLServerConnectionManager, error) {
	db, err := sqlx.Connect("postgres", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.MaxConnLifeTime)
	db.SetConnMaxIdleTime(cfg.MaxConnIdleTime)

	log.Println("Connected to database")

	return &SQLServerConnectionManager{db: db}, nil
}

func (c *SQLServerConnectionManager) Close() error {
	log.Println("closing sql server connection")
	return c.db.Close()
}

func (m *SQLServerConnectionManager) GetDB() *sqlx.DB {
	return m.db
}

package database

import (
	"fmt"
	"log"

	"market_service/internal/infrastructure/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewDB(cfg config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return db, nil
}

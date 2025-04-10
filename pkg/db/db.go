package db

import (
	"fmt"
	"log"

	"github.com/Egorpalan/api-pvz/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg config.DBConfig) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}

	log.Println("Connected to Postgres")
	return db
}

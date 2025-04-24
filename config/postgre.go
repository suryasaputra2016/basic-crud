package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "surya"
	password = "abc123"
	dbName   = "basic-crud"
)

func OpebDBConnection() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("connecting database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("connecting database: %w", err)
	}

	log.Println("database connected.")
	return db, nil
}

func CloseDBConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return fmt.Errorf("closing database connection: %w", err)
	}
	return nil
}

package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func Connect() (*sql.DB, error) {
	dns := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

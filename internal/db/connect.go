package db

import (
	"database/sql"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func Connect(dns string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	maxOpen := 20
	if v := os.Getenv("DB_MAX_OPEN_CONNS"); v != "" {
		if parsed, perr := strconv.Atoi(v); perr == nil && parsed > 0 {
			maxOpen = parsed
		}
	}
	maxIdle := 5
	if v := os.Getenv("DB_MAX_IDLE_CONNS"); v != "" {
		if parsed, perr := strconv.Atoi(v); perr == nil && parsed >= 0 {
			maxIdle = parsed
		}
	}
	connLifetime := 30 * time.Minute
	if v := os.Getenv("DB_CONN_MAX_LIFETIME"); v != "" {
		if d, derr := time.ParseDuration(v); derr == nil {
			connLifetime = d
		}
	}

	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxIdle)
	db.SetConnMaxLifetime(connLifetime)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

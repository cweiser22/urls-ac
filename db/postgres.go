package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver
)

// NewPostgresDB initializes and returns a *sqlx.DB for PostgreSQL.
func NewPostgresDB(connString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}
	// Optional: configure connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(0) // or time.Duration
	return db, nil
}

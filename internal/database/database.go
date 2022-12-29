package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Must is returning a *sqlx.DB instance or panic if error occurs.
func Must(db *sqlx.DB, err error) *sqlx.DB {
	if err != nil {
		panic(err)
	}

	return db
}

// New returns a new database connection or an error if one occurs.
func New(cfg Config) (*sqlx.DB, error) {
	if len(cfg.Name) == 0 && len(cfg.DriverName) == 0 {
		return nil, ErrInvalidConfiguration
	}

	db, err := sqlx.Open(cfg.DriverName, cfg.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection: %w", err)
	}

	return db, nil
}

// StatusCheck returns nil if it can successfully talk to the database. It
// returns a non-nil error otherwise.
func StatusCheck(ctx context.Context, db *sqlx.DB) error {
	// First check we can ping the database.
	var pingError error
	for attempts := 1; ; attempts++ {
		pingError = db.Ping()
		if pingError == nil {
			break
		}
		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	// Make sure we didn't time out or be cancelled.
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// Run a simple query to determine connectivity. Running this query forces a
	// round trip through the database.
	const q = `SELECT true`
	var tmp bool
	return db.QueryRowContext(ctx, q).Scan(&tmp)
}

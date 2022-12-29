package schema

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/jmoiron/sqlx"

	"githib.com/igomonov88/game-service/internal/database"

	"github.com/ardanlabs/darwin"
)

var (
	//go:embed migrations/schema.sql
	schemaDoc string

	//go:embed migrations/seed.sql
	seedDoc string
)

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(ctx context.Context, db *sqlx.DB) error {
	if err := database.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	driver, err := darwin.NewGenericDriver(db.DB, darwin.SqliteDialect{})
	if err != nil {
		return fmt.Errorf("construct darwin driver: %w", err)
	}

	return darwin.New(driver, darwin.ParseMigrations(schemaDoc)).Migrate()
}

// Seed runs the set of seed-data queries against db. The queries are ran in a
// transaction and rolled back if any fail.
func Seed(ctx context.Context, db *sqlx.DB) error {
	if err := database.StatusCheck(ctx, db); err != nil {
		return fmt.Errorf("status check database: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seedDoc); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

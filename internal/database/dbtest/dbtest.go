// Package dbtest contains supporting code for running tests that hit the DB.
package dbtest

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"githib.com/igomonov88/game-service/internal/database"
	"githib.com/igomonov88/game-service/internal/database/schema"
)

// NewUnit creates a test database. It creates the required table structure but
// the database is otherwise empty. It returns the database to use as well as a
// function to call at the end of the test.
func NewUnit(t *testing.T) (*zap.SugaredLogger, *sqlx.DB, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := database.New(database.Config{
		DriverName: "sqlite3",
		Name:       "billups_challenge_test.db",
	})
	if err != nil {
		t.Fatalf("Opening database connection: %v", err)
	}

	t.Log("Migrate and seed database ...")

	if err := schema.Migrate(ctx, db); err != nil {
		t.Fatalf("Migrating error: %s", err)
	}

	if err := schema.Seed(ctx, db); err != nil {
		t.Fatalf("Seeding error: %s", err)
	}

	t.Log("Ready for testing ...")

	var buf bytes.Buffer
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	writer := bufio.NewWriter(&buf)
	log := zap.New(
		zapcore.NewCore(encoder, zapcore.AddSync(writer), zapcore.DebugLevel)).
		Sugar()

	// teardown is the function that should be invoked when the caller is done
	// with the database.
	teardown := func() {
		t.Helper()
		db.Close()

		log.Sync()

		writer.Flush()
		fmt.Println("******************** LOGS ********************")
		fmt.Print(buf.String())
		fmt.Println("******************** LOGS ********************")
	}

	return log, db, teardown
}

// Test owns state for running and shutting down tests.
type Test struct {
	DB       *sqlx.DB
	Log      *zap.SugaredLogger
	Teardown func()

	t *testing.T
}

// NewIntegration creates a database, seeds it, constructs an authenticator.
func NewIntegration(t *testing.T) *Test {
	log, db, teardown := NewUnit(t)

	test := Test{
		DB:       db,
		Log:      log,
		t:        t,
		Teardown: teardown,
	}

	return &test
}

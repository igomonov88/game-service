package storage_test

import (
	"context"
	"testing"

	"githib.com/igomonov88/game-service/business/storage"
	"githib.com/igomonov88/game-service/internal/database"
	"githib.com/igomonov88/game-service/internal/database/schema"
)

// Success and failure markers.
const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func Test_Storage(t *testing.T) {
	db := database.Must(database.New(database.Config{
		DriverName: "sqlite3",
		Name:       "billups_challenge_test.db",
	}))
	svc := storage.NewService(db)
	if err := schema.Migrate(context.Background(), db); err != nil {
		t.Fatalf("failed to migrate database")
	}
	if err := schema.Seed(context.Background(), db); err != nil {
		t.Fatalf("Should be able to seed the database:.")
	}
	t.Run("store", func(t *testing.T) {
		t.Logf("Given the need to test how we store results in storage.")
		{
			if err := svc.Store(context.Background(), 1, 1, "tie"); err != nil {
				t.Fatalf("\t%s\tShould be able to store a game result: %s.", Failed, err)
			}
			t.Logf("\t%s\tShould be able to store a game result.", Success)
		}
	})

	t.Run("retrieve", func(t *testing.T) {
		t.Logf("Given the need to test how we receive results.")
		{
			limit := 1
			results, err := svc.GetResults(context.Background(), limit)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to retrieve game results: %s.", Failed, err)
			}

			if len(results) != limit {
				t.Fatalf("\t%s\tShould be able to retrieve %d game results: %d.", Failed, limit, len(results))
			}

			t.Logf("\t%s\tShould be able to retrieve %d game results.", Success, limit)
		}
	})
	t.Run("clear", func(t *testing.T) {
		t.Logf("Given the need to clear results.")
		{
			if err := svc.ClearResults(context.Background()); err != nil {
				t.Fatalf("\t%s\tShould be able to clear game results: %s.", Failed, err)
			}
			t.Logf("\t%s\tShould be able to clear game results.", Success)
		}
	})
}

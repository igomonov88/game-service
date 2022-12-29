package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"githib.com/igomonov88/game-service/internal/database"
	"githib.com/igomonov88/game-service/internal/database/schema"
)

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")

// Migrate creates the schema in the database.
func Migrate(cfg database.Config) error {
	db, err := database.New(cfg)
	if err != nil {
		return errors.Wrap(err, "connect database")
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := schema.Migrate(ctx, db); err != nil {
		return errors.Wrap(err, "migrate database")
	}

	fmt.Println("migrations complete")
	return nil
}

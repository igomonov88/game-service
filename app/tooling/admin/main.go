package main

import (
	"fmt"
	"os"

	"githib.com/igomonov88/game-service/app/tooling/admin/commands"
	"githib.com/igomonov88/game-service/internal/config"
	"githib.com/igomonov88/game-service/internal/database"
	"githib.com/igomonov88/game-service/internal/logger"
)

func main() {
	// Construct the application logger.
	log := logger.Must(logger.New("game-service"))

	// Perform the startup and shutdown sequence.
	if err := run(); err != nil {
		log.Errorw("startup ERROR: %v", err)
		log.Sync()
		os.Exit(1)
	}
}

func run() error {

	// =========================================================================
	// Configuration
	cfg := config.Must(config.ReadConfig())
	db := database.Must(database.New(cfg.Database))
	defer func() {
		db.Close()
	}()

	// =========================================================================
	// Commands
	cmds := os.Args[1:]
	if len(cmds) > 0 {
		switch cmds[0] {
		case "migrate":
			if err := commands.Migrate(cfg.Database); err != nil {
				return fmt.Errorf("failed to migrate database: %w", err)
			}
			return nil
		default:
			fmt.Println("migrate: create the schema in the database")
			return fmt.Errorf("unknown command: %s", cmds[0])
		}
	}

	return nil
}

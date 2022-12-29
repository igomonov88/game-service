package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/business/services/game"
	"githib.com/igomonov88/game-service/business/services/health"
	"githib.com/igomonov88/game-service/business/services/play"
	playStorage "githib.com/igomonov88/game-service/business/storage"
	"githib.com/igomonov88/game-service/internal/clients/codechallenge"
	"githib.com/igomonov88/game-service/internal/config"
	"githib.com/igomonov88/game-service/internal/database"
	"githib.com/igomonov88/game-service/internal/logger"
	"githib.com/igomonov88/game-service/internal/server"

	"githib.com/igomonov88/game-service/app/service/api/handlers"
	"githib.com/igomonov88/game-service/business/services/choice"
)

func main() {
	// Construct the application logger.
	log := logger.Must(logger.New("game-service"))

	// Perform the startup and shutdown sequence.
	if err := run(log); err != nil {
		log.Errorw("startup", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}

func run(logger *zap.SugaredLogger) error {

	// =========================================================================
	// Configuration

	logger.Infow("Starting service.")
	defer logger.Infow("Shutdown the game service completed.")

	// Read provided configuration.
	cfg := config.Must(config.ReadConfig())
	logger.Infow("Configuration read successfully.")

	// =========================================================================
	// Start Database
	db := database.Must(database.New(cfg.Database))
	defer func() {
		logger.Infof("Closing database connection.")
		db.Close()
	}()
	logger.Infof("Database connection established.")

	// =========================================================================
	// Start Services
	storage := playStorage.NewService(db)
	client := codechallenge.Must(codechallenge.NewClient(logger, cfg.CodeChallenge))
	gameService := game.NewService(logger)
	playService := play.NewService(logger, storage, gameService, client)
	choiceService := choice.NewService(logger, client, gameService)
	healthService := health.NewService(db)
	handler := handlers.Handler(logger, choiceService, playService, healthService)
	logger.Infof("Services and handler started.")

	// =========================================================================
	// Start API Service

	logger.Infof("Starting API service.")

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	srv := server.New(cfg.Server, handler, zap.NewStdLog(logger.Desugar()))
	// Make a channel to listen for errors coming from the listener. Use a
	// buffered channel so the goroutine can exit if we don't collect this error.
	serverErrors := make(chan error, 1)

	// Start the service listening for config requests.
	go func() {
		logger.Infof("API service listening on %s", srv.Addr)
		serverErrors <- srv.ListenAndServe()
	}()

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case _ = <-shutdown:
		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*cfg.Server.ShutdownTimeout)
		defer cancel()

		// Asking listener to shut down and shed load.
		if err := srv.Shutdown(ctx); err != nil {
			srv.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}

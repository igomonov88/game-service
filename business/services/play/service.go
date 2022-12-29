package play

import (
	"context"

	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/business/services/game"
	"githib.com/igomonov88/game-service/internal/database"
)

// Storage is the interface that decouple the storage methods. We can mock this
// interface in tests.
type Storage interface {
	Store(ctx context.Context, userChoice int, computerChoice int, result string) error
	GetResults(ctx context.Context, limit int) ([]database.Result, error)
	ClearResults(ctx context.Context) error
}

type RandomGenerator interface {
	GenerateRandomNumber(ctx context.Context) int
}

type Service struct {
	storage     Storage
	generator   RandomGenerator
	gameService *game.Service
	logger      *zap.SugaredLogger
}

func NewService(logger *zap.SugaredLogger, storage Storage, gameService *game.Service, generator RandomGenerator) *Service {
	return &Service{
		logger:      logger,
		storage:     storage,
		generator:   generator,
		gameService: gameService,
	}
}

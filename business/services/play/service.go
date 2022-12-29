package play

import (
	"context"

	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/internal/database"
)

// Storage is the interface that decouple the storage methods. We can mock this
// interface in tests.
type Storage interface {
	Store(ctx context.Context, userChoice int, computerChoice int, result string) error
	GetResults(ctx context.Context, limit int) ([]database.Result, error)
	ClearResults(ctx context.Context) error
}

// Player is the interface that decouple game package and operate only with methods
// which play service needs. This interface also can be useful to mock game
// package logic in tests.
type Player interface {
	RandomChoiceID(ctx context.Context) int
	Play(ctx context.Context, userChoice int, computerChoice int) string
}

type Service struct {
	storage Storage
	player  Player
	logger  *zap.SugaredLogger
}

func NewService(logger *zap.SugaredLogger, storage Storage, player Player) *Service {
	return &Service{
		logger:  logger,
		storage: storage,
		player:  player,
	}
}

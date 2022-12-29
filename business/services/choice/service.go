package choice

import (
	"context"

	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/business/services/game"
)

type RandomGenerator interface {
	GenerateRandomNumber(ctx context.Context) int
}

type Service struct {
	logger      *zap.SugaredLogger
	generator   RandomGenerator
	gameService *game.Service
}

func NewService(logger *zap.SugaredLogger, generator RandomGenerator, gameService *game.Service) *Service {
	return &Service{
		logger:      logger,
		generator:   generator,
		gameService: gameService,
	}
}

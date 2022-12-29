package choice

import (
	"go.uber.org/zap"

	"githib.com/igomonov88/game-service/business/services/game"
)

type Service struct {
	logger      *zap.SugaredLogger
	gameService *game.Service
}

func NewService(logger *zap.SugaredLogger, gameService *game.Service) *Service {
	return &Service{
		logger:      logger,
		gameService: gameService,
	}
}

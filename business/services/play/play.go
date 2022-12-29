package play

import (
	"context"
	"fmt"

	"githib.com/igomonov88/game-service/business/contract"
	"githib.com/igomonov88/game-service/internal/middlewaries"
)

// Play plays a game of rock, paper, scissors game with the provided player
// choice and random computer choice.
func (s *Service) Play(ctx context.Context, userChoice int) (*contract.PlayResponse, error) {
	reqID := middlewaries.RequestIDFromContext(ctx)
	computerChoice := s.player.RandomChoiceID(ctx)
	middlewaries.RequestIDFromContext(ctx)
	s.logger.Infof("request_id: %v, playing game with user choice: %d, and computer choice: %d", reqID, userChoice, computerChoice)
	result := s.player.Play(ctx, userChoice, computerChoice)
	if err := s.storage.Store(ctx, userChoice, computerChoice, result); err != nil {
		s.logger.Errorf("request_id: %v, failed to store play result: %s", reqID, err)
		return nil, fmt.Errorf("failed to store result: %w", err)
	}

	return &contract.PlayResponse{Player: userChoice, Computer: computerChoice, Results: result}, nil
}

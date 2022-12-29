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
	computerChoice := s.fromGeneratedRandomNumberToChoiceID(s.generator.GenerateRandomNumber(ctx))
	s.logger.Infof("request_id: %v, playing game with user choice: %v, and computer choice: %v", reqID, userChoice, computerChoice)
	result := s.gameService.Play(ctx, userChoice, computerChoice)
	if err := s.storage.Store(ctx, userChoice, computerChoice, result); err != nil {
		s.logger.Errorf("request_id: %v, failed to store play result: %s", reqID, err)
		return nil, fmt.Errorf("failed to store result: %w", err)
	}

	return &contract.PlayResponse{Player: userChoice, Computer: computerChoice, Results: result}, nil
}

// fromGeneratedRandomNumberToChoiceID gets a random number from the client
// response and converts it to a choice ID.
func (s *Service) fromGeneratedRandomNumberToChoiceID(number int) int {
	switch {
	case number <= 20:
		return 1
	case number > 20 && number <= 40:
		return 2
	case number > 40 && number <= 60:
		return 3
	case number > 60 && number <= 80:
		return 4
	default:
		return 5
	}
}

package choice

import (
	"context"
	"fmt"

	"githib.com/igomonov88/game-service/business/contract"
	"githib.com/igomonov88/game-service/internal/middlewaries"
)

// GetChoice returns a random choice from the game or error if choiceID is
// out of range.
func (s *Service) GetChoice(ctx context.Context) (*contract.ChoiceResponse, error) {
	return s.fromChoiceIdToResponse(ctx, s.gameService.RandomChoiceID(ctx))
}

// GetChoices returns a list of all possible choices.
func (s *Service) GetChoices(ctx context.Context) []contract.ChoiceResponse {
	s.logger.Info("getting choices")
	choices := s.gameService.AvailableChoices(ctx)
	response := make([]contract.ChoiceResponse, 0, len(choices))

	for id, name := range choices {
		response = append(response, contract.ChoiceResponse{ID: id, Name: name})
	}

	return response
}

func (s *Service) fromChoiceIdToResponse(ctx context.Context, choiceID int) (*contract.ChoiceResponse, error) {
	reqID := middlewaries.RequestIDFromContext(ctx)
	s.logger.Info("request_id: %v getting choice", reqID)
	choices := s.gameService.AvailableChoices(ctx)
	choiceName, exist := choices[choiceID]
	if !exist {
		s.logger.Errorf("request_id: %v, choiceID %d is out of range", reqID, choiceID)
		return nil, fmt.Errorf("provided random choiceID out of expected range")
	}

	return &contract.ChoiceResponse{ID: choiceID, Name: choiceName}, nil
}

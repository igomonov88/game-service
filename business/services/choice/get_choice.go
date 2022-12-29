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
	s.logger.Info("getting choice")
	choiceID := s.generator.GenerateRandomNumber(ctx)

	return s.fromChoiceIdToResponse(ctx, s.fromGeneratedRandomNumberToChoiceID(choiceID))
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

// fromGeneratedRandomNumberToChoiceID gets a random number from the generator
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

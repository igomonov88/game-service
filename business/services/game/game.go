package game

import (
	"context"

	"githib.com/igomonov88/game-service/internal/middlewaries"
	"githib.com/igomonov88/game-service/internal/random"
)

const maxChoiceID = 5

// RandomChoiceID returns a random choice ID.
func (s *Service) RandomChoiceID(ctx context.Context) int {
	choiceID := random.Int(maxChoiceID)
	reqID := middlewaries.RequestIDFromContext(ctx)
	s.logger.Infof("request_id: %v, generated random choice ID: %d", reqID, choiceID)
	return choiceID
}

// AvailableChoices returns a map of all possible choices.
func (s *Service) AvailableChoices(ctx context.Context) map[int]string {
	reqID := middlewaries.RequestIDFromContext(ctx)
	s.logger.Infof("request_id: %v, choices: %v", reqID, s.choices)
	return s.choices
}

// Play returns the result of the game based on the provided player and
// computer choice.
func (s *Service) Play(ctx context.Context, userChoice int, computerChoice int) string {
	reqID := middlewaries.RequestIDFromContext(ctx)

	switch {
	case userChoice == computerChoice:
		s.logger.Infof("request_id: %v, user's choice: %d, and computer's choice: %d, resulted in a tie", reqID, userChoice, computerChoice)
		return "tie"
	case userChoice == rock && (computerChoice == scissors || computerChoice == lizard):
		s.logger.Infof("request_id: %v, user's choice: %d, and computer's choice: %d, resulted in a win", reqID, userChoice, computerChoice)
		return "win"
	case userChoice == paper && (computerChoice == rock || computerChoice == spock):
		s.logger.Infof("request_id: %v, user's choice: %d, and computer's choice: %d, resulted in a win", reqID, userChoice, computerChoice)
		return "win"
	case userChoice == scissors && (computerChoice == paper || computerChoice == lizard):
		s.logger.Infof("request_id: %v, user's choice: %d, and computer's choice: %d, resulted in a win", reqID, userChoice, computerChoice)
		return "win"
	case userChoice == lizard && (computerChoice == paper || computerChoice == spock):
		s.logger.Infof("request_id: %v, user's choice: %d, and computer's choice: %d, resulted in a win", reqID, userChoice, computerChoice)
		return "win"
	case userChoice == spock && (computerChoice == rock || computerChoice == scissors):
		s.logger.Infof("request_id: %v, user's choice: %d, and computer's choice: %d, resulted in a win", reqID, userChoice, computerChoice)
		return "win"
	default:
		s.logger.Warnf("request_id: %v, user's choice: %d, and computer's choice: %d, resulted in a lose", reqID, userChoice, computerChoice)
		return "lose"
	}
}

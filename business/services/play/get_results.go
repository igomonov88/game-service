package play

import (
	"context"
	"errors"
	"fmt"

	"githib.com/igomonov88/game-service/business/contract"
	"githib.com/igomonov88/game-service/internal/database"
	"githib.com/igomonov88/game-service/internal/middlewaries"
)

const resultLimitations = 10

// GetResults returns the last 10 results.
func (s *Service) GetResults(ctx context.Context) ([]contract.PlayResponse, error) {
	reqID := middlewaries.RequestIDFromContext(ctx)
	s.logger.Infof("request_id: %v, getting results", reqID)
	results, err := s.storage.GetResults(ctx, resultLimitations)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			var response []contract.PlayResponse
			return response, nil // return empty slice instead of error
		}
		s.logger.Errorf("request_id: %v failed to get results: %s", reqID, err)
		return nil, fmt.Errorf("failed to get results: %w", err)
	}

	response := make([]contract.PlayResponse, 0, len(results))
	for _, result := range results {
		response = append(response, contract.PlayResponse{
			Player:   result.UsersChoice,
			Computer: result.ComputersChoice,
			Results:  result.Result,
		})
	}

	return response, nil
}

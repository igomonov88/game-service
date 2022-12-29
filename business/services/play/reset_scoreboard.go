package play

import (
	"context"
	"fmt"

	"githib.com/igomonov88/game-service/internal/middlewaries"
)

// ResetScoreboard resets the scoreboard by add to
func (s *Service) ResetScoreboard(ctx context.Context) error {
	reqID := middlewaries.RequestIDFromContext(ctx)
	s.logger.Infof("request_id: %v, resetting scoreboard", reqID)
	if err := s.storage.ClearResults(ctx); err != nil {
		s.logger.Errorf("request_id: %v, failed to clear results: %s", reqID, err)
		return fmt.Errorf("failed to clear results: %w", err)
	}
	return nil
}

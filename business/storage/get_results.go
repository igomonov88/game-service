package storage

import (
	"context"

	"githib.com/igomonov88/game-service/internal/database"
)

// GetResults returns results from database with provided count limitation.
func (s *Service) GetResults(ctx context.Context, limit int) ([]database.Result, error) {
	const query = `SELECT * FROM results WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $2`
	var results []database.Result
	if err := s.db.SelectContext(ctx, &results, query, limit); err != nil {
		return nil, database.WrapError(err)
	}

	return results, nil
}

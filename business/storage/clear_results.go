package storage

import (
	"context"
	"time"

	"githib.com/igomonov88/game-service/internal/database"
)

// ClearResults set deleted_at parameter in all records. I know that this is
// absolutely not the best wat to handle this situation, but since we do not
// have unique identifier for each player, I decided to use this approach.
// Alternative approach is to implement some user authentication and store it
// in database. But I think that this is out of scope of this task.
func (s *Service) ClearResults(ctx context.Context) error {
	const query = `UPDATE results SET deleted_at = $1 WHERE deleted_at IS NULL`
	res, err := s.db.ExecContext(ctx, query, time.Now())
	if err != nil {
		return err
	}
	return database.CheckAffectedRows(res)
}

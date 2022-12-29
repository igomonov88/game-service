package storage

import (
	"context"
	"time"

	"githib.com/igomonov88/game-service/internal/database"
)

// Store stores the result of the game in database.
func (s *Service) Store(ctx context.Context, userChoice int, computerChoice int, result string) error {
	const query = `INSERT INTO results (user_choice, computer_choice, result, created_at) VALUES (:user_choice, :computer_choice, :result, :created_at)`
	res, err := s.db.NamedExecContext(ctx, query, database.Result{
		UsersChoice:     userChoice,
		ComputersChoice: computerChoice,
		Result:          result,
		CreatedAt:       time.Now(),
	})
	if err != nil {
		return err
	}

	return database.CheckAffectedRows(res)
}

package health

import (
	"context"

	"github.com/jmoiron/sqlx"

	"githib.com/igomonov88/game-service/internal/database"
)

type Service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) HealthCheck(ctx context.Context) error {
	return database.StatusCheck(ctx, s.db)
}

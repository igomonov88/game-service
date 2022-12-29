package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrInvalidConfiguration = errors.New("invalid configuration provided for database")
	ErrNotFound             = errors.New("not found")
	ErrNoAffectedRows       = errors.New("no affected rows")
)

func CheckAffectedRows(rs sql.Result) error {
	c, err := rs.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}

	if c == 0 {
		return ErrNoAffectedRows
	}

	return nil
}

func WrapError(err error) error {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return ErrNotFound
	default:
		return err
	}
}

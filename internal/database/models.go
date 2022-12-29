package database

import "time"

type Result struct {
	UsersChoice     int        `db:"user_choice"`
	ComputersChoice int        `db:"computer_choice"`
	Result          string     `db:"result"`
	CreatedAt       time.Time  `db:"created_at"`
	DeletedAt       *time.Time `db:"deleted_at"`
}

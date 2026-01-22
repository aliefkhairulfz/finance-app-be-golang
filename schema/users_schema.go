package schema

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type UsersSchema struct{}

type User struct {
	ID        int    `db:"id"`
	Email     string `db:"email"`
	IsActive  bool   `db:"is_active"`
	Avatar    string `db:"avatar"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (s *UsersSchema) Up(db *sqlx.DB) error {
	schema := `
		CREATE TABLE users (
			id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			email VARCHAR(100) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			avatar TEXT,
			is_active BOOLEAN DEFAULT FALSE,
			refresh_token TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	if err := db.MustExec(schema); err != nil {
		return errors.New("failed to create users table")
	}

	return nil
}

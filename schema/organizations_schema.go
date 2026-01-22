package schema

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type OrganizationsSchema struct{}

type Organization struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	AuthorID    int    `db:"author_id"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

func (s *OrganizationsSchema) Up(db *sqlx.DB) error {
	schema := `
		CREATE TABLE organizations (
			id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			name VARCHAR(100) NOT NULL UNIQUE,
			description TEXT,
			author_id BIGINT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

			CONSTRAINT fk_organizations_author
				FOREIGN KEY(author_id)
				REFERENCES users(id)
				ON DELETE CASCADE,

			CONSTRAINT unique_org_per_author
				UNIQUE (author_id, name)
		);

		CREATE INDEX idx_organizations_author_id
		ON organizations(author_id);
	`

	if err := db.MustExec(schema); err != nil {
		return errors.New("failed to create organizations table")
	}

	return nil
}

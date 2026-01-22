package schema

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type BranchesSchema struct{}

type Branches struct {
	ID             int    `db:"id"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	AuthorID       int    `db:"author_id"`
	OrganizationID int    `db:"organization_id"`
	CreatedAt      string `db:"created_at"`
	UpdatedAt      string `db:"updated_at"`
}

func (s *BranchesSchema) Up(db *sqlx.DB) error {
	schema := `
		CREATE TABLE branches (
			id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			name VARCHAR(100) NOT NULL UNIQUE,
			description TEXT,
			author_id BIGINT NOT NULL,
			organization_id BIGINT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

			CONSTRAINT fk_branches_author
				FOREIGN KEY(author_id)
				REFERENCES users(id)
				ON DELETE CASCADE,

			CONSTRAINT fk_branches_organization
				FOREIGN KEY(organization_id)
				REFERENCES organizations(id)
				ON DELETE CASCADE,

			CONSTRAINT unique_branch_per_org
				UNIQUE (organization_id, name)
		);

		CREATE INDEX idx_branches_organization_id
		ON branches(organization_id);

		CREATE INDEX idx_branches_author_id
		ON branches(author_id);
	`

	if err := db.MustExec(schema); err != nil {
		return errors.New("failed to create branches table")
	}

	return nil
}

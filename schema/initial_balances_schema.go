package schema

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type InitialBalancesSchema struct{}

type InitialBalance struct {
	ID        int    `db:"id"`
	Amount    int    `db:"amount"`
	Notes     string `db:"notes"`
	BranchID  int    `db:"branch_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

func (s *InitialBalancesSchema) Up(db *sqlx.DB) error {
	schema := `
		CREATE TABLE initial_balances (
			id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
			amount INT NOT NULL DEFAULT 0,
			notes TEXT,
			branch_id BIGINT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

			CONSTRAINT fk_initial_balances_branch
				FOREIGN KEY (branch_id)
				REFERENCES branches (id)
				ON DELETE CASCADE,

			CONSTRAINT unique_initial_balance_per_branch
				UNIQUE (branch_id)
		);

		CREATE INDEX idx_initial_balances_branch_id
		ON initial_balances(branch_id);
	`

	if err := db.MustExec(schema); err != nil {
		return errors.New("failed to create initial_balances table")
	}

	return nil
}

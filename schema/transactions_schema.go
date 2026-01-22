package schema

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type TransactionsSchema struct{}

type Transaction struct {
	ID               int    `db:"id"`
	TransactionType  string `db:"transaction_type"`
	Amount           int    `db:"amount"`
	Notes            string `db:"notes"`
	EditHistory      string `db:"edit_history"`
	InitialBalanceID int    `db:"initial_balance_id"`
	UserID           int    `db:"user_id"`
	BranchID         int    `db:"branch_id"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}

func (s *TransactionsSchema) Up(db *sqlx.DB) error {
	schema := `
		CREATE TABLE transactions (
		    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		    transaction_type transaction_type_enum NOT NULL,
		    amount INT NOT NULL,
		    notes TEXT,
		    edit_history TEXT,
		    initial_balance_id BIGINT NOT NULL,
		    user_id BIGINT NOT NULL,
		    branch_id BIGINT NOT NULL,
		    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

		    CONSTRAINT fk_transactions_initial_balance
		        FOREIGN KEY (initial_balance_id)
		        REFERENCES initial_balances(id)
		        ON DELETE CASCADE,

		    CONSTRAINT fk_transactions_user
		        FOREIGN KEY (user_id)
		        REFERENCES users(id)
		        ON DELETE CASCADE,

		    CONSTRAINT fk_transactions_branch
		        FOREIGN KEY (branch_id)
		        REFERENCES branches(id)
		        ON DELETE CASCADE
		);

		CREATE INDEX idx_transactions_initial_balance_id
		ON transactions(initial_balance_id);

		CREATE INDEX idx_transactions_user_id
		ON transactions(user_id);

		CREATE INDEX idx_transactions_branch_id
		ON transactions(branch_id);
	`

	if err := db.MustExec(schema); err != nil {
		return errors.New("failed to create transactions table")
	}

	return nil
}

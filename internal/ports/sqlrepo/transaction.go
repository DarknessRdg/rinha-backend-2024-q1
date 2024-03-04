package sqlrepo

import (
	"database/sql"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
)

type SqlTransactionRepo struct {
	db *sql.DB
}

func NewSqlTransactionRepo(db *sql.DB) *SqlTransactionRepo {
	return &SqlTransactionRepo{db: db}
}

func (s *SqlTransactionRepo) Insert(transaction domain.Transaction) error {
	stmt, err := s.db.Prepare(`
		INSERT INTO transaction (account_id, amount, type, description, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		transaction.AccountId,
		transaction.Amount,
		transaction.Type,
		transaction.Description,
		transaction.CreatedAt,
	)
	return err
}

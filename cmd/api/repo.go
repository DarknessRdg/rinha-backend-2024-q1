package main

import (
	"database/sql"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/ports/transaction/sqlrepo"
	transaction_repo "github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
	account_repo "github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/repo"
)

func NewIAccountRepo(db *sql.DB) account_repo.IAccountRepo {
	return &sqlrepo.SqlAccountRepo{Db: db}
}

func NewITransactionRepo(db *sql.DB) transaction_repo.ITransactionRepo {
	return sqlrepo.NewSqlTransactionRepo(db)
}

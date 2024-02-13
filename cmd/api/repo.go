package main

import (
	"database/sql"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/ports/transaction/sqlrepo"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
)

func NewIAccountRepo(db *sql.DB) repo.IAccountRepo {
	return &sqlrepo.SqlAccountRepo{Db: db}
}

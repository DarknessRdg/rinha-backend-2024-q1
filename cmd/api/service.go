package main

import (
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/service"
)

func NewITransactionService(accountRepo repo.IAccountRepo, transactionRepo repo.ITransactionRepo) service.ITransactionService {
	return service.NewTransactionService(accountRepo, transactionRepo)
}

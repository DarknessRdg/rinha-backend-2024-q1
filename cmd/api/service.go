package main

import (
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
	transaction_service "github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/service"
	account_service "github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/service"
)

func NewITransactionService(
	accountService account_service.IAccountService, 
	transactionRepo repo.ITransactionRepo,
) transaction_service.ITransactionService {
	return transaction_service.NewTransactionService(
		accountService,
		transactionRepo,
	)
}

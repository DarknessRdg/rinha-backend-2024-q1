package service

import "github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"

type ITransactionService interface {
	TransactionClient(clientId string, transactionDto dto.TransactionDto) (dto.TransactionResult, error)
}

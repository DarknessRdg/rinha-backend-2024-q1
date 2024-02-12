package service

import (
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
)

type TransactionService struct {
	accountRepo repo.IAccountRepo
}

func (service *TransactionService) PostTransaction(
	accountId string,
	transactionDto dto.TransactionDto,
) (dto.TransactionResult, error) {
	var result dto.TransactionResult

	if !service.accountRepo.ExistsById(accountId) {
		return result, errs.NotFound("account id does not exists")
	}
	return result, nil
}

package service

import (
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
)

type AccountRepo interface {
	IncrementBalance(accountId string, amount int) error
	ExistsById(accountId string) bool
}

type TranasctionService struct {
	accountRepo AccountRepo
}

func (service *TranasctionService) PostTransaction(
	accountId string,
	transactionDto dto.TransactionDto,
) (dto.TransactionResult, error) {
	var result dto.TransactionResult

	if !service.accountRepo.ExistsById(accountId) {
		return result, errs.NotFound("account id does not exists")
	}
	return result, nil
}

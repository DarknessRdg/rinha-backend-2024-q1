package service

import (
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
)

type TransactionService struct {
	accountRepo     repo.IAccountRepo
	transactionRepo repo.ITransactionRepo
}

func (service *TransactionService) PostTransaction(
	id string,
	transactionDto dto.TransactionDto,
) (dto.TransactionResult, error) {
	account, err := service.accountRepo.GetByIdAndLock(domain.AccountId(id))
	if err != nil {
		return dto.TransactionResult{}, err
	}

	err = account.Debit(domain.MoneyCents(transactionDto.AmountCents))
	if err != nil {
		return dto.TransactionResult{}, err
	}

	err = service.accountRepo.Update(account)
	if err != nil {
		return dto.TransactionResult{}, err
	}

	err = service.transactionRepo.Insert()
	if err != nil {
		return dto.TransactionResult{}, err
	}

	return dto.TransactionResult{
		LimitCents:   account.Limit,
		BalanceCents: int(account.Balance),
	}, nil
}

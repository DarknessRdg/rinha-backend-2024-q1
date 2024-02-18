package service

import (
	"strings"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
)

type TransactionService struct {
	accountRepo     repo.IAccountRepo
	// transactionRepo repo.ITransactionRepo
}

func NewTransactionService(accountRepo repo.IAccountRepo) *TransactionService {
	return &TransactionService{accountRepo: accountRepo}
}

func (service *TransactionService) PostTransaction(
	id int,
	transactionDto dto.TransactionDto,
) (dto.TransactionResult, error) {
	account, err := service.getAccountLocked(domain.AccountId(id))
	if err != nil {
		return dto.TransactionResult{}, err
	}

	err = service.creditOrDebit(account, transactionDto)
	if err != nil {
		return dto.TransactionResult{}, err
	}

	err = service.accountRepo.Update(account)
	if err != nil {
		return dto.TransactionResult{}, err
	}

	// err = service.transactionRepo.Insert()
	// if err != nil {
	// 	return dto.TransactionResult{}, err
	// }

	return dto.TransactionResult{
		LimitCents:   account.Limit,
		BalanceCents: int(account.Balance),
	}, nil
}

func (service *TransactionService) getAccountLocked(id domain.AccountId) (*domain.Account, error) {
	account, err := service.accountRepo.GetByIdAndLock(id)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errs.NotFound("Account not found.")
	}
	return account, nil
}

func (service *TransactionService) creditOrDebit(account *domain.Account, transaction dto.TransactionDto) error {
	amount := domain.MoneyCents(transaction.AmountCents)

	switch strings.ToLower(transaction.Type) {
	case "d":
		return account.Debit(amount)
	case "c":
		return account.Credit(amount)
	}

	return errs.UnprocessableEntity("Invalid transaction type. It should be either 'c' or 'd'")
}

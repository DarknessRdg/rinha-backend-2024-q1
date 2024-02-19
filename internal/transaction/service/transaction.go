package service

import (
	"strings"
	"time"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
)

type TransactionService struct {
	accountRepo     repo.IAccountRepo
	transactionRepo repo.ITransactionRepo
}

func NewTransactionService(accountRepo repo.IAccountRepo, transactionRepo repo.ITransactionRepo) *TransactionService {
	return &TransactionService{accountRepo: accountRepo, transactionRepo: transactionRepo}
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

	transaction := domain.Transaction{
		AccountId:   account.Id,
		Amount:      domain.MoneyCents(transactionDto.AmountCents),
		Description: transactionDto.Description,
		Type:        transactionDto.Type,
		CreatedAt:   time.Now(),
	}
	err = service.transactionRepo.Insert(transaction)
	if err != nil {
		return dto.TransactionResult{}, err
	}

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

package service

import (
	"strings"
	"time"

	account_service "github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/service"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
)

type TransactionService struct {
	accountService  account_service.IAccountService
	transactionRepo repo.ITransactionRepo
}

func NewTransactionService(
	accountService account_service.IAccountService,
	transactionRepo repo.ITransactionRepo,
) *TransactionService {
	return &TransactionService{
		accountService:  accountService,
		transactionRepo: transactionRepo,
	}
}

func (service *TransactionService) PostTransaction(
	id int,
	transactionDto dto.TransactionDto,
) (dto.TransactionResult, error) {
	account, err := service.accountService.CreditOrDebit(id, transactionDto.AmountCents, transactionDto.Type)
	if err != nil {
		return dto.TransactionResult{}, err
	}

	transaction := domain.Transaction{
		AccountId:   domain.AccountId(account.Id),
		Amount:      domain.MoneyCents(transactionDto.AmountCents),
		Description: transactionDto.Description,
		Type:        domain.Operation(transactionDto.Type),
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

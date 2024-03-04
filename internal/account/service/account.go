package service

import (
	"fmt"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/repo"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
)

type AccountService struct {
	accountRepo repo.IAccountRepo
}

func (service *AccountService) CreditOrDebit(accountId int, amountCents int, operation string) (dto.AccountDto, error) {
	id, amount := domain.AccountId(accountId), domain.MoneyCents(amountCents)

	op, err := domain.OperationFromString(operation)
	if err != nil {
		return dto.AccountDto{}, err
	}

	account, err := service.getAccountLocked(id)
	if err != nil {
		return dto.AccountDto{}, err
	}

	err = account.UpdateBalance(amount, op)
	if err != nil {
		return dto.AccountDto{}, err
	}

	err = service.accountRepo.Update(account)
	if err != nil {
		return dto.AccountDto{}, err
	}

	return dto.AccountDto{
		Limit:   account.Limit,
		Balance: int(account.Balance),
	}, nil
}

func (service *AccountService) getAccountLocked(id domain.AccountId) (*domain.Account, error) {
	account, err := service.accountRepo.GetByIdAndLock(id)

	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errs.NotFound(fmt.Sprintf("account with id %v not found", id))
	}

	return account, nil
}

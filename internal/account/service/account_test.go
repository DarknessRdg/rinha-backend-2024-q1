package service

import (
	"errors"
	"testing"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/repo"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/stretchr/testify/require"
)

func TestAccountService_C(t *testing.T) {
	t.Run("When account does not exists, Then return Not Found error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)
		service := NewAccountService(accountRepo)

		accountId := domain.AccountId(1)
		accountRepo.EXPECT().GetByIdAndLock(accountId).Once().Return(nil, nil)

		result, err := service.CreditOrDebit(int(accountId), 1, "c")

		require.Empty(t, result)
		require.ErrorIs(t, err, errs.NotFound(""))
	})

	t.Run("When error happens finding account, Then return the error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)
		service := NewAccountService(accountRepo)

		accountId := domain.AccountId(1)
		dbErr := errors.New("database error")
		accountRepo.EXPECT().GetByIdAndLock(accountId).Once().Return(nil, dbErr)

		result, err := service.CreditOrDebit(int(accountId), 1, "c")

		require.Error(t, err)
		require.Equal(t, err, dbErr)
		require.Empty(t, result)
	})

	t.Run("When find account but amount the account limit to debit, Then return error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewAccountService(accountRepo)

		account := &domain.Account{
			Id:      domain.AccountId(1),
			Limit:   0,
			Balance: 0,
		}

		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)

		result, err := service.CreditOrDebit(int(account.Id), 1, "d")
		require.Error(t, err)
		require.ErrorIs(t, err, errs.UnprocessableEntity(""))
		require.Empty(t, result)
	})

	t.Run("When debit operation succeed but update account fails, Then return error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewAccountService(accountRepo)

		account := &domain.Account{
			Id:      domain.AccountId(1),
			Limit:   123,
			Balance: 0,
		}

		dbErr := errors.New("database error")
		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)
		accountRepo.EXPECT().Update(account).Once().Return(dbErr)

		result, err := service.CreditOrDebit(int(account.Id), 123, "d")
		require.Error(t, err)
		require.Equal(t, dbErr, err)
		require.Empty(t, result)
	})

	t.Run("When invalid operation, Return an error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewAccountService(accountRepo)

		result, err := service.CreditOrDebit(1, 123, "unknown")
		require.ErrorIs(t, err, errs.UnprocessableEntity(""))
		require.Empty(t, result)
	})

	t.Run("When successfully debit, Then return the account limit and balance updated", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewAccountService(accountRepo)

		account := &domain.Account{
			Id:      domain.AccountId(1),
			Limit:   123,
			Balance: 0,
		}

		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)
		accountRepo.EXPECT().Update(account).Once().Return(nil)

		result, err := service.CreditOrDebit(int(account.Id), 123, "d")
		require.NoError(t, err)

		expectedResult := dto.AccountDto{
			Id:      int(account.Id),
			Limit:   account.Limit,
			Balance: -123,
		}
		require.Equal(t, expectedResult, result)
	})

	t.Run("When successfully credit, Then return the account limit and balance updated", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewAccountService(accountRepo)

		account := &domain.Account{
			Id:      domain.AccountId(1),
			Limit:   123,
			Balance: 0,
		}

		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)
		accountRepo.EXPECT().Update(account).Once().Return(nil)

		result, err := service.CreditOrDebit(int(account.Id), 123, "c")
		require.NoError(t, err)

		expectedResult := dto.AccountDto{
			Id:      int(account.Id),
			Limit:   account.Limit,
			Balance: 123,
		}
		require.Equal(t, expectedResult, result)
	})
}

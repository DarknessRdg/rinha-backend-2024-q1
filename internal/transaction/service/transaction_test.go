package service

import (
	"errors"
	"testing"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
	"github.com/stretchr/testify/require"
)

func TestTransactionService_PostTransaction(t *testing.T) {
	t.Run("When account does not exists, Then return Not Found error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewTransactionService(accountRepo)

		accountId := domain.AccountId(1)
		accountRepo.EXPECT().GetByIdAndLock(accountId).Once().Return(nil, nil)

		result, err := service.PostTransaction(int(accountId), dto.TransactionDto{})
		require.Error(t, err)
		require.ErrorIs(t, err, errs.NotFound(""))
		require.Empty(t, result)
	})

	t.Run("When error happens finding account, Then return the error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewTransactionService(accountRepo)

		accountId := domain.AccountId(1)
		dbErr := errors.New("database error")

		accountRepo.EXPECT().GetByIdAndLock(accountId).Once().Return(nil, dbErr)

		result, err := service.PostTransaction(int(accountId), dto.TransactionDto{})
		require.Error(t, err)
		require.Equal(t, err, dbErr)
		require.Empty(t, result)
	})

	t.Run("When find account but transaction operation exceed the account limit, Then return error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewTransactionService(accountRepo)

		account := &domain.Account{
			Id:      domain.AccountId(1),
			Limit:   0,
			Balance: 0,
		}

		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)

		transaction := dto.TransactionDto{
			AmountCents: 1,
			Type:        "d",
		}

		result, err := service.PostTransaction(int(account.Id), transaction)
		require.Error(t, err)
		require.ErrorIs(t, err, errs.UnprocessableEntity(""))
		require.Empty(t, result)
	})

	t.Run("When transaction operation succeed but update account fails, Then return error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewTransactionService(accountRepo)

		account := &domain.Account{
			Id:      domain.AccountId(1),
			Limit:   123,
			Balance: 0,
		}

		dbErr := errors.New("database error")
		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)
		accountRepo.EXPECT().Update(account).Once().Return(dbErr)

		transaction := dto.TransactionDto{
			AmountCents: 1,
			Type:        "c",
		}

		result, err := service.PostTransaction(int(account.Id), transaction)
		require.Error(t, err)
		require.Equal(t, dbErr, err)
		require.Empty(t, result)
	})

	t.Run("When transaction operation is valid, Then return the result", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)

		service := NewTransactionService(accountRepo)

		account := &domain.Account{
			Id:      domain.AccountId(1),
			Limit:   123,
			Balance: 0,
		}

		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)
		accountRepo.EXPECT().Update(account).Once().Return(nil)

		transaction := dto.TransactionDto{
			AmountCents: 1,
			Type:        "c",
		}

		result, err := service.PostTransaction(int(account.Id), transaction)
		require.NoError(t, err)

		expectedResult := dto.TransactionResult{
			LimitCents:   123,
			BalanceCents: 1,
		}
		require.Equal(t, expectedResult, result)
	})
}

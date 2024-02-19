package service

import (
	"errors"
	"testing"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTransactionService_PostTransaction(t *testing.T) {
	t.Run("When account does not exists, Then return Not Found error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)
		transactionRepo := repo.NewMockITransactionRepo(t)

		service := NewTransactionService(accountRepo, transactionRepo)

		accountId := domain.AccountId(1)
		accountRepo.EXPECT().GetByIdAndLock(accountId).Once().Return(nil, nil)

		result, err := service.PostTransaction(int(accountId), dto.TransactionDto{})
		require.Error(t, err)
		require.ErrorIs(t, err, errs.NotFound(""))
		require.Empty(t, result)
	})

	t.Run("When error happens finding account, Then return the error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)
		transactionRepo := repo.NewMockITransactionRepo(t)

		service := NewTransactionService(accountRepo, transactionRepo)

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
		transactionRepo := repo.NewMockITransactionRepo(t)

		service := NewTransactionService(accountRepo, transactionRepo)

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
		transactionRepo := repo.NewMockITransactionRepo(t)

		service := NewTransactionService(accountRepo, transactionRepo)

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

	t.Run("When error inserting the new transaction, Then return error", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)
		transactionRepo := repo.NewMockITransactionRepo(t)

		service := NewTransactionService(accountRepo, transactionRepo)

		account := &domain.Account{
			Id:      domain.AccountId(1),
			Limit:   123,
			Balance: 0,
		}

		dbError := errors.New("database error")
		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)
		accountRepo.EXPECT().Update(account).Once().Return(nil)
		transactionRepo.EXPECT().Insert(mock.Anything).Once().Return(dbError)

		transaction := dto.TransactionDto{
			AmountCents: 1,
			Type:        "c",
		}

		result, err := service.PostTransaction(int(account.Id), transaction)
		require.Error(t, err)
		require.Equal(t, dbError, err)
		require.Empty(t, result)
	})

	t.Run("When transaction operation is valid, Then return the result", func(t *testing.T) {
		accountRepo := repo.NewMockIAccountRepo(t)
		transactionRepo := repo.NewMockITransactionRepo(t)

		service := NewTransactionService(accountRepo, transactionRepo)

		account := &domain.Account{
			Id:      1,
			Limit:   123,
			Balance: 0,
		}

		transactionDto := dto.TransactionDto{
			AmountCents: 1,
			Type:        "c",
		}

		matchesTransactionData := func(inserted domain.Transaction) bool {
			return assert.Equal(t, domain.MoneyCents(transactionDto.AmountCents), inserted.Amount) &&
				assert.Equal(t, transactionDto.Description, inserted.Description) &&
				assert.Equal(t, transactionDto.Type, inserted.Type) &&
				assert.Equal(t, account.Id, inserted.AccountId) &&
				assert.NotEmpty(t, inserted.CreatedAt) &&
				assert.Empty(t, inserted.Id)
		}

		accountRepo.EXPECT().GetByIdAndLock(account.Id).Once().Return(account, nil)
		accountRepo.EXPECT().Update(account).Once().Return(nil)
		transactionRepo.EXPECT().Insert(mock.MatchedBy(matchesTransactionData)).Once().Return(nil)

		result, err := service.PostTransaction(int(account.Id), transactionDto)
		require.NoError(t, err)

		expectedResult := dto.TransactionResult{
			LimitCents:   123,
			BalanceCents: 1,
		}
		require.Equal(t, expectedResult, result)
	})
}

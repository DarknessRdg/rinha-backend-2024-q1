package service

import (
	"errors"
	"testing"

	account_dto "github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/dto"
	account_service "github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/service"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTransactionService_PostTransaction(t *testing.T) {
	t.Run("When error inserting the new transaction, Then return error", func(t *testing.T) {
		transactionRepo := repo.NewMockITransactionRepo(t)
		accountService := account_service.NewMockIAccountService(t)

		service := NewTransactionService(accountService, transactionRepo)

		accountId := 1

		transaction := dto.TransactionDto{
			AmountCents: 1,
			Type:        "c",
		}

		accountService.EXPECT().
			CreditOrDebit(accountId, transaction.AmountCents, transaction.Type).
			Once().
			Return(account_dto.AccountDto{}, nil)

		dbError := errors.New("database error")
		transactionRepo.EXPECT().Insert(mock.Anything).Once().Return(dbError)

		result, err := service.PostTransaction(accountId, transaction)
		require.Error(t, err)
		require.Equal(t, dbError, err)
		require.Empty(t, result)
	})

	t.Run("When error during account credit/debit, Then return error", func(t *testing.T) {
		transactionRepo := repo.NewMockITransactionRepo(t)
		accountService := account_service.NewMockIAccountService(t)

		service := NewTransactionService(accountService, transactionRepo)

		accountId := 1

		transaction := dto.TransactionDto{
			AmountCents: 1,
			Type:        "c",
		}

		dbError := errors.New("account service business error")
		accountService.EXPECT().
			CreditOrDebit(accountId, transaction.AmountCents, transaction.Type).
			Once().
			Return(account_dto.AccountDto{}, dbError)


		result, err := service.PostTransaction(accountId, transaction)
		require.Error(t, err)
		require.Equal(t, dbError, err)
		require.Empty(t, result)
	})

	t.Run("When transaction operation is valid, Then return the result", func(t *testing.T) {
		transactionRepo := repo.NewMockITransactionRepo(t)
		accountService := account_service.NewMockIAccountService(t)

		service := NewTransactionService(accountService, transactionRepo)

		transactionDto := dto.TransactionDto{
			AmountCents: 1,
			Type:        "c",
		}

		accountId := 1

		matchesTransactionData := func(inserted domain.Transaction) bool {
			return assert.Equal(t, domain.MoneyCents(transactionDto.AmountCents), inserted.Amount) &&
				assert.Equal(t, transactionDto.Description, inserted.Description) &&
				assert.Equal(t, domain.Operation(transactionDto.Type), inserted.Type) &&
				assert.Equal(t, domain.AccountId(accountId), inserted.AccountId) &&
				assert.NotEmpty(t, inserted.CreatedAt) &&
				assert.Empty(t, inserted.Id)
		}

		accountDto := account_dto.AccountDto{
			Id:      accountId,
			Limit:   999,
			Balance: 99,
		}

		accountService.EXPECT().
			CreditOrDebit(accountId, transactionDto.AmountCents, transactionDto.Type).
			Once().
			Return(accountDto, nil)

		transactionRepo.EXPECT().Insert(mock.MatchedBy(matchesTransactionData)).Once().Return(nil)

		result, err := service.PostTransaction(accountId, transactionDto)
		require.NoError(t, err)

		expectedResult := dto.TransactionResult{
			LimitCents:   accountDto.Limit,
			BalanceCents: accountDto.Balance,
		}
		require.Equal(t, expectedResult, result)
	})
}

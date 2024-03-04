package domain

import (
	"testing"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
	"github.com/stretchr/testify/require"
)

func TestAccount_Debit(t *testing.T) {

	suite := []struct {
		name           string
		account        *Account
		debitAmount    MoneyCents
		expectedAmount MoneyCents
		expectError    error
	}{
		{
			name: "When debit exceed the limit, Then return error",
			account: &Account{
				Balance: MoneyCents(-100),
				Limit:   101,
			},
			debitAmount:    MoneyCents(2),
			expectedAmount: MoneyCents(-100),
			expectError:    errs.UnprocessableEntity(""),
		},
		{
			name: "When debit does not exceed the limit, Then return return nil and decrement amount",
			account: &Account{
				Balance: MoneyCents(-90),
				Limit:   100,
			},
			debitAmount:    MoneyCents(1),
			expectedAmount: MoneyCents(-91),
		},
		{
			name: "When balance is positive and debit does not exceed the limit, Then return return nil and decrement amount",
			account: &Account{
				Balance: MoneyCents(110),
				Limit:   100,
			},
			debitAmount:    MoneyCents(100),
			expectedAmount: MoneyCents(10),
		},
	}

	for _, test := range suite {
		t.Run(test.name, func(t *testing.T) {
			err := test.account.Debit(test.debitAmount)

			if test.expectError != nil {
				require.ErrorIs(t, test.expectError, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, test.expectedAmount, test.account.Balance)
		})
	}
}

func TestAccount_Credit(t *testing.T) {

	t.Run("When add any value to the account, Then add the balance and return no error", func(t *testing.T) {
		account := Account{
			Balance: 1,
		}

		require.NoError(t, account.Credit(2))
		require.Equal(t, account.Balance, MoneyCents(3))
	})
}

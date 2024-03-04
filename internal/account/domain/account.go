package domain

import (
	"fmt"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
)

type AccountId int

type Account struct {
	Id      AccountId
	Balance MoneyCents
	Limit   int
}

func (a *Account) UpdateBalance(amount MoneyCents, operation Operation) error {
	switch operation {
	case OperationCredit:
		return a.Credit(amount)
	case OperationDebit:
		return a.Debit(amount)
	}
	return fmt.Errorf("invalid operation")
}

func (a *Account) Debit(debitAmount MoneyCents) error {
	newAmount := a.Balance - debitAmount

	if a.exceedLimit(newAmount) {
		return errs.UnprocessableEntity("Not enough limit")
	}

	a.Balance = newAmount
	return nil
}

func (a *Account) Credit(debitAmount MoneyCents) error {
	// since credit always add, it will never be "lower"
	// than the limit
	a.Balance += debitAmount
	return nil
}

func (a *Account) exceedLimit(amount MoneyCents) bool {
	// since we use negative balance as "amount debited", we need
	// to compare with negative limit as well
	lowestLimit := -a.Limit
	return amount.LowerThan(lowestLimit)
}

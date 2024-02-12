package domain

import "github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"

type AccountId string

type Account struct {
	Id      AccountId
	Balance MoneyCents
	Limit   int
}

func (a *Account) Debit(debitAmount MoneyCents) error {
	newAmount := a.Balance - debitAmount

	if a.exceedLimit(newAmount) {
		return errs.UnprocessableEntity("Not enough limit")
	}

	a.Balance = newAmount
	return nil
}

func (a *Account) exceedLimit(amount MoneyCents) bool {
	// since we use negative balance as "amount debited", we need
	// to compare with negative limit as well
	lowestLimit := -a.Limit
	return amount.LowerThan(lowestLimit)
}

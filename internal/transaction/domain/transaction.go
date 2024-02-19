package domain

import "time"

type TransactionId int

type Transaction struct {
	Id          TransactionId
	AccountId   AccountId
	Amount      MoneyCents
	Type        string
	Description string
	CreatedAt   time.Time
}

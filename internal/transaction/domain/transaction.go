package domain

import "time"

type TransactionId int
type AccountId int
type Operation string

const (
	CreditOperation Operation = "c"
	DebitOperation  Operation = "d"
)

type Transaction struct {
	Id          TransactionId
	AccountId   AccountId
	Amount      MoneyCents
	Type        Operation
	Description string
	CreatedAt   time.Time
}

package dto

type TransactionDto struct {
	AmountCents int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
}

type TransactionResult struct {
	LimitCents   int `json:"limite"`
	BalanceCents int `json:"saldo"`
}

package service

import (
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/dto"
)

type IAccountService interface {
	CreditOrDebit(accountId int, amountCents int, operation string) (dto.AccountDto, error)
}

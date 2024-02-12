package repo

import "github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"

type IAccountRepo interface {
	Update(account *domain.Account) error
	GetByIdAndLock(id domain.AccountId) (*domain.Account, error)
}

type ITransactionRepo interface {
	Insert() error
}

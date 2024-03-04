package repo

import "github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"

type ITransactionRepo interface {
	Insert(domain.Transaction) error
}

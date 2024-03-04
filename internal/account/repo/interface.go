package repo

import "github.com/DarknessRdg/rinha-backend-2024-q1/internal/account/domain"

type IAccountRepo interface {
	Update(account *domain.Account) error
	// GetByIdAndLock retrieves an account from the database based on the provided ID
	// and applies a lock on the selected account for exclusive access.
	// If no account  is found with the given ID, both a nil account and nil error will be returned.
	// In case of other errors, such as a connection error, an error will be returned.
	//
	// This function is designed to be used when exclusive access to the account is required
	// to perform atomic operations or prevent data inconsistencies.
	GetByIdAndLock(id domain.AccountId) (*domain.Account, error)
}

package sqlrepo

import (
	"database/sql"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
)

type SqlAccountRepo struct {
	Db *sql.DB
}

func (s *SqlAccountRepo) GetByIdAndLock(id domain.AccountId) (*domain.Account, error) {
	sttmt, err := s.Db.Prepare(`SELECT id, balance, limit FROM account WHERE id = ? FOR UPDATE`)
	if err != nil {
		return nil, err
	}
	defer sttmt.Close()

	row := sttmt.QueryRow(string(id))
	if err != nil {
		return nil, err
	}

	return s.rowToAccount(row)
}

func (s *SqlAccountRepo) Update(account *domain.Account) error {
	sttm, err := s.Db.Prepare("UPDATE account SET balance = ?, limit = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer sttm.Close()
	_, err = sttm.Exec(account.Balance, account.Limit, account.Id)
	return err
}

func (s *SqlAccountRepo) rowToAccount(row *sql.Row) (*domain.Account, error) {
	account := &domain.Account{}
	err := row.Scan(&account.Id, &account.Balance, &account.Limit)

	// NotFound aren't thread as error. Error should be used for unexpected scenarios
	// When item is not found, we should return `nil` as the element, instead.
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		// prevent to return account with trash data
		account = nil
	}

	return account, err
}

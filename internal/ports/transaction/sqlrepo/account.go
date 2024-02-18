package sqlrepo

import (
	"database/sql"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/domain"
)

type SqlAccountRepo struct {
	Db *sql.DB
}

func (s *SqlAccountRepo) GetByIdAndLock(id domain.AccountId) (*domain.Account, error) {
	sttmt, err := s.Db.Prepare(`SELECT * FROM account WHERE id = $1 FOR UPDATE`)
	if err != nil {
		return nil, err
	}
	defer sttmt.Close()

	row := sttmt.QueryRow(int(id))
	if err != nil {
		return nil, err
	}

	return s.rowToAccount(row)
}

func (s *SqlAccountRepo) Update(account *domain.Account) error {
	sttm, err := s.Db.Prepare(`UPDATE account SET balance = $1, "limit" = $2 WHERE id = $3`)
	if err != nil {
		return err
	}
	defer sttm.Close()
	_, err = sttm.Exec(account.Balance, account.Limit, account.Id)
	return err
}

func (s *SqlAccountRepo) rowToAccount(row *sql.Row) (*domain.Account, error) {
	account := &domain.Account{}
	var id int
	var balance int

	err := row.Scan(&id, &balance, &account.Limit)
	account.Id = domain.AccountId(id)
	account.Balance = domain.MoneyCents(balance)

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

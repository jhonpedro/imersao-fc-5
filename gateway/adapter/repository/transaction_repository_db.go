package repository

import (
	"database/sql"
	"time"
)

type TransactionRepositoryDb struct {
	db *sql.DB
}

func NewTransactionRepositoryDb(db *sql.DB) *TransactionRepositoryDb {
	return &TransactionRepositoryDb{db: db}
}

func (t *TransactionRepositoryDb) Insert(evaluation_id string, id string, accountId string, amount float64, status string, errorMessage string) error {
	stmt, err := t.db.Prepare(`
		insert into transactions (evaluation_id, id, account_id, amount, status, error_message, created_at, updated_at)
		values($1, $2, $3, $4, $5, $6, $7, $8)
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		evaluation_id,
		id,
		accountId,
		amount,
		status,
		errorMessage,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

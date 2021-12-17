package sqlite

import (
	"database/sql"
	"time"
)

type TransactionSqliteRepository struct {
	db *sql.DB
}

func NewTransactionSqliteRepository(db *sql.DB) *TransactionSqliteRepository {
	return &TransactionSqliteRepository{db: db}
}

func (t *TransactionSqliteRepository) Save(id string, account string, amount float64, status string, errorMessage string) error {
	stmt, err := t.db.Prepare(`
		INSERT INTO transactions (id, account_id, amount, status, error_message, created_at, updated_at)
		VALUES (:id, :account_id, :amount, :status, :error_message, :created_at, :updated_at);
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, account, amount, status, errorMessage, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

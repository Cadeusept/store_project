package repository

import (
	"errors"
	"fmt"
	"store-project/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	statusNew       = "NEW"
	statusError     = "ERROR"
	statusCancelled = "CANCELLED"
	statusSuccess   = "SUCCESS"
	statusFailed    = "FAILED"
)

type TransactionsPostgres struct {
	db *sqlx.DB
}

func NewTransactionsPostgres(db *sqlx.DB) *TransactionsPostgres {
	return &TransactionsPostgres{db: db}
}

func (r *TransactionsPostgres) Create(t models.Transaction) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, user_email, amount, currency, created, changed, stat) values ($1, $2, $3, $4, $5, $6, $7) RETURNING id", transactionsTable)
	row := r.db.QueryRow(query, t.UserId, t.UserEmail, t.Amount, t.Currency, time.Now(), time.Now(), statusNew)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TransactionsPostgres) ChangeStatus(id int64, status string) error {
	var condition string
	switch {
	case status == statusError:
		condition = "tt.stat='" + statusNew + "'"
	case status == statusFailed:
		condition = "tt.stat='" + statusNew + "' OR tt.stat='" + statusError + "'"
	case status == statusSuccess:
		condition = "tt.stat='" + statusNew + "'"
	case status == statusCancelled:
		return errors.New("to cancel transaction use specialized method")
	default:
		return errors.New("wrong status")
	}

	query := fmt.Sprintf("UPDATE %s tt SET changed=$1, stat=$2 WHERE tt.id=$3 AND (%s)",
		transactionsTable, condition)

	_, err := r.db.Exec(query, time.Now(), status, id)

	return err
}

func (r *TransactionsPostgres) CancelTransactionById(id int64) error {
	condition := "tt.stat='" + statusNew + "' OR tt.stat='" + statusError + "'"

	query := fmt.Sprintf("UPDATE %s tt SET changed=$1, stat=$2 WHERE tt.id=$3 AND (%s)",
		transactionsTable, condition)

	_, err := r.db.Exec(query, time.Now(), statusCancelled, id)
	return err
}

func (r *TransactionsPostgres) CheckStatusById(id int64) (string, error) {
	var status string
	query := fmt.Sprintf("SELECT tt.stat FROM %s tt WHERE id=$1",
		transactionsTable)
	if err := r.db.Get(&status, query, id); err != nil {
		return status, err
	}

	return status, nil
}

func (r *TransactionsPostgres) GetTransactionsByUserId(uId int64) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := fmt.Sprintf("SELECT * FROM %s tt WHERE tt.user_id=$1", transactionsTable)
	if err := r.db.Select(&transactions, query, uId); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionsPostgres) GetTransactionsByUserEmail(email string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := fmt.Sprintf("SELECT * FROM %s tt WHERE tt.user_email=$1", transactionsTable)
	if err := r.db.Select(&transactions, query, email); err != nil {
		return nil, err
	}

	return transactions, nil
}

package repository

import (
	"store-project/internal/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepo interface {
	Create(models.Transaction) (int, error)
	ChangeStatus(id int64, status string) error
	CheckStatusById(id int64) (string, error)
	GetTransactionsByUserId(uId int64) ([]models.Transaction, error)
	GetTransactionsByUserEmail(email string) ([]models.Transaction, error)
	CancelTransactionById(id int64) error
}

type Repository struct {
	TransactionRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TransactionRepo: NewTransactionsPostgres(db),
	}
}

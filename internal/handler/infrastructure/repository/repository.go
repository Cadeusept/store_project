package repository

import (
	"store-project/internal/models"

	"github.com/jmoiron/sqlx"
)

type TransactionRepoI interface {
	Create(models.Transaction) (int, error)
	ChangeStatus(int64, string) error
	CheckStatusById(int64) (string, error)
	GetTransactionsByUserId(int64) ([]models.Transaction, error)
	GetTransactionsByUserEmail(string) ([]models.Transaction, error)
	CancelTransactionById(int64) error
}

type Repository struct {
	TransactionRepoI
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TransactionRepoI: NewTransactionsPostgres(db),
	}
}

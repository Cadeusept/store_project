package usecase

import (
	"store-project/internal/handler/infrastructure/repository"
	"store-project/internal/models"
)

type HandlerUseCaseI interface {
	Create(models.Transaction) (int, error)
	ChangeStatus(id int64, status string) error
	CheckStatusById(id int64) (string, error)
	GetTransactionsByUserId(uId int64) ([]models.Transaction, error)
	GetTransactionsByUserEmail(email string) ([]models.Transaction, error)
	CancelTransactionById(id int64) error
}

type handlerUC struct {
	transactionRepo repository.TransactionRepo
}

func newHandlerUC(repos *repository.Repository) *handlerUC {
	return &handlerUC{
		transactionRepo: repos.TransactionRepo,
	}
}

package usecase

import (
	"store-project/internal/handler/infrastructure/repository"
	"store-project/internal/models"
)

type HandlerUseCaseI interface {
	Create(models.Transaction) (int, error)
	ChangeStatus(int64, string) error
	CheckStatusById(int64) (string, error)
	GetTransactionsByUserId(int64) ([]models.Transaction, error)
	GetTransactionsByUserEmail(string) ([]models.Transaction, error)
	CancelTransactionById(int64) error
}

type HandlerUC struct {
	transactionRepo repository.TransactionRepoI
}

func NewHandlerUC(repos *repository.Repository) *HandlerUC {
	return &HandlerUC{
		transactionRepo: repos.TransactionRepoI,
	}
}

func (uc *HandlerUC) Create(t models.Transaction) (int, error) {
	return uc.transactionRepo.Create(t)
}

func (uc *HandlerUC) ChangeStatus(id int64, status string) error {
	return uc.transactionRepo.ChangeStatus(id, status)
}

func (uc *HandlerUC) CheckStatusById(id int64) (string, error) {
	return uc.transactionRepo.CheckStatusById(id)
}

func (uc *HandlerUC) GetTransactionsByUserId(uId int64) ([]models.Transaction, error) {
	return uc.transactionRepo.GetTransactionsByUserId(uId)
}

func (uc *HandlerUC) GetTransactionsByUserEmail(email string) ([]models.Transaction, error) {
	return uc.transactionRepo.GetTransactionsByUserEmail(email)
}

func (uc *HandlerUC) CancelTransactionById(id int64) error {
	return uc.transactionRepo.CancelTransactionById(id)
}

package auth

import "store-project/internal/models"

type UseCaseI interface {
	CreateSession(uID uint64) (*models.Session, error)
	DeleteSession(sessionID string) error
	GetSession(sessionID string) (*models.Session, error)
}

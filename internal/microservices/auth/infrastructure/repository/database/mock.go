package _authRepo

import (
	"fmt"
	"store-project/internal/microservices/auth/infrastructure/repository"
	"store-project/internal/models"
)

type sessionsDB struct {
	empty bool
}

func NewSessionRepo() repository.SessionRepositoryI {
	return &sessionsDB{
		empty: true,
	}
}

func (sDb *sessionsDB) CreateSession(uID uint64) (*models.Session, error) {
	fmt.Println("authRepo CreateSession called")
	return &models.Session{
		UID:       0,
		SessionID: "",
	}, nil
}

func (sDb *sessionsDB) DeleteSession(sessionID string) error {
	fmt.Println("authRepo DeleteSession called")
	return nil
}

func (sDb *sessionsDB) GetSession(sessionID string) (*models.Session, error) {
	fmt.Println("authRepo GetSession called")
	return &models.Session{
		UID:       0,
		SessionID: "",
	}, nil
}

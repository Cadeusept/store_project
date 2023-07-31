package usecase

import (
	"fmt"
	"store-project/internal/microservices/auth/infrastructure/repository"
	"store-project/internal/models"
)

type authUC struct {
	sessionRepo repository.SessionRepositoryI
}

func NewAuthUC(r repository.SessionRepositoryI) *authUC {
	return &authUC{
		sessionRepo: r,
	}
}

func (u *authUC) CreateSession(uID uint64) (*models.Session, error) {
	newSession, err := u.sessionRepo.CreateSession(uID)
	if err != nil {
		return nil, fmt.Errorf("%w create session", err) // pkgErrors.Wrap(err, "create session")
	}

	return newSession, nil
}

func (u *authUC) DeleteSession(sessionID string) error {
	err := u.sessionRepo.DeleteSession(sessionID)
	if err != nil {
		return fmt.Errorf("%w delete session", err) // pkgErrors.Wrap(err, "delete avatar")
	}

	return nil
}

func (u *authUC) GetSession(sessionID string) (*models.Session, error) {
	s, err := u.sessionRepo.GetSession(sessionID)
	if err != nil {
		return nil, fmt.Errorf("get session: %w", err)
	}

	return s, nil
}

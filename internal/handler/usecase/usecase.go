package usecase

import (
	"fmt"
	authClient "store-project/internal/microservices/auth/client/grpc"
	"store-project/internal/models"

	"google.golang.org/grpc"
)

type handlerUC struct {
	authClient *authClient.AuthClientGRPC
}

func newHandlerUC(cc *grpc.ClientConn) *handlerUC {
	return &handlerUC{
		authClient: authClient.NewAuthClientGRPC(cc).(*authClient.AuthClientGRPC),
	}
}

func (u *handlerUC) CreateSession(uID uint64) (*models.Session, error) {
	newSession, err := u.authClient.CreateSession(uID)
	if err != nil {
		return nil, fmt.Errorf("%w create session", err) // pkgErrors.Wrap(err, "create session")
	}

	return newSession, nil
}

func (u *handlerUC) DeleteSession(sessionID string) error {
	err := u.authClient.DeleteSession(sessionID)
	if err != nil {
		return fmt.Errorf("%w delete session", err) // pkgErrors.Wrap(err, "delete avatar")
	}

	return nil
}

func (u *handlerUC) GetSession(sessionID string) (*models.Session, error) {
	s, err := u.authClient.GetSession(sessionID)
	if err != nil {
		return nil, fmt.Errorf("get session: %w", err)
	}

	return s, nil
}

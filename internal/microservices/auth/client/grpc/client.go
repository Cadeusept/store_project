package client

import (
	"context"
	"fmt"
	"store-project/internal/microservices/auth"
	auth_proto "store-project/internal/microservices/auth/proto"
	"store-project/internal/microservices/auth/utils"
	"store-project/internal/models"

	"google.golang.org/grpc"
)

type AuthClientGRPC struct {
	authClient auth_proto.AuthServiceClient
}

func NewAuthClientGRPC(cc *grpc.ClientConn) auth.UseCaseI {
	return &AuthClientGRPC{
		authClient: auth_proto.NewAuthServiceClient(cc),
	}
}

func (g AuthClientGRPC) CreateSession(uID uint64) (*models.Session, error) {
	protoSession, err := g.authClient.CreateSession(context.TODO(), &auth_proto.UID{UID: uID})
	if err != nil {
		return nil, fmt.Errorf("auth client CreateSession error: %w", err)
	}

	return utils.SessionModelByProto(protoSession), nil
}

func (g AuthClientGRPC) DeleteSession(sessionID string) error {
	_, err := g.authClient.DeleteSession(context.TODO(), &auth_proto.SessionId{Value: sessionID})
	if err != nil {
		return fmt.Errorf("auth client DeleteSession error: %w", err)
	}

	return nil
}

func (g AuthClientGRPC) GetSession(sessionID string) (*models.Session, error) {
	protoSession, err := g.authClient.GetSession(context.TODO(), &auth_proto.SessionId{Value: sessionID})
	if err != nil {
		return nil, fmt.Errorf("auth client GetSession error: %w", err)
	}

	return utils.SessionModelByProto(protoSession), nil
}

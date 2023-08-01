package controller

import (
	"context"
	"fmt"
	"net"
	"store-project/internal/microservices/auth"
	auth_proto "store-project/internal/microservices/auth/proto"
	"store-project/internal/microservices/auth/utils"

	"google.golang.org/grpc"
)

type GrpcController struct {
	auth_proto.UnimplementedAuthServiceServer

	grpcServer *grpc.Server
	AuthUC     auth.UseCaseI
}

func NewGrpcController(gs *grpc.Server, uc auth.UseCaseI) *GrpcController {
	return &GrpcController{
		grpcServer: gs,
		AuthUC:     uc,
	}
}

func (g *GrpcController) Start(url string) error {
	lis, err := net.Listen("tcp", url)
	if err != nil {
		return err
	}
	auth_proto.RegisterAuthServiceServer(g.grpcServer, g)
	return g.grpcServer.Serve(lis)
}

func (g *GrpcController) CreateSession(ctx context.Context, protoUID *auth_proto.UID) (*auth_proto.Session, error) {
	session, err := g.AuthUC.CreateSession(protoUID.UID)

	if err != nil {
		return nil, fmt.Errorf("auth server CreateSession error: %w", err)
	}
	return utils.ProtoBySessionModel(session), nil
}

func (g *GrpcController) DeleteSession(ctx context.Context, protoSID *auth_proto.SessionId) (*auth_proto.Nothing, error) {
	err := g.AuthUC.DeleteSession(protoSID.Value)

	if err != nil {
		return nil, fmt.Errorf("auth server DeleteSession error: %w", err)
	}

	return &auth_proto.Nothing{}, nil
}

func (g *GrpcController) GetSession(ctx context.Context, protoSID *auth_proto.SessionId) (*auth_proto.Session, error) {
	session, err := g.AuthUC.GetSession(protoSID.Value)

	if err != nil {
		return nil, fmt.Errorf("auth server error: %w", err)
	}

	return utils.ProtoBySessionModel(session), nil
}

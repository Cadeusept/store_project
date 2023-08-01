package app

import (
	controller "store-project/internal/microservices/auth/controller/grpc"
	_authRepo "store-project/internal/microservices/auth/infrastructure/repository/database"
	"store-project/internal/microservices/auth/usecase"
)

func main() {
	sessionRepo := _authRepo.NewSessionRepo()

	authUC := usecase.NewAuthUC(sessionRepo)

	serverGRPC := controller.NewGrpcController()

	// TODO: доделать
}

package app

import (
	"log"
	controller "store-project/internal/microservices/auth/controller/grpc"
	_authRepo "store-project/internal/microservices/auth/infrastructure/repository/database"
	"store-project/internal/microservices/auth/usecase"

	"google.golang.org/grpc"
)

func main() {
	sessionRepo := _authRepo.NewSessionRepo()

	/*
		userServiceCon, err := grpc.Dial(
			"http://127.0.0.1"+":"+"8002",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatal("failed connect to file microservice", err)
		}

		userServiceClient := client.NewAuthClientGRPC(userServiceCon)
	*/

	authUC := usecase.NewAuthUC(sessionRepo) // userServiceClient, sessionRepo)

	grpcServer := grpc.NewServer()

	controllerGRPC := controller.NewGrpcController(grpcServer, authUC)

	err := controllerGRPC.Start("http://127.0.0.1:8001")
	if err != nil {
		log.Fatal(err)
	}
}

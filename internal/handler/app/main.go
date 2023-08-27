package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	controllerhttp "store-project/internal/handler/controller/http"
	"store-project/internal/handler/httpserver"
	"store-project/internal/handler/infrastructure/repository"
	"store-project/internal/handler/usecase"
	"syscall"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.SetFormatter(new(log.JSONFormatter))

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("error connecting to database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	uc := usecase.NewHandlerUC(repos)
	httpController := controllerhttp.NewHandler(uc)

	webSrv := new(httpserver.HandlerServer)
	go func() {
		if err := webSrv.Run(viper.GetString("handler.port"), httpController.InitRoutes()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error running web server: %s", err.Error())
		}
	}()

	log.Print("transactions app started")

	c_quit := make(chan os.Signal, 1)
	signal.Notify(c_quit, syscall.SIGTERM, syscall.SIGINT)
	sig := <-c_quit

	log.Printf("catched signal: %s. Notes app shutting down...", sig.String())

	if err := webSrv.Shutdown(context.Background()); err != nil {
		log.Errorf("error occured during shutdown: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Errorf("error occured during database connection closure: %s", err.Error())
	}

	log.Print("transactions app shut down successfully")
}

func initConfig() error {
	viper.AddConfigPath("internal/handler/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

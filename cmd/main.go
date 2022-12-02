package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asadbek21coder/eleanshop"
	"github.com/asadbek21coder/eleanshop/pkg/handler"
	"github.com/asadbek21coder/eleanshop/pkg/repository"
	"github.com/asadbek21coder/eleanshop/pkg/service"
	"github.com/joho/godotenv"

	_ "github.com/asadbek21coder/eleanshop/docs"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title       Eleanshop
// @version     1.0
// @description API server for eleanshop website.

// @host     localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in                         header
// @name                       Authorization

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		fmt.Println("this")
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(eleanshop.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())

	}
	// logrus.Print("TodoApp Started")

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// <-quit

	// logrus.Print("TodoApp Shutting Down")

	// if err := srv.Shutdown(context.Background()); err != nil {
	// 	logrus.Errorf("error occured on server shutting down: %s", err.Error())
	// }

	// if err := db.Close(); err != nil {
	// 	logrus.Errorf("error occured on db connection close: %s", err.Error())
	// }

	// fmt.Println("Bismillah")
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

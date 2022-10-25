package main

import (
	"L3WB"
	"L3WB/pkg/handler"
	"L3WB/pkg/repository"
	"L3WB/pkg/service"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           Weather forecast
// @version         1.0
// @description     This is a simple api that shows a weather forecast.

// @host      localhost:8000
// @BasePath  /

func main() {
	//Init config file
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//Init env file
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//Init db
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	//Dependency injection
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	/* Main part */
	services.Geocoding.GetAllStartData()
	/* Update temperature info for every city every minute */
	go services.Geocoding.BackgroundUpdatingProcess()

	//Run server
	srv := new(L3WB.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while server running: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

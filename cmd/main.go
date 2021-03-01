package main

import (
	"github.com/spf13/viper"
	"log"
	"todo"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/service"
)

func main(){

	//logrus.SetFormatter(new(logrus.JSONFormatter))
	//if err := initConfig(); err!= nil{
	//	log.Fatalf("error initializing configs: %s", err.Error())
	//}

	db, err := repository.NewClient("mongodb://127.0.0.1:27017")
	if err != nil{
		log.Fatalf("Mongo db error: %s",err.Error())
	}

	client := db.Database("testdb")


	repos := repository.NewRepository(client)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	//handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil{
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}


func initConfig() error{
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
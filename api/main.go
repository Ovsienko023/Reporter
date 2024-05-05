package main

import (
	"flag"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/Ovsienko023/reporter/server"
	"github.com/joho/godotenv"
	"log"
)

// @title Reporter API
// @version 0.0.1
// @description This is a report server.

// @BasePath /api/v1
func main() {
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Println("Not found locale .env file", err.Error())
	}

	cfg, err := configuration.New() // TODO: Добавить полечение пути к конфигу через flag
	if err != nil {
		log.Fatalf("Could not read configuration file with error: %+v", err)
	}

	// TODO: Добавить инициализацию логера

	log.Printf("-- Running on: %s:%s \n", cfg.Api.Host, cfg.Api.Port)

	app := server.NewApp(cfg)
	if err := app.Run(cfg); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

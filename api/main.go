package main

import (
	"embed"
	"flag"
	"fmt"
	"log"

	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/Ovsienko023/reporter/server"
)

//go:embed store_fs/*
var embedFS embed.FS

// @title Reporter API
// @version 0.0.1
// @description This is a report server.

// @BasePath /api/v1
func main() {
	flag.Parse()

	cfg, err := configuration.NewConfig()

	if err != nil {
		log.Fatalf("Could not read configuration file with error: %+v", err)
	}

	fmt.Printf("Running on: %s:%s \n", cfg.Api.Host, cfg.Api.Port)

	app := server.NewApp(cfg, embedFS)
	if err := app.Run(&cfg.Api); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/repository"
	transportHttp "github.com/Ovsienko023/reporter/app/transport/http"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/Ovsienko023/reporter/server/static"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	httpServer *http.Server

	recordCore *core.Core
}

func NewApp(cnf *configuration.Config) *App {
	db, _ := repository.New(&cnf.Db)
	recordCore := core.NewCore(*cnf, db)

	return &App{
		recordCore: recordCore,
	}
}

func (a *App) Run(apiConfig *configuration.Api) error {
	router := chi.NewRouter()

	staticServer := static.New(apiConfig.Static.FilesPath)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(staticServer.Handler)

	r := transportHttp.RegisterHTTPEndpoints(router, *a.recordCore, apiConfig)

	a.httpServer = &http.Server{
		Addr:           apiConfig.Host + ":" + apiConfig.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if apiConfig.Tls.Enable {
		a.startTls(apiConfig)
	} else {
		a.startWithoutTls()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func (a *App) startTls(cfg *configuration.Api) {
	go func() {
		if err := a.httpServer.ListenAndServeTLS(
			cfg.Tls.CertFilePath,
			cfg.Tls.KeyFilePath,
		); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
}

func (a *App) startWithoutTls() {
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()
}

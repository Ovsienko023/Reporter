package server

import (
	"context"
	"github.com/Ovsienko023/reporter/pkg/report"
	"github.com/Ovsienko023/reporter/pkg/report/repository/localstore"
	reporthttp "github.com/Ovsienko023/reporter/pkg/report/transport/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	recordCore report.Core
}

func NewApp() *App {
	//db := initDB()
	recordRepo := localstore.NewReportLocalStorage()
	//bookmarkRepo := bmmongo.NewBookmarkRepository(db, viper.GetString("mongo.bookmark_collection"))

	return &App{
		recordCore: recordRepo,
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	//router := gin.Default()
	//router.Use(
	//	gin.Recovery(),
	//	gin.Logger(),
	//)
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	r := reporthttp.RegisterHTTPEndpoints(router, a.recordCore)

	// Set up http handlers
	// SignUp/SignIn endpoints
	//authhttp.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	//authMiddleware := authhttp.NewAuthMiddleware(a.authUC)
	//api := router.Group("/api", authMiddleware)

	//bmhttp.RegisterHTTPEndpoints(api, a.bookmarkUC)
	//r := app.SetupRoutes(router)
	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

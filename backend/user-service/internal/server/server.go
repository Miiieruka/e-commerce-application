package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-service/config"
	"user-service/internal/server/handler"
	"user-service/internal/service/domain"
	"user-service/internal/storage/postgre"
)

type APIserver struct {
	cfg *config.Config
}

func NewAPIserver(cfg *config.Config) *APIserver {
	return &APIserver{
		cfg: cfg,
	}
}

func (srv *APIserver) run() {
	db := postgre.ConnectPostgre(srv.cfg.DB)
	repo := postgre.NewPostgreRepo(db)
	serv := domain.NewService(repo)
	handler := handler.NewHandler(serv)

	serverAddress := srv.cfg.HttpServer.Address
	serverPort := fmt.Sprintf(":%s", srv.cfg.HttpServer.Port)
	fmt.Printf("Running server on %s", serverAddress)

	httpServer := &http.Server{
		Addr:    serverAddress + serverPort,
		Handler: handler.InitRoutes(),
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Running server error: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Printf("Shutdown signal detected...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err.Error())
	}
	fmt.Printf("Server gracefully shut down")
}

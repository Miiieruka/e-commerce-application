package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-service/config"
	"product-service/internal/server/handler"
	"product-service/internal/service"
	"product-service/internal/storage/storage"
	"syscall"
	"time"
)

type APIserver struct {
	cfg *config.Config
}

func NewApiServer(cfg *config.Config) *APIserver {
	return &APIserver{
		cfg: cfg,
	}
}

func (serv *APIserver) Run() {
	db := storage.ConnectPostgre(serv.cfg.DB)
	redis := storage.ConnectRedis(serv.cfg.Redis)

	repo := storage.NewRepository(db, redis)

	cld := InitCloudinary(serv.cfg.CloudyConfig)
	service := service.NewService(repo, cld)

	handler := handler.NewHandler(service, serv.cfg.JwtSecret)

	serverAddr := serv.cfg.HttpServer.Address
	port := fmt.Sprintf(":%s", serv.cfg.HttpServer.Port)
	httpServer := &http.Server{
		Addr:    serverAddr + port,
		Handler: handler.InitRoutes(),
	}
	fmt.Printf("Running server on %s\n", serverAddr+port)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Running server error: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Printf("Shutdown signal...\n")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err.Error())
	}
	fmt.Printf("Server gracefully shutdown")
}

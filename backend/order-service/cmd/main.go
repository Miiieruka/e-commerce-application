package main

import (
	"context"
	"database/sql"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"order-service/config"
	grpcclient "order-service/internal/adapters/grpc"
	"order-service/internal/adapters/handlers"
	"order-service/internal/adapters/storage/postgres"
	"order-service/internal/adapters/storage/postgres/repository"
	"order-service/internal/core/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.NewConfig()
	db := postgres.ConnectPostgres(cfg.DB)
	paymentAddr := cfg.PayServiceAddr
	// Initialize gRPC client for payment service
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic("failed to close database connection: " + err.Error())
		} else {
			println("Database connection closed successfully")
		}
	}(db)

	conn, err := grpc.NewClient(paymentAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("failed to connect to payment service: " + err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic("failed to close gRPC connection: " + err.Error())
		}
	}(conn)

	paymentService := grpcclient.NewPaymentGRPCClient(conn)
	repo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(repo, paymentService)
	orderHandler := handlers.NewOrderHandler(orderService)

	router := handlers.NewRouter(orderHandler, cfg.JwtSecret)
	srvAddr := cfg.HttpServer.Address + ":" + cfg.HttpServer.Port
	httpServer := &http.Server{
		Addr:    srvAddr,
		Handler: router,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("failed to start HTTP server: " + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		panic("failed to shutdown HTTP server: " + err.Error())
	}
	println("Server gracefully stopped")

}

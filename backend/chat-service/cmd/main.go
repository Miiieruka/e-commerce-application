package main

import (
	"chat-service/config"
	"chat-service/internal/adapters/handlers"
	"chat-service/internal/adapters/storage"
	"chat-service/internal/adapters/storage/repository"
	"chat-service/internal/core/services"
	"chat-service/logger/handlers/slogpretty"
	"context"
	"errors"
	godotenv "github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	cfg := config.MustLoad()
	logger := setupPrettySlog()
	logger.Info("Chat Service started", slog.Any("config", cfg))
	db := storage.ConnectDB(cfg.DB)
	redisClient := storage.ConnectRedis(cfg.REDIS)
	logger.Info("Postgre and Redis connected", slog.Any("Postgre", db.Stats()), slog.Any("Redis", redisClient.String()))
	address := cfg.Server.Addr + ":" + strconv.Itoa(cfg.Server.Port)
	logger.Info("Server is running", slog.String("address", address))

	msgRepo := repository.NewMessageRepository(db)
	roomRepo := repository.NewRoomRepository(db)
	upr := storage.NewUserPresenceRepository(redisClient)

	msgSrv := services.NewMessageService(msgRepo)
	roomSrv := services.NewRoomService(roomRepo)

	chatHandler := handlers.NewChatHandler(msgSrv, roomSrv, upr, logger, cfg.JWTSecret)
	router := handlers.NewRouter(chatHandler)

	server := http.Server{
		Addr:    address,
		Handler: router,
	}
	logger.Info("Server is ready to accept requests", slog.String("address", address))
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("failed to start HTTP server: " + err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic("failed to shutdown HTTP server: " + err.Error())
	}

	logger.Info("Server gracefully shutdown", slog.String("address", address))
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}

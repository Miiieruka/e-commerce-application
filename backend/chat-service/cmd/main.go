package main

import (
	"chat-service/config"
	"chat-service/internal/adapters/storage"
	"chat-service/logger/handlers/slogpretty"
	godotenv "github.com/joho/godotenv"
	"log/slog"
	"os"
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

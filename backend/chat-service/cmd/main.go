package main

import (
	"chat-service/config"
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

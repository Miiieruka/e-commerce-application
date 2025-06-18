package main

import (
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"payment-service/app"
	"payment-service/config"
	"payment-service/internal/lib/logger/handlers/slogpretty"
	"syscall"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	cfg := config.MustLoad()
	logger := setupPrettySlog()
	logger.Info("starting up", slog.Any("config", cfg))

	application := app.NewGRPCServer(logger, cfg.GRPCServer.Port)
	go application.GRPCServer.MustStart()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	application.GRPCServer.Stop()
	logger.Info("server stopped")
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

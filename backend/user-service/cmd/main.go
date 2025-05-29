package main

import (
	"user-service/config"
	"user-service/internal/server"
)

func main() {
	cfg := config.NewConfig()
	srv := server.NewAPIserver(cfg)
	srv.Run()
}

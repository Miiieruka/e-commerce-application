package main

import (
	"product-service/config"
	"product-service/internal/server"
)

func main() {
	cfg := config.NewConfig()
	srv := server.NewApiServer(cfg)
	srv.Run()
}

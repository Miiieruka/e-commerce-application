package main

import (
	"auth-service/internal/http-server/handler"
	"auth-service/internal/http-server/repository"
	"auth-service/internal/http-server/service"
	"auth-service/internal/initializer"
	"log"
	"net/http"
)

func init() {
	initializer.LoadEnv()
}

func main() {

	config := initializer.LoadConfig()

	db, err := initializer.ConnectDB(*config)

	defer initializer.CloseConnection(db)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	rep := repository.NewUserRepository(db)
	serv := service.NewAuthService(rep, config.JwtSecret)
	service.InitGoogleOAuth(config.GoogleOAuth)
	handler := handler.NewAuthHandler(serv)

	router := handler.SetupRoutes()
	log.Printf("Server running at %s", config.HTTPServer.Address)
	if err := http.ListenAndServe(config.HTTPServer.Address, router); err != nil {
		log.Fatalf("Server error %v", err)
	}

}

package handler

import (
	"auth-service/internal/http-server/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	srv *service.AuthService
}

func NewAuthHandler(authservice *service.AuthService) *AuthHandler {
	return &AuthHandler{srv: authservice}
}

func (handler *AuthHandler) SetupRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
	protected := router.Group("/auth")
	protected.Use()


	return router
}

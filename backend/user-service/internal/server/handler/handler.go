package handler

import (
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		serv: serv,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(JwtMiddleware())

	router.Group("/user")
	{
		router.GET("/me")
	}

	return router
}

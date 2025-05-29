package handler

import (
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	serv      *service.Service
	jwtsecret string
}

func NewHandler(serv *service.Service, jwtsecret string) *Handler {
	return &Handler{
		serv:      serv,
		jwtsecret: jwtsecret,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(h.JwtMiddleware())

	router.GET("/user/me", h.GetProfile)

	return router
}

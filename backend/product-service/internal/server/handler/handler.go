package handler

import (
	"product-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services  *service.Service
	jwtsecret string
}

func NewHandler(srv *service.Service, jwtsecret string) *Handler {
	return &Handler{
		services:  srv,
		jwtsecret: jwtsecret,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	protected := router.Group("/products", h.JwtMiddleware())
	protected.POST("/add", h.SaveProduct)
	protected.PUT("/:id", h.UpdateProduct)
	protected.DELETE("/:id", h.DeleteProduct)

	router.GET("/products", h.GetAllProducts)
	router.GET("/products/:id", h.GetProductById)

	return router
}

package handlers

import "github.com/gin-gonic/gin"

func NewRouter(orderHandler *OrderHandler, jwtsecret string) *gin.Engine {
	router := gin.Default()

	// Order routes
	orderRoutes := router.Group("/orders", JwtMiddleware(jwtsecret))
	{
		orderRoutes.POST("/", orderHandler.CreateOrder)
		orderRoutes.GET("/history", orderHandler.GetOrderHistory)
	}

	return router
}

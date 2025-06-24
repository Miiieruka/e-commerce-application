package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *ChatHandler) *gin.Engine {
	router := gin.Default()

	chatRoutes := router.Group("/chat", JwtMiddleware(handler.jwtSecret))
	{
		chatRoutes.POST("/rooms", handler.CreateRoom)
		chatRoutes.GET("/rooms/messages/:room_id", handler.GetRoomMessages)
		chatRoutes.GET("/rooms/:room_id", handler.GetRoomByID)
		chatRoutes.POST("/message", handler.SendMessage)
		chatRoutes.GET("/rooms", handler.GetUserRooms)
	}

	return router
}

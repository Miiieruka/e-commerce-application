package handlers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *ChatHandler) *gin.Engine {
	router := gin.Default()

	return router
}

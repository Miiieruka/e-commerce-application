package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(c *gin.Context) {
	id := c.GetInt64("user_id")
	fmt.Print(id)
	ctxWithDeadline, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	user, err := h.serv.Serv.GetUserByID(ctxWithDeadline, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Printf("HERE")
	c.IndentedJSON(http.StatusOK, user)
}

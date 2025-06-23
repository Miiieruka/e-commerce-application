package handlers

import (
	"chat-service/internal/core/domain"
	"chat-service/internal/core/ports/repositories"
	"chat-service/internal/core/ports/services"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

type ChatHandler struct {
	ms     services.MessageService
	rs     services.RoomService
	upr    repositories.UserPresenceRepository
	logger *slog.Logger
}

func NewChatHandler(
	ms services.MessageService,
	rs services.RoomService,
	upr repositories.UserPresenceRepository,
	logger *slog.Logger,
) *ChatHandler {
	return &ChatHandler{
		ms:     ms,
		rs:     rs,
		upr:    upr,
		logger: logger,
	}
}

type CreateRoomRequest struct {
	UserID      int    `json:"user_id" binding:"required"`
	RoomName    string `json:"room_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	BuyerID     int    `json:"buyer_id" binding:"required"`
	SellerID    int    `json:"seller_id" binding:"required"`
}

func (h *ChatHandler) CreateRoom(c *gin.Context) {
	const op = "ChatHandler.CreateRoom"
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request data"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}
	room := &domain.Room{
		Name:        req.RoomName,
		Description: req.Description,
		BuyerID:     req.BuyerID,
		SellerID:    req.SellerID,
		CreatedAt:   time.Now(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	room, err := h.rs.CreateRoom(ctx, room)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to create room"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}
	c.JSON(201, gin.H{"room_id": room.ID, "room_name": room.Name})
}

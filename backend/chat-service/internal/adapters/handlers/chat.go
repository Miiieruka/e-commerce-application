package handlers

import (
	"chat-service/internal/core/domain"
	"chat-service/internal/core/ports/repositories"
	"chat-service/internal/core/ports/services"
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"strconv"
	"time"
)

type ChatHandler struct {
	ms        services.MessageService
	rs        services.RoomService
	upr       repositories.UserPresenceRepository
	logger    *slog.Logger
	jwtSecret string
}

func NewChatHandler(
	ms services.MessageService,
	rs services.RoomService,
	upr repositories.UserPresenceRepository,
	logger *slog.Logger,
	jwtSecret string,
) *ChatHandler {
	return &ChatHandler{
		ms:        ms,
		rs:        rs,
		upr:       upr,
		logger:    logger,
		jwtSecret: jwtSecret,
	}
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

func (h *ChatHandler) GetRoomMessages(c *gin.Context) {
	const op = "ChatHandler.GetRoomMessages"
	roomID := c.Param("room_id")
	if roomID == "" {
		c.JSON(400, gin.H{"error": "Room ID is required"})
		h.logger.Error(op, slog.String("error", "Room ID is required"))
		return
	}
	id, err := strconv.Atoi(roomID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Room ID"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	messages, err := h.ms.GetMessagesByRoomID(ctx, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve messages"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}
	c.JSON(200, messages)
}

func (h *ChatHandler) GetUserRooms(c *gin.Context) {
	const op = "ChatHandler.GetUserRooms"
	userID := c.GetInt("user_id")
	h.logger.Debug(op, slog.Int("user_id", userID))
	if userID == 0 {
		c.JSON(400, gin.H{"error": "User ID is required"})
		h.logger.Error(op, slog.String("error", "User ID is required"))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rooms, err := h.rs.GetUserRooms(ctx, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve rooms"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}
	c.JSON(200, rooms)
}

func (h *ChatHandler) SendMessage(c *gin.Context) {
	const op = "ChatHandler.SendMessage"
	var req SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid request data"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}
	userID := c.GetInt("user_id")
	var message = &domain.Message{
		RoomID:    req.RoomID,
		SenderID:  userID,
		Content:   req.Message,
		CreatedAt: time.Now(),
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	messageResponse, err := h.ms.SendMessage(ctx, message)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "Failed to send message"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}

	c.IndentedJSON(200, messageResponse)
}

func (h *ChatHandler) GetRoomByID(c *gin.Context) {
	const op = "ChatHandler.GetRoomByID"
	roomID := c.Param("room_id")
	if roomID == "" {
		c.JSON(400, gin.H{"error": "Room ID is required"})
		h.logger.Error(op, slog.String("error", "Room ID is required"))
		return
	}
	id, err := strconv.Atoi(roomID)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Room ID"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	room, err := h.rs.GetRoomByID(ctx, id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve room"})
		h.logger.Error(op, slog.String("error", err.Error()))
		return
	}
	c.JSON(200, room)
}

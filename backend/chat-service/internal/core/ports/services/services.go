package services

import (
	"chat-service/internal/core/domain"
	"context"
)

type RoomService interface {
	CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
	GetRoomByID(ctx context.Context, roomID int) (*domain.Room, error)
	GetUserRooms(ctx context.Context, userID int) ([]*domain.Room, error)
}

type MessageService interface {
	SendMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	GetMessagesByRoomID(ctx context.Context, roomID int) ([]*domain.Message, error)
}

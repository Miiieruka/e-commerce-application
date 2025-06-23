package repositories

import (
	"chat-service/internal/core/domain"
	"context"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error)
	GetRoomByID(ctx context.Context, roomID int) (*domain.Room, error)
	GetUserRooms(ctx context.Context, userID int) ([]*domain.Room, error)
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
	GetMessagesByRoomID(ctx context.Context, roomID int) ([]*domain.Message, error)
	GetMessageByID(ctx context.Context, messageID int) (*domain.Message, error)
	UpdateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error)
}

type UserPresenceRepository interface {
	SetOnline(ctx context.Context, userID string) error
	SetOffline(ctx context.Context, userID string) error
	IsOnline(ctx context.Context, userID string) (bool, error)
}

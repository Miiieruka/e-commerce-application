package repository

import (
	"chat-service/internal/core/domain"
	"chat-service/internal/core/ports/repositories"
	"context"
	"database/sql"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) repositories.RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r *RoomRepository) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	return nil, nil
}
func (r *RoomRepository) GetRoomByID(ctx context.Context, roomID int) (*domain.Room, error) {
	return nil, nil
}
func (r *RoomRepository) GetUserRooms(ctx context.Context, userID int) ([]*domain.Room, error) {
	return nil, nil
}

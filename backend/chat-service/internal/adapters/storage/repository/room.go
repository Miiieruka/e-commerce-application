package repository

import (
	"chat-service/internal/core/domain"
	"chat-service/internal/core/ports/repositories"
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	const op = "RoomRepository.CreateRoom"

	query := `INSERT INTO rooms (name, description, created_at, buyer_id, seller_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, room.Name, room.Description, room.CreatedAt, room.BuyerID, room.SellerID).Scan(&room.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return room, nil
}
func (r *RoomRepository) GetRoomByID(ctx context.Context, roomID int) (*domain.Room, error) {
	const op = "RoomRepository.GetRoomByID"
	query := `SELECT * FROM rooms WHERE id = $1`
	var room domain.Room
	err := r.db.QueryRowContext(ctx, query, roomID).Scan(&room.ID, &room.Name, &room.Description, &room.CreatedAt, &room.UpdatedAt, &room.BuyerID, &room.SellerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: room not found", op)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &room, nil
}
func (r *RoomRepository) GetUserRooms(ctx context.Context, userID int) ([]*domain.Room, error) {
	const op = "RoomRepository.GetUserRooms"
	query := `SELECT * FROM rooms WHERE buyer_id = $1 OR seller_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: no rooms found for user", op)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	var rooms []*domain.Room
	for rows.Next() {
		var room domain.Room
		err := rows.Scan(&room.ID, &room.Name, &room.Description, &room.CreatedAt, &room.UpdatedAt, &room.BuyerID, &room.SellerID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		rooms = append(rooms, &room)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return rooms, nil
}

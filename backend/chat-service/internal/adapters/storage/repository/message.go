package repository

import (
	"chat-service/internal/core/domain"
	"chat-service/internal/core/ports/repositories"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) repositories.MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (r *MessageRepository) CreateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	const op = "MessageRepository.CreateMessage"
	query := `INSERT INTO messages (room_id, user_id, content, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, message.RoomID, message.SenderID, message.Content, message.CreatedAt).Scan(&message.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return message, nil
}
func (r *MessageRepository) GetMessagesByRoomID(ctx context.Context, roomID int) ([]*domain.Message, error) {
	const op = "MessageRepository.GetMessagesByRoomID"
	query := `SELECT id, room_id, user_id, content, created_at FROM messages WHERE room_id = $1 ORDER BY created_at ASC`
	rows, err := r.db.QueryContext(ctx, query, roomID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: No messages sent yet", op)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var messages []*domain.Message
	for rows.Next() {
		var message domain.Message
		if err := rows.Scan(&message.ID, &message.RoomID, &message.SenderID, &message.Content, &message.CreatedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		messages = append(messages, &message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return messages, nil
}
func (r *MessageRepository) GetMessageByID(ctx context.Context, messageID int) (*domain.Message, error) {
	const op = "MessageRepository.GetMessageByID"
	query := `SELECT id, room_id, user_id, content, created_at FROM messages WHERE id = $1`
	var message domain.Message
	err := r.db.QueryRowContext(ctx, query, messageID).Scan(&message.ID, &message.RoomID, &message.SenderID, &message.Content, &message.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: message not found", op)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &message, nil
}
func (r *MessageRepository) UpdateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	const op = "MessageRepository.UpdateMessage"
	query := `UPDATE messages SET content = $1, updated_at = $2 WHERE id = $3 RETURNING room_id, user_id, created_at`
	err := r.db.QueryRowContext(ctx, query, message.Content, message.UpdatedAt, message.ID).Scan(&message.RoomID, &message.SenderID, &message.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: message not found", op)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return message, nil
}

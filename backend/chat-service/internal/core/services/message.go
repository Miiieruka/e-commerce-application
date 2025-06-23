package services

import (
	"chat-service/internal/core/domain"
	"chat-service/internal/core/ports/repositories"
	"chat-service/internal/core/ports/services"
	"context"
	"fmt"
)

type MessageService struct {
	repo repositories.MessageRepository
}

func NewMessageService(repo repositories.MessageRepository) services.MessageService {
	return &MessageService{
		repo: repo,
	}
}

func (s *MessageService) SendMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
	const op = "MessageService.SendMessage"
	createdMessage, err := s.repo.CreateMessage(ctx, message)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return createdMessage, nil
}
func (s *MessageService) GetMessagesByRoomID(ctx context.Context, roomID int) ([]*domain.Message, error) {
	const op = "MessageService.GetMessagesByRoomID"
	messages, err := s.repo.GetMessagesByRoomID(ctx, roomID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return messages, nil
}

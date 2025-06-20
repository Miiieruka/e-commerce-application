package services

import (
	"chat-service/internal/core/domain"
	"chat-service/internal/core/ports/repositories"
	"chat-service/internal/core/ports/services"
	"context"
	"fmt"
)

type RoomService struct {
	roomRepo repositories.RoomRepository
}

func NewRoomService(roomRepo repositories.RoomRepository) services.RoomService {
	return &RoomService{
		roomRepo: roomRepo,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, room *domain.Room) (*domain.Room, error) {
	const op = "RoomService.CreateRoom"
	createdRoom, err := s.roomRepo.CreateRoom(ctx, room)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return createdRoom, nil
}

func (s *RoomService) GetRoomByID(ctx context.Context, roomID int) (*domain.Room, error) {
	const op = "RoomService.GetRoomByID"
	room, err := s.roomRepo.GetRoomByID(ctx, roomID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return room, nil
}

func (s *RoomService) GetUserRooms(ctx context.Context, userID int) ([]*domain.Room, error) {
	const op = "RoomService.GetUserRooms"
	rooms, err := s.roomRepo.GetUserRooms(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return rooms, nil
}

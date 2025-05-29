package domain

import (
	"context"
	"fmt"
	"user-service/internal/entities"
	"user-service/internal/storage"
)

type Userservice struct {
	repo *storage.Repository
}

func NewUserService(repo *storage.Repository) *Userservice {
	return &Userservice{
		repo: repo,
	}
}

func (s *Userservice) GetUserByID(ctx context.Context, id int64) (*entities.User, error) {
	const op = "service.domain.user.getuserbyid"

	us, err := s.repo.UserRepo.GetUserByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return us, nil
}

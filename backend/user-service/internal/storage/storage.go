package storage

import (
	"context"
	"user-service/internal/entities"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int64) (*entities.User, error)
}

type Repository struct {
	UserRepo UserRepository
}

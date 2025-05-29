package service

import (
	"context"
	"user-service/internal/entities"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int64) (*entities.User, error)
}

type Service struct {
	Serv      UserService
	JwtSecret string
}

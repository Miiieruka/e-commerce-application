package domain

import (
	"user-service/internal/service"
	"user-service/internal/storage"
)

func NewService(repo *storage.Repository) *service.Service {
	return &service.Service{
		Serv: NewUserService(repo),
	}
}

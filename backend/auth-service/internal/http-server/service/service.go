package service

import (
	"auth-service/internal/http-server/repository"
)

type AuthService struct {
	repo repository.Repository
	jwt  string
}

func NewAuthService(repo repository.Repository, jwt string) *AuthService {
	return &AuthService{repo: repo, jwt: jwt}
}

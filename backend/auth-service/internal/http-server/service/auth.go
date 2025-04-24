package service

import (
	"auth-service/internal/http-server/repository"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidUsername = errors.New("invalid username")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidRole     = errors.New("invalid role")
)

func (as *AuthService) Register(req RegisterRequest) (*repository.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &repository.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Email:    req.Email,
		Role:     req.Role,
	}

	err = as.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (as *AuthService) Login(req LoginRequest) (string, error) {
	user, err := as.repo.GetUserByUsername(req.Username)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", ErrInvalidUsername
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return "", ErrInvalidPassword
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(as.jwt))
}

package service

import (
	"auth-service/internal/http-server/repository"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
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

func (as *AuthService) GoogleOAuthLogin(token *oauth2.Token) (string, error) {
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		return "", fmt.Errorf("Google.Auth.login: %w", err)
	}
	defer resp.Body.Close()

	var us GoogleUser

	if err := json.NewDecoder(resp.Body).Decode(&us); err != nil {
		return "", fmt.Errorf("Google.Auth.login: %w", err)
	}
	user, err := as.repo.GetUserByUsername(us.Email)

	if err != nil {
		newUser := &repository.User{
			Username: us.Name,
			Email:    us.Email,
			Password: "",
			Role:     "buyer",
		}
		err = as.repo.CreateUser(newUser)
		if err != nil {
			return "", fmt.Errorf("Google.Auth.login: %w", err)
		}
		user = newUser
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	}
	tokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenJwt.SignedString([]byte(as.jwt))
}

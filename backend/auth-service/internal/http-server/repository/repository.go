package repository

import (
	"database/sql"
	"errors"
)

var (
	ErrUserNotFound = errors.New("User not found")
	ErrDB           = errors.New("DB error")
)

type Repository interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) CreateUser(user *User) error {
	query := "INSERT INTO users (username, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at;"
	return ur.db.QueryRow(query, user.Username, user.Email, user.Password, user.Role).Scan(&user.ID, &user.CreatedAt)
}

func (ur *UserRepository) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	query := `
		SELECT id, username, email, password, created_at, role
		FROM users
		WHERE username = $1;
	`

	err := ur.db.QueryRow(query, username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.Role)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, ErrDB
	}

	return user, nil

}

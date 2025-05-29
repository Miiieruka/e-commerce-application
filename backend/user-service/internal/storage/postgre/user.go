package postgre

import (
	"context"
	"database/sql"
	"fmt"
	"user-service/internal/entities"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int64) (*entities.User, error) {
	const op = "repository.postgre.user.getuserbyid"
	query := `
	SELECT id, username, email, role, created_at 
	FROM users
	WHERE id = $1
	`
	ru := &entities.User{}

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&ru.ID,
		&ru.Username,
		&ru.Email,
		&ru.Role,
		&ru.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return ru, nil
}

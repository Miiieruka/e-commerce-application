package storage

import (
	"chat-service/internal/core/ports/repositories"
	"github.com/redis/go-redis/v9"
)

type UserPresenceRepository struct {
	redisClient *redis.Client
}

func NewUserPresenceRepository(redisClient *redis.Client) repositories.UserPresenceRepository {
	return &UserPresenceRepository{
		redisClient: redisClient,
	}
}

func (r *UserPresenceRepository) SetOnline(userID string) error {
	return nil
}
func (r *UserPresenceRepository) SetOffline(userID string) error {
	return nil
}
func (r *UserPresenceRepository) IsOnline(userID string) (bool, error) {
	return true, nil
}

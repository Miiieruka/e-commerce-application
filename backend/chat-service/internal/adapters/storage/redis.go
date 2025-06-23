package storage

import (
	"chat-service/internal/core/ports/repositories"
	"context"
	"fmt"
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

func (r *UserPresenceRepository) SetOnline(ctx context.Context, userID string) error {
	const op = "UserPresenceRepository.SetOnline"
	key := "user_presence:" + userID
	if err := r.redisClient.Set(ctx, key, "online", 0).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (r *UserPresenceRepository) SetOffline(ctx context.Context, userID string) error {
	const op = "UserPresenceRepository.SetOffline"
	key := "user_presence:" + userID
	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
func (r *UserPresenceRepository) IsOnline(ctx context.Context, userID string) (bool, error) {
	const op = "UserPresenceRepository.IsOnline"
	key := "user_presence:" + userID
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	if val == "online" {
		return true, nil
	}
	return false, nil
}

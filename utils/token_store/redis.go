package token_store_util

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisImpl struct {
	RedisClient *redis.Client
}

func (u RedisImpl) First(token string) (string, error) {
	val, err := u.RedisClient.Get(context.Background(), fmt.Sprintf("token-%s", token)).Result()
	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}

func (u RedisImpl) Create(token string, userID string) error {
	_, err := u.RedisClient.Set(context.Background(), fmt.Sprintf("token-%s", token), userID, 3600 * time.Second).Result()
	if err != nil {
		return err 
	}

	return nil
}

package redis

import (
	"context"
	"time"

	"github.com/brkss/dextrace-server/internal/adapter/config"
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/redis/go-redis/v9"
)


type Redis struct {
	client *redis.Client
}


func New(ctx context.Context, config *config.Redis)(port.CacheRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr: 		config.Addr,
		Password: 	config.Password,
		DB: 		0,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err;
	}

	return &Redis{client}, nil
}

func (r *Redis)Set(ctx context.Context, key string, value []byte, ttl time.Duration) (error) {
	return r.client.Set(ctx, key, value, ttl).Err()
}


func (r *Redis)Get(ctx context.Context, key string) ([]byte, error) {
	res, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err;
	}
	bytes := []byte(res)
	return bytes, nil
}

func (r *Redis)Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Redis)Close() error {
	return r.client.Close()
}

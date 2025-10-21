package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	ctx := context.Background()

	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.PassWord,
		DB:           0,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,

		PoolSize:    cfg.Redis.Pollsize,
		PoolTimeout: cfg.Redis.PoolTimeout,
	})

	_, err := redisClient.Ping(ctx).Result()

	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client {

	return redisClient
}
func CloseRedis() {
	redisClient.Close()
}

func Set[T any](c *redis.Client, key string, value T, duration time.Duration) error {

	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(context.Background(), key, v, duration).Err()
}

func Get[T any](c *redis.Client, key string) (T, error) {
	var dest T = *new(T)

	v, err := c.Get(context.Background(), key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(v), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}

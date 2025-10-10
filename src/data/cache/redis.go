package cache

import (
	"fmt"
	"time"
	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/redis/go-redis/v9"
)




var redisClient *redis.Client


func InitRedis(cfg *config.Config){

	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.PassWord,
		DB: 0,
		DialTimeout: cfg.Redis.DialTimeout * time.Second,
		ReadTimeout: cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
	
		PoolSize: cfg.Redis.Pollsize,
		PoolTimeout: cfg.Redis.PoolTimeout,

	})

}


func GetRedis() *redis.Client{

	return redisClient
}
func CloseRedis(){
	redisClient.Close()
}
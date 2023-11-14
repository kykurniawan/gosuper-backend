package redis

import (
	"fmt"
	"gosuper/config"

	"github.com/redis/go-redis/v9"
)

func CreateClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
		Protocol: 3,
	})

	return client
}

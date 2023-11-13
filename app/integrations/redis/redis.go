package redis

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func CreateClient() *redis.Client {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
		Protocol: 3,
	})

	return client
}

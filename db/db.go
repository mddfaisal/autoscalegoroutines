package db

import (
	"time"

	"github.com/go-redis/redis"
)

const (
	DAYS_30 time.Duration = 2592000
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

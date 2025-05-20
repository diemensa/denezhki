package config

import "github.com/redis/go-redis/v9"

func NewRedisClient(host string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: "",
		DB:       0,
	})
	return rdb
}

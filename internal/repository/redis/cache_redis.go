package redislocal

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type CacheRedisRepo struct {
	client *redis.Client
}

func NewCacheRedisRepo(client *redis.Client) *CacheRedisRepo {
	return &CacheRedisRepo{client: client}
}

func (repo *CacheRedisRepo) Get(c context.Context, key string) (string, error) {
	return repo.client.Get(c, key).Result()
}

func (repo *CacheRedisRepo) Set(c context.Context, key string, value interface{}, ttl time.Duration) error {
	return repo.client.Set(c, key, value, ttl).Err()
}

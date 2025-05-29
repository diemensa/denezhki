package repository

import (
	"context"
	"time"
)

type CacheRepo interface {
	Get(c context.Context, key string) (string, error)
	Set(c context.Context, key string, value interface{}, ttl time.Duration) error
}

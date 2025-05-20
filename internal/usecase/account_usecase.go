package usecase

import (
	"github.com/diemensa/denezhki/internal/repository"
	"github.com/redis/go-redis/v9"
	"time"
)

type AccountService struct {
	accountRepo repository.AccountRepo
	redisClient *redis.Client
	cacheTTL    time.Duration
}

func NewAccountService(a repository.AccountRepo, redisClient *redis.Client, ttl time.Duration) *AccountService {
	return &AccountService{
		accountRepo: a,
		redisClient: redisClient,
		cacheTTL:    ttl,
	}
}

package usecase

import (
	"context"
	"fmt"
	"github.com/diemensa/denezhki/internal/repository"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"strconv"
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

func (s *AccountService) GetAccBalance(c context.Context, id uuid.UUID) (float64, error) {
	var balance float64

	key := "balance" + id.String()
	cache, err := s.redisClient.Get(c, key).Result()
	if err == nil {
		balance, err = strconv.ParseFloat(cache, 64)
		if err == nil {
			return balance, nil
		}
	}

	balance, err = s.accountRepo.GetAccBalance(c, id)
	if err != nil {
		return 0, nil
	}
	err = s.redisClient.Set(c, key, fmt.Sprintf("%f", balance), s.cacheTTL).Err()
	return balance, nil
}

// TODO другие методы по аккаунтам

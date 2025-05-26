package usecase

import (
	"context"
	"fmt"
	"github.com/diemensa/denezhki/internal/repository"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
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
func (s *AccountService) GetAccByID(c context.Context, id uuid.UUID) (*model.Account, error) {

	return s.accountRepo.GetAccByID(c, id)

}

func (s *AccountService) GetUserByAccID(c context.Context, id uuid.UUID) (*model.User, error) {

	return s.accountRepo.GetUserByAccID(c, id)

}

func (s *AccountService) GetAccBalanceByID(c context.Context, id uuid.UUID) (float64, error) {
	var balance float64

	key := "balance" + id.String()
	cache, err := s.redisClient.Get(c, key).Result()
	if err == nil {
		balance, err = strconv.ParseFloat(cache, 64)
		if err == nil {
			return balance, nil
		}
	}

	balance, err = s.accountRepo.GetAccBalanceByID(c, id)
	if err != nil {
		return 0, fmt.Errorf("cache error: %w", err)
	}
	err = s.redisClient.Set(c, key, fmt.Sprintf("%f", balance), s.cacheTTL).Err()
	return balance, nil
}

func (s *AccountService) GetAccByAliasOwner(c context.Context, alias, owner string) (*model.Account, error) {

	return s.accountRepo.GetAccByAliasOwner(c, alias, owner)
}

func (s *AccountService) UpdateAccBalance(c context.Context, id uuid.UUID, newBal float64) error {

	err := s.accountRepo.UpdateAccBalance(c, id, newBal)

	if err != nil {
		return fmt.Errorf("failed to update balance for %s: %w", id, err)
	}

	key := "balance" + id.String()
	err = s.redisClient.Set(c, key, fmt.Sprintf("%f", newBal), s.cacheTTL).Err()

	if err != nil {
		fmt.Printf("warning: failed to update balance cache for %s: %v\n", id, err)

	}

	return nil

}

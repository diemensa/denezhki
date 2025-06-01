package usecase

import (
	"context"
	"fmt"
	"github.com/diemensa/denezhki/internal/repository"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type AccountService struct {
	accountRepo repository.AccountRepo
	cacheClient repository.CacheRepo
	cacheTTL    time.Duration
}

func NewAccountService(a repository.AccountRepo, cacheClient repository.CacheRepo, ttl time.Duration) *AccountService {
	return &AccountService{
		accountRepo: a,
		cacheClient: cacheClient,
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
	cache, err := s.cacheClient.Get(c, key)
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
	err = s.cacheClient.Set(c, key, fmt.Sprintf("%f", balance), s.cacheTTL)
	return balance, nil
}

func (s *AccountService) GetAccByAliasUsername(c context.Context, alias, username string) (*model.Account, error) {

	return s.accountRepo.GetAccByAliasUsername(c, alias, username)
}

func (s *AccountService) UpdateAccBalance(c context.Context, id uuid.UUID, newBal float64) error {

	err := s.accountRepo.UpdateAccBalance(c, id, newBal)

	if err != nil {
		return fmt.Errorf("failed to update balance for %s: %w", id, err)
	}

	key := "balance" + id.String()
	err = s.cacheClient.Set(c, key, fmt.Sprintf("%f", newBal), s.cacheTTL)

	if err != nil {
		fmt.Printf("warning: failed to update balance cache for %s: %v\n", id, err)

	}

	return nil

}

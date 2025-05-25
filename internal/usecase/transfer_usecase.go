package usecase

import (
	"context"
	"fmt"
	"github.com/diemensa/denezhki/internal/repository"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type TransferService struct {
	accountRepo     repository.AccountRepo
	transactionRepo repository.TransRepo
	redisClient     *redis.Client
	cacheTTL        time.Duration
}

func NewTransferService(a repository.AccountRepo, t repository.TransRepo, redis *redis.Client, ttl time.Duration) *TransferService {
	return &TransferService{
		accountRepo:     a,
		transactionRepo: t,
		redisClient:     redis,
		cacheTTL:        ttl,
	}
}

func (s *TransferService) PerformTransfer(c context.Context,
	transactionID, fromID, toID uuid.UUID,
	amount float64) error {

	fromAccount, err := s.accountRepo.GetAccByID(c, fromID)
	if err != nil {
		return fmt.Errorf("failed to get the sender's account: %w", err)
	}

	toAccount, err := s.accountRepo.GetAccByID(c, toID)
	if err != nil {
		return fmt.Errorf("failed to get the recipient's account: %w", err)
	}

	if fromAccount.Balance < amount {
		return fmt.Errorf("non-sufficient funds")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	err = s.transactionRepo.PerformTransfer(c,
		transactionID, fromAccount.ID, toAccount.ID,
		fromAccount.Balance, toAccount.Balance,
		amount)

	if err != nil {
		return fmt.Errorf("couldn't send money: %w", err)
	}

	updateBalanceCache(c, s.redisClient, fromID, toID, fromAccount.Balance, toAccount.Balance, s.cacheTTL)

	return nil
}

func (s *TransferService) GetTransferByID(c context.Context,
	transactionID uuid.UUID) (*model.Transaction, error) {

	return s.transactionRepo.GetTransferByID(c, transactionID)
}

func (s *TransferService) GetAllAccountTransfers(
	c context.Context,
	accountID uuid.UUID) ([]model.Transaction, error) {

	return s.transactionRepo.GetAllAccountTransfers(c, accountID)
}

func updateBalanceCache(c context.Context, rdb *redis.Client,
	fromID, toID uuid.UUID,
	fromBal, toBal float64,
	ttl time.Duration) {

	fromKey := fmt.Sprintf("balance:%s", fromID.String())
	toKey := fmt.Sprintf("balance:%s", toID.String())

	_ = rdb.Set(c, fromKey, fromBal, ttl)
	_ = rdb.Set(c, toKey, toBal, ttl)
}

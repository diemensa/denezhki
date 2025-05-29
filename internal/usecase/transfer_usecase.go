package usecase

import (
	"context"
	"fmt"
	"github.com/diemensa/denezhki/internal/repository"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
	"time"
)

type TransferService struct {
	accountRepo     repository.AccountRepo
	transactionRepo repository.TransRepo
	cacheRepo       repository.CacheRepo
	cacheTTL        time.Duration
}

func NewTransferService(a repository.AccountRepo,
	t repository.TransRepo,
	cache repository.CacheRepo,
	ttl time.Duration) *TransferService {
	return &TransferService{
		accountRepo:     a,
		transactionRepo: t,
		cacheRepo:       cache,
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

	updateBalanceCache(c, s.cacheRepo, fromID, toID, fromAccount.Balance, toAccount.Balance, s.cacheTTL)

	return nil
}

func (s *TransferService) GetTransferByID(c context.Context,
	transactionID uuid.UUID) (*model.Transaction, error) {

	return s.transactionRepo.GetTransferByID(c, transactionID)
}

func (s *TransferService) GetAllAccountTransfers(
	c context.Context,
	alias, owner string) ([]model.Transaction, error) {

	account, err := s.accountRepo.GetAccByAliasOwner(c, alias, owner)

	if err != nil {
		return nil, err
	}

	return s.transactionRepo.GetAllAccountTransfers(c, account.ID)
}

func (s *TransferService) LogTransaction(c context.Context,
	transactionID, fromID, toID uuid.UUID,
	amount float64,
	success bool) {
	s.transactionRepo.LogTransaction(c, transactionID, fromID, toID, amount, success)
}

func updateBalanceCache(c context.Context, cacheRepo repository.CacheRepo,
	fromID, toID uuid.UUID,
	fromBal, toBal float64,
	ttl time.Duration) {

	fromKey := fmt.Sprintf("balance:%s", fromID.String())
	toKey := fmt.Sprintf("balance:%s", toID.String())

	_ = cacheRepo.Set(c, fromKey, fromBal, ttl)
	_ = cacheRepo.Set(c, toKey, toBal, ttl)
}

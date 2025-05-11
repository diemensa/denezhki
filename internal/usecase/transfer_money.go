package usecase

import (
	"context"
	"fmt"
	"github.com/diemensa/denezhki/internal/entity"
	"github.com/diemensa/denezhki/internal/repository"
	"github.com/google/uuid"
)

type TransferService struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

func NewTransferService(a repository.AccountRepository, t repository.TransactionRepository) *TransferService {
	return &TransferService{
		accountRepo:     a,
		transactionRepo: t,
	}
}

func (s *TransferService) Transfer(c context.Context, fromID, toID uuid.UUID, amount float64) error {
	fromAccount, err := s.accountRepo.GetByID(c, fromID)
	if err != nil {
		return fmt.Errorf("failed to get the sender's account: %w", err)
	}

	toAccount, err := s.accountRepo.GetByID(c, toID)
	if err != nil {
		return fmt.Errorf("failed to get the recipient's account: %w", err)
	}

	if fromAccount.Balance < amount {
		return fmt.Errorf("non-sufficient funds")
	}

	fromAccount.Balance -= amount
	toAccount.Balance += amount

	tran := entity.NewTransaction(fromID, toID, amount)

	err = s.transactionRepo.PerformTransfer(c, fromAccount.ID, toAccount.ID, fromAccount.Balance, toAccount.Balance)

	if err != nil {
		logErr := s.transactionRepo.LogTransaction(c, tran)
		if logErr != nil {
			return fmt.Errorf("couldn't log failed transaction: %w", err)
		}
		return fmt.Errorf("couldn't send money: %w", err)
	}

	tran.Success = true

	err = s.transactionRepo.LogTransaction(c, tran)
	if err != nil {
		return fmt.Errorf("couldn't log successful transaction: %w", err)
	}

	return nil
}

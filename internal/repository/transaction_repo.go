package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/entity"
	"github.com/google/uuid"
)

type TransactionRepository interface {
	LogTransaction(
		c context.Context,
		tran *entity.Transaction,
	) error

	PerformTransfer(
		c context.Context,
		fromID, toID uuid.UUID,
		fromNewBalance, toNewBalance float64,
	) error
}

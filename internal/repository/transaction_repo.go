package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
)

type TransRepo interface {
	PerformTransfer(c context.Context,
		transactionID, fromID, toID uuid.UUID,
		fromNewBalance, toNewBalance, amount float64,
	) error

	LogTransaction(c context.Context,
		transactionID, fromID, toID uuid.UUID,
		amount float64,
		success bool)

	GetTransferByID(c context.Context, transactionID uuid.UUID) (*model.Transaction, error)
	GetAllAccountTransfers(c context.Context, accountID uuid.UUID) ([]model.Transaction, error)
}

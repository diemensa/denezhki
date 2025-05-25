package repository

import (
	"context"
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
}

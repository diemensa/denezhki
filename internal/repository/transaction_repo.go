package repository

import (
	"context"
	"github.com/google/uuid"
)

type TransRepo interface {
	PerformTransfer(
		c context.Context,
		fromID, toID uuid.UUID,
		fromNewBalance, toNewBalance, amount float64,
	) error
}

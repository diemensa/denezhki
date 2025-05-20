package entity

import (
	"github.com/google/uuid"
	"time"
)

type TransactionEntity struct {
	ID        uuid.UUID
	FromAccID uuid.UUID
	ToAccID   uuid.UUID
	Amount    float64
	CreatedAt time.Time
}

func NewTransactionEntity(fromID, toID uuid.UUID, amount float64) *TransactionEntity {
	return &TransactionEntity{
		ID:        uuid.New(),
		FromAccID: fromID,
		ToAccID:   toID,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
}

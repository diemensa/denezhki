package entity

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID        uuid.UUID
	FromAccID uuid.UUID
	ToAccID   uuid.UUID
	Amount    float64
	Success   bool
	CreatedAt time.Time
}

func NewTransaction(fromID, toID uuid.UUID, amount float64) *Transaction {
	return &Transaction{
		ID:        uuid.New(),
		FromAccID: fromID,
		ToAccID:   toID,
		Amount:    amount,
		Success:   false,
		CreatedAt: time.Now(),
	}
}

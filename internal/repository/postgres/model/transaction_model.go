package model

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	FromAccID uuid.UUID `gorm:"type:uuid;not null"`
	ToAccID   uuid.UUID `gorm:"type:uuid;not null"`
	Amount    float64   `gorm:"not null"`
	Success   bool      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
}

func NewTransaction(transID, fromID, toID uuid.UUID, amount float64, success bool) *Transaction {
	return &Transaction{
		ID:        transID,
		FromAccID: fromID,
		ToAccID:   toID,
		Amount:    amount,
		Success:   success,
	}
}

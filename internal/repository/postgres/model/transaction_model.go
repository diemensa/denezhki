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
	CreatedAt time.Time `gorm:"not null"`
}

func NewTransaction(fromID, toID uuid.UUID, amount float64) *Transaction {
	return &Transaction{
		ID:        uuid.New(),
		FromAccID: fromID,
		ToAccID:   toID,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
}

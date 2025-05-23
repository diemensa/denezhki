package model

import (
	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID  uuid.UUID `gorm:"type:uuid;not null"`
	Balance float64   `gorm:"not null"`
}

func NewAccount(id uuid.UUID) *Account {
	return &Account{
		ID:      uuid.New(),
		UserID:  id,
		Balance: 0,
	}
}

package model

import (
	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID  uuid.UUID `gorm:"type:uuid;not null"`
	Alias   string    `gorm:"not null;uniqueIndex:idx_owner_alias"`
	Owner   string    `gorm:"not null;uniqueIndex:idx_owner_alias"`
	Balance float64   `gorm:"not null"`
}

func NewAccount(id uuid.UUID, username, alias string) *Account {
	return &Account{
		ID:      uuid.New(),
		UserID:  id,
		Alias:   alias,
		Owner:   username,
		Balance: 0,
	}
}

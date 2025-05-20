package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
}

type Account struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID  uuid.UUID `gorm:"type:uuid;not null"`
	Balance float64   `gorm:"not null"`
}

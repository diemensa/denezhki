package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Password string
}

type Account struct {
	ID      uuid.UUID
	UserID  uuid.UUID
	Balance float64
}

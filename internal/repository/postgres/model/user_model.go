package model

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string    `gorm:"uniqueIndex;not null"`
	Password string    `gorm:"not null"`
}

func NewUser(username, password string) *User {

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return &User{
		ID:       uuid.New(),
		Username: username,
		Password: string(hashPassword),
	}
}

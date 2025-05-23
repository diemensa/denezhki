package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
)

type UserRepo interface {
	GetUser(c context.Context, username string) (*model.User, error)
	GetUserAccounts(c context.Context, userID uuid.UUID) ([]model.Account, error)
	CreateUser(c context.Context, username, password string) error
	CreateAccount(c context.Context, userID uuid.UUID) error
	ValidatePassword(c context.Context, username, password string) error
}

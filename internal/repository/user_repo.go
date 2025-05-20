package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
)

type UserRepo interface {
	GetByUsername(c context.Context, username string) (*model.User, error)
	Create(c context.Context, user model.User) error
}

package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
)

type AccountRepo interface {
	GetAccByID(c context.Context, id uuid.UUID) (*model.Account, error)
	GetUserByAccID(c context.Context, id uuid.UUID) (*model.User, error)
	GetAccBalance(c context.Context, id uuid.UUID) (float64, error)
}

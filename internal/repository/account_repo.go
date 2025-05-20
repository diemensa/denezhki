package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
)

type AccountRepo interface {
	GetByID(c context.Context, id uuid.UUID) (*model.Account, error)
	GetBalance(c context.Context, id uuid.UUID) (float64, error)
}

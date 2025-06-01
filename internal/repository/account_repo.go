package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
)

type AccountRepo interface {
	GetAccByID(c context.Context, id uuid.UUID) (*model.Account, error)
	GetAccByAliasUsername(c context.Context, alias, owner string) (*model.Account, error)
	GetUserByAccID(c context.Context, id uuid.UUID) (*model.User, error)
	GetAccBalanceByID(c context.Context, id uuid.UUID) (float64, error)
	UpdateAccBalance(c context.Context, id uuid.UUID, newBal float64) error
}

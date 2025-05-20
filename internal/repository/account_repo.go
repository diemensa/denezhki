package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
)

type AccountRepository interface {
	GetByID(c context.Context, id uuid.UUID) (*model.Account, error)
}

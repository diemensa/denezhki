package repository

import (
	"context"
	"github.com/diemensa/denezhki/internal/entity"
	"github.com/google/uuid"
)

type AccountRepository interface {
	GetByID(c context.Context, id uuid.UUID) (*entity.Account, error)
}

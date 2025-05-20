package postgres

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccPostgresRepo struct {
	db *gorm.DB
}

func NewAccPostgresRepo(db *gorm.DB) *AccPostgresRepo {
	return &AccPostgresRepo{db: db}
}

func (repo *AccPostgresRepo) GetByID(c context.Context, id uuid.UUID) (*model.Account, error) {
	var account model.Account
	err := repo.db.WithContext(c).First(&account, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil

}

func (repo *AccPostgresRepo) GetBalance(c context.Context, id uuid.UUID) (float64, error) {
	var balance float64

	err := repo.db.WithContext(c).Select("balance").Where("id = ?", id).Take(&balance).Error
	if err != nil {
		return 0, err
	}

	return balance, nil
}

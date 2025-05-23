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

func (repo *AccPostgresRepo) GetAccByID(c context.Context, id uuid.UUID) (*model.Account, error) {
	var account model.Account
	err := repo.db.WithContext(c).First(&account, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil

}

func (repo *AccPostgresRepo) GetUserByAccID(c context.Context, id uuid.UUID) (*model.User, error) {
	var user model.User

	err := repo.db.WithContext(c).
		Table("users").
		Joins("JOIN accounts ON accounts.user_id = users.id").
		Where("accounts.id = ?", id).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (repo *AccPostgresRepo) GetAccBalance(c context.Context, id uuid.UUID) (float64, error) {
	var balance float64

	err := repo.db.WithContext(c).Select("balance").Where("id = ?", id).Take(&balance).Error
	if err != nil {
		return 0, err
	}

	return balance, nil
}

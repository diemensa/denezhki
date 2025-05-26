package postgres

import (
	"context"
	"errors"
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

func (repo *AccPostgresRepo) GetUserByAccOwner(c context.Context, owner string) (*model.User, error) {
	var user model.User

	if owner == "" {
		return nil, errors.New("owner parameter is empty")
	}

	err := repo.db.WithContext(c).
		Table("users").
		Joins("JOIN accounts ON accounts.user_id = users.id").
		Where("accounts.owner = ?", owner).
		First(&user).Error

	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (repo *AccPostgresRepo) GetAccBalanceByID(c context.Context, id uuid.UUID) (float64, error) {
	var balance float64

	err := repo.db.WithContext(c).
		Model(model.Account{}).
		Where("id = ?", id).
		Pluck("balance", &balance).Error

	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (repo *AccPostgresRepo) UpdateAccBalance(c context.Context, id uuid.UUID, newBal float64) error {

	if newBal < 0 {
		return errors.New("new balance is negative")
	}

	err := repo.db.WithContext(c).
		Model(&model.Account{}).
		Where("id = ?", id).
		Update("balance", newBal).Error

	if err != nil {
		return err
	}

	return nil

}

func (repo *AccPostgresRepo) GetAccByAliasOwner(c context.Context, alias, owner string) (*model.Account, error) {
	var account model.Account

	if err := repo.db.WithContext(c).
		Where("alias = ? AND owner = ?", alias, owner).First(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}

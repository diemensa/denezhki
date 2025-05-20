package postgres

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransPostgresRepo struct {
	db *gorm.DB
}

func NewTransPostgresRepo(db *gorm.DB) *TransPostgresRepo {
	return &TransPostgresRepo{db: db}
}

func (repo *TransPostgresRepo) PerformTransfer(c context.Context,
	fromID, toID uuid.UUID,
	fromNewBalance, toNewBalance, amount float64) error {

	return repo.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.WithContext(c).Model(&model.Account{}).
			Where("id = ?", fromID).
			Update("balance", fromNewBalance).Error; err != nil {
			return err
		}

		if err := tx.WithContext(c).Model(&model.Account{}).
			Where("id = ?", toID).
			Update("balance", toNewBalance).Error; err != nil {
			return err
		}

		transaction := model.NewTransaction(fromID, toID, amount)

		if err := tx.WithContext(c).Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})

}

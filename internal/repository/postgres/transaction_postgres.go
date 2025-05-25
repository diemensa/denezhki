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
	transactionID, fromID, toID uuid.UUID,
	fromNewBalance, toNewBalance, amount float64) error {

	return repo.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.WithContext(c).Model(&model.Account{}).
			Where("id = ?", fromID).
			Update("balance", fromNewBalance).Error; err != nil {
			repo.LogTransaction(c, transactionID, fromID, toID, amount, false)
			return err
		}

		if err := tx.WithContext(c).Model(&model.Account{}).
			Where("id = ?", toID).
			Update("balance", toNewBalance).Error; err != nil {
			repo.LogTransaction(c, transactionID, fromID, toID, amount, false)
			return err
		}

		repo.LogTransaction(c, transactionID, fromID, toID, amount, true)
		return nil
	})

}

func (repo *TransPostgresRepo) LogTransaction(c context.Context,
	transactionID, fromID, toID uuid.UUID,
	amount float64,
	success bool) {

	transaction := model.NewTransaction(transactionID, fromID, toID, amount, success)

	repo.db.WithContext(c).Create(&transaction)
}

func (repo *TransPostgresRepo) GetTransferByID(c context.Context,
	transactionID uuid.UUID) (*model.Transaction, error) {

	var transfer model.Transaction

	if err := repo.db.WithContext(c).First(&transfer, "id = ?", transactionID).Error; err != nil {
		return nil, err
	}

	return &transfer, nil
}

func (repo *TransPostgresRepo) GetAllAccountTransfers(
	c context.Context,
	accountID uuid.UUID) ([]model.Transaction, error) {

	var transfer []model.Transaction

	if err := repo.db.WithContext(c).Where("from_acc_id = ?", accountID).
		Or("to_acc_id = ?", accountID).
		Find(&transfer).Error; err != nil {
		return nil, err
	}

	return transfer, nil

}

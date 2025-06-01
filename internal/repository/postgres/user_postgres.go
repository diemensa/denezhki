package postgres

import (
	"context"
	"fmt"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

type UserPostgresRepo struct {
	db *gorm.DB
}

func NewUserPostgresRepo(db *gorm.DB) *UserPostgresRepo {
	return &UserPostgresRepo{db: db}
}

func (repo *UserPostgresRepo) GetUserByID(c context.Context, userID uuid.UUID) (*model.User, error) {
	var user model.User

	err := repo.db.WithContext(c).Where("id = ?", userID).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserPostgresRepo) GetUserByUsername(c context.Context, username string) (*model.User, error) {
	var user model.User

	err := repo.db.WithContext(c).Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserPostgresRepo) GetUserAccounts(c context.Context, userID uuid.UUID) ([]model.Account, error) {
	var accounts []model.Account

	err := repo.db.WithContext(c).Where("user_id = ?", userID).Find(&accounts).Error

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (repo *UserPostgresRepo) CreateUser(c context.Context, username, password string) error {
	user := model.NewUser(username, password)
	err := repo.db.WithContext(c).Create(&user).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") &&
			strings.Contains(err.Error(), "idx_users_username") {
			return fmt.Errorf(`username "%s" is already taken`, username)
		}

		return err
	}

	return nil

}

func (repo *UserPostgresRepo) CreateAccount(c context.Context, userID uuid.UUID, username, alias string) error {
	account := model.NewAccount(userID, username, alias)
	err := repo.db.WithContext(c).Create(&account).Error

	if err != nil {

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") &&
			strings.Contains(err.Error(), "idx_owner_alias") {
			return fmt.Errorf(`user "%s" already has account named "%s"`, username, alias)

		}

		return err
	}

	return nil

}

func (repo *UserPostgresRepo) ValidatePassword(c context.Context, username, password string) error {
	var user model.User

	err := repo.db.WithContext(c).Where("username = ?", username).First(&user).Error

	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

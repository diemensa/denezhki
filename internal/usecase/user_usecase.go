package usecase

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(u repository.UserRepo) *UserService {
	return &UserService{
		userRepo: u,
	}
}

func (u *UserService) GetUser(c context.Context, username string) (*model.User, error) {

	return u.userRepo.GetUser(c, username)

}

func (u *UserService) GetUserAccounts(c context.Context, userID uuid.UUID) ([]model.Account, error) {

	return u.userRepo.GetUserAccounts(c, userID)

}

func (u *UserService) CreateUser(c context.Context, username, password string) error {

	return u.userRepo.CreateUser(c, username, password)

}

func (u *UserService) CreateAccount(c context.Context, userID uuid.UUID) error {

	return u.userRepo.CreateAccount(c, userID)

}

func (u *UserService) ValidatePassword(c context.Context, username, password string) error {

	return u.userRepo.ValidatePassword(c, username, password)

}

package usecase

import (
	"github.com/diemensa/denezhki/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepo
}

func NewUserService(u repository.UserRepo) *UserService {
	return &UserService{
		userRepo: u,
	}
}

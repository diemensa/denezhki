package usecase_test

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/diemensa/denezhki/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_Login(t *testing.T) {
	c := context.Background()
	username := "rizzler"
	password := "gyatt"

	user := &model.User{
		Username: username,
		Password: password,
	}

	mockUserRepo := mocks.NewUserRepo(t)
	mockUserService := usecase.NewUserService(mockUserRepo)

	mockUserRepo.On("GetUserByUsername", c, username).Return(user, nil)
	mockUserRepo.On("ValidatePassword", c, username, password).Return(nil)

	service := usecase.NewAuthService(mockUserService, "secrettest")

	token, err := service.Login(c, username, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	mockUserRepo.AssertExpectations(t)

}

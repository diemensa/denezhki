package usecase_test

import (
	"context"
	"testing"

	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/diemensa/denezhki/internal/usecase/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserService_GetUserByID(t *testing.T) {
	c := context.Background()
	userID := uuid.New()
	expectedUser := &model.User{
		ID:       userID,
		Username: "skibidi",
	}

	mockRepo := mocks.NewUserRepo(t)
	mockRepo.On("GetUserByID", c, userID).Return(expectedUser, nil)

	service := usecase.NewUserService(mockRepo)

	result, err := service.GetUserByID(c, userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserByUsername(t *testing.T) {
	c := context.Background()
	expectedUser := &model.User{
		ID:       uuid.New(),
		Username: "skibidi",
	}

	mockRepo := mocks.NewUserRepo(t)
	mockRepo.On("GetUserByUsername", c, expectedUser.Username).Return(expectedUser, nil)

	service := usecase.NewUserService(mockRepo)

	result, err := service.GetUserByUsername(c, expectedUser.Username)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)

	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserAccounts(t *testing.T) {

	userID := uuid.New()

	c := context.Background()
	expectedAccounts := []model.Account{
		{ID: uuid.New(), Alias: "savings", Owner: "Sigma"},
		{ID: uuid.New(), Alias: "spendings", Owner: "Gronk"},
	}

	mockRepo := mocks.NewUserRepo(t)
	mockRepo.On("GetUserAccounts", c, userID).Return(expectedAccounts, nil)

	service := usecase.NewUserService(mockRepo)

	result, err := service.GetUserAccounts(c, userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccounts, result)

	mockRepo.AssertExpectations(t)

}

func TestUserService_CreateUser(t *testing.T) {
	c := context.Background()
	username, password := "skibidi", "gyatt"

	mockRepo := mocks.NewUserRepo(t)
	mockRepo.On("CreateUser", c, username, password).Return(nil)

	service := usecase.NewUserService(mockRepo)

	err := service.CreateUser(c, username, password)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserService_CreateAccount(t *testing.T) {
	c := context.Background()
	userID, username, alias := uuid.New(), "skibidi", "savings"

	mockRepo := mocks.NewUserRepo(t)
	mockRepo.On("CreateAccount", c, userID, username, alias).Return(nil)

	service := usecase.NewUserService(mockRepo)

	err := service.CreateAccount(c, userID, username, alias)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestUserService_ValidatePassword(t *testing.T) {
	c := context.Background()
	username, password := "skibidi", "gyatt"

	mockRepo := mocks.NewUserRepo(t)
	mockRepo.On("ValidatePassword", c, username, password).Return(nil)

	service := usecase.NewUserService(mockRepo)

	err := service.ValidatePassword(c, username, password)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)

}

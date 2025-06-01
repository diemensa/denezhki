package usecase_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/diemensa/denezhki/internal/usecase/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountService_GetAccByID(t *testing.T) {
	c := context.Background()
	accID, userID := uuid.New(), uuid.New()
	expectedAccount := &model.Account{
		ID:      accID,
		UserID:  userID,
		Alias:   "savings",
		Owner:   "skibidi",
		Balance: 1500,
	}

	mockAccRepo := mocks.NewAccountRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)

	mockAccRepo.On("GetAccByID", c, accID).Return(expectedAccount, nil)

	service := usecase.NewAccountService(mockAccRepo, mockCacheRepo, 3*time.Minute)

	result, err := service.GetAccByID(c, accID)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, result)

	mockAccRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
}

func TestAccountService_GetUserByAccID(t *testing.T) {
	c := context.Background()
	accID, userID := uuid.New(), uuid.New()
	expectedUser := &model.User{
		ID:       userID,
		Username: "sigma",
	}

	mockAccRepo := mocks.NewAccountRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)
	mockAccRepo.On("GetUserByAccID", c, accID).Return(expectedUser, nil)

	service := usecase.NewAccountService(mockAccRepo, mockCacheRepo, 3*time.Minute)

	result, err := service.GetUserByAccID(c, accID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, result)

	mockAccRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
}

func TestAccountService_GetAccBalanceByID(t *testing.T) {
	c := context.Background()
	expectedBalance := float64(1500)
	accID := uuid.New()

	mockAccRepo := mocks.NewAccountRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)

	mockCacheRepo.On("Get", c, mock.Anything).Return("", nil)
	mockCacheRepo.On("Set", c, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	mockAccRepo.On("GetAccBalanceByID", c, accID).Return(expectedBalance, nil)

	service := usecase.NewAccountService(mockAccRepo, mockCacheRepo, 3*time.Minute)

	result, err := service.GetAccBalanceByID(c, accID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBalance, result)

	mockAccRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
}

func TestAccountService_GetAccByAliasOwner(t *testing.T) {
	c := context.Background()
	alias, owner := "skibidi", "sigma"
	expectedAcc := &model.Account{
		ID:      uuid.New(),
		UserID:  uuid.New(),
		Alias:   alias,
		Owner:   owner,
		Balance: 0,
	}

	mockAccRepo := mocks.NewAccountRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)

	mockAccRepo.On("GetAccByAliasUsername", c, alias, owner).Return(expectedAcc, nil)

	service := usecase.NewAccountService(mockAccRepo, mockCacheRepo, 3*time.Minute)

	result, err := service.GetAccByAliasUsername(c, alias, owner)

	assert.NoError(t, err)
	assert.Equal(t, expectedAcc, result)

	mockAccRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
}

func TestAccountService_UpdateAccBalance(t *testing.T) {

	c := context.Background()
	accID := uuid.New()
	newBal := 1500.0

	mockAccRepo := mocks.NewAccountRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)

	mockAccRepo.On("UpdateAccBalance", c, accID, newBal).Return(nil)
	mockCacheRepo.On("Set", c, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	service := usecase.NewAccountService(mockAccRepo, mockCacheRepo, 3*time.Minute)

	err := service.UpdateAccBalance(c, accID, newBal)

	assert.NoError(t, err)

	mockAccRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
}

package usecase_test

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/diemensa/denezhki/internal/usecase/mocks"
	"github.com/google/uuid"
)

func TestTransferService_PerformTransfer(t *testing.T) {
	c := context.Background()
	transactionID, fromID, toID := uuid.New(), uuid.New(), uuid.New()
	amount := 300.0

	mockTransRepo := mocks.NewTransRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)
	mockAccountRepo := mocks.NewAccountRepo(t)
	cacheTTL := 5 * time.Minute

	fromAcc := &model.Account{
		ID:      fromID,
		Balance: 1000.0,
	}
	toAcc := &model.Account{
		ID:      toID,
		Balance: 100.0,
	}

	mockTransRepo.On("PerformTransfer", c,
		transactionID, fromID, toID,
		700.0, 400.0, amount).Return(nil)
	mockAccountRepo.On("GetAccByID", c, fromAcc.ID).Return(fromAcc, nil)
	mockAccountRepo.On("GetAccByID", c, toAcc.ID).Return(toAcc, nil)
	mockCacheRepo.On("Set", c, mock.Anything, mock.Anything, cacheTTL).Return(nil)

	service := usecase.NewTransferService(mockAccountRepo, mockTransRepo, mockCacheRepo, cacheTTL)

	err := service.PerformTransfer(c, transactionID, fromID, toID, amount)

	assert.NoError(t, err)

	mockTransRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)

}

func TestTransferService_LogTransaction(t *testing.T) {
	c := context.Background()
	transactionID, fromID, toID := uuid.New(), uuid.New(), uuid.New()
	amount := 300.0

	mockTransRepo := mocks.NewTransRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)
	mockAccountRepo := mocks.NewAccountRepo(t)
	cacheTTL := 5 * time.Minute

	mockTransRepo.On("LogTransaction", c, transactionID, fromID, toID, amount, true).Return(nil)

	service := usecase.NewTransferService(mockAccountRepo, mockTransRepo, mockCacheRepo, cacheTTL)

	err := service.LogTransaction(c, transactionID, fromID, toID, amount, true)
	assert.NoError(t, err)

	mockTransRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)
}

func TestTransferService_GetTransferByID(t *testing.T) {

	c := context.Background()

	transaction := &model.Transaction{
		ID:        uuid.New(),
		FromAccID: uuid.New(),
		ToAccID:   uuid.New(),
		Amount:    5,
		Success:   true,
		CreatedAt: time.Now(),
	}

	mockTransRepo := mocks.NewTransRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)
	mockAccountRepo := mocks.NewAccountRepo(t)
	cacheTTL := 5 * time.Minute
	mockTransRepo.On("GetTransferByID", c, transaction.ID).Return(transaction, nil)

	service := usecase.NewTransferService(mockAccountRepo, mockTransRepo, mockCacheRepo, cacheTTL)
	ret, err := service.GetTransferByID(c, transaction.ID)

	assert.NoError(t, err)
	assert.Equal(t, transaction, ret)

	mockTransRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)

}

func TestTransferService_GetAllAccountTransfers(t *testing.T) {
	mockTransRepo := mocks.NewTransRepo(t)
	mockCacheRepo := mocks.NewCacheRepo(t)
	mockAccountRepo := mocks.NewAccountRepo(t)
	cacheTTL := 5 * time.Minute

	var expectedRet []model.Transaction

	c := context.Background()
	alias, owner := "sigma", "test"
	account := &model.Account{
		ID:      uuid.New(),
		UserID:  uuid.New(),
		Alias:   "sigma",
		Owner:   "test",
		Balance: 200,
	}

	mockAccountRepo.On("GetAccByAliasOwner", c, alias, owner).Return(account, nil)
	mockTransRepo.On("GetAllAccountTransfers", c, account.ID).Return(expectedRet, nil)

	service := usecase.NewTransferService(mockAccountRepo, mockTransRepo, mockCacheRepo, cacheTTL)
	ret, err := service.GetAllAccountTransfers(c, alias, owner)

	assert.NoError(t, err)
	assert.Equal(t, expectedRet, ret)
	mockTransRepo.AssertExpectations(t)
	mockCacheRepo.AssertExpectations(t)
	mockAccountRepo.AssertExpectations(t)

}

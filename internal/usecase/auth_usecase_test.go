package usecase_test

import (
	"context"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/diemensa/denezhki/internal/usecase/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestAuthService_ValidateToken(t *testing.T) {
	expectedUsername, secret := "rizzler", "megasecret1337"

	mockUserRepo := mocks.NewUserRepo(t)
	mockUserService := usecase.NewUserService(mockUserRepo)

	service := usecase.NewAuthService(mockUserService, secret)
	testToken, err := generateTestToken(secret, expectedUsername)
	assert.NoError(t, err)

	gotUsername, err := service.ValidateToken(testToken)

	assert.NoError(t, err)
	assert.Equal(t, expectedUsername, gotUsername)

}

func generateTestToken(secret, username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

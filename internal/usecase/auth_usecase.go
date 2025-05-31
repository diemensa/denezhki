package usecase

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type AuthService struct {
	userService *UserService
	jwtSecret   string
}

func NewAuthService(u *UserService, secret string) *AuthService {
	return &AuthService{
		userService: u,
		jwtSecret:   secret,
	}
}

func (a *AuthService) Login(c context.Context, username, password string) (string, error) {
	user, err := a.userService.GetUserByUsername(c, username)
	if err != nil {
		return "", fmt.Errorf("incorrect username or password")
	}

	err = a.userService.ValidatePassword(c, username, password)
	if err != nil {
		return "", fmt.Errorf("incorrect username or password")
	}

	token, err := a.generateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("couldn't generate JWT token")
	}

	return token, nil
}

func (a *AuthService) generateToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtSecret))
}

func (a *AuthService) ValidateToken(tokenString string) (userID uuid.UUID, err error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.jwtSecret), nil
	})

	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid token claims")
	}

	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user_id not found in token claims")
	}

	userID, err = uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user_id in token claims")
	}

	return userID, nil
}

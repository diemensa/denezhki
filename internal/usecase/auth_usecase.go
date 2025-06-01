package usecase

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
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
	_, err := a.userService.GetUserByUsername(c, username)
	if err != nil {
		return "", fmt.Errorf("incorrect username or password")
	}

	err = a.userService.ValidatePassword(c, username, password)
	if err != nil {
		return "", fmt.Errorf("incorrect username or password")
	}

	token, err := a.generateToken(username)
	if err != nil {
		return "", fmt.Errorf("couldn't generate JWT token")
	}

	return token, nil
}

func (a *AuthService) generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtSecret))
}

func (a *AuthService) ValidateToken(tokenString string) (username string, err error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.jwtSecret), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	tokenUser, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("username not found in token claims")
	}

	return tokenUser, nil
}

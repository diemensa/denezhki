package handler

import "github.com/diemensa/denezhki/internal/usecase"

type UserHandler struct {
	service *usecase.UserService
}

func NewUserHandler(s *usecase.UserService) *UserHandler {
	return &UserHandler{service: s}
}

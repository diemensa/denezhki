package handler

import "github.com/diemensa/denezhki/internal/usecase"

type AccountHandler struct {
	service *usecase.AccountService
}

func NewAccountHandler(s *usecase.AccountService) *AccountHandler {
	return &AccountHandler{service: s}
}

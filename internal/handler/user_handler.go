package handler

import (
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *usecase.UserService
}

func NewUserHandler(s *usecase.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) HandleGetUserByUsername(c *gin.Context) {

}

func (h *UserHandler) HandleGetUserAccounts(c *gin.Context) {

}

func (h *UserHandler) HandleCreateUser(c *gin.Context) {

}

func (h *UserHandler) HandleCreateAccount(c *gin.Context) {

}

func (h *UserHandler) HandleValidatePassword(c *gin.Context) {

}

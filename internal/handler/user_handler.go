package handler

import (
	"github.com/diemensa/denezhki/internal/handler/dto"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service *usecase.UserService
}

func NewUserHandler(s *usecase.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) HandleGetUserAccounts(c *gin.Context) {
	username := c.Param("username")
	user, err := h.service.GetUserByUsername(c, username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	accounts, err := h.service.GetUserAccounts(c, user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	} else if len(accounts) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "user has no accounts",
		})
		return
	}

	c.JSON(http.StatusOK, accounts)

}

func (h *UserHandler) HandleCreateUser(c *gin.Context) {
	var newUser dto.CreateUserRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.service.CreateUser(c, newUser.Username, newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created"})

}

func (h *UserHandler) HandleCreateAccount(c *gin.Context) {
	username := c.Param("username")

	var req dto.CreateAccountRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "something's wrong with the alias",
		})
		return
	}

	user, err := h.service.GetUserByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
		return
	}

	err = h.service.CreateAccount(c, user.ID, username, req.Alias)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "account created"})
}

//func (h *UserHandler) HandleValidatePassword(c *gin.Context) {
//
//}

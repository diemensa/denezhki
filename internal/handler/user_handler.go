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

// HandleGetUserAccounts
// @Summary Get all accounts of a user
// @Tags User
// @Param username path string true "Username"
// @Success 200 {array} dto.AccountResponse
// @Failure 400 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users/{username}/accounts [get]
func (h *UserHandler) HandleGetUserAccounts(c *gin.Context) {
	username := c.Param("username")

	err := CheckUserMatch(c, username)
	if err != nil {
		RespondWithError(c, http.StatusForbidden, err.Error())
		return
	}

	user, err := h.service.GetUserByUsername(c, username)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	accounts, err := h.service.GetUserAccounts(c, user.ID)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, accounts)

}

// HandleCreateUser
// @Summary Create a new user
// @Tags User
// @Param user body dto.CreateUserRequest true "New user data"
// @Success 201 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users [post]
func (h *UserHandler) HandleCreateUser(c *gin.Context) {
	var newUser dto.CreateUserRequest

	if err := c.ShouldBindJSON(&newUser); err != nil {
		RespondWithError(c, http.StatusBadRequest, "invalid request")
		return
	}

	err := h.service.CreateUser(c, newUser.Username, newUser.Password)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithMessage(c, http.StatusCreated, "user created")

}

// HandleCreateAccount
// @Summary Create a new account for a user
// @Tags User
// @Param username path string true "Username"
// @Param account body dto.CreateAccountRequest true "New account data"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /users/{username}/accounts [post]
func (h *UserHandler) HandleCreateAccount(c *gin.Context) {
	username := c.Param("username")

	err := CheckUserMatch(c, username)
	if err != nil {
		RespondWithError(c, http.StatusForbidden, err.Error())
		return
	}

	var req dto.CreateAccountRequest

	err = c.ShouldBindJSON(&req)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "something's wrong with the alias")
		return
	}

	user, err := h.service.GetUserByUsername(c, username)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "user does not exist")
		return
	}

	err = h.service.CreateAccount(c, user.ID, username, req.Alias)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithMessage(c, http.StatusOK, "account created")
}

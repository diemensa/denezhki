package handler

import (
	"github.com/diemensa/denezhki/internal/handler/dto"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type AccountHandler struct {
	service *usecase.AccountService
}

func NewAccountHandler(s *usecase.AccountService) *AccountHandler {
	return &AccountHandler{service: s}
}

func (h *AccountHandler) HandleGetAccByID(c *gin.Context) {

	idParam := c.Param("id")
	accountID, err := uuid.Parse(idParam)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "invalid UUID")
		return
	}

	acc, err := h.service.GetAccByID(c, accountID)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, acc)
}

// HandleGetAccBalance
// @Summary Get account balance
// @Tags Account
// @Param username path string true "Username"
// @Param alias path string true "Account Alias"
// @Success 200 {object} dto.BalanceResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{username}/accounts/{alias}/balance [get]
func (h *AccountHandler) HandleGetAccBalance(c *gin.Context) {
	alias, owner := extractAliasOwner(c)
	account, err := h.service.GetAccByAliasOwner(c, alias, owner)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.BalanceResponse{Balance: account.Balance})
}

// HandleGetAccByAliasOwner
// @Summary Get account by alias and owner username
// @Tags Account
// @Param username path string true "Username"
// @Param alias path string true "Account Alias"
// @Success 200 {object} dto.AccountResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{username}/accounts/{alias} [get]
func (h *AccountHandler) HandleGetAccByAliasOwner(c *gin.Context) {
	alias, owner := extractAliasOwner(c)

	account, err := h.service.GetAccByAliasOwner(c, alias, owner)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.NewAccountResponse(account)

	c.JSON(http.StatusOK, response)
}

// HandleUpdateBalance
// @Summary Update account balance
// @Tags Account
// @Param username path string true "Username"
// @Param alias path string true "Account Alias"
// @Param balance body dto.BalanceRequest true "Balance update payload"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{username}/accounts/{alias}/balance [put]
func (h *AccountHandler) HandleUpdateBalance(c *gin.Context) {
	alias, owner := extractAliasOwner(c)
	var req dto.BalanceRequest

	account, err := h.service.GetAccByAliasOwner(c, alias, owner)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, "invalid request")
		return
	}

	err = h.service.UpdateAccBalance(c, account.ID, req.Balance)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "couldn't update balance")
		return
	}

	RespondWithMessage(c, http.StatusOK, "balance updated")
}

func extractAliasOwner(c *gin.Context) (alias, owner string) {
	return c.Param("alias"), c.Param("username")
}

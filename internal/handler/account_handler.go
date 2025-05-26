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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid UUID",
		})
		return
	}

	acc, err := h.service.GetAccByID(c, accountID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, acc)
}

func (h *AccountHandler) HandleGetAccBalance(c *gin.Context) {
	alias, owner := extractAliasOwner(c)
	account, err := h.service.GetAccByAliasOwner(c, alias, owner)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.BalanceResponse{Balance: account.Balance})
}

func (h *AccountHandler) HandleGetAccByAliasOwner(c *gin.Context) {
	alias, owner := extractAliasOwner(c)

	account, err := h.service.GetAccByAliasOwner(c, alias, owner)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, account)
}

func (h *AccountHandler) HandleUpdateBalance(c *gin.Context) {
	alias, owner := extractAliasOwner(c)
	var req dto.BalanceRequest

	account, err := h.service.GetAccByAliasOwner(c, alias, owner)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err = h.service.UpdateAccBalance(c, account.ID, req.Balance)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "couldn't update balance"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "balance updated"})
}

func extractAliasOwner(c *gin.Context) (alias, owner string) {
	return c.Param("alias"), c.Param("username")
}

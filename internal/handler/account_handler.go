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
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, acc)
}

func (h *AccountHandler) HandleGetAccBalance(c *gin.Context) {

	idParam := c.Param("id")
	accountID, err := uuid.Parse(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid UUID",
		})
		return
	}

	balance, err := h.service.GetAccBalance(c, accountID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, dto.BalanceResponse{Balance: balance})
}

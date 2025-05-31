package handler

import (
	"github.com/diemensa/denezhki/internal/handler/dto"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type TransferHandler struct {
	service *usecase.TransferService
}

func NewTransferHandler(s *usecase.TransferService) *TransferHandler {
	return &TransferHandler{service: s}
}

// HandleTransfer
// @Summary Perform a transfer between accounts
// @Tags Transfer
// @Param transfer body dto.TransferRequest true "Transfer details"
// @Success 200 {object} dto.TransferResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{username}/accounts/{alias}/transfers [post]
func (h *TransferHandler) HandleTransfer(c *gin.Context) {
	var req dto.TransferRequest

	alias, owner := extractAliasOwner(c)
	account, err := h.service.GetAccByAliasOwner(c, alias, owner)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "such alias/username doesn't exist")
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "amount must be >= 1 or ID's aren't of UUID type")
		return
	}

	if account.ID == req.ToID {
		RespondWithError(c, http.StatusBadRequest, "you can't transfer money to the same account")
		return
	}

	transferID := uuid.New()
	err = h.service.PerformTransfer(c, transferID, account.ID, req.ToID, req.Amount)
	if err != nil {
		err = h.service.LogTransaction(c, transferID, account.ID, req.ToID, req.Amount, false)
		if err != nil {
			RespondWithError(c, http.StatusBadRequest, "couldn't log failed transaction")
		}
		c.JSON(http.StatusBadRequest, dto.NewTransferResponse(transferID, false))
		return
	}

	c.JSON(http.StatusOK, dto.NewTransferResponse(transferID, true))

}

// HandleGetTransferByID
// @Summary Get transfer details by ID
// @Tags Transfer
// @Param id path string true "Transfer UUID"
// @Success 200 {object} dto.TransferResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /transfers/{id} [get]
func (h *TransferHandler) HandleGetTransferByID(c *gin.Context) {
	idParam := c.Param("id")
	transferID, err := uuid.Parse(idParam)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "invalid UUID")
		return
	}

	transfer, err := h.service.GetTransferByID(c, transferID)

	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, transfer)
}

// HandleGetAllAccTransfers
// @Summary Get all transfers for an account
// @Tags Transfer
// @Param username path string true "Username"
// @Param alias path string true "Account Alias"
// @Success 200 {array} dto.TransferResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{username}/accounts/{alias}/transfers [get]
func (h *TransferHandler) HandleGetAllAccTransfers(c *gin.Context) {

	alias, owner := extractAliasOwner(c)

	transfers, err := h.service.GetAllAccountTransfers(c, alias, owner)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, transfers)
}

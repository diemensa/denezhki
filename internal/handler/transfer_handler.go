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

func (h *TransferHandler) HandleTransfer(c *gin.Context) {
	var req dto.TransferRequest
	transferID := uuid.New()

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewTransferResponse(transferID, false))
		return
	}

	err = h.service.PerformTransfer(c, transferID, req.FromID, req.ToID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewTransferResponse(transferID, false))
		return
	}

	c.JSON(http.StatusOK, dto.NewTransferResponse(transferID, true))

}

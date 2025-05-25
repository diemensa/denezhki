package handler

import (
	"github.com/diemensa/denezhki/internal/handler/dto"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
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

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.PerformTransfer(c, req.FromID, req.ToID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "transfer is completed",
	})

}

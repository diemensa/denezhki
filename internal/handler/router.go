package handler

import (
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupRouter(s *usecase.TransferService) *gin.Engine {
	r := gin.Default()

	r.POST("/transfer", NewTransferHandler(s).HandleTransfer)

	return r
}

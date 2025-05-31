package handler

import (
	"github.com/diemensa/denezhki/internal/handler/dto"
	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, msg string) {
	c.JSON(code, dto.ErrorResponse{Error: msg})
}

func RespondWithMessage(c *gin.Context, code int, msg string) {
	c.JSON(code, dto.MessageResponse{Message: msg})
}

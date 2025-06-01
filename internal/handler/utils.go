package handler

import (
	"fmt"
	"github.com/diemensa/denezhki/internal/handler/dto"
	"github.com/gin-gonic/gin"
)

func CheckUserMatch(c *gin.Context, username string) error {
	usernameFromToken, exists := c.Get("username")
	if !exists {
		return fmt.Errorf("username not found in context")
	}

	if username != usernameFromToken.(string) {
		return fmt.Errorf("access denied")
	}

	return nil
}

func ExtractAliasUsername(c *gin.Context) (alias, owner string) {
	return c.Param("alias"), c.Param("username")
}

func RespondWithError(c *gin.Context, code int, msg string) {
	c.JSON(code, dto.ErrorResponse{Error: msg})
}

func RespondWithMessage(c *gin.Context, code int, msg string) {
	c.JSON(code, dto.MessageResponse{Message: msg})
}

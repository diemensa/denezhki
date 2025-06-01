package handler

import (
	"fmt"
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

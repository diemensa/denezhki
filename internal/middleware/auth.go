package middleware

import (
	"net/http"
	"strings"

	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware(s *usecase.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		username, err := s.ValidateToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		c.Set("username", username)
		c.Next()
	}
}

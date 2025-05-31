package handler

import (
	"github.com/diemensa/denezhki/internal/handler/dto"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	authService *usecase.AuthService
}

func NewAuthHandler(s *usecase.AuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

func (h *AuthHandler) HandleLogin(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		RespondWithError(c, http.StatusBadRequest, "invalid request payload")
		return
	}

	token, err := h.authService.Login(c, req.Username, req.Password)
	if err != nil {
		RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{Token: token})
}

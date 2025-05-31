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

// HandleLogin godoc
// @Summary      User login
// @Description  Authenticate user and return JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        loginRequest  body      dto.LoginRequest  true  "Login credentials"
// @Success      200           {object}  dto.LoginResponse
// @Failure      400           {object}  dto.ErrorResponse  "Invalid request payload"
// @Failure      401           {object}  dto.ErrorResponse  "Unauthorized - invalid username or password"
// @Router       /auth/login [post]
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

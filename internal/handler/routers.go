package handler

import (
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupTransferRouters(r *gin.Engine, s *usecase.TransferService) {

	handler := NewTransferHandler(s)

	r.POST("/transfer", handler.HandleTransfer)
	r.GET("/transfer/:id", handler.GetTransfer)

}

func SetupUserRouters(r *gin.Engine, s *usecase.UserService) {

	handler := NewUserHandler(s)

	r.GET("/users/:username", handler.GetUserByUsername)
	r.GET("/users/:username/accounts", handler.GetUserAccounts)
	r.POST("/users", handler.CreateUser)
	r.POST("/users/:username/accounts", handler.CreateAccount)
	r.POST("/auth/login", handler.ValidatePassword)

}

func SetupAccountRouters(r *gin.Engine, s *usecase.AccountService) {

	handler := NewAccountHandler(s)

	r.GET("/accounts/:accountID", handler.GetAccByID)
	r.GET("/accounts/:accountID/user", handler.GetUserByAccID)
	r.GET("/accounts/:accountID/balance", handler.GetAccBalance)

}

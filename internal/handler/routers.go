package handler

import (
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
)

func SetupTransferRouters(r *gin.Engine, s *usecase.TransferService) {

	handler := NewTransferHandler(s)

	r.POST("/transfers", handler.HandleTransfer)
	r.GET("/transfers/:id", handler.HandleGetTransferByID)
	r.GET("/transfers", handler.HandleGetAllTransfers)

}

func SetupUserRouters(r *gin.Engine, s *usecase.UserService) {

	handler := NewUserHandler(s)

	r.GET("/users/:username", handler.HandleGetUserByUsername)
	r.GET("/users/:username/accounts", handler.HandleGetUserAccounts)
	r.POST("/users", handler.HandleCreateUser)
	r.POST("/users/:username/accounts", handler.HandleCreateAccount)
	r.POST("/auth/login", handler.HandleValidatePassword)

}

func SetupAccountRouters(r *gin.Engine, s *usecase.AccountService) {

	handler := NewAccountHandler(s)

	r.GET("/accounts/:id", handler.HandleGetAccByID)
	r.GET("/accounts/:id/balance", handler.HandleGetAccBalance)

}

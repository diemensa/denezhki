package handler

import (
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupTransferRouters(r *gin.Engine, s *usecase.TransferService) {

	handler := NewTransferHandler(s)

	r.POST("/transfers", handler.HandleTransfer)
	r.GET("/transfers/:id", handler.HandleGetTransferByID)
	r.GET("/users/:username/accounts/:alias/transfers", handler.HandleGetAllAccTransfers)
	// r.POST("/users/:username/accounts/:alias/transfers", handler.HandleTransfer) с авторизацией

}

func SetupUserAccRouters(r *gin.Engine, userServ *usecase.UserService, accountServ *usecase.AccountService) {

	handlerUser := NewUserHandler(userServ)
	handlerAccount := NewAccountHandler(accountServ)

	// User Handlers
	r.GET("/users/:username/accounts", handlerUser.HandleGetUserAccounts)
	r.POST("/users", handlerUser.HandleCreateUser)
	r.POST("/users/:username/accounts/", handlerUser.HandleCreateAccount)
	// r.POST("/auth/login", handler.HandleValidatePassword) ручка для авторизации, сделать позже

	// Account Handlers
	r.GET("/users/:username/accounts/:alias", handlerAccount.HandleGetAccByAliasOwner)
	r.GET("/users/:username/accounts/:alias/balance", handlerAccount.HandleGetAccBalance)
	r.PUT("/users/:username/accounts/:alias/balance", handlerAccount.HandleUpdateBalance)

}

func SetupDocsRouters(r *gin.Engine) {
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Router init for someday-i'll-add-this AuthService
//func SetupAuthRouters(r *gin.Engine, authServ *usecase.AuthService) {
//
//	handler := NewAuthHandler(authServ)
//
//	r.POST("/auth/login", handler.HandleLogin)
//	r.POST("/auth/me", handler.HandleGetMe)
//  r.POST("/auth/refresh", handler.HandleRefresh)
//}

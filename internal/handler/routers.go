package handler

import (
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupTransferRouters(rg *gin.RouterGroup, s *usecase.TransferService) {

	handler := NewTransferHandler(s)

	// rg.GET("/transfers/:id", handler.HandleGetTransferByID)
	rg.GET("/:username/accounts/:alias/transfers", handler.HandleGetAllAccTransfers)
	rg.POST("/:username/accounts/:alias/transfers", handler.HandleTransfer)

}

func SetupUserAccRouters(rg *gin.RouterGroup,
	userServ *usecase.UserService, accountServ *usecase.AccountService) {

	handlerUser := NewUserHandler(userServ)
	handlerAccount := NewAccountHandler(accountServ)

	// User Handlers
	rg.GET("/:username/accounts", handlerUser.HandleGetUserAccounts)
	rg.POST("/:username/accounts/", handlerUser.HandleCreateAccount)

	// Account Handlers
	rg.GET("/:username/accounts/:alias", handlerAccount.HandleGetAccByAliasOwner)
	rg.GET("/:username/accounts/:alias/balance", handlerAccount.HandleGetAccBalance)
	rg.PUT("/:username/accounts/:alias/balance", handlerAccount.HandleUpdateBalance)

}

func SetupAuthRouters(r *gin.Engine, s *usecase.AuthService, us *usecase.UserService) {
	handlerAuth := NewAuthHandler(s)
	handlerUser := NewUserHandler(us)

	r.POST("/auth/login", handlerAuth.HandleLogin)
	r.POST("/users", handlerUser.HandleCreateUser)
}

func SetupDocsRouters(r *gin.Engine) {
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

package handler

import (
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func SetupTransferRoutes(rg *gin.RouterGroup, s *usecase.TransferService) {

	handler := NewTransferHandler(s)

	rg.GET("/:username/accounts/:alias/transfers", handler.HandleGetAllAccTransfers)
	rg.POST("/:username/accounts/:alias/transfers", handler.HandleTransfer)

}

func SetupUserAccRoutes(rg *gin.RouterGroup,
	userServ *usecase.UserService, accountServ *usecase.AccountService) {

	handlerUser := NewUserHandler(userServ)
	handlerAccount := NewAccountHandler(accountServ)

	// User Handlers
	rg.GET("/:username/accounts", handlerUser.HandleGetUserAccounts)
	rg.POST("/:username/accounts/", handlerUser.HandleCreateAccount)

	// Account Handlers
	rg.GET("/:username/accounts/:alias", handlerAccount.HandleGetAccByAliasUsername)
	rg.GET("/:username/accounts/:alias/balance", handlerAccount.HandleGetAccBalance)
	rg.PUT("/:username/accounts/:alias/balance", handlerAccount.HandleUpdateBalance)

}

func SetupPublicRoutes(r *gin.Engine, s *usecase.AuthService, us *usecase.UserService, ts *usecase.TransferService) {
	handlerAuth := NewAuthHandler(s)
	handlerUser := NewUserHandler(us)
	handlerTransfer := NewTransferHandler(ts)

	r.POST("/auth/login", handlerAuth.HandleLogin)
	r.POST("/users", handlerUser.HandleCreateUser)
	r.GET("/transfers/:id", handlerTransfer.HandleGetTransferByID) // это типа блокчейн где все видят трансферы))
}

func SetupDocsRoutes(r *gin.Engine) {
	r.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/docs/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

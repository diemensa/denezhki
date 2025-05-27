package main

import (
	"github.com/diemensa/denezhki/config"
	_ "github.com/diemensa/denezhki/docs"
	"github.com/diemensa/denezhki/internal/handler"
	"github.com/diemensa/denezhki/internal/repository/postgres"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"time"
)

// @title Denezhki API
// @version 1.0
// @description Banking-like API for managing users and transfers
func main() {
	gin.SetMode(gin.ReleaseMode)
	cfg := config.LoadEnv()
	db, err := config.InitPostgres(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	rdb := config.NewRedisClient(cfg.RedisHost + ":" + cfg.RedisPort)

	if err != nil {
		log.Fatal("Couldn't connect to database:", err.Error())
	}

	transRepo := postgres.NewTransPostgresRepo(db)
	accRepo := postgres.NewAccPostgresRepo(db)
	userRepo := postgres.NewUserPostgresRepo(db)

	transferService := usecase.NewTransferService(accRepo, transRepo, rdb, 10*time.Minute)
	accountService := usecase.NewAccountService(accRepo, rdb, 10*time.Minute)
	userService := usecase.NewUserService(userRepo)

	r := gin.Default()

	handler.SetupTransferRouters(r, transferService)
	handler.SetupUserAccRouters(r, userService, accountService)

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err = r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

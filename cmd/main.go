package main

import (
	"github.com/diemensa/denezhki/config"
	_ "github.com/diemensa/denezhki/docs"
	"github.com/diemensa/denezhki/internal/handler"
	"github.com/diemensa/denezhki/internal/repository/postgres"
	redislocal "github.com/diemensa/denezhki/internal/repository/redis"
	"github.com/diemensa/denezhki/internal/usecase"
	"github.com/gin-gonic/gin"
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
	cacheRepo := redislocal.NewCacheRedisRepo(rdb)

	transferService := usecase.NewTransferService(accRepo, transRepo, cacheRepo, 10*time.Minute)
	accountService := usecase.NewAccountService(accRepo, cacheRepo, 10*time.Minute)
	userService := usecase.NewUserService(userRepo)

	r := gin.Default()

	handler.SetupTransferRouters(r, transferService)
	handler.SetupUserAccRouters(r, userService, accountService)
	handler.SetupDocsRouters(r)

	if err = r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

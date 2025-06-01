package main

import (
	"github.com/diemensa/denezhki/config"
	_ "github.com/diemensa/denezhki/docs"
	"github.com/diemensa/denezhki/internal/handler"
	"github.com/diemensa/denezhki/internal/middleware"
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
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token
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

	accountService := usecase.NewAccountService(accRepo, cacheRepo, 10*time.Minute)
	transferService := usecase.NewTransferService(accountService, transRepo, cacheRepo, 10*time.Minute)
	userService := usecase.NewUserService(userRepo)
	authService := usecase.NewAuthService(userService, cfg.SecretJWT)

	authMiddleware := middleware.JWTMiddleware(authService)

	r := gin.Default()

	authorized := r.Group("/users/")
	authorized.Use(authMiddleware)

	handler.SetupTransferRoutes(authorized, transferService)
	handler.SetupUserAccRoutes(authorized, userService, accountService)
	handler.SetupPublicRoutes(r, authService, userService, transferService)
	handler.SetupDocsRoutes(r)

	if err = r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

package main

import (
	"github.com/diemensa/denezhki/config"
	"github.com/diemensa/denezhki/internal/handler"
	"github.com/diemensa/denezhki/internal/repository/postgres"
	"github.com/diemensa/denezhki/internal/usecase"
	"log"
	"time"
)

func main() {
	cfg := config.LoadEnv()
	db, err := config.InitPostgres(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	rdb := config.NewRedisClient(cfg.RedisHost + ":" + cfg.RedisPort)

	if err != nil {
		log.Fatal("Couldn't connect to database:", err.Error())
	}

	transRepo := postgres.NewTransPostgresRepo(db)
	accRepo := postgres.NewAccPostgresRepo(db)

	transferService := usecase.NewTransferService(accRepo, transRepo)
	accountService := usecase.NewAccountService(accRepo, rdb, 10*time.Minute)

	r := handler.SetupRouter(transferService)
	r.Run(":" + cfg.Port)
}

package main

import (
	"github.com/diemensa/denezhki/config"
	"github.com/diemensa/denezhki/internal/handler"
	"github.com/diemensa/denezhki/internal/repository/postgres"
	"github.com/diemensa/denezhki/internal/usecase"
	"log"
)

func main() {
	cfg := config.LoadEnv()
	db, err := config.InitPostgres(cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	if err != nil {
		log.Fatal("Couldn't connect to database:", err.Error())
	}

	transRepo := postgres.NewTransPostgresRepo(db)
	accRepo := postgres.NewAccPostgresRepo(db)

	service := usecase.NewTransferService(accRepo, transRepo)

	r := handler.SetupRouter(service)
	r.Run(":" + cfg.Port)
}

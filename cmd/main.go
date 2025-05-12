package main

import (
	"github.com/diemensa/denezhki/config"
	"github.com/diemensa/denezhki/internal/handler"
	"github.com/diemensa/denezhki/internal/usecase"
)

func main() {
	service := usecase.NewTransferService(nil, nil)
	cfg := config.Load()

	r := handler.SetupRouter(service)
	r.Run(cfg.Port)
}

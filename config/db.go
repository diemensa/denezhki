package config

import (
	"fmt"
	"github.com/diemensa/denezhki/internal/repository/postgres/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(host, user, pw, dbname, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host, user, pw, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(
		&model.User{},
		&model.Account{},
		&model.Transaction{},
	)

	return db, nil

}

package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"url-shortening-service/internal/config"
	"url-shortening-service/internal/model"
)

func InitDB() (*gorm.DB, error) {
	var DB *gorm.DB
	var err error
	c := config.Cfg
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = DB.AutoMigrate(model.ShortURL{})
	if err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}
	return DB, nil
}

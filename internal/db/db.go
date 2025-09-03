package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"url-shortening-service/internal/config"
	"url-shortening-service/internal/model"
)

var DB *gorm.DB

func InitDB() {
	var err error
	c := config.Cfg
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	_ = DB.AutoMigrate(model.ShortURL{})
}

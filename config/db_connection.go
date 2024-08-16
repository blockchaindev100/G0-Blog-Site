package config

import (
	"fmt"
	"os"

	logger "github.com/blockchaindev100/Go-Blog-Site/logger"
	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	godotenv.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logging().Error(err)
		fmt.Println("DB LINK", dsn)
		panic(err)
	}
	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Post{}, &models.Command{}, &models.Category{})
	if err != nil {
		logger.Logging().Error(err)
		panic(err)
	}
}

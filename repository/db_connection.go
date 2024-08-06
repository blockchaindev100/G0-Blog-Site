package repository

import (
	"fmt"
	"os"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *Repository {
	godotenv.Load()
	var repository Repository
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	repository.DB = db
	return &repository
}

func (repo *Repository) AutoMigrate() {
	err := repo.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Command{}, &models.Category{})
	if err != nil {
		panic(err)
	}
}

package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository struct {
	DB     *gorm.DB
	Logger *logrus.Logger
}

type Database interface {
	User
	Post
	Command
	Category
}

func AquireDatabase(db *gorm.DB) Database {
	return &Repository{DB: db}
}

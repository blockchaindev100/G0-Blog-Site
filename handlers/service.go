package handlers

import (
	"github.com/blockchaindev100/Go-Blog-Site/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	Repo      repository.Database
	Validator *validator.Validate
	Logger    *logrus.Logger
}

func InitHandler(db repository.Database, logger *logrus.Logger) *Handlers {
	return &Handlers{db, validator.New(), logger}
}

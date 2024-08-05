package handlers

import (
	"github.com/blockchaindev100/Go-Blog-Site/repository"
	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	Repo      *repository.Repository
	Validator *validator.Validate
}

func InitHandler(db *repository.Repository) *Handlers {
	return &Handlers{db, validator.New()}
}

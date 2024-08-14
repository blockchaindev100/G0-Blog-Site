package repository

import (
	"errors"
	"time"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/blockchaindev100/Go-Blog-Site/service"
)

type User interface {
	CreateUser(*models.User) error
	GetUserByEmail(string) (models.User, error)
	GetUserById(string) (models.User, error)
}

func (repo *Repository) GetUserById(id string) (models.User, error) {
	var user models.User
	result := repo.DB.First(&user, "user_id=?", id)
	if err := result.Error; err != nil {
		repo.Logger.Error(err)
		return user, err
	}
	if result.RowsAffected == 0 {
		err := errors.New("no data found")
		repo.Logger.Error(err)
		return user, err
	}
	return user, nil
}

func (repo *Repository) CreateUser(user *models.User) error {
	user.Created_at = time.Now()
	hashedPassword, err := service.HashPassword(user.Password)
	if err != nil {
		repo.Logger.Error(err)
		return err
	}
	user.Password = hashedPassword
	user.Created_at = time.Now()
	result := repo.DB.Create(user)
	if err := result.Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := repo.DB.Where("email=?", email).First(&user)
	if err := result.Error; err != nil {
		repo.Logger.Error(err)
		return user, err
	}
	if result.RowsAffected == 0 {
		err := errors.New("no data found")
		repo.Logger.Error(err)
		return user, err
	}
	return user, nil
}

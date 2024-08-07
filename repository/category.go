package repository

import (
	"time"

	"github.com/blockchaindev100/Go-Blog-Site/models"
)

type Category interface {
	GetCategories() ([]models.Category, error)
	AddCategory(*models.Category) error
	UpdateCategory(*models.Category, string) error
	DeleteCategory(string) error
	GetCategoriesById(string) (models.Category, error)
}

func (repo *Repository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	if err := repo.DB.Find(&categories).Error; err != nil {
		repo.Logger.Error(err)
		return categories, err
	}
	return categories, nil
}

func (repo *Repository) GetCategoriesById(id string) (models.Category, error) {
	var category models.Category
	if err := repo.DB.First(&category, "category_id=?", id).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (repo *Repository) AddCategory(category *models.Category) error {
	category.Created_at = time.Now()
	if err := repo.DB.Create(category).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) UpdateCategory(category *models.Category, id string) error {
	var oldCategory models.Category
	if err := repo.DB.First(&oldCategory, "category_id=?", id).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	oldCategory.Category_Name = category.Category_Name
	oldCategory.Description = category.Description
	oldCategory.Updated_at = time.Now()
	if err := repo.DB.Save(&oldCategory).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteCategory(id string) error {
	if err := repo.DB.Delete(&models.Category{}, "category_id=?", id).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

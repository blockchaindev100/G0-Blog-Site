package repository

import (
	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/google/uuid"
)

type Post interface {
	GetPosts() ([]models.Post, error)
	CreatePosts(*models.Post, string) error
	UpdatePost(*models.Post, string) error
	DeletePost(string) error
}

func (repo *Repository) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := repo.DB.Find(&posts).Error; err != nil {
		repo.Logger.Error(err)
		return nil, err
	}
	return posts, nil
}

func (repo *Repository) CreatePosts(post *models.Post, id string) error {
	parsed_user_id, err := uuid.Parse(id)
	if err != nil {
		repo.Logger.Error(err)
		return err
	}
	post.User_id = parsed_user_id
	if err := repo.DB.Create(post).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) UpdatePost(post *models.Post, id string) error {
	var oldPost models.Post
	if err := repo.DB.First(&oldPost, "post_id=?", id).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	oldPost.Title = post.Title
	oldPost.Body = post.Body
	if err := repo.DB.Save(&oldPost).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) DeletePost(id string) error {
	if err := repo.DB.Delete(&models.Post{}, "post_id=?", id).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

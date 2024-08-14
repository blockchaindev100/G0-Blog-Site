package repository

import (
	"errors"
	"time"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/google/uuid"
)

type Post interface {
	GetPosts() ([]models.Post, error)
	CreatePosts(*models.Post, string) error
	UpdatePost(*models.Post, string) error
	DeletePost(string) error
	TotalPostCount() (int64, error)
	FirstPost() (models.Post, error)
}

func (repo *Repository) GetPosts() ([]models.Post, error) {
	var posts []models.Post
	result := repo.DB.Find(&posts)
	if err := result.Error; err != nil {
		repo.Logger.Error(err)
		return nil, err
	}
	if result.RowsAffected == 0 {
		err := errors.New("no data found")
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
	post.Created_at = time.Now()
	result := repo.DB.Create(post)
	if err := result.Error; err != nil {
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
	oldPost.Updated_at = time.Now()
	if err := repo.DB.Save(&oldPost).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) DeletePost(id string) error {
	result := repo.DB.Delete(&models.Post{}, "post_id=?", id)
	if err := result.Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	if result.RowsAffected == 0 {
		err := errors.New("no data deleted")
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) TotalPostCount() (int64, error) {
	var count int64
	if err := repo.DB.Find(&[]models.Post{}).Count(&count).Error; err != nil {
		repo.Logger.Error(err)
		return count, err
	}
	return count, nil
}

func (repo *Repository) FirstPost() (models.Post, error) {
	var post models.Post
	result := repo.DB.Order("created_at").First(&post)
	if result.Error != nil {
		repo.Logger.Error(result.Error.Error())
		return models.Post{}, errors.New("record not found")
	} else if result.RowsAffected == 0 {
		repo.Logger.Error("no rows affected")
		return models.Post{}, errors.New("no rows affected")
	}

	return post, nil
}

package repository

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/google/uuid"
)

type Command interface {
	GetCommandById(string) (models.Command, error)
	AddCommand(*models.Command, string, string) error
	DeleteCommand(string) error
	GetCommandsByPostId(string) ([]models.Command, error)
	UpdateCommand(string, string, models.Command) error
	TotalCommands() (int64, error)
}

func (repo *Repository) GetCommandById(id string) (models.Command, error) {
	var command models.Command
	if err := repo.DB.First(&command, "command_id=?", id).Error; err != nil {
		repo.Logger.Error(err)
		return command, err
	}
	return command, nil
}

func (repo *Repository) AddCommand(command *models.Command, post_id string, user_id string) error {
	parsed_post_id, err := uuid.Parse(post_id)
	if err != nil {
		repo.Logger.Error(err)
		return errors.New("invalid id")
	}
	parsed_user_id, err := uuid.Parse(user_id)
	if err != nil {
		repo.Logger.Error(err)
		return errors.New("invalid user id")
	}
	user, err := repo.GetUserById(user_id)
	if err != nil {
		repo.Logger.Error(err)
		return errors.New("invalid user")
	}
	command.UserName = user.Username
	command.Post_id = parsed_post_id
	command.User_id = parsed_user_id
	if err := repo.DB.Create(&command).Error; err != nil {
		repo.Logger.Error(err)
		return errors.New("command creation failed")
	}
	return nil
}

func (repo *Repository) UpdateCommand(id string, user_id string, command models.Command) error {
	exist_command, err := repo.GetCommandById(id)
	if err != nil {
		repo.Logger.Error(err)
		return err
	}
	if exist_command.User_id.String() != user_id {
		repo.Logger.Error(err)
		return errors.New("invalid user")
	}
	exist_command.Content = command.Content
	if err := repo.DB.Save(&exist_command).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) DeleteCommand(id string) error {
	if err := repo.DB.Delete(&models.Command{}, "command_id=?", id).Error; err != nil {
		repo.Logger.Error(err)
		return err
	}
	return nil
}

func (repo *Repository) GetCommandsByPostId(id string) ([]models.Command, error) {
	var commands []models.Command
	if err := repo.DB.Where("post_id=?", id).Find(&commands).Error; err != nil {
		repo.Logger.Error(err)
		return nil, err
	}
	return commands, nil
}

func (repo *Repository) TotalCommands() (int64, error) {
	var count int64
	if err := repo.DB.Find(&[]models.Command{}).Count(&count).Error; err != nil {
		repo.Logger.Error(err)
		return count, err
	}
	return count, nil
}

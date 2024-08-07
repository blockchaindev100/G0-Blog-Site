package handlers

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetPosts(c *fiber.Ctx) error {
	posts, err := h.Repo.GetPosts()
	if err != nil {
		h.Logger.Error(err)
		return errors.New("fetching failed")
	}
	for i := 0; i < len(posts); i++ {
		str := posts[i].Post_id.String()
		commands, err := h.Repo.GetCommandsByPostId(str)
		if err != nil {
			h.Logger.Error(err)
		} else {
			posts[i].Commands = commands
		}
		categories := []models.Category{}
		for j := 0; j < len(posts[i].Categories); j++ {
			category, err := h.Repo.GetCategoriesById(posts[i].Categories[j])
			if err != nil {
				h.Logger.Error(err)
			} else {
				categories = append(categories, category)
			}
		}
		posts[i].Category = categories
	}
	return c.JSON(posts)
}

func (h *Handlers) CreatePost(c *fiber.Ctx) error {
	var post models.Post
	id := c.Get("user_id")
	if err := c.BodyParser(&post); err != nil {
		h.Logger.Error(err)
		return errors.New("parsing failed")
	}
	if err := h.Validator.Struct(post); err != nil {
		h.Logger.Error(err)
		return errors.New("invalid payload")
	}
	if err := h.Repo.CreatePosts(&post, id); err != nil {
		h.Logger.Error(err)
		return errors.New("post creation failed")
	}
	return c.JSON(fiber.Map{
		"message": "post created successfully",
	})
}

func (h *Handlers) UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := c.BodyParser(&post); err != nil {
		h.Logger.Error(err)
		return errors.New("parsing failed")
	}
	if err := h.Validator.Struct(post); err != nil {
		h.Logger.Error(err)
		return errors.New("invalid payload")
	}
	if err := h.Repo.UpdatePost(&post, id); err != nil {
		h.Logger.Error(err)
		return err
	}
	return c.JSON(fiber.Map{
		"message": "updated successful",
	})
}

func (h *Handlers) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Repo.DeletePost(id); err != nil {
		h.Logger.Error(err)
		return errors.New("post deletion failed")
	}
	return c.JSON(fiber.Map{
		"message": "post deleted successful",
	})
}

package handlers

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) GetCategories(c *fiber.Ctx) error {
	categories, err := h.Repo.GetCategories()
	if err != nil {
		h.Logger.Error(err)
		return errors.New("fetching failed")
	}
	return c.JSON(categories)
}

func (h *Handlers) AddCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		h.Logger.Error(err)
		return errors.New("fetching payload failed")
	}
	if err := h.Validator.Struct(category); err != nil {
		h.Logger.Error(err)
		return errors.New("invalid payload")
	}
	err := h.Repo.AddCategory(&category)
	if err != nil {
		h.Logger.Error(err)
		return errors.New("category creation failed")
	}
	return c.JSON(fiber.Map{
		"message": "created successful",
	})
}

func (h *Handlers) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		h.Logger.Error(err)
		return errors.New("fetching payload failed")
	}
	if err := h.Validator.Struct(category); err != nil {
		h.Logger.Error(err)
		return errors.New("invalid payload")
	}
	if err := h.Repo.UpdateCategory(&category, id); err != nil {
		h.Logger.Error(err)
		return errors.New("category update failed")
	}
	return c.JSON(fiber.Map{
		"message": "updated successful",
	})
}

func (h *Handlers) DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Repo.DeleteCategory(id); err != nil {
		h.Logger.Error(err)
		return errors.New("deletion failed")
	}
	return c.JSON(fiber.Map{
		"message": "deleted successful",
	})
}

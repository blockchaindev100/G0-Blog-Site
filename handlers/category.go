package handlers

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get Category
// @Schemes http
// @Description Get Category from the blog site
// @Tags Category
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /category [get]
func (h *Handlers) GetCategories(c *fiber.Ctx) error {
	categories, err := h.Repo.GetCategories()
	if err != nil {
		h.Logger.Error(err)
		return errors.New("fetching failed")
	}
	return c.JSON(categories)
}

// @Summary Add Category
// @Schemes http
// @Description Add Category in the blog site
// @Tags Category
// @Accept json
// @Produce json
// @Param user body models.Category true "Category details"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /category [post]
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

// @Summary Update Category
// @Schemes http
// @Description Update Category in the blog site
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param user body models.Category true "Category details"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /category [put]
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

// @Summary Delete Category
// @Schemes http
// @Description Delete Category in the blog site
// @Tags Category
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /category [Delete]
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

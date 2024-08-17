package handlers

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/gofiber/fiber/v2"
)

// @Summary Add Command
// @Schemes http
// @Description Add Command for the post in the blog site
// @Tags Command
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param user body models.Command true "Command details"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /command/{id} [post]
func (h *Handlers) AddCommand(c *fiber.Ctx) error {
	var command models.Command
	post_id := c.Params("id")
	user_id := c.Get("user_id")
	if err := c.BodyParser(&command); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Validator.Struct(command); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Repo.AddCommand(&command, post_id, user_id); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{
		"message": "command added successful",
	})
}

// @Summary Update Command
// @Schemes http
// @Description Update Command for the post in the blog site
// @Tags Command
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Param user body models.Command true "Command details"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /command/{id} [put]
func (h *Handlers) UpdateCommand(c *fiber.Ctx) error {
	id := c.Params("id")
	user_id := c.Get("user_id")
	var command models.Command
	if err := c.BodyParser(&command); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	err := h.Repo.UpdateCommand(id, user_id, command)
	if err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{
		"message": "updated successful",
	})
}

// @Summary Delete Command
// @Schemes http
// @Description Delete Command for the post in the blog site
// @Tags Command
// @Accept json
// @Produce json
// @Param id path string true "Command ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.Response
// @Router /command/{id} [delete]
func (h *Handlers) DeleteCommand(c *fiber.Ctx) error {
	id := c.Params("id")
	user_id := c.Get("user_id")
	command, err := h.Repo.GetCommandById(id)
	if err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	if user_id != command.User_id.String() {
		h.Logger.Error(errors.New("unauthorized"))
		return fiber.ErrUnauthorized
	}
	if err := h.Repo.DeleteCommand(id); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(fiber.Map{
		"message": "command deleted successful",
	})
}

// @Summary Get Command By Post ID
// @Schemes http
// @Description Get Command for the post in the blog site
// @Tags Command
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Security ApiKeyAuth
// @Success 200 {object} []models.Command
// @Router /command/{id} [get]
func (h *Handlers) GetCommandsByPostId(c *fiber.Ctx) error {
	id := c.Params("id")
	commands, err := h.Repo.GetCommandsByPostId(id)
	if err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	return c.JSON(commands)
}

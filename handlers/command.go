package handlers

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) AddCommand(c *fiber.Ctx) error {
	var command models.Command
	post_id := c.Params("id")
	user_id := c.Get("user_id")
	if err := c.BodyParser(&command); err != nil {
		h.Logger.Error(err)
		return errors.New("parsing failed")
	}
	if err := h.Validator.Struct(command); err != nil {
		h.Logger.Error(err)
		return errors.New("invalid payload")
	}
	if err := h.Repo.AddCommand(&command, post_id, user_id); err != nil {
		h.Logger.Error(err)
		return err
	}
	return c.JSON(fiber.Map{
		"message": "command added successful",
	})
}

func (h *Handlers) UpdateCommand(c *fiber.Ctx) error {
	id := c.Params("id")
	user_id := c.Get("user_id")
	var command models.Command
	if err := c.BodyParser(&command); err != nil {
		return errors.New("parsing failed")
	}
	err := h.Repo.UpdateCommand(id, user_id, command)
	if err != nil {
		h.Logger.Error(err)
		return errors.New("updation failed")
	}
	return c.JSON(fiber.Map{
		"message": "updated successful",
	})
}

func (h *Handlers) DeleteCommand(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.Repo.DeleteCommand(id); err != nil {
		h.Logger.Error(err)
		return errors.New("command deletion failed")
	}
	return c.JSON(fiber.Map{
		"message": "command deleted successful",
	})
}

func (h *Handlers) GetCommandsByPostId(c *fiber.Ctx) error {
	id := c.Params("id")
	commands, err := h.Repo.GetCommandsByPostId(id)
	if err != nil {
		h.Logger.Error(err)
		return errors.New("fetching failed")
	}
	return c.JSON(commands)
}

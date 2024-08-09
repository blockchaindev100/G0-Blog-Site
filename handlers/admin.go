package handlers

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) Overview(c *fiber.Ctx) error {
	var overview models.Overview
	if val, err := h.Repo.TotalPostCount(); err != nil {
		h.Logger.Error(err)
		return errors.New("failed to get total post")
	} else {
		overview.Total_Posts = val
	}
	if val, err := h.Repo.TotalCommands(); err != nil {
		h.Logger.Error(err)
		return err
	} else {
		overview.Total_Commands = val
	}
	if val, err := h.Repo.FirstPost(); err != nil {
		h.Logger.Error(err)
		return err
	} else {
		overview.First_Blog = val.Created_at
	}
	return c.JSON(&overview)
}

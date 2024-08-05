package handlers

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/blockchaindev100/Go-Blog-Site/service"
	"github.com/gofiber/fiber/v2"
)

func (h *Handlers) Signup(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return errors.New("invalid payload")
	}
	if err := h.Validator.Struct(user); err != nil {
		return errors.New("invalid payload")
	}
	if err := h.Repo.CreateUser(&user); err != nil {
		return errors.New("user creation failed")
	}
	return c.JSON(fiber.Map{
		"message": "user created",
	})
}

func (h *Handlers) Login(c *fiber.Ctx) error {
	var login models.Login
	if err := c.BodyParser(&login); err != nil {
		return errors.New("invalid payload")
	}
	if err := h.Validator.Struct(login); err != nil {
		return errors.New("invalid payload")
	}
	user, err := h.Repo.GetUserByEmail(login.Email)
	if err != nil {
		return errors.New("user not exists")
	}
	isMatch := service.ComparePassword(login.Password, user.Password)
	if !isMatch {
		return errors.New("invalid credentials")
	}
	token, err := service.CreateToken(&user)
	if err != nil {
		return errors.New("authentication failed")
	}
	c.Response().Header.Add("Authentication", token)
	return c.JSON(fiber.Map{
		"message": "login successful",
	})
}

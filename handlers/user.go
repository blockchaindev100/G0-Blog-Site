package handlers

import (
	"net/http"

	"github.com/blockchaindev100/Go-Blog-Site/models"
	"github.com/blockchaindev100/Go-Blog-Site/service"
	"github.com/gofiber/fiber/v2"
)

// @Summary User SignUp
// @Schemes http
// @Description Create a user account
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} models.Response
// @Router /signup [post]
func (h *Handlers) Signup(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Validator.Struct(user); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Repo.CreateUser(&user); err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	c.SendStatus(http.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "user created",
	})
}

// @Summary User Login
// @Schemes http
// @Description Login into the blog site
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.Login true "User details"
// @Success 200 {object} models.Response
// @Router /login [post]
func (h *Handlers) Login(c *fiber.Ctx) error {
	var login models.Login
	if err := c.BodyParser(&login); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	if err := h.Validator.Struct(login); err != nil {
		h.Logger.Error(err)
		return fiber.ErrBadRequest
	}
	user, err := h.Repo.GetUserByEmail(login.Email)
	if err != nil {
		h.Logger.Error(err)
		return fiber.ErrNotFound
	}
	isMatch := service.ComparePassword(login.Password, user.Password)
	if !isMatch {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	token, err := service.CreateToken(&user)
	if err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	uuidToken := service.GenerateUUIDString()
	err = service.SetData(uuidToken, token)
	if err != nil {
		h.Logger.Error(err)
		return fiber.ErrInternalServerError
	}
	c.Response().Header.Add("Authorization", uuidToken)
	return c.JSON(fiber.Map{
		"message": "login successful",
	})
}

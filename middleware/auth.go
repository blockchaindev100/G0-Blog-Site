package middleware

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/service"
	"github.com/gofiber/fiber/v2"
)

func UserAuth(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	is_admin, user_id, err := service.VerifyToken(token[7:])
	if err != nil {
		return errors.New("authentication failed")
	}
	admin := "false"
	if is_admin {
		admin = "true"
	}
	c.Request().Header.Set("user_id", user_id)
	c.Request().Header.Set("is_admin", admin)
	return c.Next()
}

func AdminAuth(c *fiber.Ctx) error {
	admin := c.Get("is_admin")
	if admin != "true" {
		return errors.New("user is not a admin")
	}
	return c.Next()
}

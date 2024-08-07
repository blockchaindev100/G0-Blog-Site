package middleware

import (
	"errors"

	"github.com/blockchaindev100/Go-Blog-Site/service"
	"github.com/gofiber/fiber/v2"
)

func (midd *Middleware) UserAuth(c *fiber.Ctx) error {
	uuidToken := c.Get("Authorization")
	token, err := service.GetData(uuidToken)
	if token == "" {
		midd.Logger.Error(err)
		return errors.New("invalid token")
	}
	if err != nil {
		midd.Logger.Error(err)
		return errors.New("token is expired")
	}
	is_admin, user_id, err := service.VerifyToken(token)
	if err != nil {
		midd.Logger.Error(err)
		return errors.New("not authorized")
	}
	admin := "false"
	if is_admin {
		admin = "true"
	}
	c.Request().Header.Set("user_id", user_id)
	c.Request().Header.Set("is_admin", admin)
	return c.Next()
}

func (midd *Middleware) AdminAuth(c *fiber.Ctx) error {
	admin := c.Get("is_admin")
	if admin != "true" {
		midd.Logger.Error(errors.New("user is not a admin"))
		return errors.New("user is not a admin")
	}
	return c.Next()
}

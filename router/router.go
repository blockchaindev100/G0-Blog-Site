package routers

import (
	"github.com/blockchaindev100/Go-Blog-Site/handlers"
	"github.com/blockchaindev100/Go-Blog-Site/middleware"
	"github.com/blockchaindev100/Go-Blog-Site/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func InitRouter(app *fiber.App, db repository.Database, logger *logrus.Logger) {
	handler := handlers.InitHandler(db, logger)
	middleware := middleware.AcquireMiddleware(logger)
	app.Post("/signup", handler.Signup)
	app.Post("/login", handler.Login)
	blog := app.Group("/blog", middleware.UserAuth)
	{
		blog.Get("/", handler.GetPosts)
		blog.Use(middleware.AdminAuth)
		blog.Post("/", handler.CreatePost)
		blog.Put("/:id", handler.UpdatePost)
		blog.Delete("/:id", handler.DeletePost)
	}
	category := app.Group("/category", middleware.UserAuth)
	{
		category.Get("/", handler.GetCategories)
		category.Use(middleware.AdminAuth)
		category.Post("/", handler.AddCategory)
		category.Put("/:id", handler.UpdateCategory)
		category.Delete("/:id", handler.DeleteCategory)
	}
	command := app.Group("/command", middleware.UserAuth)
	{
		command.Get("/:id", handler.GetCommandsByPostId)
		command.Post("/:id", handler.AddCommand)
		command.Delete("/:id", handler.DeleteCommand)
	}
}

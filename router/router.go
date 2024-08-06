package routers

import (
	"github.com/blockchaindev100/Go-Blog-Site/handlers"
	"github.com/blockchaindev100/Go-Blog-Site/middleware"
	"github.com/blockchaindev100/Go-Blog-Site/repository"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, db *repository.Repository) {
	handler := handlers.InitHandler(db)
	app.Post("/signup", handler.Signup)
	app.Post("/login", handler.Login)
	app.Use(middleware.UserAuth)
	category := app.Group("/category")
	{
		category.Get("/", handler.GetCategories)
		category.Use(middleware.AdminAuth)
		category.Post("/", handler.AddCategory)
		category.Put("/:id", handler.UpdateCategory)
		category.Delete("/:id", handler.DeleteCategory)
	}
	blog := app.Group("/blog")
	{
		blog.Get("/", handler.GetPosts)
		blog.Use(middleware.AdminAuth)
		blog.Post("/", handler.CreatePost)
		blog.Put("/:id", handler.UpdatePost)
		blog.Delete("/:id", handler.DeletePost)
	}
	command := app.Group("/command")
	{
		command.Get("/:id", handler.GetCommandsByPostId)
		command.Post("/:id", handler.AddCommand)
		command.Delete("/:id", handler.DeleteCommand)
	}
}

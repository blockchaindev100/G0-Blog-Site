package routers

import (
	"github.com/blockchaindev100/Go-Blog-Site/handlers"
	"github.com/blockchaindev100/Go-Blog-Site/repository"
	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, db *repository.Repository) {
	handler := handlers.InitHandler(db)
	app.Post("/signup", handler.Signup)
	app.Post("/login", handler.Login)
	category := app.Group("/category")
	category.Get("/", handler.GetCategories)
	category.Post("/", handler.AddCategory)
	category.Put("/:id", handler.UpdateCategory)
	category.Delete("/:id", handler.DeleteCategory)
}

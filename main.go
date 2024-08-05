package main

import (
	"log"

	"github.com/blockchaindev100/Go-Blog-Site/repository"
	routers "github.com/blockchaindev100/Go-Blog-Site/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := repository.InitDB()
	db.AutoMigrate()
	app := fiber.New()
	routers.InitRouter(app, db)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Failed to start the server\n", err)
	}
}

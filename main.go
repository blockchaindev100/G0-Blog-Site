package main

import (
	"log"

	"github.com/blockchaindev100/Go-Blog-Site/config"
	logger "github.com/blockchaindev100/Go-Blog-Site/logger"
	"github.com/blockchaindev100/Go-Blog-Site/repository"
	routers "github.com/blockchaindev100/Go-Blog-Site/router"
	"github.com/blockchaindev100/Go-Blog-Site/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	service.RedisInit()
	db := config.InitDB()
	app := fiber.New()
	logger := logger.Logging()
	routers.InitRouter(app, repository.AquireDatabase(db), logger)
	err := app.Listen(":8080")
	if err != nil {
		logger.Error(err)
		log.Fatal("Failed to start the server\n", err)
	}
}

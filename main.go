package main

import (
	"pinjemlah-fiber/configs"
	"pinjemlah-fiber/databases"
	"pinjemlah-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	configs.LoadEnv()
	databases.Connection()
	app := fiber.New()
	app.Use(logger.New())
	configs.SetupCORS(app)
	app.Use(configs.RateLimiter())
	routes.SetupRoutes(app)

	app.Listen(":8888")
}

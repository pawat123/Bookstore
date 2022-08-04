package main

import (
	"Bookstore/Routes"
	"Bookstore/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Config
	config.ConnectSQL()

	// Router
	Routes.Routes(app)

	// Run
	log.Fatal(app.Listen(":3000"))
}

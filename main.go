package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	os.Mkdir("./static", os.ModePerm)

	app.Static("/", "./static")

	app.Use(cors.New())
	app.Use(favicon.New())
	app.Use(logger.New())

	app.Listen(":3000")
}

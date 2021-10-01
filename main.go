package main

import (
	"log"
	"os"

	"server_deploy/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := SetupFiber()
	app.Listen(os.Getenv("APP_PORT"))
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file, ", err)
	}
}

func SetupFiber() *fiber.App {
	app := fiber.New(fiber.Config{})
	app.Use(logger.New(logger.ConfigDefault))
	route.SetupRoutes(app)
	return app
}

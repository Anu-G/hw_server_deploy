package route

import (
	"server_deploy/api"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("")
	route.Get("/", api.Home)
	route.Get("/follower/:username", api.Follower)
	route.Get("/:userid/detail", api.Userid)
}

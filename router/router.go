package router

import (
	"github.com/forceki/sso-go/src/apps"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	apps.AppController(api)
}

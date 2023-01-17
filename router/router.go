package router

import (
	"github.com/forceki/sso-go/src/apps"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB, app *fiber.App) {
	api := app.Group("/api")
	apps.AppRouter(db, api)
}

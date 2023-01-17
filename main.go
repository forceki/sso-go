package main

import (
	"github.com/forceki/sso-go/database"
	"github.com/forceki/sso-go/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	db := database.ConnectDB()

	router.SetupRoutes(db, app)

	// Listen on PORT 3000
	app.Listen(":3000")
}

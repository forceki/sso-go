package apps

import "github.com/gofiber/fiber/v2"

func AppController(router fiber.Router) {
	app := router.Group("apps")

	app.Get("/", GetApps)
	app.Get("/by", GetAppsById)

}

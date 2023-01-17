package apps

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AppRouter(db *gorm.DB, router fiber.Router) {
	repository := NewRepositoryApps(db)
	service := NewServiceApps(repository)
	controller := NewAppController(service)

	app := router.Group("apps")

	app.Get("/", controller.GetAll)
	app.Get("/module", controller.GetById)
	app.Post("/", controller.Create)
	app.Put("/", controller.UpdateApp)
	app.Get("/by-code/:code", controller.FindByCode)
	app.Get("/total", controller.CountApps)

}

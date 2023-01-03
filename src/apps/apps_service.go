package apps

import (
	"github.com/forceki/sso-go/database"
	"github.com/gofiber/fiber/v2"
)

func GetApps(c *fiber.Ctx) error {
	var get []Get

	err := database.DB.Model(&Apps{}).Find(&get)

	if err.Error != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Found", "data": get})
}

func GetAppsById(c *fiber.Ctx) error {
	id := c.Query("id")

	var get []Get

	err := database.DB.Model(&Apps{}).First(&get, "app_id = ?", id)

	if err.Error != nil {
		return c.SendStatus(fiber.StatusNoContent)
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Found", "data": get})
}

package apps

import (
	"github.com/forceki/sso-go/src/apps/entities"
	"github.com/gofiber/fiber/v2"
)

type appsController struct {
	serviceApps ServiceApps
}

func NewAppController(serviceApps ServiceApps) *appsController {
	return &appsController{serviceApps: serviceApps}
}

func (a *appsController) GetAll(f *fiber.Ctx) error {

	apps, err := a.serviceApps.FindAllApp()

	// var data []GetApp

	// for _, b := range apps {
	// 	app := GetApp{
	// 		AppId:      b.AppId,
	// 		AppName:    b.AppName,
	// 		AppUrl:     b.AppUrl,
	// 		AppDesc:    b.AppDesc,
	// 		SecretCode: b.SecretCode,
	// 	}

	// 	data = append(data, app)
	// }

	if err != nil {

		return f.Status(501).JSON(fiber.Map{
			"status":  0,
			"message": "not found",
			"data":    nil,
		})

	}

	return f.Status(201).JSON(fiber.Map{
		"status":  1,
		"message": "found",
		"data":    apps,
	})

}

func (a *appsController) GetById(f *fiber.Ctx) error {
	Id := f.Query("app_id")

	apps, err := a.serviceApps.FindById(Id)

	if err != nil {

		return f.Status(501).JSON(fiber.Map{
			"status":  0,
			"message": "not found",
			"data":    nil,
		})

	}

	return f.Status(201).JSON(fiber.Map{
		"status":  1,
		"message": "found",
		"data":    apps,
	})

}

func (a *appsController) Create(f *fiber.Ctx) error {

	var app entities.AppRequest
	err := f.BodyParser(&app)
	if err != nil {
		return err
	}

	res, err := a.serviceApps.Create(app)

	if err != nil {

		return f.Status(501).JSON(fiber.Map{
			"status":  0,
			"message": "not found",
			"data":    nil,
		})

	}

	return f.Status(201).JSON(fiber.Map{
		"status":  1,
		"message": "found",
		"data":    res,
	})
}

func (a *appsController) FindByCode(f *fiber.Ctx) error {
	Code := f.Params("code")

	data, err := a.serviceApps.FindByCode(Code)

	if err != nil {
		return f.Status(501).JSON(fiber.Map{
			"status":  0,
			"message": "not found",
			"data":    nil,
		})
	}

	return f.Status(200).JSON(fiber.Map{
		"status":  1,
		"message": "found",
		"data":    data,
	})
}

func (a *appsController) CountApps(f *fiber.Ctx) error {
	data, err := a.serviceApps.CountApps()

	if err != nil {
		return f.Status(501).JSON(fiber.Map{
			"status":  0,
			"message": "not found",
			"data":    nil,
		})
	}

	return f.Status(200).JSON(fiber.Map{
		"status":  1,
		"message": "found",
		"data":    data,
	})
}

func (a *appsController) UpdateApp(f *fiber.Ctx) error {
	Id := f.Query("app_id")
	var data entities.UpdateApp

	err := f.BodyParser(&data)
	if err != nil {
		return err
	}

	app, err := a.serviceApps.UpdateApps(Id, data)

	if err != nil {
		return f.Status(501).JSON(fiber.Map{
			"status":  0,
			"message": err,
			"data":    nil,
		})
	}

	return f.Status(201).JSON(fiber.Map{
		"status":  1,
		"message": "found",
		"data":    app,
	})
}

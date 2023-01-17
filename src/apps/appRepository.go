package apps

import (
	"github.com/forceki/sso-go/src/apps/entities"
	"gorm.io/gorm"
)

type RepositoryApps interface {
	FindAllApp() ([]entities.Apps, error)
	FindById(Id string) (entities.Apps, error)
	FindModuleByApp(Id string) ([]entities.Modules, error)
	CreateApps(Data entities.CreateApp) (entities.CreateApp, error)
	FindAppByCode(Code string) (entities.Apps, error)
	CountApps() (int64, error)
	UpdateApp(Id string, App entities.Apps) (entities.Apps, error)
}

type repositoryApps struct {
	db *gorm.DB
}

func NewRepositoryApps(db *gorm.DB) *repositoryApps {
	return &repositoryApps{db: db}
}

func (r *repositoryApps) FindAllApp() ([]entities.Apps, error) {
	var apps []entities.Apps
	err := r.db.Find(&apps).Error

	return apps, err

}

func (r *repositoryApps) FindById(Id string) (entities.Apps, error) {
	var app entities.Apps

	err := r.db.Where("app_id = ?", Id).First(&app).Error
	return app, err
}

func (r *repositoryApps) FindModuleByApp(Id string) ([]entities.Modules, error) {
	var module []entities.Modules

	err := r.db.Where("app_id = ?", Id).Find(&module).Error

	return module, err
}

func (r *repositoryApps) CreateApps(Data entities.CreateApp) (entities.CreateApp, error) {

	apps := Data.App
	app := entities.Apps{
		AppId:          apps.AppId,
		AppName:        apps.AppName,
		AppUrl:         apps.AppUrl,
		AppUrlCallback: apps.AppUrlCallback,
		AppDesc:        apps.AppDesc,
		SecretCode:     apps.SecretCode,
	}

	var module []entities.Modules

	for _, item := range Data.Module {
		moodul := entities.Modules{
			ModuleId:   item.ModuleId,
			AppId:      apps.AppId,
			ModuleName: item.ModuleName,
		}

		module = append(module, moodul)
	}

	tx := r.db.Begin()

	err := tx.Create(&app).Error

	if err != nil {
		tx.Rollback()
	}

	err = tx.Create(&module).Error
	if err != nil {
		tx.Rollback()
	}

	err = tx.Commit().Error

	return Data, err
}

func (r *repositoryApps) FindAppByCode(Code string) (entities.Apps, error) {
	var apps entities.Apps

	err := r.db.Where("secret_code = ?", Code).Find(&apps).Error

	return apps, err
}

func (r *repositoryApps) CountApps() (int64, error) {
	var count int64

	err := r.db.Table("tbm_apps").Count(&count).Error

	return count, err
}

func (r *repositoryApps) UpdateApp(Id string, Data entities.Apps) (entities.Apps, error) {
	apps := Data
	err := r.db.Model(&apps).Where("app_id = ?", Id).Updates(
		entities.Apps{
			AppName:        apps.AppName,
			AppUrl:         apps.AppUrl,
			AppDesc:        apps.AppDesc,
			AppUrlCallback: apps.AppUrlCallback,
		}).Error

	return Data, err
}

package apps

import (
	"encoding/hex"
	"time"

	"crypto/md5"

	"github.com/forceki/sso-go/src/apps/entities"
	"github.com/google/uuid"
)

type ServiceApps interface {
	FindAllApp() ([]entities.Apps, error)
	FindById(Id string) (*entities.SendApps, error)
	Create(Data entities.AppRequest) (*entities.AppRequest, error)
	FindByCode(Code string) (*entities.Apps, error)
	CountApps() (*int64, error)
	UpdateApps(Id string, Data entities.UpdateApp) (*entities.UpdateApp, error)
}

type serviceApps struct {
	repositoryApps RepositoryApps
}

func NewServiceApps(repositoryApps RepositoryApps) *serviceApps {
	return &serviceApps{repositoryApps}
}

func (s *serviceApps) FindAllApp() ([]entities.Apps, error) {
	apps, err := s.repositoryApps.FindAllApp()

	return apps, err
}

func (s *serviceApps) FindById(Id string) (*entities.SendApps, error) {
	apps, err := s.repositoryApps.FindById(Id)

	if err != nil {
		return nil, err
	}

	module, err := s.repositoryApps.FindModuleByApp(Id)

	if err != nil {
		return nil, err
	}

	data := entities.SendApps{
		Apps:   apps,
		Module: module,
	}

	return &data, nil
}

func (s *serviceApps) Create(Data entities.AppRequest) (*entities.AppRequest, error) {
	id := uuid.New()

	var secret = []byte(Data.SecretCode)
	pass := md5.Sum(secret)

	apps := entities.Apps{
		AppId:          id.String(),
		AppName:        Data.AppName,
		AppUrl:         Data.AppUrl,
		AppDesc:        Data.AppDesc,
		SecretCode:     hex.EncodeToString(pass[:]),
		AppUrlCallback: Data.AppUrlCallback,
		CreatedAt:      time.Now(),
	}
	var module []entities.Modules
	for _, i := range Data.Modules {
		moduleId := uuid.New()
		item := entities.Modules{
			ModuleId:   moduleId.String(),
			ModuleName: i.ModuleName,
			AppId:      id.String(),
		}

		module = append(module, item)
	}

	var data entities.CreateApp

	data.Module = module
	data.App = apps

	_, err := s.repositoryApps.CreateApps(data)

	if err != nil {
		return nil, err
	}

	return &Data, nil
}

func (s *serviceApps) FindByCode(Code string) (*entities.Apps, error) {
	apps, err := s.repositoryApps.FindAppByCode(Code)

	if err != nil {
		return nil, err
	}

	return &apps, nil
}

func (s *serviceApps) CountApps() (*int64, error) {
	data, err := s.repositoryApps.CountApps()

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *serviceApps) UpdateApps(Id string, Data entities.UpdateApp) (*entities.UpdateApp, error) {

	app := entities.Apps{
		AppName:        Data.AppName,
		AppUrl:         Data.AppUrl,
		AppUrlCallback: Data.AppUrlCallback,
		AppDesc:        Data.AppDesc,
	}
	_, err := s.repositoryApps.UpdateApp(Id, app)

	if err != nil {
		return nil, err
	}

	return &Data, nil
}

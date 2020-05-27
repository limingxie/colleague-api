package colleagues

import (
	"context"
	"errors"
	"time"

	"github.com/hublabs/colleague-api/factory"
	"github.com/hublabs/common/api"
)

type App struct {
	Id         int64     `json:"id"`
	TenantCode string    `json:"tenantCode"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	Enable     bool      `json:"enable"`
	CreatedAt  time.Time `json:"-" xorm:"created"`
	UpdatedAt  time.Time `json:"-" xorm:"updated"`
}

type AppJsonView struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func (app *App) CreateApp(ctx context.Context) error {
	a, err := App{}.GetByCode(ctx, app.TenantCode, app.Code)
	if err != nil {
		return err
	}

	if a.Id != 0 {
		return api.ErrorInvalidFields.New(errors.New("appCode already exists"))
	}

	_, err = factory.DB(ctx).Insert(app)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) UpdateApp(ctx context.Context) error {
	s, err := App{}.GetById(ctx, app.TenantCode, app.Id)
	if err != nil {
		return err
	}

	if s.Id == 0 {
		return api.ErrorHasExisted.New(errors.New("Invalid app(appId)"))
	}

	if _, err := factory.DB(ctx).ID(app.Id).
		Cols("name").
		Update(app); err != nil {
		return err
	}

	return nil
}

func (App) DeleteApp(ctx context.Context, tenantCode string, appId int64) error {
	s, err := App{}.GetById(ctx, tenantCode, appId)
	if err != nil {
		return err
	}

	if s.Id == 0 {
		return api.ErrorHasExisted.New(errors.New("Invalid app(appId)"))
	}

	s.Enable = false
	if _, err := factory.DB(ctx).ID(s.Id).Cols("enable").
		Update(&s); err != nil {
		return err
	}

	return nil
}

func (App) GetByCode(ctx context.Context, tenantCode, appCode string) (App, error) {
	var app App
	if has, err := factory.DB(ctx).Table("app").
		Where("app.code = ? ", appCode).
		And("app.tenant_code = ? ", tenantCode).
		And("app.enable = ? ", true).
		Get(&app); err != nil {
		return App{}, err
	} else if has {
		return app, nil
	} else {
		return App{}, nil
	}
}

func (App) GetById(ctx context.Context, tenantCode string, appId int64) (App, error) {
	var app App
	if has, err := factory.DB(ctx).
		Where("app.id = ? ", appId).
		And("app.tenant_code = ? ", tenantCode).
		And("app.enable = ? ", true).
		Get(&app); err != nil {
		return App{}, err
	} else if has {
		return app, nil
	} else {
		return App{}, nil
	}

}

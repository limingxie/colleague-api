package controllers

import (
	"net/http"

	"github.com/hublabs/colleague-api/colleagues"
	"github.com/hublabs/colleague-api/factory"
	"github.com/hublabs/colleague-api/tenants"
	"github.com/labstack/echo"
)

type MigrationController struct{}

func (c MigrationController) Init(g *echo.Echo) {
	g.POST("/seed", c.Seed)
}

func (MigrationController) Seed(c echo.Context) error {
	if err := colleagues.DropTables(factory.XormEngine()); err != nil {
		return renderFail(c, err)
	}

	if err := colleagues.Init(factory.XormEngine()); err != nil {
		return renderFail(c, err)
	}
	if err := colleagues.Seed(factory.XormEngine()); err != nil {
		return renderFail(c, err)
	}

	if err := tenants.DropTables(factory.XormEngine()); err != nil {
		return renderFail(c, err)
	}
	if err := tenants.Init(factory.XormEngine()); err != nil {
		return renderFail(c, err)
	}
	if err := tenants.Seed(factory.XormEngine()); err != nil {
		return renderFail(c, err)
	}

	return renderSucc(c, http.StatusOK, nil)
}

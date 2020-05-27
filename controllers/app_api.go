package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/hublabs/colleague-api/colleagues"
	"github.com/hublabs/common/api"
	"github.com/hublabs/common/auth"
	"github.com/labstack/echo"
)

type AppApiController struct {
}

func (c AppApiController) Init(g *echo.Echo) {
	g.GET("/v1/apps/:id", c.GetById)
	g.POST("/v1/apps", c.PostApp)
	g.PUT("/v1/apps", c.PutApp)
	g.DELETE("/v1/apps/:id", c.DeleteApp)
}

func (c AppApiController) GetById(ctx echo.Context) error {
	appId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || appId == 0 {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())
	tenantCode := userClaim.TenantCode
	if len(strings.TrimSpace(userClaim.TenantCode)) == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	/*=======================> Main Function GetAppAndBrandsByAppId <=======================*/
	result, err := colleagues.App{}.GetById(ctx.Request().Context(), tenantCode, appId)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, result)
}

func (c AppApiController) PostApp(ctx echo.Context) error {
	var app colleagues.App
	if err := ctx.Bind(&app); err != nil {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())

	tenantCode := userClaim.TenantCode
	if len(strings.TrimSpace(userClaim.TenantCode)) == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	app.TenantCode = tenantCode

	/*=======================> Main Function app.CreateApp <=======================*/
	if err := app.CreateApp(ctx.Request().Context()); err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, app)
}

func (c AppApiController) PutApp(ctx echo.Context) error {
	var app colleagues.App
	if err := ctx.Bind(&app); err != nil {

		return renderFail(ctx, api.ErrorParameter.New(err))
	}
	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())

	tenantCode := userClaim.TenantCode
	if len(strings.TrimSpace(userClaim.TenantCode)) == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	app.TenantCode = tenantCode

	/*=======================> Main Function app.UpdateApp <=======================*/
	if err := app.UpdateApp(ctx.Request().Context()); err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}
	return renderSucc(ctx, http.StatusOK, app)
}

func (c AppApiController) DeleteApp(ctx echo.Context) error {
	appId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || appId == 0 {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())

	tenantCode := userClaim.TenantCode
	if len(strings.TrimSpace(userClaim.TenantCode)) == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	/*=======================> Main Function app.DeleteApp <=======================*/
	if err := (colleagues.App{}).DeleteApp(ctx.Request().Context(), tenantCode, appId); err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}
	return renderSucc(ctx, http.StatusOK, appId)
}

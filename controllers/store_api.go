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

type StoreApiController struct {
}

func (c StoreApiController) Init(g *echo.Echo) {
	g.GET("/v1/store/:id", c.GetStoreAndBrandsByStoreId)
	g.POST("/v1/store", c.PostStore)
	g.PUT("/v1/store", c.PutStore)
	g.DELETE("/v1/stores/:id", c.DeleteStore)
}

func (c StoreApiController) GetStoreAndBrandsByStoreId(ctx echo.Context) error {
	storeId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || storeId == 0 {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	/*=======================> Main Function Colleague.Authentication <=======================*/
	result, err := colleagues.Store{}.GetStoreAndBrandsByStoreId(ctx.Request().Context(), storeId)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, result)
}

func (c StoreApiController) PostStore(ctx echo.Context) error {
	var store colleagues.Store
	if err := ctx.Bind(&store); err != nil {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())

	tenantCode := userClaim.TenantCode
	if len(strings.TrimSpace(userClaim.TenantCode)) == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	store.TenantCode = tenantCode

	/*=======================> Main Function store.CreateStore <=======================*/
	if err := store.CreateStore(ctx.Request().Context()); err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, store)
}

func (c StoreApiController) PutStore(ctx echo.Context) error {
	var store colleagues.Store
	if err := ctx.Bind(&store); err != nil {

		return renderFail(ctx, api.ErrorParameter.New(err))
	}
	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())

	tenantCode := userClaim.TenantCode
	if len(strings.TrimSpace(userClaim.TenantCode)) == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	store.TenantCode = tenantCode

	/*=======================> Main Function store.UpdateStore <=======================*/
	if err := store.UpdateStore(ctx.Request().Context()); err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}
	return renderSucc(ctx, http.StatusOK, store)
}

func (c StoreApiController) DeleteStore(ctx echo.Context) error {
	storeId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || storeId == 0 {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())

	tenantCode := userClaim.TenantCode
	if len(strings.TrimSpace(userClaim.TenantCode)) == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	/*=======================> Main Function store.UpdateStore <=======================*/
	if err := (colleagues.Store{}).DeleteStore(ctx.Request().Context(), tenantCode, storeId); err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}
	return renderSucc(ctx, http.StatusOK, storeId)
}

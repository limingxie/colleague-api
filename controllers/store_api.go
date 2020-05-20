package controllers

import (
	"net/http"
	"strconv"

	"github.com/hublabs/colleague-api/colleagues"
	"github.com/hublabs/common/api"
	"github.com/labstack/echo"
)

type StoreApiController struct {
}

func (c StoreApiController) Init(g *echo.Echo) {
	g.GET("/v1/store/:storeId/address", c.GetStoreAddressById)

}

func (c StoreApiController) GetStoreAddressById(ctx echo.Context) error {
	storeId, err := strconv.ParseInt(ctx.Param("storeId"), 10, 64)
	if err != nil || storeId == 0 {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	/*=======================> Main Function Colleague.Authentication <=======================*/
	result, err := colleagues.Store{}.GetStoreAddressById(ctx.Request().Context(), storeId)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, result)
}

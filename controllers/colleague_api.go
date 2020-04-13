package controllers

import (
	"net/http"
	"strconv"

	"github.com/hublabs/colleague-api/models"
	"github.com/hublabs/common/api"

	"github.com/labstack/echo"
)

type ColleagueApiController struct {
}

func (c ColleagueApiController) Init(g *echo.Echo) {
	g.GET("/v1/colleagues/:id", c.GetColleagueById)
}

func (c ColleagueApiController) GetColleagueById(ctx echo.Context) error {
	colleagueId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || colleagueId == 0 {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	/*=======================> Main Function GetColleagueById <=======================*/
	result, err := models.Colleague{}.GetColleagueById(ctx.Request().Context(), colleagueId)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, result)

}

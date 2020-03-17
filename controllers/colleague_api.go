package controllers

import (
	"net/http"
	"strconv"

	"github.com/hublabs/colleague-api/models"

	"github.com/labstack/echo"
)

type ColleagueApiController struct {
}

func (c ColleagueApiController) Init(g *echo.Group) {
	g.GET("/api/v1/colleagues/:id", c.GetColleagueById)
}

func (c ColleagueApiController) GetColleagueById(ctx echo.Context) error {
	parameters := []string{}

	colleagueId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || colleagueId == 0 {
		parameters = append(parameters, "id")
	}

	if len(parameters) > 0 {
		return ReturnApiParameterWarn(ctx, parameters)
	}
	/*=======================> Main Function GetColleagueById <=======================*/
	result, err := models.Colleague{}.GetColleagueById(ctx.Request().Context(), colleagueId)
	if err != nil {
		return ReturnApiFail(ctx, ApiErrorDB, err, map[string]interface{}{"colleagueId": colleagueId})
	}

	return ReturnResultApiSucc(ctx, http.StatusOK, result)

}

package controllers

import (
	"net/http"

	"github.com/hublabs/colleague-api/models"
	"github.com/hublabs/common/api"
	"github.com/labstack/echo"
)

type AuthenticationApiController struct {
}

func (c AuthenticationApiController) Init(g *echo.Echo) {
	//用户账号密码登录验证
	g.POST("/v1/colleague/authentication", c.ColleagueAuthentication)
}

func (c AuthenticationApiController) ColleagueAuthentication(ctx echo.Context) error {
	var v struct {
		Mode      string `json:"mode"`
		IdentiKey string `json:"identiKey"`
		Password  string `json:"password"`
	}
	if err := ctx.Bind(&v); err != nil {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	/*=======================> Main Function Colleague.Authentication <=======================*/
	tokenDetail, err := models.Authentication(ctx.Request().Context(), v.Mode, v.IdentiKey, v.Password)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, tokenDetail)
}

package controllers

import (
	"net/http"
	"strings"

	"github.com/hublabs/colleague-api/colleagues"
	"github.com/hublabs/common/api"
	"github.com/hublabs/common/auth"
	"github.com/labstack/echo"
)

type LoginApiController struct {
}

func (c LoginApiController) Init(g *echo.Echo) {
	//用户账号密码登录验证
	g.POST("/v1/login/token-detail", c.GetTokenDetail)
	g.GET("/v1/login/colleague-info", c.GetColleagueInfos)

}

func (c LoginApiController) GetTokenDetail(ctx echo.Context) error {
	var v struct {
		Mode      string `json:"mode"`
		IdentiKey string `json:"identiKey"`
		Password  string `json:"password"`
	}
	if err := ctx.Bind(&v); err != nil {
		return renderFail(ctx, api.ErrorParameter.New(err))
	}

	/*=======================> Main Function Colleague.Authentication <=======================*/
	tokenDetail, err := colleagues.Login{}.GetTokenDetail(ctx.Request().Context(), v.Mode, v.IdentiKey, v.Password)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, tokenDetail)
}

func (c LoginApiController) GetColleagueInfos(ctx echo.Context) error {
	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())
	colleagueId := userClaim.ColleagueId
	if userClaim.ColleagueId == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	tenantCode := userClaim.TenantCode
	if len(strings.TrimSpace(userClaim.TenantCode)) == 0 {
		return renderFail(ctx, api.ErrorTokenInvaild.New(nil))
	}

	/*=======================> Main Function Colleague.Authentication <=======================*/
	result, err := colleagues.Login{}.GetColleagueInfos(ctx.Request().Context(), tenantCode, colleagueId)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, result)
}

package controllers

import (
	"net/http"
	"strconv"

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
	g.GET("/v1/login/colleague-info", c.GetColleagueAndStores)

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

func (c LoginApiController) GetColleagueAndStores(ctx echo.Context) error {
	var colleagueId int64
	userClaim := auth.UserClaim{}.FromCtx(ctx.Request().Context())
	if userClaim.ColleagueId != 0 {
		colleagueId = userClaim.ColleagueId
	} else {
		var err error
		colleagueId, err = strconv.ParseInt(ctx.QueryParams().Get("colleagueId"), 10, 64)
		if err != nil || colleagueId == 0 {
			return renderFail(ctx, api.ErrorParameter.New(err))
		}
	}

	/*=======================> Main Function Colleague.Authentication <=======================*/
	result, err := colleagues.Login{}.GetColleagueAndStores(ctx.Request().Context(), colleagueId)
	if err != nil {
		return renderFail(ctx, api.ErrorDB.New(err))
	}

	return renderSucc(ctx, http.StatusOK, result)
}

package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HomeApiController struct {
}

func (c HomeApiController) Init(g *echo.Group) {
	g.GET("", c.Ping)
}

func (c HomeApiController) Ping(ctx echo.Context) error {
	return ReturnResultApiSucc(ctx, http.StatusOK, "colleague-api-ping")
}

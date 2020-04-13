package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HomeApiController struct {
}

func (c HomeApiController) Init(g *echo.Echo) {
	g.GET("/ping", c.Ping)
}

func (c HomeApiController) Ping(ctx echo.Context) error {
	return renderSucc(ctx, http.StatusOK, "colleague-api-ping")
}

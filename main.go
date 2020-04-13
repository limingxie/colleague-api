package main

import (
	"os"
	"sort"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/kafka"
	"github.com/urfave/cli/v2"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	configutil "github.com/hublabs/colleague-api/config"
	"github.com/hublabs/colleague-api/controllers"
	"github.com/hublabs/colleague-api/models"
)

func main() {
	configutil.Read()

	xormEngine, err := xorm.NewEngine(configutil.DataBaseDriver, configutil.ColleagueApiConnection)
	if err != nil {
		panic(err)
	}

	defer xormEngine.Close()
	models.SetXormEngineSync(xormEngine)

	app := cli.NewApp()
	app.Name = "colleague"
	app.Commands = []*cli.Command{
		{
			Name:  "api-server",
			Usage: "run api server",
			Action: func(c *cli.Context) error {
				if err := initEchoApp(xormEngine).Start(":" + configutil.Httpport); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "seed",
			Usage: "create seed data",
			Action: func(c *cli.Context) error {
				if err := models.Seed(xormEngine); err != nil {
					return err
				}
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)

}

func initEchoApp(xormEngine *xorm.Engine) *echo.Echo {
	xormEngine.SetMaxOpenConns(50)
	xormEngine.SetMaxIdleConns(50)
	xormEngine.SetConnMaxLifetime(60 * time.Second)

	e := echo.New()

	InitControllers(e)

	e.Static("/static", "static")
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	e.Use(echomiddleware.ContextDB(configutil.Service, xormEngine, kafka.Config{}))

	// 초기에 token 인증을 처리하지 않고 후에는 처리 되여야 함.
	// e.Use(auth.UserClaimMiddelware())

	return e
}

func InitControllers(e *echo.Echo) {
	controllers.HomeApiController{}.Init(e)
	controllers.ColleagueApiController{}.Init(e)
	controllers.AuthenticationApiController{}.Init(e)
}

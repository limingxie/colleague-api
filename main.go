package main

import (
	"flag"
	"os"
	"sort"
	"time"

	"github.com/hublabs/colleague-api/colleagues"
	"github.com/hublabs/colleague-api/config"
	"github.com/hublabs/colleague-api/controllers"
	"github.com/hublabs/colleague-api/tenants"
	"github.com/hublabs/common/api"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/kafka"
	"github.com/urfave/cli/v2"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	appEnv = flag.String("app-env", os.Getenv("APP_ENV"), "app env")
)

func main() {
	c := config.Init(*appEnv)
	api.SetErrorMessagePrefix(c.ServiceName)

	colleagues.SetColleagueConfig(&colleagues.ColleagueConfig{
		AppEnv: *appEnv,
	})

	xormEngine, err := xorm.NewEngine(c.Database.Driver, c.Database.Connection)
	if err != nil {
		panic(err)
	}

	defer xormEngine.Close()
	SetXormEngineSync(xormEngine)

	app := cli.NewApp()
	app.Name = "colleague"
	app.Commands = []*cli.Command{
		{
			Name:  "api-server",
			Usage: "run api server",
			Action: func(cliContext *cli.Context) error {
				if err := initEchoApp(xormEngine, c.ServiceName).Start(":" + c.HttpPort); err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "seed",
			Usage: "create seed data",
			Action: func(c *cli.Context) error {
				if err := colleagues.Seed(xormEngine); err != nil {
					return err
				}
				if err := tenants.Seed(xormEngine); err != nil {
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

func initEchoApp(xormEngine *xorm.Engine, serviceName string) *echo.Echo {
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

	e.Use(echomiddleware.ContextDB(serviceName, xormEngine, kafka.Config{}))

	// 초기에 token 인증을 처리하지 않고 후에는 처리 되여야 함.
	// e.Use(auth.UserClaimMiddelware())

	return e
}

func InitControllers(e *echo.Echo) {
	controllers.HomeApiController{}.Init(e)
	controllers.ColleagueApiController{}.Init(e)
	controllers.LoginApiController{}.Init(e)
}

func SetXormEngineSync(xormEngine *xorm.Engine) {
	//xormEngine.ShowSQL(true)

	xormEngine.Sync(new(tenants.Tenant))
	xormEngine.Sync(new(tenants.Brand))

	xormEngine.Sync(new(colleagues.Colleague))
	xormEngine.Sync(new(colleagues.Store))
	xormEngine.Sync(new(colleagues.StoreBrand))
	xormEngine.Sync(new(colleagues.StoreColleague))
}

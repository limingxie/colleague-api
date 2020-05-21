package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/hublabs/colleague-api/colleagues"
	"github.com/hublabs/colleague-api/config"
	"github.com/hublabs/colleague-api/controllers"
	"github.com/hublabs/colleague-api/factory"
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
	fmt.Println("----------appEnv---------------")
	fmt.Println(*appEnv)
	fmt.Println("-----------appEnv--------------")
	api.SetErrorMessagePrefix(c.ServiceName)

	colleagues.SetColleagueConfig(&colleagues.ColleagueConfig{
		AppEnv: *appEnv,
	})

	xormEngine := initXormEngine(c.Database.Driver, c.Database.Connection)
	factory.InitDB(xormEngine)

	defer xormEngine.Close()

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
	controllers.MigrationController{}.Init(e)
}

func initXormEngine(driver, connection string) *xorm.Engine {
	fmt.Println("-------------------------")
	fmt.Println(driver)
	fmt.Println(connection)
	fmt.Println("-------------------------")

	xormEngine, err := xorm.NewEngine(driver, connection)
	if err != nil {
		panic(err)
	}
	xormEngine.SetMaxIdleConns(5)
	xormEngine.SetMaxOpenConns(20)
	xormEngine.SetConnMaxLifetime(time.Minute * 10)
	//xormEngine.ShowSQL()

	if err := tenants.Init(xormEngine); err != nil {
		log.Fatal(err)
	}

	if err := colleagues.Init(xormEngine); err != nil {
		log.Fatal(err)
	}

	return xormEngine
}

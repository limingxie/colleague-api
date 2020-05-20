package controllers

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/hublabs/colleague-api/tenants"
	"github.com/hublabs/common/auth"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pangpanglabs/goutils/echomiddleware"
	"github.com/pangpanglabs/goutils/jwtutil"
	"github.com/pangpanglabs/goutils/kafka"

	"github.com/hublabs/colleague-api/colleagues"
)

var (
	appEnv           = flag.String("app-env", os.Getenv("APP_ENV"), "app env")
	ctx              context.Context
	echoApp          *echo.Echo
	handleWithFilter func(handlerFunc echo.HandlerFunc, c echo.Context) error
	xormEngine       *xorm.Engine
)

func init() {
	runtime.GOMAXPROCS(1)
	var err error
	xormEngine, err = xorm.NewEngine("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	colleagues.SetColleagueConfig(&colleagues.ColleagueConfig{
		AppEnv: "test",
	})

	SetXormEngineSync(xormEngine)
	if err := colleagues.Seed(xormEngine); err != nil {
		fmt.Println("seed err:", err)
	}

	if err := tenants.Seed(xormEngine); err != nil {
		fmt.Println("seed err:", err)
	}

	echoApp = echo.New()
	handleWithFilter = func(handlerFunc echo.HandlerFunc, c echo.Context) error {
		return echomiddleware.ContextDB("colleague-api", xormEngine, kafka.Config{})(handlerFunc)(c)
	}
	ctx = context.WithValue(context.Background(), echomiddleware.ContextDBName, xormEngine.NewSession())
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

func SetContext(req *http.Request) (echo.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	c := echoApp.NewContext(req, rec)
	c.SetRequest(req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, xormEngine.NewSession())))

	return c, rec
}
func SetContextWithSession(req *http.Request, session *xorm.Session) (echo.Context, *httptest.ResponseRecorder) {

	rec := httptest.NewRecorder()

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := echoApp.NewContext(req, rec)
	c.SetRequest(req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, session)))

	return c, rec
}

func GetTokenForTest() string {
	token, _ := jwtutil.NewTokenWithSecret(map[string]interface{}{
		"aud": "colleague", "tenantCode": "hublabs", "colleagueId": 1, "iss": "colleague",
		"nbf": time.Now().Add(-5 * time.Minute).Unix(),
	}, os.Getenv("JWT_SECRET"))
	return token
}

func SetContextWithToken(req *http.Request, token string) (echo.Context, *httptest.ResponseRecorder) {
	rec := httptest.NewRecorder()

	if len(strings.TrimSpace(token)) == 0 {
		token = GetTokenForTest()
	}
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, token)
	userClaims := auth.UserClaim{
		TenantCode:  "hublabs",
		ColleagueId: 1,
		Username:    "system",
	}
	req = req.WithContext(context.WithValue(req.Context(), "userClaim", userClaims))
	req = req.WithContext(context.WithValue(req.Context(), "token", token))
	req = req.WithContext(context.WithValue(req.Context(), echomiddleware.ContextDBName, xormEngine.NewSession()))

	c := echoApp.NewContext(req, rec)
	c.SetRequest(req)

	return c, rec
}

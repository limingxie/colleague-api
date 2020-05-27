package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hublabs/colleague-api/colleagues"
	"github.com/hublabs/colleague-api/factory"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/test"
)

func Test_AppApiController_GetAppAndBrandsByAppId(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/v1/apps/:id", nil)

	c, rec := SetContextWithToken(req, "")
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	c.SetParamNames("id")
	c.SetParamValues("2")

	test.Ok(t, AppApiController{}.GetById(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  colleagues.App         `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Id, int64(2))
	test.Equals(t, v.Result.TenantCode, "hublabs")
	test.Equals(t, v.Result.Code, "OHUB")
	test.Equals(t, v.Result.Name, "在线结算")
}

func Test_AppApiController_PostApp(t *testing.T) {
	param := map[string]interface{}{
		"tenantCode": "hublabs",
		"code":       "A001",
		"name":       "TestApp",
		"enable":     true,
	}

	body, err := json.Marshal(param)
	test.Ok(t, err)

	req := httptest.NewRequest(echo.POST, "/v1/apps", bytes.NewBuffer(body))

	c, rec := SetContextWithToken(req, "")
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, AppApiController{}.PostApp(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  colleagues.App         `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Code, "A001")
	test.Equals(t, v.Result.TenantCode, "hublabs")

	var aa []colleagues.App
	err = factory.DB(c.Request().Context()).Find(&aa)
	if err != nil {
		fmt.Println(err)
	}

	app, err := colleagues.App{}.GetByCode(c.Request().Context(), "hublabs", "A001")
	test.Ok(t, err)

	test.Equals(t, app.TenantCode, "hublabs")
	test.Equals(t, app.Code, "A001")
	test.Equals(t, app.Name, "TestApp")
	test.Equals(t, app.Enable, true)

}

func Test_AppApiController_PutApp(t *testing.T) {
	param := map[string]interface{}{
		"id":         int64(2),
		"tenantCode": "hublabs",
		"name":       "TestApp",
		"enable":     true,
	}

	body, err := json.Marshal(param)
	test.Ok(t, err)

	req := httptest.NewRequest(echo.PUT, "/v1/apps", bytes.NewBuffer(body))

	c, rec := SetContextWithToken(req, "")
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, AppApiController{}.PutApp(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  colleagues.App         `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.TenantCode, "hublabs")

	app, err := colleagues.App{}.GetById(c.Request().Context(), "hublabs", 2)
	test.Ok(t, err)

	test.Equals(t, app.TenantCode, "hublabs")
	test.Equals(t, app.Code, "OHUB")
	test.Equals(t, app.Name, "TestApp")
	test.Equals(t, app.Enable, true)
}

func Test_AppApiController_DeleteApp(t *testing.T) {
	req := httptest.NewRequest(echo.DELETE, "/v1/apps/:id", nil)

	c, rec := SetContextWithToken(req, "")
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	c.SetParamNames("id")
	c.SetParamValues("2")

	test.Ok(t, AppApiController{}.DeleteApp(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  int64                  `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result, int64(2))

	app, err := colleagues.App{}.GetById(c.Request().Context(), "hublabs", 2)
	test.Ok(t, err)

	test.Equals(t, app.Id, int64(0))

}

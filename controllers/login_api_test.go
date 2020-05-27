package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hublabs/colleague-api/factory"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/test"
)

func Test_LoginApiController_GetTokenDetail(t *testing.T) {

	param := map[string]interface{}{
		"mode":      "email",
		"identiKey": "system@email.com",
		"password":  "1111",
	}

	body, err := json.Marshal(param)
	test.Ok(t, err)

	req := httptest.NewRequest(echo.POST, "/v1/login/token-detail", bytes.NewBuffer(body))

	c, rec := SetContext(req)
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, LoginApiController{}.GetTokenDetail(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  map[string]interface{} `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result["colleagueId"].(float64), float64(1))
}

func Test_LoginApiController_GetColleagueInfos(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/v1/login/colleague-info", nil)

	c, rec := SetContextWithToken(req, "")
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, LoginApiController{}.GetColleagueInfos(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  map[string]interface{} `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result["id"].(float64), float64(1))
	test.Equals(t, v.Result["name"].(string), "系统管理员")
	test.Equals(t, v.Result["username"].(string), "system")

	stores := v.Result["stores"].([]interface{})
	test.Equals(t, stores[0].(map[string]interface{})["code"].(string), "C001")
	test.Equals(t, stores[0].(map[string]interface{})["name"].(string), "北京Nike店")
	test.Equals(t, stores[0].(map[string]interface{})["role"].(string), "admin")

	brands0 := stores[0].(map[string]interface{})["brands"].([]interface{})
	test.Equals(t, brands0[0].(map[string]interface{})["tenantCode"].(string), "hublabs")
	test.Equals(t, brands0[0].(map[string]interface{})["code"].(string), "NK")
	test.Equals(t, brands0[0].(map[string]interface{})["name"].(string), "Nike")

	test.Equals(t, stores[1].(map[string]interface{})["code"].(string), "C002")
	test.Equals(t, stores[1].(map[string]interface{})["name"].(string), "北京Adidas店")
	test.Equals(t, stores[1].(map[string]interface{})["role"].(string), "admin")

	brands1 := stores[1].(map[string]interface{})["brands"].([]interface{})
	test.Equals(t, brands1[0].(map[string]interface{})["tenantCode"].(string), "hublabs")
	test.Equals(t, brands1[0].(map[string]interface{})["code"].(string), "AD")
	test.Equals(t, brands1[0].(map[string]interface{})["name"].(string), "Adidas")

	test.Equals(t, stores[2].(map[string]interface{})["code"].(string), "C003")
	test.Equals(t, stores[2].(map[string]interface{})["name"].(string), "上海运动专卖店")
	test.Equals(t, stores[2].(map[string]interface{})["role"].(string), "admin")

	brands20 := stores[2].(map[string]interface{})["brands"].([]interface{})
	test.Equals(t, brands20[0].(map[string]interface{})["tenantCode"].(string), "hublabs")
	test.Equals(t, brands20[0].(map[string]interface{})["code"].(string), "NK")
	test.Equals(t, brands20[0].(map[string]interface{})["name"].(string), "Nike")

	brands21 := stores[2].(map[string]interface{})["brands"].([]interface{})
	test.Equals(t, brands21[1].(map[string]interface{})["tenantCode"].(string), "hublabs")
	test.Equals(t, brands21[1].(map[string]interface{})["code"].(string), "AD")
	test.Equals(t, brands21[1].(map[string]interface{})["name"].(string), "Adidas")

	apps := v.Result["apps"].([]interface{})
	test.Equals(t, apps[0].(map[string]interface{})["id"].(float64), float64(1))
	test.Equals(t, apps[0].(map[string]interface{})["code"].(string), "O2O")
	test.Equals(t, apps[0].(map[string]interface{})["name"].(string), "在线抢单")
	test.Equals(t, apps[0].(map[string]interface{})["role"].(string), "admin")

	test.Equals(t, apps[1].(map[string]interface{})["id"].(float64), float64(2))
	test.Equals(t, apps[1].(map[string]interface{})["code"].(string), "OHUB")
	test.Equals(t, apps[1].(map[string]interface{})["name"].(string), "在线结算")
	test.Equals(t, apps[1].(map[string]interface{})["role"].(string), "admin")
}

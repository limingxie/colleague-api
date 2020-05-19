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
		"identiKey": "xiao_ming@email.com",
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

func Test_LoginApiController_GetColleagueAndStores(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/v1/stores/authorization?colleagueId=1&tenantCode=hublabs", nil)

	c, rec := SetContext(req)
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, LoginApiController{}.GetColleagueAndStores(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  map[string]interface{} `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result["id"].(float64), float64(1))
	test.Equals(t, v.Result["name"].(string), "xiao_ming")

	stores := v.Result["stores"].([]interface{})
	test.Equals(t, stores[0].(map[string]interface{})["code"].(string), "C001")
	test.Equals(t, stores[0].(map[string]interface{})["name"].(string), "北京朝阳门店")
	test.Equals(t, stores[0].(map[string]interface{})["role"].(string), "admin")

	brands0 := stores[0].(map[string]interface{})["brands"].([]interface{})
	test.Equals(t, brands0[0].(map[string]interface{})["tenantCode"].(string), "hublabs")
	test.Equals(t, brands0[0].(map[string]interface{})["code"].(string), "NK")
	test.Equals(t, brands0[0].(map[string]interface{})["name"].(string), "Nike")
	test.Equals(t, brands0[1].(map[string]interface{})["tenantCode"].(string), "hublabs")
	test.Equals(t, brands0[1].(map[string]interface{})["code"].(string), "AD")
	test.Equals(t, brands0[1].(map[string]interface{})["name"].(string), "Adidas")

	test.Equals(t, stores[1].(map[string]interface{})["code"].(string), "C002")
	test.Equals(t, stores[1].(map[string]interface{})["name"].(string), "北京新世界百货店")
	test.Equals(t, stores[1].(map[string]interface{})["role"].(string), "member")

	brands1 := stores[1].(map[string]interface{})["brands"].([]interface{})
	test.Equals(t, brands1[0].(map[string]interface{})["tenantCode"].(string), "hublabs")
	test.Equals(t, brands1[0].(map[string]interface{})["code"].(string), "NK")
	test.Equals(t, brands1[0].(map[string]interface{})["name"].(string), "Nike")

	test.Equals(t, stores[2].(map[string]interface{})["code"].(string), "C003")
	test.Equals(t, stores[2].(map[string]interface{})["name"].(string), "上海西单店")
	test.Equals(t, stores[2].(map[string]interface{})["role"].(string), "guest")

	brands2 := stores[2].(map[string]interface{})["brands"].([]interface{})
	test.Equals(t, brands2[0].(map[string]interface{})["tenantCode"].(string), "hublabs")
	test.Equals(t, brands2[0].(map[string]interface{})["code"].(string), "AD")
	test.Equals(t, brands2[0].(map[string]interface{})["name"].(string), "Adidas")
}

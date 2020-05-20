package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hublabs/colleague-api/factory"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/goutils/test"
)

func Test_StoreApiController_GetStoreAddressById(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/v1/store/:storeId/address", nil)

	c, rec := SetContext(req)
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	c.SetParamNames("storeId")
	c.SetParamValues("1")

	test.Ok(t, StoreApiController{}.GetStoreAddressById(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  map[string]interface{} `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result["id"].(float64), float64(1))
	test.Equals(t, v.Result["province"].(string), "北京市")
	test.Equals(t, v.Result["city"].(string), "北京市")
	test.Equals(t, v.Result["district"].(string), "朝阳区")
	test.Equals(t, v.Result["detail"].(string), "酒仙桥中路恒通商务园B37")
}

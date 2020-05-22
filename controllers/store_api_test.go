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

func Test_StoreApiController_PostStore(t *testing.T) {
	param := map[string]interface{}{
		"tenantCode": "hublabs",
		"code":       "A001",
		"name":       "测试A卖场",
		"province":   "河北省",
		"city":       "石家庄",
		"district":   "朝阳区",
		"detail":     "恒通商务园B37",
		"enable":     true,
	}

	body, err := json.Marshal(param)
	test.Ok(t, err)

	req := httptest.NewRequest(echo.POST, "/v1/store", bytes.NewBuffer(body))

	c, rec := SetContextWithToken(req, "")
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, StoreApiController{}.PostStore(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  colleagues.Store       `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Code, "A001")
	test.Equals(t, v.Result.TenantCode, "hublabs")

	var aa []colleagues.Store
	err = factory.DB(c.Request().Context()).Find(&aa)
	if err != nil {
		fmt.Println(err)
	}

	store, err := colleagues.Store{}.GetStoreByCode(c.Request().Context(), "hublabs", "A001")
	test.Ok(t, err)

	test.Equals(t, store.TenantCode, "hublabs")
	test.Equals(t, store.Code, "A001")
	test.Equals(t, store.Name, "测试A卖场")
	test.Equals(t, store.Province, "河北省")
	test.Equals(t, store.City, "石家庄")
	test.Equals(t, store.District, "朝阳区")
	test.Equals(t, store.Detail, "恒通商务园B37")
	test.Equals(t, store.Enable, true)

}

func Test_StoreApiController_PutStore(t *testing.T) {
	param := map[string]interface{}{
		"id":         int64(3),
		"tenantCode": "hublabs",
		"name":       "测试A卖场",
		"province":   "河北省",
		"city":       "石家庄",
		"district":   "朝阳区",
		"detail":     "恒通商务园B37",
		"enable":     true,
	}

	body, err := json.Marshal(param)
	test.Ok(t, err)

	req := httptest.NewRequest(echo.PUT, "/v1/store", bytes.NewBuffer(body))

	c, rec := SetContextWithToken(req, "")
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, StoreApiController{}.PutStore(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  colleagues.Store       `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.TenantCode, "hublabs")

	store, err := colleagues.Store{}.GetStoreById(c.Request().Context(), "hublabs", 3)
	test.Ok(t, err)

	test.Equals(t, store.TenantCode, "hublabs")
	test.Equals(t, store.Name, "测试A卖场")
	test.Equals(t, store.Province, "河北省")
	test.Equals(t, store.City, "石家庄")
	test.Equals(t, store.District, "朝阳区")
	test.Equals(t, store.Detail, "恒通商务园B37")
	test.Equals(t, store.Enable, true)
}

func Test_StoreApiController_DeleteStore(t *testing.T) {
	req := httptest.NewRequest(echo.DELETE, "/v1/stores/:id", nil)

	c, rec := SetContextWithToken(req, "")
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	c.SetParamNames("id")
	c.SetParamValues("3")

	test.Ok(t, StoreApiController{}.DeleteStore(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  int64                  `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result, int64(3))

	store, err := colleagues.Store{}.GetStoreById(c.Request().Context(), "hublabs", 3)
	test.Ok(t, err)

	test.Equals(t, store.Id, int64(0))

}

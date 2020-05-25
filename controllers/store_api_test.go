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

func Test_StoreApiController_GetStoreAndBrandsByStoreId(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/v1/stores/:id", nil)

	c, rec := SetContext(req)
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	c.SetParamNames("id")
	c.SetParamValues("3")

	test.Ok(t, StoreApiController{}.GetStoreAndBrandsByStoreId(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  colleagues.Store       `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result.Id, int64(3))
	test.Equals(t, v.Result.TenantCode, "hublabs")
	test.Equals(t, v.Result.Code, "C003")
	test.Equals(t, v.Result.Name, "上海运动专卖店")
	test.Equals(t, v.Result.AddrProvince, "上海市")
	test.Equals(t, v.Result.AddrCity, "上海市")
	test.Equals(t, v.Result.AddrDistrict, "浦东区")
	test.Equals(t, v.Result.AddrDetail, "漕宝路23号")
	test.Equals(t, v.Result.Mobile, "15031056672")
	test.Equals(t, v.Result.Tel, "3333333330")
	test.Equals(t, v.Result.Brands[0].Code, "NK")
	test.Equals(t, v.Result.Brands[0].Name, "Nike")
	test.Equals(t, v.Result.Brands[1].Code, "AD")
	test.Equals(t, v.Result.Brands[1].Name, "Adidas")

}

func Test_StoreApiController_PostStore(t *testing.T) {
	param := map[string]interface{}{
		"tenantCode":   "hublabs",
		"code":         "A001",
		"name":         "测试A卖场",
		"addrProvince": "河北省",
		"addrCity":     "石家庄",
		"addrDistrict": "朝阳区",
		"addrDetail":   "恒通商务园B37",
		"mobile":       "111",
		"tel":          "222",
		"enable":       true,
	}

	body, err := json.Marshal(param)
	test.Ok(t, err)

	req := httptest.NewRequest(echo.POST, "/v1/stores", bytes.NewBuffer(body))

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
	test.Equals(t, store.AddrProvince, "河北省")
	test.Equals(t, store.AddrCity, "石家庄")
	test.Equals(t, store.AddrDistrict, "朝阳区")
	test.Equals(t, store.AddrDetail, "恒通商务园B37")
	test.Equals(t, store.Mobile, "111")
	test.Equals(t, store.Tel, "222")
	test.Equals(t, store.Enable, true)

}

func Test_StoreApiController_PutStore(t *testing.T) {
	param := map[string]interface{}{
		"id":           int64(3),
		"tenantCode":   "hublabs",
		"name":         "测试A卖场",
		"addrProvince": "河北省",
		"addrCity":     "石家庄",
		"addrDistrict": "朝阳区",
		"addrDetail":   "恒通商务园B37",
		"mobile":       "111",
		"tel":          "222",
		"enable":       true,
	}

	body, err := json.Marshal(param)
	test.Ok(t, err)

	req := httptest.NewRequest(echo.PUT, "/v1/stores", bytes.NewBuffer(body))

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
	test.Equals(t, store.AddrProvince, "河北省")
	test.Equals(t, store.AddrCity, "石家庄")
	test.Equals(t, store.AddrDistrict, "朝阳区")
	test.Equals(t, store.AddrDetail, "恒通商务园B37")
	test.Equals(t, store.Mobile, "111")
	test.Equals(t, store.Tel, "222")
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

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

func Test_AuthenticationApiController_ColleagueAuthentication(t *testing.T) {

	param := map[string]interface{}{
		"mode":      "email",
		"identiKey": "xiao_ming@email.com",
		"password":  "1111",
	}

	body, err := json.Marshal(param)
	test.Ok(t, err)

	req := httptest.NewRequest(echo.POST, "/v1/colleague/authentication", bytes.NewBuffer(body))

	c, rec := SetContext(req)
	dbSession := factory.DB(c.Request().Context())
	dbSession.Begin()
	defer func() {
		factory.DB(c.Request().Context()).Close()
		factory.DB(c.Request().Context()).Rollback()
	}()

	test.Ok(t, AuthenticationApiController{}.ColleagueAuthentication(c))
	test.Equals(t, http.StatusOK, rec.Code)

	var v struct {
		Result  map[string]interface{} `json:"result"`
		Success bool                   `json:"success"`
		Errors  map[string]interface{} `json:"error"`
	}

	test.Ok(t, json.Unmarshal(rec.Body.Bytes(), &v))
	test.Equals(t, v.Result["colleagueId"].(float64), float64(1))
}

package common

import (
	"logistics/constants"
	myerrors "logistics/libs/errors"
	"logistics/libs/logs"
	"reflect"
	"strconv"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func GetRequestUserId(c echo.Context) int64 {
	userId := c.Request().Header.Get("user_id")
	if userId == "" {
		return 0
	}
	if id, err := strconv.ParseInt(userId, 10, 64); err == nil {
		return id
	}
	return 0
}

func ValidateRequest[Type any](c echo.Context) (*Type, *myerrors.StackError) {
	req := new(Type)
	err := c.Bind(req)
	if err != nil {
		t := reflect.TypeOf(req).Elem()
		logs.Errorf("bind type: %v err: %s", t, err.Error())
		return nil, constants.ErrorDataBindError
	}
	err = c.Validate(req)
	if err != nil {
		t := reflect.TypeOf(req).Elem()
		logs.Errorf("validate type: %v err: %s", t, err.Error())
		return nil, constants.ErrorDataBindError
	}
	uid := c.Request().Header.Get("user_id")
	if uid == "" {
		logs.Error("account_id header is missing")
		return nil, constants.ErrorAccountIdMissing
	}
	role := c.Request().Header.Get("role")
	if role == "" {
		logs.Error("role header is missing")
		return nil, constants.ErrorRoleMissing
	}
	return req, nil
}

// CustomValidator 自定义 Validator
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate 实现 Validate 方法
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

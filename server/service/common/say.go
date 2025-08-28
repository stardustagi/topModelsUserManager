package common

import (
	"logistics/libs/errors"

	"github.com/labstack/echo/v4"
)

func newReply(err *errors.StackError, data interface{}) (int, BaseResponse) {
	if err == nil {
		if data != nil {
			return 200, BaseResponse{
				ErrCode: 0,
				ErrMsg:  "操作成功!",
				Data:    data,
			}
		}
		return 200, BaseResponse{
			ErrCode: 0,
			ErrMsg:  "操作成功!",
		}
	}
	return 200, BaseResponse{
		ErrCode: err.Code(),
		ErrMsg:  err.Msg(),
		Data:    nil,
	}
}

func _response(err *errors.StackError, data interface{}) (int, BaseResponse) {
	return newReply(err, data)
}

func Response(c echo.Context, err *errors.StackError, data any) error {
	return c.JSON(_response(err, data))
}

func ResponseEx(c echo.Context, err error, data any) error {
	serr, ok := err.(*errors.StackError)
	if ok {
		return Response(c, serr, data)
	} else {
		return Response(c, errors.New(err.Error(), 20001), data)
	}
}

func NativeToStack(err error) *errors.StackError {
	serr, ok := err.(*errors.StackError)
	if ok {
		return serr
	} else {
		return errors.New(err.Error(), 20001)
	}
}

func newPageReply(err *errors.StackError, data interface{}, pageResp PageResp) (int, BasePageResponse) {
	if err == nil {
		if data != nil {
			return 200, BasePageResponse{
				ErrCode: 0,
				ErrMsg:  "操作成功!",
				Data:    data,
				Page:    pageResp,
			}
		}
		return 200, BasePageResponse{
			// Replace ErrCode with the correct field name, e.g., Code or Status
			ErrCode: 0,
		}
	}
	return 200, BasePageResponse{
		ErrCode: err.Code(),
		ErrMsg:  err.Msg(),
	}
}

func _pageResponse(err *errors.StackError, data interface{}, pageResp PageResp) (int, BasePageResponse) {
	return newPageReply(err, data, pageResp)
}

func PageResponse(c echo.Context, err *errors.StackError, data interface{}, pageResp PageResp) error {
	return c.JSON(_pageResponse(err, data, pageResp))
}

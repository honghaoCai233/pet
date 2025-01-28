package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pet/pkg/ecode"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(data any, err error) (int, *Response) {

	var (
		code     = 0
		msg      = ""
		httpCode = http.StatusOK
	)
	rv := new(Response)

	ec := ecode.FromError(err)
	if ec != nil {
		code = ec.Code
		msg = ec.Message
		httpCode = ec.HttpCode
	}
	if code == ecode.UnknownCode {
		msg = ecode.InternalErr.WithCause(err).Error()
	}

	rv.Code = code
	rv.Message = msg
	rv.Data = data

	return httpCode, rv
}

func ErrorResponse(c *gin.Context, err error) {
	c.JSON(newResponse(nil, err))
	c.Abort()
}

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(newResponse(data, nil))
}

func NewResponse(c *gin.Context) func(data any, err error) {
	return func(data any, err error) {
		if err != nil {
			ErrorResponse(c, err)
		} else {
			SuccessResponse(c, data)
		}
	}
}

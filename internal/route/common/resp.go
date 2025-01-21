package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"pet/pkg/ecode"
)

var (
	log *zap.SugaredLogger
)

func SetRespLog(logger *zap.SugaredLogger) {
	log = logger
}

/*{
"success": true,
"data": {},
"errorCode": "1001",
"errorMessage": "error message",
"showType": 2,
"traceId": "someid",
"host": "10.1.1.1"
}*/

// Resp See Doc https://pro.ant.design/zh-CN/docs/request
type Resp struct {
	Success      bool   `json:"success"`
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Data         any    `json:"data"`
}

func NewResp(data interface{}, err error) (int, *Resp) {
	var (
		code     = 0
		msg      = ""
		httpCode = http.StatusOK
	)
	ec := ecode.FromError(err)
	if ec != nil {
		code = ec.Code
		msg = ec.Message
		httpCode = ec.HttpCode
	}
	if code == ecode.UnknownCode {
		msg = ecode.InternalErr.Message
		log.Errorf("NewResp receive unknown error: %s", ec)
	}
	return httpCode, &Resp{
		Success:      code == 0,
		ErrorCode:    code,
		ErrorMessage: msg,
		Data:         data,
	}
}

func ErrorResp(c *gin.Context, err error) {
	c.JSON(NewResp(nil, err))
	c.Abort()
}

func ParamsErrorResp(c *gin.Context, err error) {
	ErrorResp(c, ecode.NewInvalidParamsErr(TranslateErr(err)))
}

func SuccessResp(c *gin.Context, data any) {
	c.JSON(NewResp(data, nil))
}

func WrapResp(c *gin.Context) func(data any, err error) {
	return func(data any, err error) {
		if err != nil {
			ErrorResp(c, err)
		} else {
			SuccessResp(c, data)
		}
	}
}

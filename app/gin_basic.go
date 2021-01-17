package app

import (
	"gin-kit-demo/exception"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type BasicController struct {
	Ctx *gin.Context
}


// 统一返回值
type ResponseData struct {
	Ret        int         `json:"ret"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
	ServerTime int64       `json:"serverTime"`
}

func (t *BasicController) Ok(d interface{}) {
	rd := &ResponseData{
		Ret:        200,
		Message:    "ok",
		Result:     d,
		ServerTime: time.Now().UnixNano() / 1000000,
	}
	t.Ctx.JSONP(http.StatusOK, rd)
	return
}

func Wrapper(handler func(c *gin.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			err error
		)
		err =handler(c)
		if err != nil {
			var apiException *exception.APIException
			if h,ok := err.(*exception.APIException); ok {
				apiException = h
			}else {
				apiException = exception.ServerError()
			}
			apiException.Request = c.Request.Method + " "+ c.Request.URL.String()
			c.JSON(apiException.Code,apiException)
			return
		}

	}
}
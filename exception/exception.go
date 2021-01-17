package exception

import "net/http"



type APIException struct {
	Code      int    `json:"-"`
	ErrorCode int    `json:"error_code"`
	Msg       string `json:"msg"`
	Request   string `json:"request"`
}

// 实现接口
func (e *APIException) Error() string {
	return e.Msg
}

func newAPIException(code int, errorCode int, msg string) *APIException {
	return &APIException{
		Code:      code,
		ErrorCode: errorCode,
		Msg:       msg,
	}
}

func ServerError() *APIException {
	return newAPIException(http.StatusInternalServerError, SERVER_ERROR, http.StatusText(http.StatusInternalServerError))
}

// 参数错误
func ParameterError(message string) *APIException {
	return newAPIException(http.StatusBadRequest,PARAMETER_ERROR,message)
}
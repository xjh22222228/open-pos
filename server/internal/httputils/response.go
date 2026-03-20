package httputils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	C *gin.Context
}

type Result struct {
	Data      any    `json:"data"`
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

func NewResponse(c *gin.Context) *Response {
	return &Response{C: c}
}

func (r *Response) Success(data any, message ...string) {
	msg := "success"
	if len(message) > 0 {
		msg = message[0]
	}
	r.C.JSON(http.StatusOK, Result{
		Data:      data,
		Message:   msg,
		ErrorCode: 0,
	})
}

func (r *Response) Error(code int, msg string) {
	r.C.AbortWithStatusJSON(http.StatusInternalServerError, Result{
		Data:      nil,
		Message:   msg,
		ErrorCode: code,
	})
}

// 参数错误
func (r *Response) BadRequest(msg string) {
	r.C.AbortWithStatusJSON(http.StatusBadRequest, Result{
		Data:      nil,
		Message:   msg,
		ErrorCode: 400,
	})
}

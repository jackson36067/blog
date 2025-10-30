package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构体
type Response struct {
	Code    int    `json:"code"`    // 状态码
	Message string `json:"message"` // 提示信息
	Data    any    `json:"data"`    // 返回数据
}

// Success 返回成功响应
func Success(c *gin.Context, data any, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: msg,
		Data:    data,
	})
}

// Fail 返回失败响应
func Fail(c *gin.Context, code int, msg string) {
	if code == 0 {
		code = 400
	}
	c.JSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

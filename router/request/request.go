package request

import (
	"explore/router/response"
	"github.com/gin-gonic/gin"
)

func CheckJsonParams[T any](c *gin.Context, body T) bool {
	if err := c.ShouldBindJSON(&body); err != nil {
		response.FailWithCode(c, "参数不正确", response.ParamsErrorCode)
		return false
	}
	return true
}

func CheckQueryParams[T any](c *gin.Context, body T) bool {
	if err := c.ShouldBindQuery(&body); err != nil {
		response.FailWithCode(c, "参数不正确", response.ParamsErrorCode)
		return false
	}
	return true
}

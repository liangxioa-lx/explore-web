package response

import (
	"contract/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success[T any](c *gin.Context, data T, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": SuccessCode,
		"data": data,
		"msg":  message,
	})
}

func SuccessPageList[T any](c *gin.Context, page common.Page, list T, total int64) {
	Success(c, gin.H{
		"list":      list,
		"pageNum":   page.PageNum,
		"pageSize":  page.PageSize,
		"total":     total,
		"totalPage": (total / int64(page.PageSize)) + 1,
	}, "查询成功")
}

func Fail(c *gin.Context, message string) {
	FailWithCode(c, message, NormalErrorCode)
}

func FailWithCode(c *gin.Context, message string, code int) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}

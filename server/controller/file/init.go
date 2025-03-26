package FileController

import (
	"explore/router/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup) {
	r := router.Group("/file")
	r.POST("/file/create", CreateFile)
	r.POST("/dir/create", CreateDirectory)
	r.POST("/rename", Rename)
	r.POST("/move", Move)
	r.POST("/delete", Delete)
	r.POST("/find", Find)
	r.GET("/drives", FindDriverList)
	
	// 系统剪贴板操作
	r.POST("/copy", Copy)      // 复制到系统剪贴板
	r.POST("/cut", Cut)        // 剪切到系统剪贴板
}

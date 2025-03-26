package FileController

import (
	"explore/router/request"
	"explore/router/response"
	"fileService"
	"github.com/gin-gonic/gin"
)

type ClipboardBody struct {
	Paths []string `json:"paths" binding:"required"`
}

// Copy 复制文件到系统剪贴板
func Copy(c *gin.Context) {
	var body ClipboardBody
	if !request.CheckJsonParams(c, &body) {
		return
	}

	err := fileService.SetClipboard(body.Paths, fileService.Copy)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, nil, "已复制到系统剪贴板")
}

// Cut 剪切文件到系统剪贴板
func Cut(c *gin.Context) {
	var body ClipboardBody
	if !request.CheckJsonParams(c, &body) {
		return
	}

	err := fileService.SetClipboard(body.Paths, fileService.Cut)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, nil, "已剪切到系统剪贴板")
}

package FileController

import (
	"explore/router/request"
	"explore/router/response"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"fileService"
)

type RenameBody struct {
	Name string `json:"name" binding:"required"`
	OldName string `json:"oldName" binding:"required"`
	Path string `json:"path" binding:"required"`
}

func Rename(c *gin.Context) {
	var u RenameBody
	if !request.CheckJsonParams(c, &u) {
		return
	}

	err := fileService.Rename(u.OldName, u.Name, u.Path)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{}, "重命名成功")
}

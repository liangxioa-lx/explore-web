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

type DeleteBody struct {
	Path string `json:"path" binding:"required"`
}

func Delete(c *gin.Context) {
	var u DeleteBody
	if !request.CheckJsonParams(c, &u) {
		return
	}

	err := fileService.Delete(u.Path)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{}, "删除成功")
}

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

type CreateBody struct {
	Name string `json:"name" binding:"required"`
	Path string `json:"path" binding:"required"`
}

func CreateFile(c *gin.Context) {
	var u CreateBody
	if !request.CheckJsonParams(c, &u) {
		return
	}

	err := fileService.CreateFile(u.Name, u.Path)	
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{}, "创建成功")
}

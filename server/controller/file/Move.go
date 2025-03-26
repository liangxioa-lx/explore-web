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

type MoveBody struct {
	Path string `json:"path" binding:"required"`
	NewPath string `json:"newPath" binding:"required"`
}

func Move(c *gin.Context) {
	var u MoveBody
	if !request.CheckJsonParams(c, &u) {
		return
	}

	err := fileService.Move(u.Path, u.NewPath)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{}, "移动成功")
}

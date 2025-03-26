package FileController

import (
	"explore/router/request"
	"explore/router/response"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"fileService"
	"syscall"
)

type FindFileListByPathBody struct {
	Path string `json:"path" binding:"required"`
}

func FindFileListByPath(c *gin.Context) {
	var u FindFileListByPathBody
	if !request.CheckJsonParams(c, &u) {
		return
	}

	fileList, err := fileService.FindFileListByPath(u.Path)	
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	response.Success(c, gin.H{"list": fileList}, "查询成功")	
}

// 查询驱动盘列表
func FindDriverList(c *gin.Context) {
	// 获取所有可用的驱动器盘符
	drives := make([]string, 0)
	
	// 遍历所有可能的驱动器字母（A-Z）
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		drivePath := string(drive) + ":\\"
		
		// 获取驱动器类型
		driveType := syscall.GetDriveType(syscall.StringToUTF16Ptr(drivePath))
		
		// 只添加固定磁盘和可移动磁盘
		if driveType == syscall.DRIVE_FIXED || driveType == syscall.DRIVE_REMOVABLE {
			// 检查驱动器是否可访问
			if _, err := os.Stat(drivePath); err == nil {
				drives = append(drives, drivePath)
			}
		}
	}

	response.Success(c, gin.H{"drives": drives}, "查询成功")
}
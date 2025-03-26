package fileService

import (
	"log"
	"os"
	"path/filepath"
	"errors"
)

/**
 * 通过路径查找文件列表
 * @param path
 * @return []os.FileInfo, error
 */
func FindFileListByPath(path string) ([]os.FileInfo, error) {
	// 清理和验证文件名
	filePath := filepath.Clean(path)
	
	// 检查目录是否存在
	if _, err := os.Stat(filePath); err != nil {
		log.Printf("目录不存在: %s", filePath)
		return nil, errors.New("目录不存在")
	}

	// 读取目录
	fileList, err := os.ReadDir(filePath)
	if err != nil {
		log.Printf("读取目录失败: %v", err)
		return nil, errors.New("读取目录失败")
	}

	// 返回文件列表
	var files []os.FileInfo
	for _, file := range fileList {
		if !file.IsDir() {
			files = append(files, file)
		}
	}
	return files, nil
}

/**
 * 获取文件详细信息
 * @param path
 * @return os.FileInfo, error
 */
func FindFileDetail(path string) (os.FileInfo, error) {
	// 清理和验证文件名
	filePath := filepath.Clean(path)
	
	// 检查文件是否存在
	if _, err := os.Stat(filePath); err != nil {
		log.Printf("文件不存在: %s", filePath)
		return nil, errors.New("文件不存在")
	}

	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Printf("获取文件信息失败: %v", err)
		return nil, errors.New("获取文件信息失败")
	}

	// 返回文件信息
	return fileInfo, nil
}
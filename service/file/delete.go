package fileService

import (
	"log"
	"os"
	"path/filepath"
	"errors"
)

/**
 * 删除文件
 * @param path
 * @return error
 */
func DeleteFile(path string) error {
	// 清理和验证文件名
	filePath := filepath.Clean(path)
	
	// 检查文件是否存在
	if _, err := os.Stat(filePath); err != nil {
		log.Printf("文件不存在: %s", filePath)
		return errors.New("文件不存在")
	}

	// 删除文件
	err := os.Remove(filePath)
	if err != nil {
		log.Printf("删除文件失败: %v", err)
		return errors.New("删除文件失败")
	}
	return nil
}

/**
 * 删除目录
 * @param path
 * @return error
 */
func DeleteDirectory(path string) error {
	// 清理和验证目录名
	dirPath := filepath.Clean(path)
	
	// 检查目录是否存在
	if _, err := os.Stat(dirPath); err != nil {
		log.Printf("目录不存在: %s", dirPath)
		return errors.New("目录不存在")
	}

	// 删除目录
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Printf("删除目录失败: %v", err)
		return errors.New("删除目录失败")
	}
	return nil
}

func Delete (path string) error {
	// 判断文件类型
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Printf("获取文件信息失败: %v", err)
		return errors.New("获取文件信息失败")
	}

	// 删除文件
	if fileInfo.IsDir() {
		return DeleteDirectory(path)
	} else {
		return DeleteFile(path)
	}
}
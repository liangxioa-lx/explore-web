package fileService

import (
	"log"
	"os"
	"path/filepath"
	"errors"
)

/**
 * 创建文件
 * @param name
 * @param path
 * @return error
 */
func CreateFile(name string, path string) error {
	// 清理和验证文件名
	filePath := filepath.Clean(path + "/" + name)
	
	// 检查文件是否已存在
	if _, err := os.Stat(filePath); err == nil {
		log.Printf("文件已存在: %s", filePath)
		return errors.New("文件已存在")
	}

	// 检查目录是否存在
	if _, err := os.Stat(filepath.Dir(filePath)); err != nil {
		log.Printf("目录不存在: %s", filePath)
		return errors.New("目录不存在")
	}

	// 检查文件名是否合法
	if !filepath.Base(filePath).IsValidUTF8() {
		log.Printf("文件名不合法: %s", filePath)
		return errors.New("文件名不合法")
	}	

	// 创建文件
	err := os.Create(filePath)
	if err != nil {
		log.Printf("创建文件失败: %v", err)
		return errors.New("创建文件失败")
	}
	return nil
}

/**
 * 创建目录
 * @param name
 * @param path
 * @return error
 */
func CreateDirectory(name string, path string) error {
	// 清理和验证目录名
	dirPath := filepath.Clean(path + "/" + name)
		
	// 检查目录是否已存在
	if _, err := os.Stat(dirPath); err == nil {
		log.Printf("目录已存在: %s", dirPath)
		return errors.New("目录已存在")
	}

	// 检查目录名是否合法
	if !filepath.Base(dirPath).IsValidUTF8() {
		log.Printf("目录名不合法: %s", dirPath)
		return errors.New("目录名不合法")
	}

	// 创建目录
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		log.Printf("创建目录失败: %v", err)
		return errors.New("创建目录失败")
	}
	return nil
}
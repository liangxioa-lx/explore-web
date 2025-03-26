package fileService

import (
	"log"
	"os"
	"path/filepath"
	"errors"
)

/**
 * 重命名文件、文件夹
 * @param oldPath
 * @param newPath
 * @return error
 */
func Rename(oldPath string, newPath string) error {
	// 清理和验证文件名
	path := filepath.Clean(oldPath)
	newPath := filepath.Clean(newPath)
	
	// 检查原路径是否存在
	if _, err := os.Stat(path); err != nil {
		log.Printf("原路径不存在: %s", path)
		return errors.New("原路径不存在")
	}

	// 检查新路径是否存在
	if _, err := os.Stat(newPath); err == nil {
		log.Printf("新路径已存在: %s", newPath)
		return errors.New("新路径已存在")
	}

	// 重命名
	err := os.Rename(path, newPath)
	if err != nil {
		log.Printf("重命名失败: %v", err)
		return errors.New("重命名失败")
	}
	return nil
}
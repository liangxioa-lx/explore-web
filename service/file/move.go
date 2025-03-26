package fileService

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

/**
 * 移动文件、文件夹
 * @param oldPath
 * @param newPath
 * @return error
 */
func Move(oldPath string, newPath string) error {
	// 清理和验证文件名
	path := filepath.Clean(oldPath)
	newPath := filepath.Clean(newPath)

	// 检查原路径是否存在
	if _, err := os.Stat(path); err != nil {
		log.Printf("原路径不存在: %s", path)
		return errors.New("原路径不存在")
	}

	// 检查目标路径的父目录是否存在，不存在则创建
	destDir := filepath.Dir(newPath)
	if err := os.MkdirAll(destDir, 0755); err != nil {
		log.Printf("创建目标目录失败: %v", err)
		return errors.New("创建目标目录失败")
	}

	// 检查目标路径是否已存在文件
	if _, err := os.Stat(newPath); err == nil {
		log.Printf("目标路径已存在文件: %s", newPath)
		return errors.New("目标路径已存在文件")
	}

	// 打开源文件
	srcFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// 复制文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	// 删除源文件
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
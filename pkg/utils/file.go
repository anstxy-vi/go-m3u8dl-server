package utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

// IsExist 是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir 是否目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CWD 获取执行目录
func CWD() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}

// MkdirWhenNotExist 不存在则创建目录
func MkdirWhenNotExist(path string) error {
	if !IsDir(path) {
		return os.MkdirAll(path, os.ModeDir|fs.ModePerm)
	}
	return nil
}

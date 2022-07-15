package utils

import (
	"errors"
	"os"
	"path/filepath"
)

// GetUserConfigPath 从HOME变量获取用户目录路径
func GetUserConfigPath(path string) (string, error) {
	var home = os.Getenv("HOME")
	if home == "" || len(home) == 0 {
		return "", errors.New("GetUserConfigPath: HOME 变量不存在")
	}

	var userConfigPath = filepath.Join(home, ".config", path)
	return userConfigPath, nil
}

// FileExists 文件是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// FileExists 文件是否存在
func FileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return !fi.IsDir()
}

// DirExists 文件夹是否存在
func DirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return !fi.IsDir()
}

// IsFile
func IsFile(path string) bool {
	if !Exists(path) {
		return false
	}
	fi, _ := os.Stat(path)
	return !fi.IsDir()
}

// IsDir
func IsDir(path string) bool {
	if !Exists(path) {
		return false
	}
	fi, _ := os.Stat(path)
	return fi.IsDir()
}

// FileUserConfigExists 配置文件是否存在
func FileUserConfigExists(path string) bool {
	userConfigPath, err := GetUserConfigPath(path)
	if err != nil {
		return false
	}
	return FileExists(userConfigPath)
}

package configmanager

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// $HOME/.config
	configPrefix = ".config"
)

type Manager interface {
	GetUserConfigPath() string
	GetUserConfigFilePath() string
	UserConfigFileExists() bool
	LoadUserConfigFile() ([]byte, error)
}

type ConfigManager struct {
	Prefix     string // 目录前缀，$HOME/.config/prefix
	ConfigFile string // 配置文件名称
}

// GetUserConfigPath 获取用户目录前缀路径
func (m *ConfigManager) GetUserConfigPath() (string, error) {
	s := os.Getenv("HOME")
	if s == "" {
		return "", errors.New("path is nil")
	}
	path := filepath.Join(s, configPrefix, m.Prefix)
	return path, os.MkdirAll(path, os.FileMode(0755))
}

// GetUserConfigFilePath 获取用户目录前缀的配置文件路径
func (m *ConfigManager) GetUserConfigFilePath() (string, error) {
	dir, err := m.GetUserConfigPath()
	if err != nil {
		return "", err
	}
	if m.ConfigFile == "" {
		return "", errors.New("configfile is empty")
	}
	return filepath.Join(dir, m.ConfigFile), nil
}

// UserConfigFileExists 获取配置文件是否存在
func (m *ConfigManager) UserConfigFileExists() (bool, error) {
	path, err := m.GetUserConfigFilePath()
	if err != nil {
		return false, err
	}
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, errors.New("file does not exist")
		}
	}
	return !info.IsDir(), nil
}

// LoadUserConfigFile 加载配置文件
func (m *ConfigManager) LoadUserConfigFile() ([]byte, error) {
	path, err := m.GetUserConfigFilePath()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(path)
}

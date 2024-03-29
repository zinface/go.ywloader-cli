package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strings"

	"gitee.com/zinface/ywloader-cli/utils"
)

const (
	RefreshFilesExplorer = "workbench.files.action.refreshFilesExplorer"
	ReloadWindow         = "workbench.action.reloadWindow"
)

type Shell struct {
	// 延时操作
	Delay int `json:"delay"`
	// commands: {} 或 commads: []
	// 全部默认为 [] 吧
	Commands []string `json:"commands"`
}

type FileItem struct {
	Name       string `json:"name"`
	Base64     string `json:"base64"`
	Permission int32  `json:"perm"`
}
type FileItems []FileItem

type Template struct {
	Message    string    `json:"message"`
	Action     string    `json:"action"`
	Shell      Shell     `json:"shell"`
	Files      FileItems `json:"files,omitempty"`
	VSCommands []string  `json:"commands,omitempty"`
}

type Youwant struct {
	Label    string   `json:"label"`
	Detail   string   `json:"detail"`
	Type     string   `json:"type"`
	Template Template `json:"template"`
}

func NewYouwant() Youwant {
	return Youwant{
		Label:  "",
		Detail: "",
		Type:   "",
		Template: Template{
			Action: "shell",
			Shell: Shell{
				Delay:    0,
				Commands: []string{},
			},
			VSCommands: []string{},
		},
	}
}

func (youwant Youwant) IsValid() bool {
	if len(youwant.Label) == 0 ||
		len(strings.TrimSpace(youwant.Label)) == 0 {
		return false
	}
	return true
}

// ToJsonString 将条目项转为 json 字符串
//
// 此函数可能会将部分内容进行转义所以很少使用
func (youwant *Youwant) ToJson() []byte {
	data, _ := json.MarshalIndent(youwant, "", "    ")
	return data
}

// ToJsonString 将条目项转为标准的 json 字符串
func (youwant *Youwant) ToJsonString() string {
	var data []byte
	var buffer = bytes.NewBuffer(data)
	var encoder = json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "    ")
	encoder.Encode(youwant)
	return string(buffer.String())
}

// === 用于处理多个 条目项容器 操作的部分
type Youwants []Youwant

// ToJsonString 将条目项转为标准的 json 字符串
func (wants Youwants) ToJsonString() string {
	var data []byte
	var buffer = bytes.NewBuffer(data)
	var encoder = json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "    ")
	encoder.Encode(&wants)
	return string(buffer.String())
}

// SaveFile 将条目项存储到文件
func (wants Youwants) SaveFile(filename string) error {
	return os.WriteFile(filename, []byte(wants.ToJsonString()), 0644)
}

// LoaderYouwantsFromFile 从配置文件中加载内容结构
func LoaderYouwantsFromFile(path string) (Youwants, error) {
	// 1. 确保配置文件存在
	if utils.FileExists(path) {
		// 2. 直接读取所有内容(内容实现自动打开、读取、关闭等操作)，并确保没有异常，否则直接停止
		data, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		// 3. 准备条目项容器结构，并将配置内容以序列化形式存储到条目项容器，确保不存在错误，否则直接停止
		var wants Youwants
		if err := json.Unmarshal(data, &wants); err != nil {
			panic(err)
		}
		return wants, nil
	}
	return Youwants{}, errors.New("LoaderYouwantsFromFile: File Not Found:" + path)
}

// bf := bytes.NewBuffer([]byte{})
// jsonEncoder := json.NewEncoder(bf)
// jsonEncoder.SetEscapeHTML(false)
// jsonEncoder.Encode(htmlJson)
// fmt.Println("第二种解决办法：", bf.String())

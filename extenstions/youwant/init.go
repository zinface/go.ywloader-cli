package youwant

import (
	"os"

	"github.com/spf13/cobra"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/utils"
)

var ilog = logs.Logs{
	Prefix: "init",
}

// initGenerateYwloaderFile 初始化配置
func initGenerateYwloaderFile(path string) bool {
	// 如果没有此文件将会创建一个 youwant.json
	if !utils.FileExists(path) {
		file, err := os.Create(path)
		if err != nil {
			ilog.FileCreatedFail(err.Error())
			return false
		}
		// 在函数结束后执行文件关闭
		defer file.Close()
		// 文件创建并打开后写入默认内容
		file.WriteString("[]")
		return true
	}
	return false
}

// InitHandler 初始化配置文件
func InitHandler(cmd *cobra.Command, args []string) {
	ilog.Println("初始化配置文件")
	// 将初始化配置文件(youwant.json)，成功将打印完成，失败将打印文件已存在
	if initGenerateYwloaderFile(youwant_json) {
		ilog.ConfigCreatedSuccess(youwant_json)
	} else {
		ilog.ConfigIsExist(youwant_json)
	}
}

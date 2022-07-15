package youwant

import (
	"os"

	"github.com/spf13/cobra"

	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/utils"
)

var ilog = logs.Logs{
	Prefix: "init",
}

// initGenerateYwloaderFile 初始化配置
func initGenerateYwloaderFile(path string) bool {
	// 没有就创建一个 youwant.json
	if !utils.FileExists(path) {
		file, err := os.Create(path)
		if err != nil {
			ilog.FileCreatedFail(err.Error())
			return false
		}
		defer file.Close()
		file.WriteString("[]")
		return true
	}
	return false
}

func InitHandler(cmd *cobra.Command, args []string) {
	ilog.Println("初始化配置文件")
	if initGenerateYwloaderFile(youwant_json) {
		ilog.ConfigCreatedSuccess(youwant_json)
	} else {
		ilog.ConfigIsExist(youwant_json)
	}
}

package youwant

import (
	"fmt"

	"github.com/spf13/cobra"

	cm "gitee.com/zinface/go.ywloader-cli/extenstions/configmanager"
	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/models"
	"gitee.com/zinface/go.ywloader-cli/utils"
)

const (
	youwant_json = "youwant.json"
)

var ylog = &logs.Logs{
	Prefix: "youwant",
}

var ycm = &cm.ConfigManager{
	Prefix:     "ywloader",
	ConfigFile: youwant_json,
}

var async bool = false
var global bool = false

// configArguments 配置内部参数
func configArguments(cmd *cobra.Command, args []string) {
	async, _ = cmd.Flags().GetBool("async")
	global, _ = cmd.Flags().GetBool("global")
}

// useConfigFilePath 获取使用的配置文件
func useConfigFilePath(cmd *cobra.Command) string {
	configArguments(cmd, []string{})
	if global {
		global_youwant_json, err := ycm.GetUserConfigFilePath()
		if err != nil {
			return youwant_json
		}
		return global_youwant_json
	}
	return youwant_json
}

// useConfigFilePathDefaultLocal 获取以本地优先的配置文件
func useConfigFilePathDefaultLocal(cmd *cobra.Command) string {
	configArguments(cmd, []string{})
	global_youwant_json, err := ycm.GetUserConfigFilePath()
	if global {
		if err != nil {
			return youwant_json
		}
		return global_youwant_json
	} else {
		if !utils.FileExists(youwant_json) {
			global = true
			return global_youwant_json
		}
	}
	return youwant_json
}

// loaderYouwants 读取配置内容
func loaderYouwants(cmd *cobra.Command) (models.Youwants, error) {
	useConfigFile := useConfigFilePathDefaultLocal(cmd)

	if !utils.FileExists(useConfigFile) {
		ylog.FileNotExits(useConfigFile)
		return nil, fmt.Errorf("配置文件不存在: %v", useConfigFile)
	}

	return models.LoaderYouwantsFromFile(useConfigFile)

}

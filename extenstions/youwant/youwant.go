package youwant

import (
	"fmt"

	"github.com/spf13/cobra"

	cm "gitee.com/zinface/ywloader-cli/extenstions/configmanager"
	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/models"
	"gitee.com/zinface/ywloader-cli/utils"
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
//
// 尝试配置 -g/--global, -a/--async 等参数， 并应用到 global/async 全局变量
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
//
// 逻辑较为复杂，在某些情况下算是有不算bug的bug。
// 例如在未指定 -g 时却因本地配置不存在而使用了全局配置，进行了除添加之外的修改操作而在本地创建了副本。
func useConfigFilePathDefaultLocal(cmd *cobra.Command) string {
	// 尝试配置 -g/--global 等参数， 并应用到 global 全局变量
	configArguments(cmd, []string{})

	// 尝试获取用户配置文件路径
	global_youwant_json, err := ycm.GetUserConfigFilePath()
	// 如果明确指定了 -g 参数，并且全局配置没有异常将使用'全局配置文件路径'
	if global {
		// 明确指定了 -g 参数，但出现了异常
		if err != nil { // 将尝试使用本地配置文件路径，这是无奈的逻辑
			return youwant_json
		}
		// 全局配置文件路径
		return global_youwant_json
	} else {
		// 如果本地配置文件不存在，将转为使用全局配置文件
		if !utils.FileExists(youwant_json) {
			global = true
			return global_youwant_json
		}
	}
	// 默认将使用本地文件路径
	return youwant_json
}

// loaderYouwants 读取配置内容
func loaderYouwants(cmd *cobra.Command) (models.Youwants, error) {
	// 获取魔法配置文件位置
	useConfigFile := useConfigFilePathDefaultLocal(cmd)

	// 如果配置文件不存在将打印配置文件不存在
	if !utils.FileExists(useConfigFile) {
		ylog.FileNotExits(useConfigFile)
		return nil, fmt.Errorf("配置文件不存在: %v", useConfigFile)
	}

	return models.LoaderYouwantsFromFile(useConfigFile)

}

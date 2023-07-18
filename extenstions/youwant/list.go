package youwant

import (
	"fmt"

	"github.com/spf13/cobra"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/models"
	"gitee.com/zinface/ywloader-cli/utils"
)

var llog = &logs.Logs{
	Prefix: "list",
}

// ListHandler 列出可用项
func ListHandler(cmd *cobra.Command, args []string) {
	// 获取魔法配置文件位置
	useConfigFile := useConfigFilePathDefaultLocal(cmd)

	// 如果配置文件不存在将打印配置文件不存在
	if !utils.FileExists(useConfigFile) {
		llog.ConfigNotExist(useConfigFile)
		return
	}

	// 获取所有条目信息
	wants, err := models.LoaderYouwantsFromFile(useConfigFile)
	if err != nil { // 出现异常将进行停止
		panic(err)
	}
	// 打印 编号，指令数量，文件数量，条目名称
	for i := 0; i < len(wants); i++ {
		var want = wants[i]
		fmt.Printf("%v: (%v个指令,%v个文件): %v \n", i, len(want.Template.Shell.Commands), len(want.Template.Files), want.Label)
	}
}

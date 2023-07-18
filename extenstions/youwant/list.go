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

// listWantFiles 仅打印出文件列表
func listWantFiles(want models.Youwant) {
	for _, file := range want.Template.Files {
		fmt.Println(file.Name)
	}
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

	// 内部支持逻辑，用于提供指定项的文件列表
	// 	如果指定了 -f 或 --files 将仅提供其文件列表
	useFiles, _ := cmd.Flags().GetBool("files")

	if useFiles && len(args) != 0 {
		// var arrayForInfomation []string
		for _, want := range wants {
			if want.Label == args[0] {
				listWantFiles(want)
				return
			}
		}

		// 为了达到某些功能的绝对性，我们在此处跳出函数
		return
	}

	// 打印 编号，指令数量，文件数量，条目名称
	for i := 0; i < len(wants); i++ {
		var want = wants[i]
		fmt.Printf("%v: (%v个指令,%v个文件): %v \n", i, len(want.Template.Shell.Commands), len(want.Template.Files), want.Label)
	}
}

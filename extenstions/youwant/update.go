package youwant

import (
	"fmt"
	"os"
	"strings"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/models"
	"gitee.com/zinface/ywloader-cli/utils"
	"github.com/spf13/cobra"
)

var ulog = &logs.Logs{
	Prefix: "update",
}

// UpdateHandler 指令实现
func UpdateHandler(cmd *cobra.Command, args []string) {
	// 使用的配置文件
	var useConfigFile = useConfigFilePathDefaultLocal(cmd)
	ulog.UseConfig(useConfigFile)

	want, err := useWant(cmd, args)
	if err != nil {
		ulog.Println(err.Error())
		os.Exit(1)
	}

	wants, err := loaderYouwants(cmd)
	if err != nil {
		ulog.Println(err.Error())
		os.Exit(1)
	}

	// 从参数列表中分析的文件列表
	var files models.FileItems = addFromCommandFileFlags(cmd, args)
	if len(files) == 0 {
		ulog.Println("警告: 匹配到 0 项可更新内容, 忽略本次操作")
		os.Exit(1)
	}

	// 提取模板中已存在的文件
	var updateFiles models.FileItems
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(want.Template.Files); j++ {
			if files[i].Name == want.Template.Files[j].Name {
				updateFiles = append(updateFiles, files[i])
			}
		}
	}

	// 处理意外被发现的文件
	for i := 0; i < len(files); i++ {
		var unstored = true
		for j := 0; j < len(want.Template.Files); j++ {
			if files[i].Name == updateFiles[j].Name {
				unstored = false
				break
			}
		}
		if unstored {
			question := fmt.Sprintf("发现意外文件: '%s' 是否加入更新?(N/y)", files[i].Name)
			var answer = utils.GetStdinStringValue(question, "")
			if strings.Contains(answer, "y") {
				updateFiles = append(updateFiles, files[i])
				ulog.Println("已加入模板文件更新队列")
			} else {
				ulog.Println(fmt.Sprintf("未确认，已忽略文件: %s", files[i].Name))
			}
		}
	}

	// 最终确认是否更新
	question := fmt.Sprintf("> NOTE: 你确定要对 '%s' 进行更新吗?(N/y)", want.Label)
	var answer = utils.GetStdinStringValue(question, "")
	if strings.Contains(answer, "y") {
		for i := 0; i < len(wants); i++ {
			if compare(want, wants[i]) {
				wants[i].Template.Files = updateFiles
				break
			}
		}
		if err = wants.SaveFile(useConfigFile); err == nil {
			ulog.Println("项目更新成功")
		}
	} else {
		ulog.Println("未确认，放弃更新.")
	}
}
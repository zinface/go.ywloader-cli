package youwant

import (
	"fmt"

	"github.com/spf13/cobra"

	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/models"
	"gitee.com/zinface/go.ywloader-cli/utils"
)

var llog = &logs.Logs{
	Prefix: "list",
}

func ListHandler(cmd *cobra.Command, args []string) {
	useConfigFile := useConfigFilePath(cmd)

	if !utils.FileExists(useConfigFile) {
		llog.ConfigNotExist(useConfigFile)
		return
	}

	wants, err := models.LoaderYouwantsFromFile(useConfigFile)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(wants); i++ {
		var want = wants[i]
		fmt.Printf("%v: (%v个指令,%v个文件): %v \n", i, len(want.Template.Shell.Commands), len(want.Template.Files), want.Label)
	}
}

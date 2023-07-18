package youwant

import (
	"os"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"github.com/spf13/cobra"
)

var clog = logs.Logs{
	Prefix: "change",
}

func ChangeHandler(cmd *cobra.Command, args []string) {
	var useConfigFile = useConfigFilePath(cmd)

	want, err := useWant(cmd, args)
	if err != nil {
		clog.Println(err.Error())
		os.Exit(1)
	}
	wants, err := loaderYouwants(cmd)
	if err != nil {
		clog.Println(err.Error())
		os.Exit(1)
	}

	for i := 0; i < len(wants); i++ {
		if compare(want, wants[i]) {
			add := addYouwantItem()
			add.Template.Files = addFromCommandFileFlags(cmd, args)
			wants[i] = add
		}
	}

	if err = wants.SaveFile(useConfigFile); err == nil {
		clog.Println("项目修改成功")
	}
}

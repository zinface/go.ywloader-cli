package youwant

import (
	"fmt"
	"strconv"

	"gitee.com/zinface/go.ywloader-cli/utils"
	"github.com/spf13/cobra"
)

func ShowHandler(cmd *cobra.Command, args []string) {
	wants, err := loaderYouwants(cmd)
	if err != nil {
		panic(err)
	}

	// 默认不处理
	var index UseIndex = UnselectUse

	// 处理可能的参数
	if len(args) == 1 {
		s := args[0]
		i, err := strconv.Atoi(s)
		if err == nil && UseIndex(i) > UnselectUse && i < len(wants) {
			index = UseIndex(i)
		}
	}

	// 如果未使用参数
	if index == UnselectUse {
		for i := 0; i < len(wants); i++ {
			var want = wants[i]
			fmt.Printf("%v: (%v个指令): %v \n", i, len(want.Template.Shell.Commands), want.Label)
		}
		index = UseIndex(utils.GetStdinNumberValue("输入你想执行的编号:", int(UnselectUse)))
	}

	want := wants[index]

	fmt.Printf("Label:  %v\n", want.Label)
	fmt.Printf("Detail: %v\n", want.Detail)
	fmt.Printf("Type:   %v\n", want.Type)
	fmt.Printf("Action: %v\n", want.Template.Action)
	fmt.Printf("Shell.Commands:  %v\n", "")
	for _, command := range want.Template.Shell.Commands {
		fmt.Printf("    command:   %v\n", command)
	}

	fmt.Printf("Files:  %v\n", "")
	for _, file := range want.Template.Files {
		fmt.Printf("    file.Name: %v\n", file.Name)
	}

}

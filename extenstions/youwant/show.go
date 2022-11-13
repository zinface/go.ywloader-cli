package youwant

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 打印指定条目的基本信息
func ShowHandler(cmd *cobra.Command, args []string) {
	// 从命令行参数中获取指定的条目信息
	want, err := useWant(cmd, args)
	if err != nil { // 出现异常将进行停止
		panic(err)
	}

	// 按一定顺序结构打印条目结构信息
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

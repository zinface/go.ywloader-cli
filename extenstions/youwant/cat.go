package youwant

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

// CatHandler 指令实现
//
// 此指令用于实现查看指定条目下的指定文件
//
// 要求传入的参数为 cat <条目> <条目文件>
func CatHandler(cmd *cobra.Command, args []string) {

	// 检查 cat 传入的参数

	if len(args) < 2 {
		fmt.Println("应提供 2 个参数")
		return
	}

	// 从传入的参数中获取指定的条目信息
	want, err := useWant(cmd, args)
	if err != nil { // 出现异常将进行停止，并且不输出任何内容
		return
	}

	// 从模板条目的文件列表中输出文件
	for _, file := range want.Template.Files {
		if args[1] == file.Name {
			b, err := base64.StdEncoding.DecodeString(file.Base64)
			if err != nil {
				fmt.Println("文件内容解析出错")
			}
			fmt.Printf("%s\n", b)
		}
	}
}

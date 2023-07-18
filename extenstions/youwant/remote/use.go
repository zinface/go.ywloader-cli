package remote

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/extenstions/youwant"
	"gitee.com/zinface/ywloader-cli/models"
	"gitee.com/zinface/ywloader-cli/utils"
	"github.com/spf13/cobra"
)

var uselog = &logs.Logs{
	Prefix: "remote-use",
}

func useWants(cmd *cobra.Command, args []string) models.Youwants {
	// 1. 如果参数中有指定多个要使用的条目项
	// if len(args) > 1 {
	// 2. 将匹配的项进行依次取出条目
	// 	for _, arg := range args {
	// 		if want, err := useWant(cmd, []string{arg}); err == nil {
	// 			wants = append(wants, want)
	// 		}
	// 	}
	// } else {
	// 2. 认为只有一个，将直接进行取出条目
	// 	if want, err := useWant(cmd, args); err == nil {
	// 		wants = append(wants, want)
	// 	}
	// }

	wants, err := getRemoteWantsWithLabels(args)
	if err != nil {
		panic(err)
	}

	// 3. 收集已取出的条目信息标题
	var labels = []string{}
	for _, want := range wants {
		labels = append(labels, fmt.Sprintf("'%v'", want.Label))
	}
	// 4. 打印为选中项日志
	uselog.Println(fmt.Sprintf("选中%v项: [%v]", len(labels), strings.Join(labels, ",")))

	return wants
}

// applyWant 准备执行执行已传入(指定)的条目信息
func applyWant(want models.Youwant) {
	// ======== 准备处理 文件 ========
	if len(want.Template.Files) != 0 {
		fmt.Println("--------------------------------")
	}

	// 1. 预打印文件信息
	for _, file := range want.Template.Files {
		s := fmt.Sprintf("文件: %v", file.Name)
		uselog.Println(s)
	}

	// 2. 当文件信息数量大于 0 时，将询问是否处理文件，如果当需要处理的文件已存在过将询问是否继续处理(覆盖)
	if len(want.Template.Files) != 0 {

		for {
			var _continue = utils.GetStdinStringValue("处理文件: 是否继续(yes/no):", "")
			if strings.Contains(_continue, "yes") {
				for i := 0; i < len(want.Template.Files); i++ {
					file := want.Template.Files[i]

					// 默认值为继续，当发生已存在过将暂停并询问继续
					var _continue = true
					if utils.FileExists(file.Name) {
						_continue = false
						// uselog.Print()
						question := fmt.Sprintf("NOTE: 文件已存在(%v), 是否继续(y/n)?", file.Name)
						var answer = utils.GetStdinStringValue(question, "y")
						if strings.Contains(answer, "y") {
							_continue = true
						}
					}

					// 当得到默认继续或询问到已继续时，将进行解出文件内容
					if _continue {
						// 1. 从条目信息中预解出文件内容，将 base64 解码为文件内容
						b, err := base64.StdEncoding.DecodeString(file.Base64)
						if err != nil {
							uselog.Print("处理失败", err.Error())
							continue
						}

						// 2. 处理文件所在目录路径，还原当时文件的目录路径
						// 如果路径不是 '.' 即可开始创建目录
						dirPath := filepath.Dir(file.Name)
						if dirPath != "." {
							os.MkdirAll(dirPath, 0755)
						}

						// 3. 处理文件，还原当时的文件
						// 以正确的文件创建方式，来创建文件: O_RDWR|O_CREATE|O_TRUNC
						f, err := os.Create(file.Name)
						if err != nil {
							uselog.Print("创建失败", err.Error())
							continue
						}
						defer f.Close() // 在函数结束前关闭文件

						f.Write(b) // 在此处将预解码的文件内容写入文件

						// 4. 处理文件权限值，还原当时创建文件时的权限
						if file.Permission != 0 { //
							f.Chmod(os.FileMode(file.Permission))
						}
						uselog.Print("已处理", file.Name)
					}
					if !_continue {
						uselog.Print("放弃处理", file.Name)
					}
				}
				break
			}
			if strings.Contains(_continue, "no") {
				uselog.Println("放弃处理文件.")
				break
			}
		}
	}

	// ======== 准备处理 指令 ========
	if len(want.Template.Shell.Commands) != 0 {
		fmt.Println("--------------------------------")
	}

	// 1. 预打印指令信息
	for _, command := range want.Template.Shell.Commands {
		s := fmt.Sprintf("指令: %v", command)
		uselog.Println(s)
	}

	// 2. 当指令信息数量不为 0 时
	if len(want.Template.Shell.Commands) != 0 {
		for {
			// 1. 询问是否处理指令集合
			var _continue = utils.GetStdinStringValue("处理指令: 是否继续(yes/no):", "")
			if strings.Contains(_continue, "yes") {
				for _, command := range want.Template.Shell.Commands {
					// 2. 由 RunCommand 来实现指令的执行
					youwant.RunCommand(command)
				}
				break
			}
			if strings.Contains(_continue, "no") {
				uselog.Println("放弃处理指令集.")
				break
			}
		}
	}
}

// UseHandler remote-use 实现
func UseHandler(cmd *cobra.Command, args []string) {

	// 1. 从远端获取所有条目项
	wants := useWants(cmd, args)
	for _, want := range wants {
		applyWant(want)
	}
}

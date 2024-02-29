package youwant

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/models"
	"gitee.com/zinface/ywloader-cli/utils"
)

type UseIndex int
type UsePosible int

const (
	UnselectUse UseIndex = -1
)

const (
	UseUnknown UsePosible = iota
	UseString
	UseNumber
)

var uselog = &logs.Logs{
	Prefix: "use",
}

// posibleUse 对于参数中不确定项类型来进行处理，返回可能的指定类型(字符串或编号)
func posibleUse(args []string) UsePosible {
	var posible UsePosible = UseUnknown

	// 参数大于0
	if len(args) > 0 {
		// 第一个参数
		arg := args[0]
		// 尝试转数字
		_, err := strconv.Atoi(arg)
		if err != nil {
			// 失败，可能是字符串
			posible = UseString
		} else {
			posible = UseNumber
		}
	}

	return posible
}

// useWant 从命令行参数配置中获取指定条目
func useWant(cmd *cobra.Command, args []string) (models.Youwant, error) {
	var err error
	// 获取所有条目信息
	wants, err := loaderYouwants(cmd)
	if err != nil {
		return models.Youwant{}, err
	}

	// 当已获取条目信息数量大于 0
	if len(wants) > 0 {
		// 默认不处理
		var index UseIndex = UnselectUse

		// 处理可能的use类型
		var posible = posibleUse(args)

		// 如果不为字符串
		if posible != UseString {
			// 处理可能的数字参数
			if len(args) == 1 {
				arg := args[0]
				i, err := strconv.Atoi(arg)
				if err == nil && UseIndex(i) > UnselectUse && i < len(wants) {
					index = UseIndex(i)
				}
			}

			// 如果未能处理为数字，将直接进行询问
			if index == UnselectUse {
				for i := 0; i < len(wants); i++ {
					var want = wants[i]
					fmt.Printf("%v: (%v个指令): %v \n", i, len(want.Template.Shell.Commands), want.Label)
				}
				index = UseIndex(utils.GetStdinNumberValue("输入你想执行的编号:", int(UnselectUse)))
				// 指定为是数字类型
				posible = UseNumber
			}
		}

		if posible == UseString ||
			posible == UseNumber &&
				index > UnselectUse &&
				int(index) < len(wants) {
			var want models.Youwant

			// 如果是已为数字
			if posible == UseNumber {
				want = wants[int(index)]
			}
			// 如果已为字符串
			if posible == UseString {
				for i := 0; i < len(wants); i++ {
					if strings.Compare(wants[i].Label, args[0]) == 0 {
						want = wants[i]
						index = UseIndex(i)
					}
				}
			}
			if index != UnselectUse {
				fmt.Fprintf(os.Stderr, "选中: %v: (%v个指令,%v个文件): %v \n", index, len(want.Template.Shell.Commands), len(want.Template.Files), want.Label)
				return want, nil
			}
			return want, fmt.Errorf("未匹配选择项: %v", args[0])
		} else {
			err = fmt.Errorf("未选中编号: %v", index)
		}
	} else {
		err = errors.New("无项目索引信息")
	}

	return models.Youwant{}, err
}

// useWants  从命令行参数配置中获取指定条目
func useWants(cmd *cobra.Command, args []string) models.Youwants {
	var wants = models.Youwants{}

	// 1. 如果参数中有指定多个要使用的条目项
	if len(args) > 1 {
		// 2. 将匹配的项进行依次取出条目
		for _, arg := range args {
			if want, err := useWant(cmd, []string{arg}); err == nil {
				wants = append(wants, want)
			}
		}
	} else {
		// 2. 认为只有一个，将直接进行取出条目
		if want, err := useWant(cmd, args); err == nil {
			wants = append(wants, want)
		}
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
						if utils.FileCompareBase64(file.Name, file.Base64) {
							uselog.Printf("跳过相同文件: %v", file.Name)
							continue
						}
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
					RunCommand(command)
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

// useHandler 指令实现
func UseHandler(cmd *cobra.Command, args []string) {
	// if _, err := useWant(cmd, args); err != nil {
	// 	uselog.Println(err.Error())
	// 	os.Exit(1)
	// }

	// 1. 从命令行参数配置中获取指定条目
	wants := useWants(cmd, args)
	for _, want := range wants {
		applyWant(want)
	}
}

// RunCommand 执行一条指令
//
// 并支持由命令行参数中提供的 --async 参数来实现指令的异步执行
func RunCommand(command string) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	if async {
		cmd.Start()
	} else {
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

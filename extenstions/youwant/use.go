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

	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/models"
	"gitee.com/zinface/go.ywloader-cli/utils"
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

func useWant(cmd *cobra.Command, args []string) (models.Youwant, error) {
	var err error
	wants, err := loaderYouwants(cmd)
	if err != nil {
		return models.Youwant{}, err
	}

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
				fmt.Printf("%v: (%v个指令,%v个文件): %v \n", index, len(want.Template.Shell.Commands), len(want.Template.Files), want.Label)
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

// useHandler 指令实现
func UseHandler(cmd *cobra.Command, args []string) {
	want, err := useWant(cmd, args)
	if err != nil {
		uselog.Println(err.Error())
		os.Exit(1)
	}

	for _, cmd := range want.Template.Shell.Commands {
		fmt.Printf("指令: %v\n", cmd)
	}
	for _, file := range want.Template.Files {
		fmt.Printf("文件: %v\n", file.Name)
	}

	if len(want.Template.Files) != 0 {
		for i := 0; i < len(want.Template.Files); i++ {
			file := want.Template.Files[i]
			fmt.Printf("> file: %v\n", file.Name)
			if utils.FileExists(file.Name) {
				uselog.Print("NOTE: 文件已存在:", file.Name)
			}
			var _continue = utils.GetStdinStringValue("处理文件: 是否继续(y/n):", "y")
			if strings.Contains(_continue, "y") {
				// 解出文件内容
				b, err := base64.StdEncoding.DecodeString(file.Base64)
				if err != nil {
					uselog.Print("处理失败", err.Error())
					continue
				}
				// 还原文件的目录路径
				dirPath := filepath.Dir(file.Name)
				if dirPath != "." {
					os.MkdirAll(dirPath, 0755)
				}
				// 还原文件
				f, err := os.Create(file.Name)
				if err != nil {
					uselog.Print("创建失败", err.Error())
					continue
				}
				defer f.Close()
				f.Write(b)

				// 还原权限
				if file.Permission != 0 {
					f.Chmod(os.FileMode(file.Permission))
				}
				uselog.Print("已处理:", file.Name)
			}
			if strings.Contains(_continue, "n") {
				uselog.Print("放弃处理", file.Name)
			}
		}
	}

	if len(want.Template.Shell.Commands) != 0 {
		for {
			var _continue = utils.GetStdinStringValue("处理指令: 是否继续(yes/no):", "")
			if strings.Contains(_continue, "yes") {
				for _, cmd := range want.Template.Shell.Commands {
					RunCommand(cmd)
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

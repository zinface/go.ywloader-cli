package remote

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func SearchHandler(cmd *cobra.Command, args []string) {
	// 0. 准备用于输出项目配置类型，检查 --all 参数
	useAll, err := cmd.Flags().GetBool("all")

	// 1. 预检查传入的参数信息
	if !useAll && len(args) < 1 {
		fmt.Fprintf(os.Stderr, "未指定参数\n")
		os.Exit(1)
	}

	// 2. 获取所有条目信息
	wants, err := getRemoteWants()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// 3. 使用 []string 存储存在关键词的条目Label
	var arrayForLabel []string

	// 以数组形式存储 Label 内容
	// 此部分用于 arrayForJson 内容
	for i := 0; i < len(wants); i++ {
		if useAll {
			arrayForLabel = append(arrayForLabel, wants[i].Label)
		} else {
			// TODO：增加多关键词搜索功能
			if strings.Contains(wants[i].Label, args[0]) {
				arrayForLabel = append(arrayForLabel, wants[i].Label)
			}
		}
	}

	// 4. 尝试在此时获取 --json/-j 参数标志
	useJson, err := cmd.Flags().GetBool("json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	// 如果 Flags 中 json 被指定将输出 json 内容
	// 否则直接输出 Label 内容
	if useJson {
		arrayForJson, _ := json.Marshal(arrayForLabel)
		fmt.Println(string(arrayForJson))
	} else {
		for i := 0; i < len(arrayForLabel); i++ {
			fmt.Println(arrayForLabel[i])
		}
	}
}

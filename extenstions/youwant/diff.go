package youwant

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/models"
	"gitee.com/zinface/ywloader-cli/utils"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/spf13/cobra"
)

var difflog = &logs.Logs{
	Prefix: "diff",
}

func DiffHandler(cmd *cobra.Command, args []string) {
	var useConfigFile = useConfigFilePathDefaultLocal(cmd)
	difflog.UseConfig(useConfigFile)

	// 加载命令行参数中指定的项目模板
	want, err := useWant(cmd, args)
	if err != nil {
		difflog.Println(err.Error())
		os.Exit(1)
	}

	// 从参数列表中分析的文件列表
	var files models.FileItems = addFromCommandFileFlags(cmd, args)
	if len(files) == 0 {
		difflog.Println("警告: 匹配到 0 项可diff内容, 忽略本次操作")
		os.Exit(1)
	}

	// 提取模板中已存在的文件
	var diffFiles models.FileItems
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(want.Template.Files); j++ {
			if files[i].Name == want.Template.Files[j].Name {
				diffFiles = append(diffFiles, files[i])
			}
		}
	}

	var fileMissing = false
	for i := 0; i < len(want.Template.Files); i++ {
		var isExist = false
		for j := 0; j < len(files); j++ {
			if want.Template.Files[i].Name == files[j].Name {
				isExist = true
				break
			}
		}

		if !isExist {
			if !fileMissing {
				fileMissing = true
			}
			difflog.Println(fmt.Sprintf("发现缺失的的文件: '%s'", want.Template.Files[i].Name))
		}
	}

	for _, v := range diffFiles {
		difflog.Println(fmt.Sprintf("将进行比较的文件: %v", v.Name))
	}

	for _, v := range diffFiles {
		data, err := ioutil.ReadFile(v.Name)
		if err != nil {
			panic(err)
		}
		base64Str := base64.StdEncoding.EncodeToString(data)

		for i := 0; i < len(want.Template.Files); i++ {
			fi := want.Template.Files[i]
			// if strings.Compare(v.Name, fi.Name) {
			if strings.Compare(v.Name, fi.Name) == 0 {
				if base64Str == fi.Base64 {
					difflog.Println(fmt.Sprintf("相同的文件: %v", fi.Name))
				} else {
					difflog.Println(fmt.Sprintf("不一致的文件: %v", fi.Name))

					str, err := base64.StdEncoding.DecodeString(fi.Base64)
					if err != nil {
						difflog.Print("处理失败", err.Error())
						continue
					}

					var question = fmt.Sprintf("> NOTE: 你要diff '%s' 文件吗?(N/y)", fi.Name)
					var answer = utils.GetStdinStringValue(question, "")
					if strings.Contains(answer, "y") {
						dmp := diffmatchpatch.New()
						d := dmp.DiffMain(string(data), string(str), false)

						fmt.Println(dmp.DiffPrettyText(d))
					}
				}
			}
		}
	}
}

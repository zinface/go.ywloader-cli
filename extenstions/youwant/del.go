package youwant

import (
	"fmt"
	"os"
	"strings"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/models"
	"gitee.com/zinface/ywloader-cli/utils"
	"github.com/spf13/cobra"
)

var dlog = logs.Logs{
	Prefix: "del",
}

func compare(want, b models.Youwant) bool {
	if strings.Compare(want.Label, b.Label) == 0 &&
		strings.Compare(want.Detail, b.Detail) == 0 &&
		strings.Compare(want.Type, b.Type) == 0 &&
		strings.Compare(want.Template.Action, b.Template.Action) == 0 &&
		strings.Compare(want.Template.Message, b.Template.Message) == 0 &&
		len(want.Template.Files) == len(b.Template.Files) &&
		len(want.Template.Shell.Commands) == len(b.Template.Shell.Commands) {
		return true
	}
	return false
}

func DelHandler(cmd *cobra.Command, args []string) {
	var useConfigFile = useConfigFilePath(cmd)

	want, err := useWant(cmd, args)
	if err != nil {
		dlog.Println(err.Error())
		os.Exit(1)
	}
	question := fmt.Sprintf("确定删除项目: '%s' ?(Yes/No)", want.Label)

	for {
		answer := utils.GetStdinStringValue(question, "")

		// 回答 yes
		if strings.Contains(answer, "yes") ||
			strings.Contains(answer, "y") {
			break
		}
		// 回答 no
		if strings.Contains(answer, "no") ||
			strings.Contains(answer, "n") {
			return
		}
	}

	wants, err := loaderYouwants(cmd)
	if err != nil {
		dlog.Println(err.Error())
		os.Exit(1)
	}

	// fmt.Printf("len(wants): %v\n", len(wants))

	var newWants = models.Youwants{}
	for i := 0; i < len(wants); i++ {
		if compare(want, wants[i]) {
			continue
		}
		newWants = append(newWants, wants[i])
	}

	// fmt.Printf("len(newWants): %v\n", len(newWants))
	// for i := 0; i < len(newWants); i++ {
	// fmt.Printf("newWants[i].Label: %v\n", newWants[i].Label)
	// }

	if err = newWants.SaveFile(useConfigFile); err == nil {
		dlog.Println("项目删除成功")
	}
}

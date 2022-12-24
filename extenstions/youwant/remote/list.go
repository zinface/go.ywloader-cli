package remote

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ListHandler
func ListHandler(cmd *cobra.Command, args []string) {
	//
	length, err := getRemoteHeaderContentLength(remoteUrl)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "已接收远端发送的 %d 个字节\n", length)

	wants, err := getRemoteWants()
	if err != nil {
		panic(err)
	}

	// fmt.Printf("远端目前有 %d 个项目\n", len(wants))
	// 打印 编号，指令数量，文件数量，条目名称
	for i := 0; i < len(wants); i++ {
		var want = wants[i]
		fmt.Printf("%v: (%v个指令,%v个文件): %v \n", i, len(want.Template.Shell.Commands), len(want.Template.Files), want.Label)
	}
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/models"
	"github.com/spf13/cobra"
)

var llog = &logs.Logs{
	Prefix: "list",
}

// remoteListCmd represents the remoteList command
var remoteListCmd = &cobra.Command{
	Use:   "list",
	Short: _desc_remote_list,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get(remoteUrl)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		// 此处逻辑无法确定应该接收多少内容，使用 ioutil 来读取内容
		// var buffer = make([]byte, 1024)
		// n, err := response.Body.Read(buffer)

		buffer, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "已接收远端发送的 %d 个字节\n", len(buffer))

		var wants models.Youwants
		err = json.Unmarshal(buffer, &wants)
		if err != nil {
			panic(err)
		}

		// fmt.Printf("远端目前有 %d 个项目\n", len(wants))
		// 打印 编号，指令数量，文件数量，条目名称
		for i := 0; i < len(wants); i++ {
			var want = wants[i]
			fmt.Printf("%v: (%v个指令,%v个文件): %v \n", i, len(want.Template.Shell.Commands), len(want.Template.Files), want.Label)
		}
	},
}

func init() {
	remoteCmd.AddCommand(remoteListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	remoteListCmd.Flags().StringVarP(&remoteUrl, "url", "u", "", "远程仓库源路径")
	remoteListCmd.MarkFlagRequired("url")
}

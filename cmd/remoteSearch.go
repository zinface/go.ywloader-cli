/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gitee.com/zinface/ywloader-cli/extenstions/youwant/remote"
	"github.com/spf13/cobra"
)

const (
	_desc_remote_search = "搜索远程仓库源项目索引"
)

// remoteSearchCmd represents the remoteSearch command
var remoteSearchCmd = &cobra.Command{
	Use:    "search",
	Short:  _desc_remote_search,
	Long:   ``,
	PreRun: remote.PreRunHandler,
	Run:    remote.SearchHandler,
}

func init() {
	remoteCmd.AddCommand(remoteSearchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteSearchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteSearchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	remoteSearchCmd.Flags().BoolP("all", "a", false, "输出全部配置项")
	remoteSearchCmd.Flags().BoolP("json", "j", false, "使用 json 格式输出结果")
	remoteSearchCmd.Flags().StringP("url", "u", "", "远程仓库源路径(http)")
	remoteSearchCmd.MarkFlagRequired("url")
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gitee.com/zinface/ywloader-cli/extenstions/youwant/remote"
	"github.com/spf13/cobra"
)

const (
	_desc_remote_list = "列出远程仓库源项目索引"
)

// remoteListCmd represents the remoteList command
var remoteListCmd = &cobra.Command{
	Use:    "list",
	Short:  _desc_remote_list,
	Long:   ``,
	PreRun: remote.PreRunHandler,
	Run:    remote.ListHandler,
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
	remoteListCmd.Flags().StringP("url", "u", "", "远程仓库源路径(http)")
	remoteListCmd.MarkFlagRequired("url")
}

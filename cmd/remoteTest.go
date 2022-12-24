/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gitee.com/zinface/go.ywloader-cli/extenstions/youwant/remote"
	"github.com/spf13/cobra"
)

const (
	_desc_remote_test = "测试远程仓库源"
)

// remoteTestCmd represents the remoteTest command
var remoteTestCmd = &cobra.Command{
	Use:    "test",
	Short:  _desc_remote_test,
	Long:   ``,
	PreRun: remote.PreRunHandler,
	Run:    remote.TestHandler,
}

func init() {
	remoteCmd.AddCommand(remoteTestCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteTestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteTestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	remoteTestCmd.Flags().StringP("url", "u", "", "远程仓库源路径(http)")
	remoteTestCmd.MarkFlagRequired("url")
}

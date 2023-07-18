/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gitee.com/zinface/ywloader-cli/extenstions/youwant/remote"
	"github.com/spf13/cobra"
)

const (
	_desc_remote_use = "使用远程服务提供的项目模板"
)

// remoteUseCmd represents the remoteUse command
var remoteUseCmd = &cobra.Command{
	Use:    "use",
	Short:  _desc_remote_use,
	Long:   ``,
	PreRun: remote.PreRunHandler,
	Run:    remote.UseHandler,
}

func init() {
	remoteCmd.AddCommand(remoteUseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteUseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteUseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	remoteUseCmd.Flags().StringP("url", "u", "", "远程仓库源路径(http)")
	remoteUseCmd.MarkFlagRequired("url")
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/extenstions/youwant"
	"github.com/spf13/cobra"
)

const (
	_desc_update = "更新一条项目的文件列表"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: _desc_update,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logs.WelcomeMessage("update", _desc_update)
		youwant.UpdateHandler(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	updateCmd.Flags().BoolP("global", "g", false, "编辑全局配置项")
	updateCmd.Flags().StringArrayP("file", "f", []string{}, "添加文件到项目")
}

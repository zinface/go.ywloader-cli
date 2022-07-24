/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/extenstions/youwant"
	"github.com/spf13/cobra"
)

var _desc_change = "修改一条项目"

// changeCmd represents the change command
var changeCmd = &cobra.Command{
	Use:   "change",
	Short: _desc_change,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logs.WelcomeMessage("change", _desc_change)
		youwant.ChangeHandler(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// changeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// changeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	changeCmd.Flags().BoolP("global", "g", false, "编辑全局配置项")
	changeCmd.Flags().StringArrayP("file", "f", []string{}, "添加文件到项目")
}

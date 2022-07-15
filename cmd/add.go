/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/extenstions/youwant"
)

const (
	_desc_add = "添加项目到仓库索引"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: _desc_add,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logs.WelcomeMessage("add", _desc_add)
		youwant.AddHandler(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolP("global", "g", false, "编辑全局配置项")
	addCmd.Flags().StringArrayP("file", "f", []string{}, "添加文件到项目")
}

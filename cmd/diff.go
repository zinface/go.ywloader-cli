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
	_desc_diff = "与指定的项目进行对比"
)

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: _desc_diff,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logs.WelcomeMessage("diff", _desc_diff)
		youwant.DiffHandler(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diffCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diffCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	diffCmd.Flags().BoolP("global", "g", false, "使用全局配置项")
	diffCmd.Flags().StringArrayP("file", "f", []string{}, "添加文件到项目")
}

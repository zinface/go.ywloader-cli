/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"gitee.com/zinface/ywloader-cli/extenstions/logs"
	"gitee.com/zinface/ywloader-cli/extenstions/youwant"
)

const (
	_desc_show = "显示项目详细信息"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: _desc_show,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logs.WelcomeMessage("youwant", _desc_show)
		youwant.ShowHandler(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	showCmd.Flags().BoolP("global", "g", false, "使用全局配置项")
}

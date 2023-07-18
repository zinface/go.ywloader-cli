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
	_desc_use = "使用快速项目模板"
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: _desc_use,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logs.WelcomeMessage("use", _desc_use)
		youwant.UseHandler(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	useCmd.Flags().BoolP("global", "g", false, "全局配置项")
	useCmd.Flags().BoolP("async", "a", false, "异步指令")
}

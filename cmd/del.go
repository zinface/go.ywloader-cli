/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gitee.com/zinface/go.ywloader-cli/extenstions/logs"
	"gitee.com/zinface/go.ywloader-cli/extenstions/youwant"
	"github.com/spf13/cobra"
)

const (
	_desc_del = "删除一条项目"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: _desc_del,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logs.WelcomeMessage("del", _desc_del)
		youwant.DelHandler(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	delCmd.Flags().BoolP("global", "g", false, "使用全局配置项")
}

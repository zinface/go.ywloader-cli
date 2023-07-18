/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"gitee.com/zinface/ywloader-cli/extenstions/youwant"
)

const (
	_desc_list = "列出仓库项目索引"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: _desc_list,
	Long:  ``,
	Run:   youwant.ListHandler,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolP("global", "g", false, "使用全局配置项")
}

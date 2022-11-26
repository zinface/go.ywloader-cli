/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gitee.com/zinface/go.ywloader-cli/extenstions/youwant"
	"github.com/spf13/cobra"
)

const (
	_desc_search = "搜索仓库项目索引"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: _desc_search,
	Long:  ``,
	Run:   youwant.SearchHandler,
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.Flags().BoolP("global", "g", false, "使用全局配置项")
	searchCmd.Flags().BoolP("all", "a", false, "输出全部配置项")
	searchCmd.Flags().BoolP("json", "j", false, "使用 json 格式输出结果")
}

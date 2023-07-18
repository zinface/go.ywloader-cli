/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"gitee.com/zinface/ywloader-cli/extenstions/youwant"
	"github.com/spf13/cobra"
)

const (
	_desc_cat = "查看指定项目的文件"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: _desc_cat,
	Long:  ``,
	Run:   youwant.CatHandler,
}

func init() {
	rootCmd.AddCommand(catCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// catCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// catCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	catCmd.Flags().BoolP("global", "g", false, "使用全局配置项")
}

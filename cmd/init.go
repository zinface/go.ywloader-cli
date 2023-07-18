/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"gitee.com/zinface/ywloader-cli/extenstions/youwant"
)

const (
	_desc_init = "初始化当前目录为 Youwant 仓库"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: _desc_init,
	Long:  ``,
	Run:   youwant.InitHandler,
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

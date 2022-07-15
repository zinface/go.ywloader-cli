/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	_desc_version = "打印版本信息"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: _desc_version,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("版本: v0.0.1")
		fmt.Println("作者: zinface <zinface@163.com>")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// versionCmd.Flags().StringP("author", "a", "unknown", "作者信息")
}

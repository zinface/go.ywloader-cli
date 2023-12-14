/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"gitee.com/zinface/ywloader-cli/extenstions/web/simpleserver"
	"github.com/spf13/cobra"
)

const (
	_desc_serve = "启动一个web服务(默认:8080)"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: _desc_serve,
	Long:  ``,
	Run:   simpleserver.SimpleServeHandler,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// serveCmd.Flags().BoolP("global", "g", false, "直接使用全局")
	serveCmd.Flags().IntP("port", "p", 8080, "端口号")
}

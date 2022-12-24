/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

const (
	_desc_remote = "远程仓库命令"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: _desc_remote,
	Long:  ``,
}

func init() {
	rootCmd.AddCommand(remoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

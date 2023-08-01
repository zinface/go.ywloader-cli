/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	_desc_completion = "bash补全脚本"
)

var BashCompletion string

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: _desc_completion,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(BashCompletion)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

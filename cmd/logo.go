/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

const (
	_desc_logo = "打印可执行程序 logo 头部"
)

// logoCmd represents the logo command
var logoCmd = &cobra.Command{
	Use:   "logo",
	Short: _desc_logo,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		str, _ := base64.RawStdEncoding.DecodeString("ICAgICAgICAgICAgICAgICAgICAsLS0uICAgICAgICAgICAgICAgICAgLC0tLiAgICAgICAgICAgICAgIAosLS0uICwtLS4sLS0uICAgLC0tLnwgIHwgLC0tLS4gICwtLSwtLS4gLC18ICB8ICwtLS0uICwtLS4tLS4gCiBcICAnICAvIHwgIHwuJy58ICB8fCAgfHwgLi0uIHwnICwtLiAgfCcgLi0uIHx8IC4tLiA6fCAgLi0tJyAKICBcICAgJyAgfCAgIC4nLiAgIHx8ICB8JyAnLScgJ1wgJy0nICB8XCBgLScgfFwgICAtLS58ICB8ICAgIAouLScgIC8gICAnLS0nICAgJy0tJ2AtLScgYC0tLScgIGAtLWAtLScgYC0tLScgIGAtLS0tJ2AtLScgICAgCmAtLS0nICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICAgICAgICAgICAgIOWfuuS6jiBjb2JyYSDlvIDlj5HnmoQgeW91d2FudCDlkb3ku6TooYzlt6XlhbcuIA==")
		fmt.Println(string(str))

		//                     ,--.                  ,--.
		// ,--. ,--.,--.   ,--.|  | ,---.  ,--,--. ,-|  | ,---. ,--.--.
		//  \  '  / |  |.'.|  ||  || .-. |' ,-.  |' .-. || .-. :|  .--'
		//   \   '  |   .'.   ||  |' '-' '\ '-'  |\ `-' |\   --.|  |
		// .-'  /   '--'   '--'`--' `---'  `--`--' `---'  `----'`--'
		// `---'
		//                         基于 cobra 开发的 youwant 命令行工具.

	},
}

func init() {
	rootCmd.AddCommand(logoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logo2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logo2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

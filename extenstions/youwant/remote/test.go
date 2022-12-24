package remote

import (
	"os"

	"github.com/spf13/cobra"
)

// TestHandler 基于严格要求情况下的验证操作
func TestHandler(cmd *cobra.Command, args []string) {
	// 貌似也不需要这个
	// os.Stderr = os.NewFile(uintptr(syscall.Stderr), os.DevNull)

	// url, err := cmd.Flags().GetString("url")
	// if err != nil || url == "" || strings.Index(url, "http") != 0 {
	// 	os.Exit(1)
	// }

	// 简单的尝试请求进行验证
	i, err := getRemoteHeaderContentLength(remoteUrl)
	if err != nil && i == 0 {
		os.Exit(1)
	}

	// buffer, err := getRemoteData(remoteUrl)
	// if err != nil || !jsonValid(buffer) {
	// 	os.Exit(1)
	// }

	os.Exit(0)
}

package youwant

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ShowHandler(cmd *cobra.Command, args []string) {
	// 尝试加载 wants
	want, err := useWant(cmd, args)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Label:  %v\n", want.Label)
	fmt.Printf("Detail: %v\n", want.Detail)
	fmt.Printf("Type:   %v\n", want.Type)
	fmt.Printf("Action: %v\n", want.Template.Action)
	fmt.Printf("Shell.Commands:  %v\n", "")
	for _, command := range want.Template.Shell.Commands {
		fmt.Printf("    command:   %v\n", command)
	}

	fmt.Printf("Files:  %v\n", "")
	for _, file := range want.Template.Files {
		fmt.Printf("    file.Name: %v\n", file.Name)
	}

}

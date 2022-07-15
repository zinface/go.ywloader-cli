package configmanager

import (
	"fmt"
	"os"
	"testing"
)

func TestConfigManager(t *testing.T) {
	tests := []struct {
		Prefix     string
		ConfigFile string
		ResultDir  string
		ResultFile string
	}{
		{
			Prefix:     "ywloader",
			ConfigFile: "youwant.json",
			ResultDir:  fmt.Sprintf("%v/.config/ywloader", os.Getenv("HOME")),
			ResultFile: fmt.Sprintf("%v/.config/ywloader/youwant.json", os.Getenv("HOME")),
		},
	}

	for _, test := range tests {
		config := &ConfigManager{
			Prefix:     test.Prefix,
			ConfigFile: test.ConfigFile,
		}

		s1, err := config.GetUserConfigPath()
		if err != nil {
			t.Fatal(err)
			continue
		}

		s2, err := config.GetUserConfigFilePath()
		if err != nil {
			t.Fatal(err)
			continue
		}

		fmt.Printf("s1:              %v\n", s1)
		fmt.Printf("test.ResultDir:  %v\n", test.ResultDir)
		fmt.Printf("s2:              %v\n", s2)
		fmt.Printf("test.ResultFile: %v\n", test.ResultFile)

		if s1 != test.ResultDir {
			t.Fatalf("should is %v, but is %v", test.ResultDir, s1)
		}
		if s2 != test.ResultFile {
			t.Fatalf("should is %v, but is %v", test.ResultFile, s2)
		}
	}
}

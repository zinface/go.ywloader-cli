package remote

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"gitee.com/zinface/go.ywloader-cli/models"
	"github.com/spf13/cobra"
)

var (
	ErrResponseInvalid = errors.New("响应无效或缺少必填字段")
	ErrJsonInvalid     = errors.New("Json数据无效或格式不正确")
)

// remote 子命令通用定义
var remoteUrl string
var remoteOrigin string

// PreRunHandler 在子命令执行前准备处理 --url 内容
func PreRunHandler(cmd *cobra.Command, args []string) {
	var err error
	url, err := cmd.Flags().GetString("url")
	if err != nil || url == "" || strings.Index(url, "http") != 0 {
		os.Exit(1)
	}
	remoteUrl = url
}

func getRemoteHeader(url string, header string) (string, error) {
	resp, err := http.DefaultClient.Head(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Header.Get(header), nil
}
func getRemoteHeaderContentLength(url string) (int, error) {
	scl, err := getRemoteHeader(url, "Content-Length")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(scl)
}

func getRemoteData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func jsonValid(data []byte) bool {
	return json.Valid(data)
}

func getRemoteWants() (models.Youwants, error) {
	buffer, err := getRemoteData(remoteUrl)
	if err != nil {
		return nil, ErrResponseInvalid
	}

	if ok := jsonValid(buffer); !ok {
		return nil, ErrJsonInvalid
	}

	var wants models.Youwants
	return wants, json.Unmarshal(buffer, &wants)
}

func getRemoteWantsWithLabels(labels []string) (models.Youwants, error) {
	remoteWants, err := getRemoteWants()
	if err != nil {
		return nil, err
	}

	var wants models.Youwants
	for i := 0; i < len(remoteWants); i++ {
		for j := 0; j < len(labels); j++ {
			if labels[j] == remoteWants[i].Label {
				wants = append(wants, remoteWants[i])
			}
		}
	}

	return wants, nil
}

// func getRemoteWants(url string) (models.Youwants, error) {

// 	return nil, nil
// }

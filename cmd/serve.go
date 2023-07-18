/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"gitee.com/zinface/ywloader-cli/utils"
	"github.com/gin-gonic/gin"
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
	Run:   serveHandler,
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

// 此处开始实现 serve 命令

// serveStart 启动一个服务
func serveStart(_rootDir string, _host string, _port int) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFS("/", gin.Dir(_rootDir, true))

	var host = fmt.Sprintf("%s:%d", _host, _port)

	// log.Println("启动服务器:", "http://"+host)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		// addr.String() -> 127.0.0.1/8
		if ipnet, ok := addr.(*net.IPNet); ok && !strings.Contains(ipnet.IP.String(), "::") {
			log.Printf("启动服务器: http://%v:%v\n", ipnet.IP.String(), _port)
		}
	}

	return r.Run(host)
}

// serveHandler 指令实现
func serveHandler(cmd *cobra.Command, args []string) {
	// 如果是 / 将进行询问
	var _rootDir string = "."
	if len(args) > 0 {
		log.Println("访问根节点:", args[0])
		if strings.HasPrefix(args[0], "/") && args[0] == "/" {
			for {
				var _val = utils.GetStdinStringValue("确定使用'/'作为根访问节点吗(yes/no): ", "")
				// 回答 yes，设置根访问节点
				if strings.Compare(_val, "yes") == 0 {
					_rootDir = args[0]
					break
				}
				// 回答 no，退出
				if strings.Compare(_val, "no") == 0 {
					os.Exit(0)
				}
			}
		} else {
			_rootDir = args[0]
		}
	}

	var port, _ = cmd.Flags().GetInt("port")
	err := serveStart(_rootDir, "0.0.0.0", port)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

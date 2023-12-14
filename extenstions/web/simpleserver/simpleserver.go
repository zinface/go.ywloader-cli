package simpleserver

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"gitee.com/zinface/ywloader-cli/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// 输出可访问地址
func printAccessAddress(port int) {
	addrs, err := GetLocalAddress(IPV4)
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		log.Printf("启动服务器: http://%v:%v\n", addr, port)
	}
}

// SimpleServer 启动一个简单服务
func SimpleServer(filepath string, host string, port int) error {
	fi, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if fi.IsDir() {
		r.StaticFS("/", gin.Dir(filepath, true))
	} else {
		// r.StaticFile("/", filepath)
		r.GET("/", func(ctx *gin.Context) {
			ctx.Header("Content-Disposition", "attachment; filename="+path.Base(filepath))
			ctx.File(filepath)
			// data, _ := os.ReadFile(filepath)
			// ctx.Data(200, "application/octet-stream", data)
		})
	}

	printAccessAddress(port)

	return r.Run(fmt.Sprintf("%s:%d", host, port))
}

// SimpleServeHandler 指令实现
func SimpleServeHandler(cmd *cobra.Command, args []string) {
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
	err := SimpleServer(_rootDir, "0.0.0.0", port)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

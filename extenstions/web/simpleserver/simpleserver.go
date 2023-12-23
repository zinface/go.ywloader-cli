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

var cmd_filename string
var cmd_cache_size uint64

// SimpleServer 启动一个简单服务
func SimpleServer(filepath string, host string, port int) error {
	var filename string
	fi, err := os.Stat(filepath)

	if filepath == "-" {
		fi, err = os.Stdin.Stat()
		filename = cmd_filename
	}

	if err != nil {
		return err
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	if filename != "-" && fi.IsDir() {
		r.StaticFS("/", gin.Dir(filepath, true))
	} else {
		filename = path.Base(filepath)
		// r.StaticFile("/", filepath)
		r.GET("/", func(ctx *gin.Context) {
			if filepath == "-" {
				handleFromStdin(ctx, cmd_cache_size, cmd_filename)
				if outOfCache {
					os.Exit(0)
				}
				ctx.Done()
			} else {
				ctx.FileAttachment(filepath, filename)
			}
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

	if _rootDir == "-" {
		cmd_filename, _ = cmd.Flags().GetString("attachment-filename")
		cmd_cache_size, _ = cmd.Flags().GetUint64("cache-size")
	}

	var port, _ = cmd.Flags().GetInt("port")
	err := SimpleServer(_rootDir, "0.0.0.0", port)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

package simpleserver

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"gitee.com/zinface/ywloader-cli/utils"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

func Handler(serve *ServeCommand) {
	if serve.isLimitMode() {
		log.Printf("启动模式为: 限速模式(LimitMode)")
		serve.printLimitInfo()
	} else {
		log.Printf("启动模式为: 简单模式(SimpleMode)")
	}

	if serve.isDir() {
		serve.RunDirServer()
	} else if serve.isIO() {
		serve.RunIOServer()
	} else {
		serve.RunFileServer()
	}
}

func (serve *ServeCommand) RunDirServer() {
	log.Println("访问根节点:", serve.Argument)
	for serve.Argument == "/" {
		var _val = utils.GetStdinStringValue("确定使用'/'作为根访问节点吗(yes/no): ", "")
		// 回答 yes，设置根访问节点
		if strings.Compare(_val, "yes") == 0 {
			break
		}
		// 回答 no，退出
		if strings.Compare(_val, "no") == 0 {
			os.Exit(0)
		}
	}

	// 创建文件服务器处理器
	fs := http.FileServer(gin.Dir(serve.Argument, true))

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		// 创建限速写入器
		var limitWriter http.ResponseWriter = ctx.Writer
		if serve.isLimitMode() {
			limitWriter = &SpeedLimitWriter{
				ResponseWriter: ctx.Writer,
				bytesPerSecond: serve.LimitSpeed,
				startTime:      time.Now(),
			}
		}

		// 使用文件服务器处理请求，并通过限速写入器输出
		fs.ServeHTTP(limitWriter, ctx.Request)

		// 如果文件服务器已经处理了请求，中止 Gin 的后续处理
		ctx.Abort()
	})

	// 经在中间件中处理了静态文件
	r.StaticFS("/", gin.Dir(serve.Argument, true))

	port := PrintLocalAddressAndGetPort(serve.Port)
	r.Run(fmt.Sprintf(":%v", port))
}

func (serve *ServeCommand) RunFileServer() {
	filename := path.Base(serve.Argument)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		// // 检查 User-Agent，如果是 wget，重定向到包含文件名的URL
		// userAgent := ctx.GetHeader("User-Agent")
		// if strings.Contains(userAgent, "Wget") {
		// 	ctx.Redirect(http.StatusFound, "/"+filename)
		// 	return
		// }
		ctx.Redirect(http.StatusMovedPermanently, "/"+filename)
	})

	r.GET("/:filename", func(ctx *gin.Context) {
		requestedFilename := ctx.Param("filename")
		if requestedFilename != filename {
			return
		}

		// 打开文件
		file, err := os.Open(serve.Argument)
		if err != nil {
			return
		}
		defer file.Close()

		// 获取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			return
		}

		// 设置下载头
		ctx.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
		ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
		ctx.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

		// 创建限速写入器
		var limitWriter http.ResponseWriter = ctx.Writer
		if serve.isLimitMode() {
			limitWriter = &SpeedLimitWriter{
				ResponseWriter: ctx.Writer,
				bytesPerSecond: serve.LimitSpeed,
				startTime:      time.Now(),
			}
		}

		http.ServeContent(limitWriter, ctx.Request, filename, fileInfo.ModTime(), file)
	})

	port := PrintLocalAddressAndGetPort(serve.Port)
	r.Run(fmt.Sprintf(":%v", port))
}

func (serve *ServeCommand) RunIOServer() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	serve.makeCacheBuffer()

	var once bool = true
	var reuseableSize int = 0

	r.GET("/", func(ctx *gin.Context) {
		// handleFromStdin(ctx, serve.CacheSize, serve.AttachementFilename)
		// 1. 先读取到缓冲区(先创建缓冲区)
		log.Println("准备缓冲中.")

		// 2. 设置响应头
		ctx.Writer.Header().Set("Content-Type", "application/octet-stream")
		ctx.Writer.Header().Set("Content-Disposition", `"attachment; filename="`+serve.AttachementFilename+`"`)
		// ctx.Header("Content-Transfer-Encoding", "binary")
		// ctx.Writer.Header().Set("Transfer-Encoding", "chunked")

		// 创建限速写入器
		var writer http.ResponseWriter = ctx.Writer
		if serve.isLimitMode() {
			writer = &SpeedLimitWriter{
				ResponseWriter: ctx.Writer,
				bytesPerSecond: serve.LimitSpeed,
				startTime:      time.Now(),
			}
		}

		if once {
			// 读取数据到缓冲区
			var bufferSize int = 0
			for {
				n, err := os.Stdin.Read(serve.cache_buffer[bufferSize:])
				bufferSize += n
				if err != nil || bufferSize == int(serve.CacheSize) {
					break
				}
			}

			// 检查缓冲区读取情况，未满则尽，响应长度
			if bufferSize < int(serve.CacheSize) {
				once = false
				reuseableSize = bufferSize
				ctx.Header("Content-Length", strconv.FormatInt(int64(bufferSize), 10))
				log.Printf("缓冲区大小: %v, 已读取: %v", humanize.Bytes(uint64(serve.CacheSize)), humanize.Bytes(uint64(bufferSize)))
			} else {
				log.Printf("缓冲区大小: %v, 已读取: %v(用尽，无法计算剩余长度)", humanize.Bytes(uint64(serve.CacheSize)), humanize.Bytes(uint64(bufferSize)))
			}

			// 将缓冲区的数据写入响应
			// var totalWritten = 0
			totalWritten, err := writer.Write(serve.cache_buffer[:bufferSize])
			if err != nil {
				return
			}

			//
			pr, pw := io.Pipe()
			go func() {
				defer pw.Close()
				defer pr.Close()
				stdinRead, err := io.Copy(pw, os.Stdin)
				log.Printf("从标准输入中读取: %v(%v)", humanize.Bytes(uint64(stdinRead)), err)
			}()
			for {
				written, err := io.Copy(writer, pr)
				totalWritten += int(written)
				if err != nil {
					break
				}
				log.Printf("转发输入到响应: %v(%v)", humanize.Bytes(uint64(written)), err)
			}
			ctx.Writer.Flush()
			log.Printf("已传输完成：%v (%v).\n", serve.AttachementFilename, humanize.Bytes(uint64(totalWritten)))

			if once {
				ctx.Done()
				os.Exit(0)
			}
		} else {
			ctx.Header("Content-Length", strconv.FormatUint(uint64(reuseableSize), 10))
			writer.Write(serve.cache_buffer[:reuseableSize])
			log.Printf("已传输完成：%v (%v).\n", serve.AttachementFilename, humanize.Bytes(uint64(reuseableSize)))
			ctx.Done()
		}

	})

	port := PrintLocalAddressAndGetPort(serve.Port)
	r.Run(fmt.Sprintf(":%v", port))
}

package simpleserver

import (
	"log"
	"os"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

var (
	outOfCache        = true
	nextUseCache      = false
	cache_buffer      []byte
	cache_size        uint64
	caceh_filename    string
	cache_data_length uint64
)

// type HandleCache struct {
// 	cache_size uint64
// }

func readStdinToCache(buffer []byte, cache_size uint64) uint64 {
	// fmt.Printf("len(buffer): %v\n", len())
	var readed uint64 = 0
	for {
		n, err := os.Stdin.Read(buffer[readed:])
		// log.Printf("缓冲区已读: %v, err: %v\n", n, err)
		readed += uint64(n)
		// 当读取结束时或首次读取已读满缓冲区时不再继续
		if err != nil || readed == cache_size {
			break
		}
	}
	log.Printf("缓冲区已读取完成，溢出状态: %v\n", readed == cache_size)
	return readed
}

func handleFromCache(ctx *gin.Context) uint64 {
	var full_written uint64 = 0
	for {
		n, err := ctx.Writer.Write(cache_buffer[full_written:cache_data_length])
		if err != nil {
			log.Println(err)
			break
		}
		full_written += uint64(n)
		ctx.Writer.Flush()

		if full_written == cache_data_length {
			break
		}
	}
	ctx.Writer.Flush()
	return full_written
}

func handleFromStdin(ctx *gin.Context, cache_size_mb uint64, file_name string) {

	if outOfCache {
		var full_written uint64 = 0

		cache_size = cache_size_mb * 1024 * 1024
		cache_buffer = make([]byte, cache_size)
		caceh_filename = file_name

		// 1. 标准输入处理，创建缓冲区
		log.Println("准备中.")

		ctx.Writer.WriteHeader(200)
		ctx.Header("Content-Disposition", `attachment; filename="`+file_name+`"`)
		ctx.Header("Content-Transfer-Encoding", "binary")

		// 2. 首次读写使用缓冲区，并读满缓冲区或读尽标准输入
		first_read := readStdinToCache(cache_buffer, cache_size)

		// 2.1 检查首次读取的数据量是否未满缓冲区，未满时认为标准输入已读完，实际文件大小将是首次读取时的大小，可以设立 'Content-Length'
		// 如缓冲区已满，则标准输入未读尽，无法读取标准输入大小，并直接发送数据
		if first_read < cache_size {
			outOfCache = false
			nextUseCache = true
			cache_data_length = first_read

			log.Printf("缓冲区大小: %v, 已读取: %v", humanize.Bytes(uint64(cache_size)), humanize.Bytes(uint64(first_read)))
			ctx.Header("Content-Length", strconv.FormatInt(int64(cache_data_length), 10))
		} else {
			log.Printf("缓冲区大小: %v, 已读取: %v(溢出，无法计算'Content-Length')", humanize.Bytes(uint64(cache_size)), humanize.Bytes(uint64(first_read)))
		}

		// 3. 将首次所读取的缓冲区内容全部写入到响应中
		ctx.Writer.Write(cache_buffer[:first_read])
		full_written += first_read

		// 4. 尝试从标准输入中读取数据，如果已读完则将数据完全发送完毕
		for {
			n, err := os.Stdin.Read(cache_buffer)
			if err != nil {
				break
			}
			full_written += uint64(n)
			_, err = ctx.Writer.Write(cache_buffer[:n])
			if err != nil {
				log.Fatalln(err)
			}
		}
		ctx.Writer.Flush()

		// 打印，退出
		log.Printf("已传输完成：%v (%v).\n", cmd_filename, humanize.Bytes(full_written))

		if outOfCache {
			os.Exit(0)
		}
	} else {
		var full_written uint64 = 0

		ctx.Writer.WriteHeader(200)
		ctx.Header("Content-Disposition", `attachment; filename="`+caceh_filename+`"`)
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Length", strconv.FormatUint(cache_data_length, 10))

		full_written = handleFromCache(ctx)

		log.Printf("已传输完成：%v (%v).\n", cmd_filename, humanize.Bytes(full_written))
	}
}

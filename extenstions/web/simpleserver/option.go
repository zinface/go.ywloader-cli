package simpleserver

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	KeyPort               = "port"
	KeyAttachmentFilename = "attachment-filename"
	KeyCacheSize          = "cache-size"
	KeyLimitSpeed         = "limit-speed"
)

type Option func(serve *ServeCommand)

type ServeCommand struct {
	*cobra.Command
	Argument string

	AttachementFilename string
	CacheSize           int
	Port                int
	LimitSpeed          int

	// flags
	port                *int
	attachment_filename *string
	cache_size          *string
	limit_speed         *string

	// cache
	cache_buffer []byte
}

func (s *ServeCommand) FlagsProvider() {
	s.Command.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalf("It should have one parameter ('.', '-', or <filepath>)")
		}
		s.Argument = args[0]
		if !s.isIO() && !s.isExists() {
			log.Fatalf("It should have one parameter ('.', '-', or <filepath>)")
		}
		s.Port = *s.port
		s.AttachementFilename = *s.attachment_filename
		s.CacheSize = SpeedStringToInt(*s.cache_size)
		s.LimitSpeed = SpeedStringToInt(*s.limit_speed)
		s.Run()
	}
	s.port = s.Command.Flags().IntP(KeyPort, "p", 8080, "端口号")
	s.attachment_filename = s.Command.Flags().StringP(KeyAttachmentFilename, "", "UnknowFile", "标准输入时的文件命名，用于响应(Content-Disposition)支持")
	s.cache_size = s.Command.Flags().StringP(KeyCacheSize, "", "50m", "标准输入时的缓冲区大小，用于响应(Content-Length)支持")
	s.limit_speed = s.Command.Flags().StringP(KeyLimitSpeed, "", "0m", "限制传输速率")
}

func (s *ServeCommand) isExists() bool {
	if s.isIO() {
		_, err := os.Stat(s.Argument)
		return os.IsExist(err)
	}
	return true
}

func (s *ServeCommand) isDir() bool {
	if !s.isIO() {
		fi, _ := os.Stat(s.Argument)
		return fi.IsDir()
	}
	return false
}

func (s *ServeCommand) isIO() bool {
	return s.Argument == "-"
}

func (s *ServeCommand) isLimitMode() bool {
	return s.LimitSpeed != 0
}

func (s *ServeCommand) printLimitInfo() {
	log.Printf("当前限速值: %s -> %d bytes/s", *s.limit_speed, s.LimitSpeed)
}

func (serve *ServeCommand) Run() {
	Handler(serve)
}

func (serve *ServeCommand) makeCacheBuffer() {
	if serve.cache_buffer == nil {
		serve.cache_buffer = make([]byte, serve.CacheSize)
	}
}

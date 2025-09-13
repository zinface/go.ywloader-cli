package simpleserver

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"testing"
	"time"
)

type LimitWriter struct {
	io.Writer
	bytesPerSecond int
	bytesWritten   int
	startTime      time.Time
}

func (tw *LimitWriter) Write(p []byte) (int, error) {
	var totalWritten = 0
	var remaining = len(p)

	for remaining > 0 {
		// 已持续时间
		elapsed := time.Since(tw.startTime)

		// 持续时间内被允许的动态字节数 (未分片 - 总量已被使用)
		allowedBytes := int(float64(tw.bytesPerSecond) * elapsed.Seconds())

		// 切片，在 100ms 内可发送的最大块
		// chunkSize := tw.bytesPerSecond / 10
		chunkSize := allowedBytes
		if chunkSize > remaining {
			chunkSize = remaining
		}
		if chunkSize > 4096 { // 限制块大小
			chunkSize = 4096
		}
		if chunkSize == 0 {
			chunkSize = 1
		}

		// 一个块大小
		n, err := tw.Writer.Write(p[totalWritten : totalWritten+chunkSize])
		if err != nil {
			log.Fatalln(err)
			return totalWritten, err
		}

		totalWritten += n
		remaining -= n
		tw.bytesWritten += n

		expectedTime := time.Duration(float64(time.Second) * float64(tw.bytesWritten) / float64(tw.bytesPerSecond))
		if expectedTime > elapsed && totalWritten < len(p) {
			time.Sleep(expectedTime - elapsed)
		}
	}

	return totalWritten, nil
}

func TestThrottledWriter_Write(t *testing.T) {
	bufferSize := 1024
	buffer := make([]byte, bufferSize)
	fmt.Printf("分配了 %d 字节 (%d MB) 的缓冲区\n", len(buffer), len(buffer)/(1024*1024))

	var data bytes.Buffer
	tw := LimitWriter{
		Writer:         &data,
		bytesPerSecond: int(float64(bufferSize) * 0.1),
		startTime:      time.Now(),
	}

	tw.Write(buffer)
}

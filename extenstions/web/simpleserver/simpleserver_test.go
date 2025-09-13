package simpleserver

import (
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	// 直接访问 http://0.0.0.0:7789 就可以访问根访问点
	// r.StaticFS("/", gin.Dir(".", true))

	// 必须访问 http://0.0.0.0:7789/prefix 才可以访问根访问点
	r.StaticFS("/prefix", gin.Dir(".", true))

	r.Run("0.0.0.0:7789")
}

func TestHttp(t *testing.T) {
	fs := http.FileServer(http.Dir("."))

	http.Handle("/", fs)

	err := http.ListenAndServe(":7789", nil)
	if err != nil {
		panic(err)
	}
}

func TestHttpLimitMiddleware(t *testing.T) {
	fs := http.FileServer(http.Dir("."))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("w:", w, "r", r)
		var bytesPerSecond = 1024

		limitWriter := &SpeedLimitWriter{
			ResponseWriter: w,
			bytesPerSecond: bytesPerSecond,
			startTime:      time.Now(),
		}

		fs.ServeHTTP(limitWriter, r)
	})

	err := http.ListenAndServe(":7789", nil)
	if err != nil {
		panic(err)
	}
}

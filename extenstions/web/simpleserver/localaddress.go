package simpleserver

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type AddressType int

const (
	IPV4 AddressType = iota
	IPV6
)

func GetLocalAddress(t AddressType) ([]string, error) {
	var address []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		// addr.String() -> 127.0.0.1/8
		if ipnet, ok := addr.(*net.IPNet); ok {
			if t == IPV4 && !strings.Contains(ipnet.IP.String(), "::") {
				address = append(address, ipnet.IP.String())
			}

			if t == IPV6 && strings.Contains(ipnet.IP.String(), "::") {
				address = append(address, ipnet.IP.String())
			}
		}
	}
	return address, nil
}

// 输出可访问地址
func PrintLocalAddressAndGetPort(port int) int {
	addrs, err := GetLocalAddress(IPV4)
	if err != nil {
		panic(err)
	}

	// 当默认的端口被占用时，将使用随机端口
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		ln, _ = net.Listen("tcp", fmt.Sprintf(":%v", 0))
	}
	defer ln.Close() // 已获取随机端口，先关闭，再交给后续使用

	_, listenPort, _ := net.SplitHostPort(ln.Addr().String())
	port, _ = strconv.Atoi(listenPort)

	for _, addr := range addrs {
		log.Printf("启动服务器: http://%v:%v\n", addr, port)
	}

	return port
}

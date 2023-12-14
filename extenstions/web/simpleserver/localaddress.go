package simpleserver

import (
	"net"
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

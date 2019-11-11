package ip

import (
	"fmt"
	"net"
)

var (
	err      error
	vLocalIP string
)

func init() {
	if vLocalIP, err = localIP(); err != nil {
		panic(err)
	}
}

func LocalIP() string {
	return vLocalIP
}

// ------- internal ---------

func localIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if nil != err {
		return "", err
	}
	for _, address := range addrs {
		// check loop back ip
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("ip is empty")
}

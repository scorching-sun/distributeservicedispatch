// serviceDispatchListener project serviceDispatchListener.go
package serviceDispatchListener

import (
	"errors"
	"net"
	"os"
	"strconv"
)

var (
	listenIP   net.IP
	listenPort int
)

var (
	errListenIPConf   error = errors.New("service dispatch listen ip not configure or error")
	errListenPortConf error = errors.New("service dispatch listen port not configure or error")
)

func init() {
	initListenConf()
}

func initListenConf() {
	listenIP = net.ParseIP(os.Getenv("ServiceDispatchListenIP"))

	if listenIP == nil {
		panic(errListenIPConf)
	}

	port := os.Getenv("ServiceDispatchListenPort")

	var err error

	listenPort, err = strconv.Atoi(port)

	if err != nil {
		panic(errListenPortConf)
	}
}

type ServiceDispatchListener struct{}

func NewServiceDispatchListener() *ServiceDispatchListener {
	this := new(ServiceDispatchListener)

	return this
}

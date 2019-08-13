// endPointServiceRegisteApply project endPointServiceRegisteApply.go
package endPointServiceRegisteApply

import (
	"errors"
	"net"
	"os"
	"serviceDispatchDataModel/modelSDS"
	"strings"
)

const (
	envKeyRegisteListenModel string = "RegisteListenModel"
)

var (
	registeListenModel string
)

var (
	errNetConnectFailure error = errors.New("net connect failure")

	errIP   error = errors.New("not valid ip with servie registe")
	errPort error = errors.New("not valid port with servie registe")
)

func init() {
	initRegisteListenModel()
}

//初始化服务注册监听方式
//默认为tcp监听
func initRegisteListenModel() {
	registeListenModel = os.Getenv(envKeyRegisteListenModel)

	if strings.TrimSpace(registeListenModel) == "" {
		registeListenModel = modelSDS.RegisteListenModelTCP
	}
}

type EndPointServiceRegisteApply struct{}

func NewEndPointServiceRegisteApply() *EndPointServiceRegisteApply {
	this := new(EndPointServiceRegisteApply)

	return this
}

func (this *EndPointServiceRegisteApply) Sender(registeServiceIP net.IP, registeServicePort int) net.Conn {
	switch registeListenModel {
	case modelSDS.RegisteListenModelUDPMulticast:
		return this.udpMulticastSender(registeServiceIP, registeServicePort)
	default:
		return this.tcpSender(registeServiceIP, registeServicePort)
	}
}

//终结点服务注册与健康检查
func (this *EndPointServiceRegisteApply) RegisteApplyAndHealthReport(sender net.Conn, text string) error {
	switch registeListenModel {
	case modelSDS.RegisteListenModelUDPMulticast:
		return this.udpMulticastRegisteApplyAndHealthReport(sender, text)
	default:
		return this.tcpRegisteApplyAndHealthReport(sender, text)
	}
}

// endPointServiceRegister project endPointOverseeInfoBroadcast.go
package endPointServiceRegisteApply

import (
	"library/Net/Udp"
	"net"
)

func (this *EndPointServiceRegisteApply) udpMulticastSender(registeServiceIP net.IP, registeServicePort int) net.Conn {
	util := Udp.NewUdpUtil(registeServiceIP, registeServicePort)
	return util.Connect()
}

//终结点服务注册与健康检查
func (this *EndPointServiceRegisteApply) udpMulticastRegisteApplyAndHealthReport(sender net.Conn, text string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errNetConnectFailure
		}
	}()

	//	log.Println(sender.RemoteAddr(), "text=\r\n", text)
	_, err = sender.Write([]byte(text))

	return err
}

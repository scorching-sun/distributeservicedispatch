package endPointServiceRegisteApply

import (
	"library/Net/Tcp"
	"net"
)

func (this *EndPointServiceRegisteApply) tcpSender(registeServiceIP net.IP, registeServicePort int) net.Conn {
	tcpUtil := Tcp.NewTcpUtil()
	conn := tcpUtil.Connect(registeServiceIP, registeServicePort)

	return conn
}

//终结点服务注册与健康检查
func (this *EndPointServiceRegisteApply) tcpRegisteApplyAndHealthReport(sender net.Conn, text string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errNetConnectFailure
		}
	}()

	//	log.Println(sender.RemoteAddr(), "text=\r\n", text)

	_, err = sender.Write([]byte(text))

	return err
}

package endPointServiceListener

import (
	"bufio"
	"library/Net/Udp"
	"net"
	"serviceDispatch"
	"time"
)

//udp监听end point service
func (this *EndPointServiceListener) UDPListenEndPointServices() {
	udpUtil := Udp.NewUdpUtil(SDSListenIP, SDSListenPort)

	conn := udpUtil.ListenMulticastConnect()

	defer func(c *net.UDPConn) {
		if c != nil {
			c.Close()
		}
	}(conn)

	if conn != nil {
		udpUtil.ListenMulticast(conn, this.udpRecieveEndPointServiceStat)
	}
}

//udp监听接收end point service状态
func (this *EndPointServiceListener) udpRecieveEndPointServiceStat(conn *net.UDPConn) {
	//	log.Println("remote address =", conn.RemoteAddr())

	var (
		buffer []byte
		err    error
	)

	//	buffer = make([]byte, Udp.MaxBufferSize)
	reader := bufio.NewReader(conn)

	for {
		time.Sleep(10 * time.Millisecond)

		buffer, err = reader.ReadBytes('\n')

		//		n, clientAddr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			continue
		}

		if buffer == nil && len(buffer) <= 0 {
			continue
		}

		event := new(serviceDispatch.Event)
		event.Data = string(buffer)

		//	log.Println("recieve data=\r\n", event.Data)

		this.Notify(event)
	}
}

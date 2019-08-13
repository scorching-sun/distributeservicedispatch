// serviceDispatchProxyLite project serviceDispatchProxyLite.go
/*
描述:
服务调度服务调用代理——轻代理。仅用于获取调度服务返回的可用终结点服务ip及port。
*/
package serviceDispatchProxyLite

import (
	"errors"
	"library/Json"
	"library/Net/Tcp"
	"log"
	"net"
	"serviceDispatchDataModel/modelProxyRequest"
	"serviceDispatchDataModel/modelSDSResponse"
	"strconv"
	"strings"
	"time"
)

const (
	bufferLimit int = 1024
)

var (
	errDispatchServiceIP error = errors.New("service dispatch service ip not configure or error")
)

type ProxyLite struct {
	request *modelProxyRequest.SDSRequestInfo
}

func NewProxyLite(request *modelProxyRequest.SDSRequestInfo) *ProxyLite {
	this := new(ProxyLite)
	this.request = request

	return this
}

//proxy tcp connect
func (this *ProxyLite) Connect() *net.TCPConn {
	tcpUtil := Tcp.NewTcpUtil()

	ip := net.ParseIP(this.request.DispatchServiceIP)

	if ip == nil {
		panic(errDispatchServiceIP)
	}

	port, err := strconv.Atoi(this.request.DispatchServicePort)

	if err != nil {
		panic(err)
	}

	return tcpUtil.Connect(ip, port)
}

//available service net address
func (this *ProxyLite) AvailableServiceAddress(conn *net.TCPConn) (data *modelSDSResponse.SDSResponseInfo) {
	if conn == nil {
		return nil
	}

	var (
		err error
		r   string
	)

	r, err = Json.StructToJsonString(this.request)

	if err != nil {
		return nil
	}

	_, err = conn.Write([]byte(strings.Join([]string{r, "\n"}, "")))

	if err != nil {
		log.Println("proxy send request failure \r\n", err)
		return nil
	}

	//	log.Println("proxy send request =\r\n", string(r))

	var (
		recv_len int
		buff     []byte = make([]byte, bufferLimit)
		response []byte
	)

	for {
		time.Sleep(10 * time.Millisecond)

		//		conn.SetReadDeadline(time.Now().Add(this.requestTimeout))
		recv_len, err = conn.Read(buff)

		if err != nil {
			log.Println("proxy recieve response failure ", err)
			response = nil
			break
		}

		response = append(response, buff[:recv_len]...)

		if recv_len < bufferLimit {
			break
		}
	}

	eventData := string(response)

	//	log.Println("eventData==>", eventData)

	if eventData == "EOF" || (response == nil && len(response) <= 0) {
		//	if response == nil && len(response) <= 0 {
		return nil
	}

	err = Json.JsonbytesToStruct(response, &data)

	if err != nil {
		return nil
	}

	return
}

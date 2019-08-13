// example project main.go
package main

import (
	"log"
	"os"
	"serviceDispatchDataModel/modelProxyRequest"
	"serviceDispatchDataModel/modelSDSResponse"
	"serviceDispatchProxy/serviceDispatchProxyLite"
	"time"
)

const (
	envKeyDispatchServiceIP   string = "DispatchServiceIP"
	envKeyDispatchServicePort string = "DispatchServicePort"
)

var (
	dispatchServiceIP   string
	dispatchServicePort string
)

func init() {
	dispatchServiceIP = os.Getenv(envKeyDispatchServiceIP)
	dispatchServicePort = os.Getenv(envKeyDispatchServicePort)
}

func main() {
	ProxyTest()
}

func ProxyTest() {
	r := new(modelProxyRequest.SDSRequestInfo)
	r.DispatchServiceIP = dispatchServiceIP
	r.DispatchServicePort = dispatchServicePort

	r.DispatchServiceIP = "52.83.185.214"
	//	r.DispatchServicePort = "50000"
	r.NameSpace = "mtdms"
	r.ServiceProcessName = "mtdmsWebApi"
	r.ServiceProvideMode = "restful"

	log.Println("dispatchServiceIP=", dispatchServiceIP, "|", "dispatchServicePort=", dispatchServicePort)

	proxy := serviceDispatchProxyLite.NewProxyLite(r)
	conn := proxy.Connect()

	//	conn.SetReadDeadline(time.Now().Add(10000 * time.Millisecond))
	//	conn.SetKeepAlive(true)
	//	conn.SetWriteDeadline(time.Now().Add(30 * time.Second))
	defer func() {
		if conn != nil {
			conn.Close()
			conn = nil
		}
	}()

	totalCyc := 50000
	timeout := time.Second * time.Duration(60)
	ch := make(chan *modelSDSResponse.SDSResponseInfo, totalCyc)

	for i := 0; i < totalCyc; i++ {
		time.Sleep(10 * time.Millisecond)

		addr := proxy.AvailableServiceAddress(conn)
		if addr != nil {
			log.Println("recieve remote response , service available address=", addr)
		} else {
			log.Println("recieve remote response failure")
		}

		ch <- addr

		//		go func() {
		//			addr := proxy.AvailableServiceAddress(conn)
		//			//					if addr != nil {
		//			//						log.Println("recieve remote response , service available address=", addr)
		//			//					} else {
		//			//						log.Println("recieve remote response failure")
		//			//					}
		//			ch <- addr
		//		}()
	}

	var count int = 0

DONE:
	for {
		select {
		case response := <-ch:
			count++

			log.Println("complate request,and recieve response->", response, "|count=", count)

			if count == totalCyc {
				log.Println("exit cycle")

				goto Finish //退出for select循环
			}
		case <-time.After(timeout):
			log.Println("timeout exit")

			break DONE //退出for select循环
		}
	}
Finish:

	log.Println("finish")
}

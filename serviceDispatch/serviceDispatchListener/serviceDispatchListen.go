package serviceDispatchListener

import (
	//	"bufio"
	"library/Json"
	"library/Net/Tcp"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"serviceDispatch/serviceDispatchStrategy"
	"serviceDispatchDataModel/modelServiceInfo"
	"strings"
	"time"
)

const (
	bufferLimit int = 1024
)

func (this *ServiceDispatchListener) ServiceDispatchListen() {
	tcpUtil := Tcp.NewTcpUtil()

	tcpUtil.Listen(listenIP, listenPort, this.serviceDispatchListenCallback)
}

func (this *ServiceDispatchListener) serviceDispatchListenCallback(listener *net.TCPListener) {
	stop_chan := make(chan os.Signal) // 接收系统中断信号
	signal.Notify(stop_chan, os.Interrupt)

	var err error

	go func() {
		<-stop_chan

		close(stop_chan)
		if err = listener.Close(); err != nil {
			log.Println("service dispacth service listen close failure.", err)
		}
	}()

	log.Println("service dispacth service listen start,address=", listener.Addr().String())

	for {
		time.Sleep(10 * time.Millisecond)

		//获取当前连接监听服务的实例
		conn, err := listener.AcceptTCP()

		if err != nil {
			continue
		}

		log.Println("proxy address :", conn.RemoteAddr())
		//		conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))

		//follow process
		go this.handleConnection(conn, 2)
	}
}

func (this *ServiceDispatchListener) availableServiceResponse(conn net.Conn, requestMessage string) {
	var (
		err              error
		responseBuff     []byte
		disPatchStrategy *serviceDispatchStrategy.ServiceDispatchStrategy
		availableService *modelServiceInfo.ServiceInfo
		request          *modelServiceInfo.ServiceInfo
	)

	request = new(modelServiceInfo.ServiceInfo)

	err = Json.JsonStringToStruct(string(requestMessage), &request)

	if err != nil {
		log.Println("service dispatch service recieve request message convert failure ", err)
		return
	}

	//	log.Println("service dispatch service recieve request=\r\n", request)

	disPatchStrategy = serviceDispatchStrategy.NewServiceDispatchStrategy()
	availableService = disPatchStrategy.DispatchAvailableService(request)
	responseBuff, err = Json.StructTobytes(availableService)

	if err == nil {
		_, err = conn.Write(responseBuff)

		if err == nil {
			log.Println("send data =>", "responseBuff=", string(responseBuff), "<to>", conn.RemoteAddr())
		}
	} else {
		conn.Write([]byte("EOF"))
	}

	availableService = nil
}

func (this *ServiceDispatchListener) handleConnection(conn net.Conn, timeout int) {
	readChan := make(chan string)
	stopChan := make(chan bool)

	defer func() {
		if conn != nil {
			log.Println(conn.RemoteAddr(), "->", conn.LocalAddr(), " goroutine exit")
			conn.Close()
		}

		close(readChan)
		close(stopChan)

		runtime.GC()
	}()

	var (
		requestMessage string
		stop           bool
	)

	go this.readRegisteApply(conn, readChan, stopChan)

endForCycle:
	for {
		time.Sleep(10 * time.Millisecond)

		select {
		case requestMessage = <-readChan:
			this.availableServiceResponse(conn, requestMessage)

			//			runtime.GC()
			//		case <-time.After(time.Second * time.Duration(timeout)):
			//			log.Println("It's really weird to get Nothing!!!")

			//			break endForCycle
		case stop = <-stopChan:
			if stop {
				break endForCycle
			}
		}
	}
}

func (this *ServiceDispatchListener) readRegisteApply(conn net.Conn, readChan chan<- string, stopChan chan<- bool) {
	var (
		request string
		err     error
		//		reader  *bufio.Reader
		bytes  []byte
		length int
	)

	//	reader = bufio.NewReader(conn)
	bytes = make([]byte, bufferLimit)

	for {
		time.Sleep(10 * time.Millisecond)

		//		request, err = reader.ReadString('\n')

		length, err = conn.Read(bytes)
		log.Println("bytes=", string(bytes[:length]))
		request = string(bytes[:length])

		if err != nil {
			log.Println("service dispatch service get request of proxy failure", err)
			break
		}

		if strings.TrimSpace(request) != "" {
			readChan <- request
		}
	}

	stopChan <- true
}

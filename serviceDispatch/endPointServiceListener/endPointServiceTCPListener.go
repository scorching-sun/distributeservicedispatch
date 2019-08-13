/*
注意：
1.使用goroutine注意其创建的数量,特别是在死循环中创建goroutine时。以及解决可能存在的内存泄露；
*/
package endPointServiceListener

import (
	"bufio"
	"library/Net/Tcp"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"serviceDispatch"
	"sync"
	"time"
)
import (
	"flag"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

const (
	bufferLimit int = 2048
)

var (
	muexLock sync.Mutex
)

//tcp监听end point service
func (this *EndPointServiceListener) TCPListenEndPointServices() {
	tcpUtil := Tcp.NewTcpUtil()

	tcpUtil.Listen(SDSListenIP, SDSListenPort, this.serviceStatListenCallback)
}

func (this *EndPointServiceListener) endPointServiceRegiste(requestMessage string) {
	event := new(serviceDispatch.Event)
	event.Data = string(requestMessage)

	//			log.Println("recieve data=\r\n", event.Data)

	this.Notify(event)

	event = nil
}

func (this *EndPointServiceListener) serviceStatListenCallback(listener *net.TCPListener) {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil { //监控cpu
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

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

	this.handleListen(listener)
}

func (this *EndPointServiceListener) handleListen(listener *net.TCPListener) {
	for {
		time.Sleep(10 * time.Millisecond)

		//获取当前连接监听服务的实例
		conn, err := listener.AcceptTCP()

		if err != nil {
			continue
		}

		//		log.Println("proxy address :", conn.RemoteAddr())
		//				conn.SetReadDeadline(time.Now().Add(time.Duration(10) * time.Second))

		//follow process
		go this.handleConnection(conn)

		goNum := runtime.NumGoroutine()
		log.Println("end point service tcp listen, goroutine number=>", goNum)
	}
}

func (this *EndPointServiceListener) handleConnection(conn net.Conn) {
	readChan := make(chan string)
	stopChan := make(chan bool)

	defer func() {

		if *memprofile != "" {
			f, err := os.Create(*memprofile)
			if err != nil {
				log.Fatal("could not create memory profile: ", err)
			}
			runtime.GC()                                      // GC，获取最新的数据信息
			if err := pprof.WriteHeapProfile(f); err != nil { // 写入内存信息
				log.Fatal("could not write memory profile: ", err)
			}
			f.Close()
		}

		if conn != nil {
			log.Println(conn.RemoteAddr(), "->", conn.LocalAddr(), " goroutine exit")
			conn.Close()
		}

		close(readChan)
		close(stopChan)
	}()

	var (
		requestMessage string
		stop           bool
		//		event          *serviceDispatch.Event
	)

	go this.readRegisteApply(conn, readChan, stopChan)

endForCycle:
	for {
		select {
		case requestMessage = <-readChan:

			//			log.Println(conn.RemoteAddr(), "->", conn.LocalAddr(), " recieve data=\r\n", requestMessage)
			this.endPointServiceRegiste(requestMessage)
		case stop = <-stopChan:
			if stop {
				runtime.GC()
				break endForCycle
			}
		}
	}

	return
}

func (this *EndPointServiceListener) readRegisteApply(conn net.Conn, readChan chan<- string, stopChan chan<- bool) {
	var (
		request string
		err     error
		reader  *bufio.Reader
	)

	reader = bufio.NewReader(conn)

	for {
		request, err = reader.ReadString('\n')

		if err != nil {
			log.Println("service dispatch service get request of proxy failure", err)
			break
		}

		readChan <- request
	}

	stopChan <- true
}

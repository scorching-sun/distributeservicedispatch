// endPointOverseeService project endPointOverseeService.go
package endPointOverseeService

import (
	"endPointOversee/endPointOverseeManager"
	"endPointOversee/endPointServiceRegisteApply"
	"errors"
	"log"
	"net"
	"strings"
	"time"
)

const (
	intervalRegisteInfoSend = time.Millisecond * 200

	msgHead                             string = "end-point oversee service:"
	msgGetEndPointAndServiceRunningInfo string = "not get information of end-point and service running"
)

var (
	errUDPMulticastAddress error = errors.New("udp multicast address split error")
)

type EndPointOverseeService struct{}

func NewEndPointOverseeService() *EndPointOverseeService {
	this := new(EndPointOverseeService)

	return this
}

func (this *EndPointOverseeService) Stop() error {
	var err error

	defer func() {
		if p := recover(); p != nil {
			if er, ok := p.(error); ok {
				err = er
			}
		}
	}()

	return err
}

func (this *EndPointOverseeService) Start() error {
	var err error

	defer func() {
		if p := recover(); p != nil {
			if er, ok := p.(error); ok {
				err = er
			}
		}
	}()

	go this.registeAndHealthCheckTask()

	return err
}

//registe and health check of end-point and service running information
func (this *EndPointOverseeService) registeAndHealthCheckTask() {
	registerApply := endPointServiceRegisteApply.NewEndPointServiceRegisteApply()

	//	manager := endPointOverseeManager.NewEndPointOverseeManager()
	targetAddrs := endPointOverseeManager.TargetAddresses()

	if targetAddrs == nil || len(targetAddrs) <= 0 {
		log.Println("can not config information of oversee end point and service")
		return
	}

	conns := make(map[string]net.Conn)

	for k, addr := range targetAddrs {
		conns[k] = registerApply.Sender(net.ParseIP(addr.RegisteServiceIP), addr.RegisteServicePort)
	}

	defer func() {
		for _, conn := range conns {
			if conn == nil {
				continue
			}

			conn.Close()
		}
	}()

	for {
		//睡眠
		time.Sleep(intervalRegisteInfoSend)

		packageGroupInfos := endPointOverseeManager.OverseeInfos()

		for k, packageInfo := range packageGroupInfos {
			if conns[k] == nil {
				if conns[k] = registerApply.Sender(net.ParseIP(targetAddrs[k].RegisteServiceIP), targetAddrs[k].RegisteServicePort); conns[k] == nil {
					continue
				}
			}

			if strings.TrimSpace(packageInfo) == "" {
				log.Println(strings.Join([]string{msgHead, msgGetEndPointAndServiceRunningInfo}, ""))
				continue
			}

			err := registerApply.RegisteApplyAndHealthReport(conns[k], strings.Join([]string{packageInfo, "\n"}, ""))

			if err != nil {
				log.Println(strings.Join([]string{msgHead, err.Error(), " ", "remote address=", k}, ""))

				conns[k] = nil

				continue
			}
		}
	}
}

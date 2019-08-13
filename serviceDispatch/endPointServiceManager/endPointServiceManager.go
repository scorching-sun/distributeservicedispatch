// endPointServiceManager project endPointServiceManager.go
/*
描述：
功能包括：
1.监听end point service发送信息；
2.注册end point service；
3.end point service可用监控；
4.end point service详细信息http restful api
5.end point service告警
*/
package endPointServiceManager

import (
	"os"
	"serviceDispatch"
	"serviceDispatch/endPointServiceAlarmer"
	"serviceDispatch/endPointServiceListener"
	"serviceDispatch/endPointServiceMonitor"
	"serviceDispatch/endPointServiceProvider"
	"serviceDispatch/endPointServiceRegister"
	"serviceDispatchDataModel/modelSDS"
	"strings"
)

const (
	envKeyRegisteListenModel string = "RegisteListenModel"
)

var (
	registeListenModel string
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

type EndPointServiceManager struct {
}

func NewEndPointServiceManager() *EndPointServiceManager {
	this := new(EndPointServiceManager)

	return this
}

//end point service管理
func (this *EndPointServiceManager) Management() {
	register := endPointServiceRegister.NewEndPointServiceRegister()
	monitor := endPointServiceMonitor.NewEndPointServiceMonitor()
	alarmer := endPointServiceAlarmer.NewEndPointServiceAlarmer()

	listener := endPointServiceListener.NewEndPointServiceListener()
	listener.Observers = make(map[serviceDispatch.Observer]struct{})

	//regist
	listener.Regist(register)
	listener.Regist(monitor)
	listener.Regist(alarmer)

	if registeListenModel == modelSDS.RegisteListenModelUDPMulticast {
		//监听end point service，并注册
		go listener.UDPListenEndPointServices()
	} else {
		go listener.TCPListenEndPointServices()
	}

	//end point service provider
	serviceProvider := endPointServiceProvider.NewEndPointServiceProvider()
	go serviceProvider.ProvideService()

	//服务可用监控
	monitor.ServiceAvailableMonitor()
}

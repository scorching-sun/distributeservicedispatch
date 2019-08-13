package serviceDispatchService

import (
	"serviceDispatch/endPointServiceManager"
	"serviceDispatch/serviceDispatchListener"
)

func (this *ServiceDispatchService) dispatchService() {
	//end point service管理
	manager := endPointServiceManager.NewEndPointServiceManager()
	go manager.Management()

	listener := serviceDispatchListener.NewServiceDispatchListener()
	listener.ServiceDispatchListen()
}

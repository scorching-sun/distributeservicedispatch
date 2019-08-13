// serviceDispatchStrategy project serviceDispatchStrategy.go
package serviceDispatchStrategy

import (
	"serviceDispatch/dalEndPointService"
	"serviceDispatchDataModel/modelServiceInfo"
)

type ServiceDispatchStrategy struct{}

func NewServiceDispatchStrategy() *ServiceDispatchStrategy {
	this := new(ServiceDispatchStrategy)

	return this
}

//调度可用服务
func (this *ServiceDispatchStrategy) DispatchAvailableService(request *modelServiceInfo.ServiceInfo) *modelServiceInfo.ServiceInfo {
	dal := dalEndPointService.NewDALEndPointService()

	availableServices := dal.AvailableServiceStrategy(request.NameSpace, request.ServiceProcessName, request.ServiceProvideMode)

	if availableServices == nil && len(availableServices) <= 0 {
		return nil
	}

	return availableServices[0]
}

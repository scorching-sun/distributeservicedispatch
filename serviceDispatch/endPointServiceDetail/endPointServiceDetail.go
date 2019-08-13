// endPointServiceDetail project endPointServiceDetail.go
package endPointServiceDetail

import (
	"serviceDispatch/dalEndPointService"
	"serviceDispatchDataModel/modelServiceInfo"
)

type EndPointServiceDetail struct{}

func NewEndPointServiceDetail() *EndPointServiceDetail {
	this := new(EndPointServiceDetail)

	return this
}

//调度服务可提供可用服务明细信息
func (this *EndPointServiceDetail) ServiceDetailInfo() (datas []*modelServiceInfo.ServiceInfo) {
	dal := dalEndPointService.NewDALEndPointService()

	datas = dal.ServiceDetailInfo()

	if datas == nil && len(datas) <= 0 {
		return nil
	}

	return
}

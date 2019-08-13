// endPointServiceRegister project endPointServiceRegister.go
package endPointServiceRegister

import (
	"serviceDispatch"
	"serviceDispatch/dalEndPointService"
	"serviceDispatchDataModel/modelEndPointOverseeInfo"
)

type EndPointServiceRegister struct {
}

func NewEndPointServiceRegister() *EndPointServiceRegister {
	this := new(EndPointServiceRegister)

	return this
}

//接收end-point oversee发送信息，注册并存储至本地数据库
func (this *EndPointServiceRegister) FollowProcess(event *serviceDispatch.Event) error {
	var (
		dal   *dalEndPointService.DALEndPointService
		datas []*modelEndPointOverseeInfo.EndPointOverseeInfo
		err   error
	)

	if datas, err = serviceDispatch.CheckRegisteData(event); err != nil {
		return err
	}

	dal = dalEndPointService.NewDALEndPointService()
	dal.RegisteAppend(datas)

	return nil
}

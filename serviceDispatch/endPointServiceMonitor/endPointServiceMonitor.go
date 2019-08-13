// endPointServiceMonitor project endPointServiceMonitor.go
/*
描述：
获取已注册的end point service，且service status为 running 的服务，
若其updatetime和service dispatch service所在host当前时间相差1 second，
则设置其service status为 stoped （这里须考虑end point host与service dispatch service host的时间同步），
并更新至数据库。
*/
package endPointServiceMonitor

import (
	"library/Format"
	"math"
	"serviceDispatch"
	"serviceDispatch/dalEndPointService"
	"serviceDispatchDataModel/modelEndPointOverseeInfo"
	"serviceDispatchDataModel/modelServiceInfo"
	"time"
)

const (
	monitorInterval = time.Millisecond * 10
	//unit : second
	monitorLimit float64 = 2
)

type EndPointServiceMonitor struct {
}

func NewEndPointServiceMonitor() *EndPointServiceMonitor {
	this := new(EndPointServiceMonitor)

	return this
}

//服务监控跟随处理
func (this *EndPointServiceMonitor) FollowProcess(event *serviceDispatch.Event) error {
	var (
		dal   *dalEndPointService.DALEndPointService
		datas []*modelEndPointOverseeInfo.EndPointOverseeInfo
		err   error
	)

	if datas, err = serviceDispatch.CheckRegisteData(event); err != nil {
		return err
	}

	dal = dalEndPointService.NewDALEndPointService()
	dal.ModifyServiceStat(datas)

	return nil
}

//服务可用监控
func (this *EndPointServiceMonitor) ServiceAvailableMonitor() {
	var (
		dal                 *dalEndPointService.DALEndPointService
		modifyServiceStats  []*modelServiceInfo.ServiceInfo
		datas               []*modelServiceInfo.ServiceInfo
		data                *modelServiceInfo.ServiceInfo
		nowTime, updateTime time.Time
		err                 error
		sub                 float64
	)

	dal = dalEndPointService.NewDALEndPointService()

	for {
		time.Sleep(monitorInterval)

		nowTime = time.Now().Local()

		datas = dal.ServiceStat()

		for _, data = range datas {
			updateTime, err = time.ParseInLocation(Format.TimeStampFormat0, data.UpdateTime, time.Local)

			if err != nil {
				panic(err)
			}

			sub = math.Abs(nowTime.Sub(updateTime).Seconds())

			//			log.Println("nowTime=", nowTime, "|", "updateTime=", data.UpdateTime, "monitor service stat discriminate=", sub, "|", monitorLimit)

			//			if data.ServiceStatus == modelServiceInfo.ServiceStatusRunning && sub >= monitorLimit {
			if sub >= monitorLimit {
				data.ServiceStatus = modelServiceInfo.ServiceStatusStoped

				modifyServiceStats = append(modifyServiceStats, data)
			}
		}

		if modifyServiceStats != nil && len(modifyServiceStats) > 0 {
			dal.ModifyServiceStatForMonitor(modifyServiceStats...)
			modifyServiceStats = nil
		}
	}
}

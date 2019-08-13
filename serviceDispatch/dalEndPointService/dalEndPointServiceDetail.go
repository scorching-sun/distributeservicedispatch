/*
描述：
使用http协议以restful风格提供api，展示当前已注册且可用的服务详细信息,以供调用者使用。
*/
package dalEndPointService

import (
	"serviceDispatchDataModel/modelServiceInfo"
)

const (
	serviceDetailSql string = `
	select 
	s.ip,
	s.namespace,
	s.serviceprocessname,
	s.port,
	s.serviceprovidemode,
	s.serviceextendinfo,
	ss.servicestatus
	from end_point_service as s
	inner join 
	end_point_servicestat as ss
	on s.ip=ss.ip and s.namespace=ss.namespace and s.serviceprocessname=ss.serviceprocessname and s.port=ss.port
	where 1=1
	and ss.servicestatus='running'
	`
)

//调度服务可提供可用服务明细信息,通过serviceextendinfo字段描述服务明细以供调用者获知
func (this *DALEndPointService) ServiceDetailInfo() (datas []*modelServiceInfo.ServiceInfo) {
	records, err := SqliteUtil.QueryByConn(Conn, serviceDetailSql)

	if err != nil {
		return nil
	}

	for _, record := range records {
		if data := this.convertService(record); data != nil {
			datas = append(datas, data)
		}
	}

	return
}

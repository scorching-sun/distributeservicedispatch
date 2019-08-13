/*
描述：
服务调度策略，获取已注册且服务当前为running状态，
按cpu总使用百分比最小，且cpu空闲百分比最大，
且memery总使用百分比最小，且memery空闲百分比最大，
且disk总使用百分比最小，且disk空闲百分比最大排序
*/
package dalEndPointService

import (
	//	"log"
	"serviceDispatchDataModel/modelServiceInfo"
	"strings"
)

const (
	availableServiceSql string = `
	select s.ip,s.namespace,s.serviceprocessname,s.port,s.serviceprovidemode,ss.servicestatus
	from end_point_service as s
	inner join 
	end_point_servicestat as ss
	on s.ip=ss.ip and s.namespace=ss.namespace and s.serviceprocessname=ss.serviceprocessname and s.port=ss.port
	inner join
	end_point_cpustat as cs
	on s.ip=cs.ip
	inner join
	end_point_memerystat as ms
	on s.ip=ms.ip
	inner join
	end_point_diskstat as ds
	on s.ip=ds.ip
	inner join
	end_point_netstat as ns
	on s.ip=ns.ip
	where 1=1
	and ss.servicestatus='running'
	`

	availableStrategySql string = `
	 
	order by 
	cs.totalused_percent,cs.idie_percent desc,
	ms.totalused_percent,ms.totalfree_percent desc,
	ds.totalused_percent,ds.totalfree_percent desc
	`
)

//可用服务策略
func (this *DALEndPointService) AvailableServiceStrategy(nameSpace, serviceProcessName, serviceProvideMode string) (datas []*modelServiceInfo.ServiceInfo) {
	sql := strings.Join([]string{availableServiceSql,
		" ", "and", " ", "s.namespace=", "'", nameSpace, "'",
		" ", "and", " ", "s.serviceprocessname=", "'", serviceProcessName, "'",
		" ", "and", " ", "s.serviceprovidemode=", "'", serviceProvideMode, "'",
		" ",
		availableStrategySql}, "")

	//	log.Println("availableStrategySql=\n", sql)

	records, err := SqliteUtil.QueryByConn(Conn, sql)

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

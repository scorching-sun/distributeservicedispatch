package dalEndPointService

import (
	"library/Format"
	"log"
	"serviceDispatchDataModel/modelEndPointOverseeInfo"
	"serviceDispatchDataModel/modelServiceInfo"
	"strings"
	"time"
)

const (
	end_point_servicestat string = "end_point_servicestat"
)

//end point servicestat添加sql
func (this *DALEndPointService) endPointServiceStatAppendContext(data *modelServiceInfo.ServiceInfo) string {
	if data == nil {
		return ""
	}

	nowTime := time.Now().Local()

	sql := strings.Join([]string{"insert into", " ", end_point_servicestat, "(",
		"ip", ",",
		"namespace", ",",
		"serviceprocessname", ",",
		"port", ",",
		"servicestatus", ",",
		"updatetime",
		")",
		"values", "(",
		"\"", data.ServiceProvideIP, "\"", ",",
		"\"", data.NameSpace, "\"", ",",
		"\"", data.ServiceProcessName, "\"", ",",
		data.ServiceProvidePort, ",",
		"\"", data.ServiceStatus, "\"", ",",
		"\"", nowTime.Format(Format.TimeStampFormat0), "\"",
		")", ";"}, "")

	return sql
}

//end point servicestat修改sql
func (this *DALEndPointService) endPointServiceStatModifyContext(data *modelServiceInfo.ServiceInfo) string {
	if data == nil {
		return ""
	}

	var updateTimeSetSql string

	if data.ServiceStatus == modelServiceInfo.ServiceStatusRunning {
		nowTime := time.Now().Local()

		updateTimeSetSql = strings.Join([]string{",",
			"updatetime=", "\"", nowTime.Format(Format.TimeStampFormat0), "\""}, "")
	}

	sql := strings.Join([]string{"update", " ", end_point_servicestat,
		" ", "set", " ",
		"servicestatus=", "\"", data.ServiceStatus, "\"", updateTimeSetSql,
		" ", "where",
		" ", "1=1", " ",
		"and", " ",
		"ip=", "\"", data.ServiceProvideIP, "\"",
		" ", "and", " ",
		"namespace=", "\"", data.NameSpace, "\"",
		" ", "and", " ",
		"serviceprocessname=", "\"", data.ServiceProcessName, "\"",
		" ", "and", " ",
		"port=", data.ServiceProvidePort,
		";"}, "")

	return sql
}

//end point servicestat修改
func (this *DALEndPointService) ModifyServiceStat(datas []*modelEndPointOverseeInfo.EndPointOverseeInfo) {
	var modifyServiceStats []*modelServiceInfo.ServiceInfo

	for _, data := range datas {
		modifyServiceStats = append(modifyServiceStats, data.Service)
	}

	if modifyServiceStats != nil && len(modifyServiceStats) > 0 {
		this.ModifyServiceStatForMonitor(modifyServiceStats...)
	}
}

//end point servicestat修改
func (this *DALEndPointService) ModifyServiceStatForMonitor(datas ...*modelServiceInfo.ServiceInfo) {
	var (
		err           error
		scriptContext string
	)

	for _, data := range datas {
		scriptContext = this.endPointServiceStatAppendContext(data)
		//		log.Println("monitor service stat append sql context=", scriptContext)

		if err = SqliteUtil.ExecuteByConn(Conn, scriptContext); err != nil {
			scriptContext = this.endPointServiceStatModifyContext(data)
			//			log.Println("monitor service stat modify sql context=", scriptContext)

			if err = SqliteUtil.ExecuteByConn(Conn, scriptContext); err != nil {
				log.Println("monitor service stat failure ", err)
			}
		}
	}
}

//end point servicestat修改
func (this *DALEndPointService) ServiceStat() (datas []*modelServiceInfo.ServiceInfo) {
	sql := strings.Join([]string{"select * from end_point_servicestat where 1=1 and servicestatus='", modelServiceInfo.ServiceStatusRunning, "';"}, "")

	records, err := SqliteUtil.QueryByConn(Conn, sql)

	if err != nil {
		return nil
	}

	if len(records) <= 0 {
		return nil
	}

	for _, record := range records {
		if data := this.convertService(record); data != nil {
			datas = append(datas, data)
		}
	}

	return
}

package dalEndPointService

import (
	"errors"
	"log"
	"serviceDispatchDataModel/modelEndPointOverseeInfo"
	"serviceDispatchDataModel/modelEndPointStat"
	"serviceDispatchDataModel/modelServiceInfo"
	"strconv"
	"strings"
	"sync"
)

const (
	//	end_point             string = "end_point"
	end_point_service    string = "end_point_service"
	end_point_cpustat    string = "end_point_cpustat"
	end_point_memerystat string = "end_point_memerystat"
	end_point_diskstat   string = "end_point_diskstat"
	end_point_netstat    string = "end_point_netstat"
)

var (
	mutexLock                 sync.Mutex
	errEndPointServiceInfoNil error = errors.New("service information is none")
)

/*

//end point添加sql
func (this *DALEndPointService) endPointAppendContext(data *modelServiceInfo.ServiceInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"insert into", " ", end_point, "(ip)",
		"values",
		"(", "\"", data.ServiceProvideIP, "\"", ")", ";"}, "")

	return sql
}

*/

//end point service添加sql
func (this *DALEndPointService) endPointServiceAppendContext(data *modelServiceInfo.ServiceInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"insert into", " ", end_point_service, "(",
		"ip", ",",
		"namespace", ",",
		"serviceprocessname", ",",
		"port", ",",
		"serviceprogrampath", ",",
		"serviceprovidemode",
		//		",", "serviceextendinfo",
		")",
		"values", "(",
		"\"", data.ServiceProvideIP, "\"", ",",
		"\"", data.NameSpace, "\"", ",",
		"\"", data.ServiceProcessName, "\"", ",",
		data.ServiceProvidePort, ",",
		"\"", data.ServiceProgramPath, "\"", ",",
		"\"", data.ServiceProvideMode, "\"",
		//		",", "\"", data.ServiceExtendInfo, "\"",
		")", ";"}, "")

	return sql
}

//end point service修改sql
func (this *DALEndPointService) endPointServiceModifyContext(data *modelServiceInfo.ServiceInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"update", " ", end_point_service,
		" ", "set", " ",
		"ip=", "\"", data.ServiceProvideIP, "\"", ",",
		"namespace=", "\"", data.NameSpace, "\"", ",",
		"serviceprocessname=", "\"", data.ServiceProcessName, "\"", ",",
		"port=", data.ServiceProvidePort, ",",
		"serviceprogrampath=", "\"", data.ServiceProgramPath, "\"", ",",
		"serviceprovidemode=", "\"", data.ServiceProvideMode, "\"",
		" ", "where", " ",
		"1=1", " ",
		"and", " ",
		"ip=", "\"", data.ServiceProvideIP, "\"",
		" ", "and", " ",
		"namespace=", "\"", data.NameSpace, "\"",
		" ", "and", " ",
		"serviceprocessname=", "\"", data.ServiceProcessName, "\"",
		" ", "and", " ",
		"port=", data.ServiceProvidePort,
		/*
			" ", "and", " ",
			"serviceprogrampath=", "\"", data.ServiceProgramPath, "\"",
			" ", "and", " ",
			"serviceprovidemode=", "\"", data.ServiceProvideMode, "\"",
		*/
		";"}, "")

	return sql
}

//end point cpustat添加sql
func (this *DALEndPointService) endPointCPUStatAppendContext(ip string, data *modelEndPointStat.CPUStatInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"insert into", " ", end_point_cpustat, "(",
		"ip", ",",
		"totalused_percent", ",",
		"used_percent", ",",
		"system_percent", ",",
		"nice_percent", ",",
		"idie_percent", ",",
		"iowait_percent", ",",
		"irq_percent", ",",
		"softirq_percent", ",",
		"steal_percent", ",",
		"guest_percent", ",",
		"guestnice_percent", ",",
		"stolen_percent",
		")",
		"values", "(",
		"\"", ip, "\"", ",",
		strconv.FormatFloat(data.TotalUsedPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.UsedPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.SystemPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.NicePercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.IDIEPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.IOWaitPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.IrqPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.SoftirqPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.StealPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.GuestPercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.GuestNicePercent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.StolenPercent, 'f', 6, 64),
		")", ";"}, "")

	return sql
}

//end point cpustat修改sql
func (this *DALEndPointService) endPointCPUStatModifyContext(ip string, data *modelEndPointStat.CPUStatInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"update", " ", end_point_cpustat,
		" ", "set", " ",
		"totalused_percent=", strconv.FormatFloat(data.TotalUsedPercent, 'f', 6, 64), ",",
		"used_percent=", strconv.FormatFloat(data.UsedPercent, 'f', 6, 64), ",",
		"system_percent=", strconv.FormatFloat(data.SystemPercent, 'f', 6, 64), ",",
		"nice_percent=", strconv.FormatFloat(data.NicePercent, 'f', 6, 64), ",",
		"idie_percent=", strconv.FormatFloat(data.IDIEPercent, 'f', 6, 64), ",",
		"iowait_percent=", strconv.FormatFloat(data.IOWaitPercent, 'f', 6, 64), ",",
		"irq_percent=", strconv.FormatFloat(data.IrqPercent, 'f', 6, 64), ",",
		"softirq_percent=", strconv.FormatFloat(data.SoftirqPercent, 'f', 6, 64), ",",
		"steal_percent=", strconv.FormatFloat(data.StealPercent, 'f', 6, 64), ",",
		"guest_percent=", strconv.FormatFloat(data.GuestPercent, 'f', 6, 64), ",",
		"guestnice_percent=", strconv.FormatFloat(data.GuestNicePercent, 'f', 6, 64), ",",
		"stolen_percent=", strconv.FormatFloat(data.StolenPercent, 'f', 6, 64),
		" ", "where", " ",
		"1=1", " ",
		"and", " ",
		"ip=", "\"", ip, "\"",
		";"}, "")

	return sql
}

//end point memerystat添加sql
func (this *DALEndPointService) endPointMemeryStatAppendContext(ip string, data *modelEndPointStat.MemeryStatInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"insert into", " ", end_point_memerystat, "(",
		"ip", ",",
		"totalused_percent", ",",
		"totalfree_percent", ",",
		"totalfree",
		")",
		"values", "(",
		"\"", ip, "\"", ",",
		strconv.FormatFloat(data.MemeryUsedPrecent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.MemeryFreePrecent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.MemeryFree, 'f', 6, 64),
		")", ";"}, "")

	return sql
}

//end point memerystat修改sql
func (this *DALEndPointService) endPointMemeryStatModifyContext(ip string, data *modelEndPointStat.MemeryStatInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"update", " ", end_point_memerystat,
		" ", "set", " ",
		"totalused_percent=", strconv.FormatFloat(data.MemeryUsedPrecent, 'f', 6, 64), ",",
		"totalfree_percent=", strconv.FormatFloat(data.MemeryFreePrecent, 'f', 6, 64), ",",
		"totalfree=", strconv.FormatFloat(data.MemeryFree, 'f', 6, 64),
		" ", "where", " ",
		"1=1", " ",
		"and", " ",
		"ip=", "\"", ip, "\"",
		";"}, "")

	return sql
}

//end point diskstat添加sql
func (this *DALEndPointService) endPointDiskStatAppendContext(ip string, data *modelEndPointStat.DiskStatInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"insert into", " ", end_point_diskstat, "(",
		"ip", ",",
		"totalused_percent", ",",
		"totalfree_percent", ",",
		"totalfree",
		")",
		"values", "(",
		"\"", ip, "\"", ",",
		strconv.FormatFloat(data.DiskUsedPrecent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.DiskFreePrecent, 'f', 6, 64), ",",
		strconv.FormatFloat(data.DiskFree, 'f', 6, 64),
		")", ";"}, "")

	return sql
}

//end point diskstat修改sql
func (this *DALEndPointService) endPointDiskStatModifyContext(ip string, data *modelEndPointStat.DiskStatInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"update", " ", end_point_diskstat,
		" ", "set", " ",
		"totalused_percent=", strconv.FormatFloat(data.DiskUsedPrecent, 'f', 6, 64), ",",
		"totalfree_percent=", strconv.FormatFloat(data.DiskFreePrecent, 'f', 6, 64), ",",
		"totalfree=", strconv.FormatFloat(data.DiskFree, 'f', 6, 64),
		" ", "where", " ",
		"1=1", " ",
		"and", " ",
		"ip=", "\"", ip, "\"",
		";"}, "")

	return sql
}

//end point netstat添加sql
func (this *DALEndPointService) endPointNetStatAppendContext(ip string, data *modelEndPointStat.NetStatInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"insert into", " ", end_point_netstat, "(",
		"ip", ",",
		"bytesrecv", ",",
		"bytessent",
		")",
		"values", "(",
		"\"", ip, "\"", ",",
		strconv.FormatFloat(data.RecieveBytesSize, 'f', 6, 64), ",",
		strconv.FormatFloat(data.SendBytesSize, 'f', 6, 64),
		")", ";"}, "")

	return sql
}

//end point netstat修改sql
func (this *DALEndPointService) endPointNetStatModifyContext(ip string, data *modelEndPointStat.NetStatInfo) string {
	if data == nil {
		return ""
	}

	sql := strings.Join([]string{"update", " ", end_point_netstat,
		" ", "set", " ",
		"bytesrecv=", strconv.FormatFloat(data.RecieveBytesSize, 'f', 6, 64), ",",
		"bytessent=", strconv.FormatFloat(data.SendBytesSize, 'f', 6, 64),
		" ", "where",
		" ", "1=1", " ",
		"and", " ",
		"ip=", "\"", ip, "\"",
		";"}, "")

	return sql
}

//添加注册
func (this *DALEndPointService) RegisteAppend(datas []*modelEndPointOverseeInfo.EndPointOverseeInfo) {
	var (
		scriptContextArr                                                                                 []string
		scriptContextService, scriptContextCpu, scriptContextMemery, scriptContextDisk, scriptContextNet string
		err                                                                                              error
	)

	for _, data := range datas {
		ip := data.Service.ServiceProvideIP
		scriptContextArr = nil

		scriptContextService = this.endPointServiceAppendContext(data.Service)
		scriptContextArr = append(scriptContextArr, scriptContextService)

		scriptContextCpu = this.endPointCPUStatAppendContext(ip, data.EndPointStat.CPUStat)
		scriptContextArr = append(scriptContextArr, scriptContextCpu)

		scriptContextMemery = this.endPointMemeryStatAppendContext(ip, data.EndPointStat.MemeryStat)
		scriptContextArr = append(scriptContextArr, scriptContextMemery)

		scriptContextDisk = this.endPointDiskStatAppendContext(ip, data.EndPointStat.DiskStat)
		scriptContextArr = append(scriptContextArr, scriptContextDisk)

		scriptContextNet = this.endPointNetStatAppendContext(ip, data.EndPointStat.NetStat)
		scriptContextArr = append(scriptContextArr, scriptContextNet)

		//		log.Println("registe append end point service sql context= \r\n", scriptContextArr)

		if err = SqliteUtil.ExecuteByConn(Conn, scriptContextArr...); err != nil {
			//		if _, err = SqliteUtil.Execute(scriptContextArr...); err != nil {
			scriptContextArr = nil

			scriptContextService = this.endPointServiceModifyContext(data.Service)
			scriptContextArr = append(scriptContextArr, scriptContextService)

			scriptContextCpu = this.endPointCPUStatModifyContext(ip, data.EndPointStat.CPUStat)
			scriptContextArr = append(scriptContextArr, scriptContextCpu)

			scriptContextMemery = this.endPointMemeryStatModifyContext(ip, data.EndPointStat.MemeryStat)
			scriptContextArr = append(scriptContextArr, scriptContextMemery)

			scriptContextDisk = this.endPointDiskStatModifyContext(ip, data.EndPointStat.DiskStat)
			scriptContextArr = append(scriptContextArr, scriptContextDisk)

			scriptContextNet = this.endPointNetStatModifyContext(ip, data.EndPointStat.NetStat)
			scriptContextArr = append(scriptContextArr, scriptContextNet)

			//			log.Println("registe modify end point service sql context= \r\n", scriptContextArr)

			if err = SqliteUtil.ExecuteByConn(Conn, scriptContextArr...); err != nil {
				//			if _, err = SqliteUtil.Execute(scriptContextArr...); err != nil {
				log.Println("registe failure ", err)
			}

		}
	}
}

// endPointHostInfoCollector project endPointHostInfoCollector.go
/*
获取宿主机资源运行信息，信息采集时间间隔200millisecond
*/
package endPointHostInfoCollector

import (
	"library/OS"
	"library/OS/CPU"
	"library/OS/Disk"
	"library/OS/Memery"
	"library/OS/Net"
	"serviceDispatchDataModel/modelEndPointStat"
	"time"
)

const (
	interval = 200 * time.Millisecond
)

func EndPointStatInfo() *modelEndPointStat.EndPointStat {
	endPointStat := new(modelEndPointStat.EndPointStat)

	//cpu
	cpuStatInfo := new(modelEndPointStat.CPUStatInfo)

	cpustats, _ := CPU.Percent(interval, false)

	if cpustats != nil && len(cpustats) > 0 {
		cpuStatInfo.CPU = cpustats[0].CPU

		cpuStatInfo.UsedPercent = cpustats[0].User
		cpuStatInfo.SystemPercent = cpustats[0].System
		cpuStatInfo.NicePercent = cpustats[0].Nice
		cpuStatInfo.IDIEPercent = cpustats[0].Idle
		cpuStatInfo.IOWaitPercent = cpustats[0].Iowait
		cpuStatInfo.IrqPercent = cpustats[0].Irq
		cpuStatInfo.SoftirqPercent = cpustats[0].Softirq
		cpuStatInfo.StealPercent = cpustats[0].Steal

		cpuStatInfo.GuestPercent = cpustats[0].Guest
		cpuStatInfo.GuestNicePercent = cpustats[0].GuestNice
		cpuStatInfo.StolenPercent = cpustats[0].Stolen

		cpuStatInfo.TotalUsedPercent = cpuStatInfo.UsedPercent + cpuStatInfo.SystemPercent + cpuStatInfo.NicePercent +
			cpuStatInfo.IOWaitPercent + cpuStatInfo.IrqPercent + cpuStatInfo.SoftirqPercent + cpuStatInfo.StealPercent +
			cpuStatInfo.GuestPercent + cpuStatInfo.GuestNicePercent + cpuStatInfo.StolenPercent

		endPointStat.CPUStat = cpuStatInfo
	}

	//memery
	memeryrunning := new(modelEndPointStat.MemeryStatInfo)

	mem := Memery.NewMemeryStat()

	if mem != nil {
		memeryrunning.MemeryUsedPrecent = mem.UsagePercent()
		memeryrunning.MemeryFreePrecent = 100 - memeryrunning.MemeryUsedPrecent
		memeryrunning.MemeryFree = mem.Free(OS.UnitMib)

		endPointStat.MemeryStat = memeryrunning
	}

	//disk
	diskrunning := new(modelEndPointStat.DiskStatInfo)

	disk := Disk.NewDiskStat()

	if disk != nil {
		diskrunning.DiskUsedPrecent = disk.UsagePercent()
		diskrunning.DiskFreePrecent = 100 - diskrunning.DiskUsedPrecent
		diskrunning.DiskFree = disk.Free(OS.UnitMib)

		endPointStat.DiskStat = diskrunning
	}

	//net
	netrunning := new(modelEndPointStat.NetStatInfo)

	netIOStats, err := Net.IOCountersStats(false)

	if err == nil {
		netrunning.RecieveBytesSize = float64(netIOStats[0].BytesRecv)
		netrunning.SendBytesSize = float64(netIOStats[0].BytesSent)

		endPointStat.NetStat = netrunning
	}

	return endPointStat
}

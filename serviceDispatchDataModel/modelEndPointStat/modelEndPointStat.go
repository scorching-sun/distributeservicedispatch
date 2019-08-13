// modelEndPointStat project modelEndPointStat.go
package modelEndPointStat

type CPUStatInfo struct {
	//cpu总占用率
	CPU              string  `json:"CPU"`
	TotalUsedPercent float64 `json:"TotalUsedPercent"`
	UsedPercent      float64 `json:"UsedPercent"`
	SystemPercent    float64 `json:"SystemPercent"`
	NicePercent      float64 `json:"NicePercent"`
	IDIEPercent      float64 `json:"IDIEPercent"`
	IOWaitPercent    float64 `json:"IOWaitPercent"`
	IrqPercent       float64 `json:"IrqPercent"`
	SoftirqPercent   float64 `json:"SoftirqPercent"`
	StealPercent     float64 `json:"StealPercent"`

	GuestPercent     float64 `json:"GuestPercent"`
	GuestNicePercent float64 `json:"GuestNicePercent"`
	StolenPercent    float64 `json:"StolenPercent"`
}

type MemeryStatInfo struct {
	//memery总占用率
	MemeryUsedPrecent float64 `json:"MemeryUsedPrecent"`
	//memery总available率
	MemeryFreePrecent float64 `json:"MemeryFreePrecent"`
	MemeryFree        float64 `json:"MemeryFree"`
}

type DiskStatInfo struct {
	//磁盘总占用率
	DiskUsedPrecent float64 `json:"DiskUsedPrecent"`
	//磁盘总available率
	DiskFreePrecent float64 `json:"DiskFreePrecent"`
	DiskFree        float64 `json:"DiskFree"`
}

type NetStatInfo struct {
	//接收字节数
	RecieveBytesSize float64 `json:"RecieveBytesSize"`
	//发送字节数
	SendBytesSize float64 `json:"SendBytesSize"`
}

type EndPointStat struct {
	//cpu运行信息
	CPUStat *CPUStatInfo `json:"CPUStat"`
	//memery运行信息
	MemeryStat *MemeryStatInfo `json:"MemeryStat"`
	//disk运行信息
	DiskStat *DiskStatInfo `json:"DiskStat"`
	//net运行信息
	NetStat *NetStatInfo `json:"NetStat"`
}

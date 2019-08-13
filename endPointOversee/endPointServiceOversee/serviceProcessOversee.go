/*
服务进程
*/
package endPointServiceOversee

import (
	"library/OSInfo"
	"log"
)

type ServiceProcessOversee struct {
}

const (
	ProcessStatusStoped  string = "stoped"
	ProcessStatusRunning string = "running"
)

func NewServiceProcessOversee() *ServiceProcessOversee {
	this := new(ServiceProcessOversee)

	return this
}

//根据服务进程配置信息获取其状态
func (this *ServiceProcessOversee) ProcessStatusByProcInfo(processName, processPath string) string {
	procUtil := OSInfo.NewProcess()

	proc := procUtil.ProcessByProcInfo(processName, processPath)

	log.Println("service process oversee :\r\n",
		"processName=", processName, "\r\n",
		"processPath=", processPath, "\r\n",
		"proc=", proc)

	if proc == nil {
		return ProcessStatusStoped
	}

	return ProcessStatusRunning
}

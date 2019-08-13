// endPointOverseeManager project endPointOverseeManager.go
package endPointOverseeManager

import (
	"endPointOversee/endPointHostInfoCollector"
	"endPointOversee/endPointServiceInfoCollector"
	"endPointOversee/endPointServiceOversee"
	"library/Json"
	"log"
	"serviceDispatchDataModel/modelEndPointOverseeInfo"
	"serviceDispatchDataModel/modelEndPointStat"
	"serviceDispatchDataModel/modelServiceInfo"
	"strconv"
	"strings"
)

type EndPointOverseeManager struct {
}

var (
	serviceConf *modelServiceInfo.EndPointServiceConf
)

func NewEndPointOverseeManager() *EndPointOverseeManager {
	this := new(EndPointOverseeManager)

	return this
}

func init() {
	serviceConf = endPointServiceInfoCollector.EndPointServiceConfig()
}

func getKey(ip string, port int) string {
	return strings.Join([]string{ip, ":", strconv.Itoa(port)}, "")
}

func buildEndPointOverseePackage(serviceInfos []*modelServiceInfo.ServiceInfo, endPointStat *modelEndPointStat.EndPointStat) *modelEndPointOverseeInfo.EndPointOverseePackage {
	msgPackage := new(modelEndPointOverseeInfo.EndPointOverseePackage)

	head := new(modelEndPointOverseeInfo.OverseeMessageHead)
	head.MessageType = "EndPointAndService"
	head.MessageVersion = "v1.0.0"
	msgPackage.Head = head

	var infos []*modelEndPointOverseeInfo.EndPointOverseeInfo

	processOversee := endPointServiceOversee.NewServiceProcessOversee()

	for _, serviceInfo := range serviceInfos {
		processStatus := processOversee.ProcessStatusByProcInfo(serviceInfo.ServiceProcessName, serviceInfo.ServiceProgramPath)
		serviceInfo.ServiceStatus = processStatus

		info := new(modelEndPointOverseeInfo.EndPointOverseeInfo)
		info.Service = serviceInfo
		info.EndPointStat = endPointStat

		infos = append(infos, info)
	}

	body := new(modelEndPointOverseeInfo.OverseeMessageBody)
	body.EndPointOverseeInfos = infos

	msgPackage.Body = body

	return msgPackage
}

//终结点监视信息（包括：宿主机资源运行信息和服务运行信息）
func overseeInfos() map[string]*modelEndPointOverseeInfo.EndPointOverseePackage {
	if serviceConf == nil {
		return nil
	}

	serviceGroup := make(map[string][]*modelServiceInfo.ServiceInfo)

	for _, relation := range serviceConf.RegisteAddrAndServices {
		key := getKey(relation.RegisteServiceIP, relation.RegisteServicePort)

		for _, serviceInfo := range serviceConf.ServiceInfos {
			if relation.ServiceProvideIP == serviceInfo.ServiceProvideIP &&
				relation.ServiceProvidePort == serviceInfo.ServiceProvidePort &&
				relation.NameSpace == serviceInfo.NameSpace &&
				relation.ServiceProcessName == serviceInfo.ServiceProcessName {

				serviceGroup[key] = append(serviceGroup[key], serviceInfo)
			}
		}
	}

	endPointStat := endPointHostInfoCollector.EndPointStatInfo()

	packageMap := make(map[string]*modelEndPointOverseeInfo.EndPointOverseePackage)

	for key, serviceInfos := range serviceGroup {
		packageMap[key] = buildEndPointOverseePackage(serviceInfos, endPointStat)
	}

	return packageMap
}

//终结点监视信息
/*
返回数据json格式，如下:
[
{
	"ServiceInfo":{},
	"EndPointStat":{}
},
{
	"ServiceInfo":{},
	"EndPointStat":{}
}
]
*/
func OverseeInfos() map[string]string {
	infos := overseeInfos()

	if infos == nil || len(infos) <= 0 {
		return nil
	}

	maps := make(map[string]string)

	for k, info := range infos {
		if info == nil {
			continue
		}

		jsons, err := Json.StructToJsonString(info)

		if err != nil {
			//			panic(err)
			log.Println("OverseeInfos", " ", " exception :", err)
		}

		maps[k] = jsons
	}

	return maps
}

func TargetAddresses() map[string]*modelServiceInfo.TargetAddress {
	if serviceConf == nil {
		return nil
	}

	addrMap := make(map[string]*modelServiceInfo.TargetAddress)

	for _, addr := range serviceConf.TargetAddresses {
		key := getKey(addr.RegisteServiceIP, addr.RegisteServicePort)

		addrMap[key] = addr
	}

	return addrMap
}

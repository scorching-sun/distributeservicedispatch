// modelEndPointOverseeInfo project modelEndPointOverseeInfo.go
package modelEndPointOverseeInfo

import (
	"serviceDispatchDataModel/modelEndPointStat"
	"serviceDispatchDataModel/modelServiceInfo"
)

type EndPointOverseeInfo struct {
	Service      *modelServiceInfo.ServiceInfo   `json:"ServiceInfo"`
	EndPointStat *modelEndPointStat.EndPointStat `json:"EndPointStat"`
}

type OverseeMessageHead struct {
	MessageType    string `json:"MessageType"`
	MessageVersion string `json:"MessageVersion"`
}

type OverseeMessageBody struct {
	EndPointOverseeInfos []*EndPointOverseeInfo `json:"EndPointOverseeInfos"`
}

type EndPointOverseePackage struct {
	Head *OverseeMessageHead `json:"head"`
	Body *OverseeMessageBody `json:"body"`
}

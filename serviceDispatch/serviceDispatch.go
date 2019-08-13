package serviceDispatch

import (
	"errors"
	"library/Json"
	"serviceDispatchDataModel/modelEndPointOverseeInfo"
)

type Event struct {
	Data string `json:"data"`
}

type Observer interface {
	//跟进处理
	FollowProcess(*Event) error
}

// 被观察的对象接口
type Subject interface {
	//注册观察者
	Regist(Observer)
	//注销观察者
	Deregist(Observer)

	//通知观察者事件
	Notify(*Event)
}

var (
	errRecievePackage     error = errors.New("recieve data exception with registe")
	errRecieveRegisteData error = errors.New("not registe data")
)

func CheckRegisteData(event *Event) ([]*modelEndPointOverseeInfo.EndPointOverseeInfo, error) {
	if event == nil {
		return nil, errRecievePackage
	}

	var (
		info *modelEndPointOverseeInfo.EndPointOverseePackage
		err  error
	)

	//
	info = new(modelEndPointOverseeInfo.EndPointOverseePackage)
	err = Json.JsonStringToStruct(event.Data, &info)

	if err != nil {
		return nil, err
	}

	if info.Body == nil && info.Body.EndPointOverseeInfos == nil && len(info.Body.EndPointOverseeInfos) <= 0 {
		return nil, errRecieveRegisteData
	}

	return info.Body.EndPointOverseeInfos, nil
}

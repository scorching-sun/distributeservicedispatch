// modelServiceInfo project modelServiceInfo.go
package modelServiceInfo

const (
	ServiceStatusRunning string = "running"
	ServiceStatusStoped  string = "stoped"
)

type ServiceInfo struct {
	//服务提供ip
	ServiceProvideIP string `json:"ServiceProvideIP"`
	//服务提供端口
	ServiceProvidePort string `json:"ServiceProvidePort"`
	//命名空间，以区分服务分类或系统
	NameSpace string `json:"NameSpace"`
	//服务进程名，用以监控服务进程运行状态
	ServiceProcessName string `json:"ServiceProcessName"`
	//服务程序路径，包含程序文件名的完整路径
	ServiceProgramPath string `json:"ServiceProgramPath"`
	//服务提供方式，包括:grpc、restful
	ServiceProvideMode string `json:"ServiceProvideMode"`
	//服务状态，包括：running(运行)|stoped(停止)
	ServiceStatus string `json:"ServiceStatus"`
	UpdateTime    string `json:"-"`
	//服务扩展信息
	ServiceExtendInfo string `json:"ServiceExtendInfo"`
}

type RegisteAddrAndService struct {
	//注册服务ip
	RegisteServiceIP string `json:"RegisteServiceIP"`
	//注册服务port
	RegisteServicePort int `json:"RegisteServicePort"`
	//服务提供ip
	ServiceProvideIP string `json:"ServiceProvideIP"`
	//服务提供端口
	ServiceProvidePort string `json:"ServiceProvidePort"`
	//命名空间，以区分服务分类或系统
	NameSpace string `json:"NameSpace"`
	//服务进程名，用以监控服务进程运行状态
	ServiceProcessName string `json:"ServiceProcessName"`
}

type TargetAddress struct {
	RegisteServiceIP   string `json:"RegisteServiceIP"`
	RegisteServicePort int    `json:"RegisteServicePort"`
}

type EndPointServiceConf struct {
	TargetAddresses        []*TargetAddress         `json:"TargetAddresses"`
	RegisteAddrAndServices []*RegisteAddrAndService `json:"RegisteAddrAndServices"`
	ServiceInfos           []*ServiceInfo           `json:"EndPointServiceInfos"`
}

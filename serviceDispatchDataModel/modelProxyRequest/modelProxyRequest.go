// modelProxyRequest project modelProxyRequest.go
package modelProxyRequest

type SDSRequestInfo struct {
	//调度服务ip
	DispatchServiceIP string `json:"ServiceProvideIP"`
	//调度服务port
	DispatchServicePort string `json:"ServiceProvidePort"`
	//命名空间，以区分服务分类或系统
	NameSpace string `json:"NameSpace"`
	//服务进程名，用以监控服务进程运行状态
	ServiceProcessName string `json:"ServiceProcessName"`
	//服务提供方式，包括:grpc、restful
	ServiceProvideMode string `json:"ServiceProvideMode"`
}

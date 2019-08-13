// endPointServiceInfoCollector project doc.go

/*
endPointServiceInfoCollector document

微服务可看作功能单一类库，以不同形式（包括:http restful、rpc、grpc、asp.net webservice、asp.net webapi）
提供对外（异构程序语言系统）服务，本功能通过系统环境变量获取一个或多个服务配置信息
配置信息包含，如下信息：
命名空间-区分不同类型（或系统）的服务
服务进程名-用于endpointmonitor监视服务运行状况及其资源占用情况
服务提供方式-自分布式服务调度版本1.0开始,支持的服务提供方式包括:grpc、restful
服务提供端口-
服务扩展信息-提供服务信息的描述（如：服务名或服务url、http请求方法（如:get、post等）、输入参数及其数据格式、返回值及其数据格式等），以json格式描述，且无约定json约定格式。

可配置多个服务，其信息数据格式举例如下：
[
	{
	"NameSpace":"NameSpace.0",
	"ServiceProcessName":"processName.0",
	"ServiceProgramPath":"/usr/local/program.0",
	"ServiceProvideMode":"grpc",
	"ServiceProvideIP":"0.0.0.0",
	"ServiceProvidePort":"8080",
	"ServiceExtendInfo":[]
	},
	{
	"NameSpace":"NameSpace.1",
	"ServiceProcessName":"processName.1",
	"ServiceProgramPath":"/usr/local/program.1",
	"ServiceProvideMode":"restful",
	"ServiceProvideIP":"0.0.0.0",
	"ServiceProvidePort":"8081",
	"ServiceExtendInfo":[]
	}
]
*/
package endPointServiceInfoCollector

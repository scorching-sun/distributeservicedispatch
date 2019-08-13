/*
服务配置信息可以配置多个服务，当以系统环境变量提供情况时，此功能通过获取系统环境变量采集服务信息，
但此功能不进行服务配置信息验证，使用系统环境变量为 ServiceConfigInfo 。
*/
package endPointServiceInfoCollector

import (
	"library/Json"
	"log"
	"os"
	"serviceDispatchDataModel/modelServiceInfo"
	"strings"
)

const (
	EndPointServiceConf string = "EndPointServiceConf"
)

func EndPointServiceConfig() (config *modelServiceInfo.EndPointServiceConf) {
	conf := os.Getenv(EndPointServiceConf)

	if strings.TrimSpace(conf) == "" {
		return nil
	}

	log.Println("service config information :\r\n", conf)

	err := Json.JsonStringToStruct(conf, &config)

	if err != nil {
		log.Println("service config information collector error=", err)
		return nil
	}

	log.Println("service config information model :\r\n", config)

	return
}

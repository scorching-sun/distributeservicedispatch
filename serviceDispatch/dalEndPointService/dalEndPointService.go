// dalEndPointService project dalEndPointService.go
package dalEndPointService

import (
	"library/Path"
	"library/sqliteHelper"
	"strconv"
	"strings"

	"serviceDispatchDataModel/modelServiceInfo"

	"github.com/gohouse/gorose"
)

var (
	//	DbConfig = map[string]interface{}{
	//		"Default":         "serviceDispatch", // 默认数据库配置
	//		"SetMaxOpenConns": 50,                // (连接池)最大打开的连接数，默认值为0表示不限制
	//		"SetMaxIdleConns": -1,                // (连接池)闲置的连接数, 默认-1
	//		"Connections": map[string]map[string]string{
	//			"serviceDispatch": {
	//				//"host":     "localhost",
	//				//"username": "root",
	//				//"password": "",
	//				//"port":     "3306",
	//				"database": strings.Join([]string{Path.GetCurrentPath(), "data/db_sds.sqlite"}, ""),
	//				"prefix":   "",
	//				//"charset":  "utf8",
	//				//"protocol": "tcp",
	//				"driver": "sqlite3",
	//			},
	//		},
	//	}

	DbConfig = &gorose.DbConfigSingle{
		Driver:          "sqlite3",                                                               // 驱动: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,                                                                    // 是否开启sql日志
		SetMaxOpenConns: 0,                                                                       // (连接池)最大打开的连接数，默认值为0表示不限制
		SetMaxIdleConns: 0,                                                                       // (连接池)闲置的连接数
		Prefix:          "",                                                                      // 表前缀
		Dsn:             strings.Join([]string{Path.GetCurrentPath(), "data/db_sds.sqlite"}, ""), // 数据库链接
	}
)

var (
	SqliteUtil *sqliteHelper.SqliteUtil
	Conn       *gorose.Connection
)

func init() {
	SqliteUtil = sqliteHelper.NewSqliteUtilNoneConf()
	Conn = SqliteUtil.Connect(DbConfig)
}

type DALEndPointService struct {
}

func NewDALEndPointService() *DALEndPointService {
	this := new(DALEndPointService)

	return this
}

//记录转换为对象
func (this *DALEndPointService) convertService(record map[string]interface{}) *modelServiceInfo.ServiceInfo {
	if len(record) <= 0 {
		return nil
	}

	data := new(modelServiceInfo.ServiceInfo)

	if _, ok := record["ip"]; ok {
		data.ServiceProvideIP = record["ip"].(string)
	}

	if _, ok := record["namespace"]; ok {
		data.NameSpace = record["namespace"].(string)
	}

	if _, ok := record["serviceprocessname"]; ok {
		data.ServiceProcessName = record["serviceprocessname"].(string)
	}

	if _, ok := record["port"]; ok {
		data.ServiceProvidePort = strconv.Itoa(int(record["port"].(int64)))
	}

	if _, ok := record["serviceprogrampath"]; ok {
		if record["serviceprogrampath"] != nil {
			data.ServiceProgramPath = record["serviceprogrampath"].(string)
		}
	}

	if _, ok := record["serviceprovidemode"]; ok {
		if record["serviceprovidemode"] != nil {
			data.ServiceProvideMode = record["serviceprovidemode"].(string)
		}
	}

	if _, ok := record["serviceextendinfo"]; ok {
		if record["serviceextendinfo"] != nil {
			data.ServiceExtendInfo = record["serviceextendinfo"].(string)
		}
	}

	if _, ok := record["servicestatus"]; ok {
		if record["servicestatus"] != nil {
			data.ServiceStatus = record["servicestatus"].(string)
		}
	}

	if _, ok := record["updatetime"]; ok {
		if record["updatetime"] != nil {
			data.UpdateTime = record["updatetime"].(string)
		}
	}

	return data
}

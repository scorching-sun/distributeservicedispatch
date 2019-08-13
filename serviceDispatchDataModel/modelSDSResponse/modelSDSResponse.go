// modelSDSResponse project modelSDSResponse.go
package modelSDSResponse

type SDSResponseInfo struct {
	//可用服务ip
	AvailableServiceIP string `json:"ServiceProvideIP"`
	//可用服务port
	AvailableServicePort string `json:"ServiceProvidePort"`
}

package endPointServiceProvider

import (
	"fmt"
	"library/Json"
	"log"
	"net/http"
	"serviceDispatch/endPointServiceDetail"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, use the service dispatching service!")
}

func ServiceDetail(w http.ResponseWriter, r *http.Request) {
	log.Println("debug--------")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	serviceDetail := endPointServiceDetail.NewEndPointServiceDetail()
	datas := serviceDetail.ServiceDetailInfo()

	jsonConvert := new(Json.JsonConvert)

	buff, err := jsonConvert.StructToBytes(datas)

	if err != nil {
		w.Write([]byte("get available service detail failure"))
	}

	w.Write(buff)
}
